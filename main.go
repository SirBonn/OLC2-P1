package main

import (
	"compiler/errors"
	parser "compiler/parser"
	"fmt"

	// "go/ast"
	"io/ioutil"
	"strings"

	"compiler/ast"

	// "compiler/semantic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

type IDE struct {
	window      fyne.Window
	codeEntry   *widget.Entry
	outputEntry *widget.Entry
	currentFile string
	app         fyne.App
	// Componentes para reportes
	errorTable  *errors.ErrorTable
	symbolTable interface{} // Por ahora interface{}, después será *semantic.SymbolTable
	astRoot     ast.Node    // Por ahora interface{}, después será ast.Node
}

func main() {
	a := app.NewWithID("com.vlang.ide")
	w := a.NewWindow("🍒 V-Lang Cherry IDE")

	ide := &IDE{
		window: w,
		app:    a,
	}

	content := ide.createMainContent()
	mainMenu := ide.createMenu()
	w.SetMainMenu(mainMenu)

	w.SetContent(content)
	w.Resize(fyne.NewSize(1000, 700))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func (ide *IDE) createMainContent() fyne.CanvasObject {
	// Editor de código
	ide.codeEntry = widget.NewMultiLineEntry()
	ide.codeEntry.Wrapping = fyne.TextTruncate
	ide.codeEntry.SetPlaceHolder("// Escribe tu código V-Lang Cherry aquí...")

	// Consola de salida
	ide.outputEntry = widget.NewMultiLineEntry()
	ide.outputEntry.Wrapping = fyne.TextWrapWord
	ide.outputEntry.Disable()

	// Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("📝 Editor", container.NewVScroll(ide.codeEntry)),
		container.NewTabItem("🖥️ Consola", container.NewVScroll(ide.outputEntry)),
	)

	// Toolbar
	toolbar := ide.createToolbar()

	return container.NewBorder(toolbar, nil, nil, nil, tabs)
}

func (ide *IDE) createMenu() *fyne.MainMenu {
	// Menú Archivo
	fileMenu := fyne.NewMenu("Archivo",
		fyne.NewMenuItem("Nuevo", func() { ide.newFile() }),
		fyne.NewMenuItem("Abrir...", func() { ide.openFile() }),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Guardar", func() { ide.saveFile() }),
		fyne.NewMenuItem("Guardar como...", func() { ide.saveFileAs() }),
	)

	// Menú Herramientas
	toolsMenu := fyne.NewMenu("Herramientas",
		fyne.NewMenuItem("Ejecutar", func() { ide.runCode() }),
	)

	// Menú Reportes - ahora con todos los métodos implementados
	reportsMenu := fyne.NewMenu("Reportes",
		fyne.NewMenuItem("Reporte de Errores", func() { ide.showErrorsReport() }),
		fyne.NewMenuItem("Tabla de Símbolos", func() { ide.showSymbolTableReport() }),
		fyne.NewMenuItem("Reporte AST", func() { ide.showASTReport() }),
	)

	return fyne.NewMainMenu(fileMenu, toolsMenu, reportsMenu)
}

func (ide *IDE) createToolbar() fyne.CanvasObject {
	runBtn := widget.NewButtonWithIcon("Ejecutar", theme.MediaPlayIcon(), func() {
		ide.runCode()
	})
	runBtn.Importance = widget.HighImportance

	return container.NewHBox(runBtn)
}

// Funciones de archivo
func (ide *IDE) newFile() {
	ide.codeEntry.SetText("")
	ide.currentFile = ""
	ide.outputEntry.SetText("")
}

func (ide *IDE) openFile() {
	dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader == nil {
			return
		}
		defer reader.Close()

		data, err := ioutil.ReadAll(reader)
		if err != nil {
			dialog.ShowError(err, ide.window)
			return
		}

		ide.codeEntry.SetText(string(data))
		ide.currentFile = reader.URI().Path()
	}, ide.window)
}

func (ide *IDE) saveFile() {
	if ide.currentFile == "" {
		ide.saveFileAs()
		return
	}

	err := ioutil.WriteFile(ide.currentFile, []byte(ide.codeEntry.Text), 0644)
	if err != nil {
		dialog.ShowError(err, ide.window)
	}
}

func (ide *IDE) saveFileAs() {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil || writer == nil {
			return
		}
		defer writer.Close()

		_, err = writer.Write([]byte(ide.codeEntry.Text))
		if err != nil {
			dialog.ShowError(err, ide.window)
			return
		}

		ide.currentFile = writer.URI().Path()
	}, ide.window)
}

