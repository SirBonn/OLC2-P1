package main

import (
	"compiler/errors"
	parser "compiler/parser"
	"compiler/repl"
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

	// tree es el CST (Concrete Syntax Tree) de ANTLR
	tree := p.Programa()

	// Combinar errores
	ide.errorTable = lexicalErrs.ErrorTable

	// Si hay errores léxicos o sintácticos, mostrar y detener
	if ide.errorTable.HasErrors() {
		ide.showErrors()
		return
	}

	ide.outputEntry.SetText("✅ Análisis léxico y sintáctico completado\n")

	// === FASE 3: CONSTRUCCIÓN DEL AST ===
	ide.outputEntry.SetText(ide.outputEntry.Text + "🔨 Construyendo AST...\n")

	// Crear el builder del AST
	astBuilder := NewASTBuilder()

	// Construir el AST visitando el CST
	astResult := astBuilder.Visit(tree)
	if astResult != nil {
		ide.astRoot = astResult.(ast.Node)
		ide.outputEntry.SetText(ide.outputEntry.Text + "✅ AST construido exitosamente\n")
	} else {
		ide.outputEntry.SetText(ide.outputEntry.Text + "❌ Error al construir el AST\n")
		return
	}

	// === FASE 4: ANÁLISIS SEMÁNTICO (si ya lo tienes implementado) ===
	/*
	   ide.outputEntry.SetText(ide.outputEntry.Text + "🔍 Realizando análisis semántico...\n")

	   semanticAnalyzer := semantic.NewSemanticAnalyzer()
	   semanticErrors := semanticAnalyzer.Analyze(ide.astRoot.(*ast.Program))

	   // Agregar errores semánticos a la tabla
	   for _, err := range semanticErrors {
	       ide.errorTable.AddError(err)
	   }

	   // Guardar la tabla de símbolos
	   ide.symbolTable = semanticAnalyzer.GetSymbolTable()

	   // Si hay errores semánticos, mostrar y detener
	   if ide.errorTable.HasErrors() {
	       ide.showErrors()
	       return
	   }

	   ide.outputEntry.SetText(ide.outputEntry.Text + "✅ Análisis semántico completado\n")
	*/

	// === FASE 5: INTERPRETACIÓN ===
	ide.outputEntry.SetText(ide.outputEntry.Text + "\n🚀 Ejecutando programa...\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	// Por ahora, usar tu visitor existente con el CST
	// Cuando tengas el intérprete del AST listo, cambiar a:
	// interpreter := repl.NewInterpreter()
	// output := interpreter.Execute(ide.astRoot)

	// Temporalmente, seguir usando el visitor del CST
	visitor := repl.NewReplVisitor(repl.NewErrorTable())
	// result := visitor.Visit(tree)

	// Mostrar la salida
	if visitor.Console != nil {
		output := visitor.Console.GetOutput()
		ide.outputEntry.SetText(ide.outputEntry.Text + output)
	}

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