func (ide *IDE) runCode() {
	code := ide.codeEntry.Text
	ide.outputEntry.SetText("🔄 Compilando...\n\n")

	// === FASE 1: ANÁLISIS LÉXICO ===
	lexicalErrs := errors.NewLexicalErrorListener()
	lexer := parser.NewVlangLexer(antlr.NewInputStream(code))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// === FASE 2: ANÁLISIS SINTÁCTICO ===
	p := parser.NewVlangParser(tokens)
	p.BuildParseTrees = true
	syntaxErrs := errors.NewSyntaxErrorListener(lexicalErrs.ErrorTable)
	p.RemoveErrorListeners()
	p.AddErrorListener(syntaxErrs)
	tree := p.Programa()

	ide.errorTable = lexicalErrs.ErrorTable

	if ide.errorTable.HasErrors() {
		ide.showErrors()
		return
	}

	ide.outputEntry.SetText("✅ Análisis léxico y sintáctico completado\n")

	// === FASE 3: CONSTRUCCIÓN DEL AST ===
	ide.outputEntry.SetText(ide.outputEntry.Text + "🔨 Construyendo AST...\n")

	astBuilder := NewASTBuilder()
	astProgram, err := astBuilder.Build(tree)

	if err != nil {
		ide.outputEntry.SetText(ide.outputEntry.Text + fmt.Sprintf("❌ Error al construir el AST: %v\n", err))
		return
	}

	// NUEVO: Realizar análisis semántico
	semanticAnalyzer := ast.NewSemanticAnalyzer()
	err = semanticAnalyzer.Analyze(astProgram)
	if err != nil {
		// Imprimir errores semánticos
		for _, e := range semanticAnalyzer.GetSymbolTable().GetErrors() {
			fmt.Println("Semantic Error:", e)
		}
		ide.outputEntry.SetText(ide.outputEntry.Text + fmt.Sprintf("❌ Error: %v\n", err))
	}

	// Imprimir tabla de símbolos (opcional, para debug)
	semanticAnalyzer.GetSymbolTable().PrintTable()

	ide.outputEntry.SetText(ide.outputEntry.Text + "✅ AST construido exitosamente\n")

	// === FASE 4: INTERPRETACIÓN ===
	ide.outputEntry.SetText(ide.outputEntry.Text + "\n🚀 Ejecutando programa...\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	interpreter := ast.NewInterpreter()
	output, err := interpreter.Interpret(astProgram)

	if err != nil {
		errorMsg := fmt.Sprintf("❌ Error durante la ejecución: %v\n", err)
		ide.outputEntry.SetText(ide.outputEntry.Text + errorMsg)
		fmt.Print(errorMsg) // También mostrar en terminal
		ide.outputEntry.SetText(ide.outputEntry.Text + "\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		return
	}

	ide.outputEntry.SetText(ide.outputEntry.Text + output)
	fmt.Print(output)
	ide.outputEntry.SetText(ide.outputEntry.Text + "\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "✅ Ejecución completada\n")
}

func (ide *IDE) showErrors() {
	var output strings.Builder
	output.WriteString("❌ Se encontraron errores:\n\n")

	for i, err := range ide.errorTable.Errors {
		output.WriteString(fmt.Sprintf("[%d] %s\n", i+1, err.String()))
	}

	ide.outputEntry.SetText(output.String())
}

func (ide *IDE) showErrorsReport() {
	if ide.errorTable == nil || !ide.errorTable.HasErrors() {
		dialog.ShowInformation("Reporte de Errores",
			"No hay errores para mostrar", ide.window)
		return
	}

	ide.showErrors()
}

func (ide *IDE) showASTReport() {
	if ide.astRoot == nil {
		dialog.ShowInformation("Reporte AST",
			"No hay AST disponible. Ejecuta el código primero.", ide.window)
		return
	}

	if ide.astRoot == nil {
		fmt.Println("AST es nil")
		return
	}

	// Imprimir tipo del nodo raíz
	fmt.Printf("AST Root Type: %T\n", ide.astRoot)

	// Si es un Program, imprimir el número de statements
	if program, ok := ide.astRoot.(*ast.Program); ok {
		fmt.Printf("Program tiene %d statements\n", len(program.Statements))
		for i, stmt := range program.Statements {
			fmt.Printf("  Statement %d: %T\n", i, stmt)
		}
	}

	// Generar el reporte del AST
	// html := reports.GenerateASTReport(ide.astRoot)

	// // Guardar el archivo HTML
	// filename := "ast_report.html"
	// reports.SaveAndOpenReport(html, filename)

	// // Opcionalmente, generar imagen con Graphviz
	// dot := reports.GenerateASTDot(ide.astRoot)
	// reports.GenerateASTImage(dot, "ast_graph.png")

	// dialog.ShowInformation("Reporte AST",
	// 	"Reporte generado exitosamente en "+filename, ide.window)
}

func (ide *IDE) showSymbolTableReport() {
	if ide.symbolTable == nil {
		dialog.ShowInformation("Tabla de Símbolos",
			"No hay tabla de símbolos disponible. Ejecuta el código primero.", ide.window)
		return
	}

	// Por ahora, mostrar un mensaje simple
	dialog.ShowInformation("Tabla de Símbolos",
		"Generación de tabla de símbolos en desarrollo...", ide.window)

	// TODO: Cuando tengas el analizador semántico listo:
	// html := reports.GenerateSymbolTableReport(ide.symbolTable)
	// reports.SaveAndOpenReport(html, "symbol_table_report.html")
}

func (ide *IDE) debugPrintAST() {
	if ide.astRoot == nil {
		fmt.Println("AST es nil")
		return
	}

	// Imprimir tipo del nodo raíz
	fmt.Printf("AST Root Type: %T\n", ide.astRoot)

	// Si es un Program, imprimir el número de statements
	if program, ok := ide.astRoot.(*ast.Program); ok {
		fmt.Printf("Program tiene %d statements\n", len(program.Statements))
		for i, stmt := range program.Statements {
			fmt.Printf("  Statement %d: %T\n", i, stmt)
		}
	}
}
