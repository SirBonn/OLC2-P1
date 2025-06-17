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

type LineNumberedEditor struct {
	widget.BaseWidget
	entry       *widget.Entry
	lineNumbers *widget.Label
	container   *fyne.Container
}

// NewLineNumberedEditor crea una nueva instancia del editor con números de línea
func NewLineNumberedEditor() *LineNumberedEditor {
	editor := &LineNumberedEditor{}
	editor.ExtendBaseWidget(editor)

	// Crear el editor de texto
	editor.entry = widget.NewMultiLineEntry()
	editor.entry.Wrapping = fyne.TextWrapOff // Desactivar wrap para mejor alineación
	editor.entry.SetPlaceHolder("// Escribe tu código V-Lang Cherry aquí...")

	// Crear el widget de números de línea usando Label simple
	editor.lineNumbers = widget.NewLabel("1")
	editor.lineNumbers.Alignment = fyne.TextAlignTrailing // Alinear a la derecha

	// Configurar el callback para actualizar los números cuando el texto cambie
	editor.entry.OnChanged = func(text string) {
		editor.updateLineNumbers(text)
	}

	// Configurar tamaño fijo para los números de línea
	editor.lineNumbers.Resize(fyne.NewSize(60, 0)) // Ancho fijo de 60 píxeles

	// Crear el contenedor usando Border para mejor control de tamaño
	editor.container = container.NewBorder(
		nil, nil, // top, bottom
		editor.lineNumbers, nil, // left, right
		editor.entry, // center (toma el espacio restante)
	)

	// Inicializar con la primera línea
	editor.updateLineNumbers("")

	return editor
}

// updateLineNumbers actualiza los números de línea basándose en el contenido
func (e *LineNumberedEditor) updateLineNumbers(text string) {
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		lines = []string{""}
	}

	// Construir los números de línea
	var lineNumbers []string
	for i := 1; i <= len(lines); i++ {
		lineNumbers = append(lineNumbers, fmt.Sprintf("%3d", i))
	}

	// Actualizar el label de números de línea
	numbersText := strings.Join(lineNumbers, "\n")
	e.lineNumbers.SetText(numbersText)
}

// CreateRenderer implementa la interfaz fyne.Widget
func (e *LineNumberedEditor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.container)
}

// Métodos para exponer la funcionalidad del Entry subyacente
func (e *LineNumberedEditor) SetText(text string) {
	e.entry.SetText(text)
}

func (e *LineNumberedEditor) Text() string {
	return e.entry.Text
}

func (e *LineNumberedEditor) SetPlaceHolder(placeholder string) {
	e.entry.SetPlaceHolder(placeholder)
}

func (e *LineNumberedEditor) Disable() {
	e.entry.Disable()
}

func (e *LineNumberedEditor) Enable() {
	e.entry.Enable()
}

// Método para establecer callback de cambio de texto
func (e *LineNumberedEditor) OnChanged(callback func(string)) {
	originalCallback := e.entry.OnChanged
	e.entry.OnChanged = func(text string) {
		originalCallback(text) // Mantener la actualización de números
		if callback != nil {
			callback(text) // Llamar al callback personalizado
		}
	}
}

type IDE struct {
	window      fyne.Window
	codeEditor  *LineNumberedEditor // Cambiar de codeEntry a codeEditor
	outputEntry *widget.Entry
	currentFile string
	app         fyne.App
	// Componentes para reportes
	errorTable  *errors.ErrorTable
	symbolTable interface{}
	astRoot     ast.Node
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
	// Usar el nuevo editor con números de línea
	ide.codeEditor = NewLineNumberedEditor()

	// Consola de salida (mantener como está)
	ide.outputEntry = widget.NewMultiLineEntry()
	ide.outputEntry.Wrapping = fyne.TextWrapWord
	ide.outputEntry.Disable()

	// Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("📝 Editor", container.NewVScroll(ide.codeEditor)),
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
	ide.codeEditor.SetText("")
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

		ide.codeEditor.SetText(string(data))
		ide.currentFile = reader.URI().Path()
	}, ide.window)
}

func (ide *IDE) saveFile() {
	if ide.currentFile == "" {
		ide.saveFileAs()
		return
	}

	err := ioutil.WriteFile(ide.currentFile, []byte(ide.codeEditor.Text()), 0644)
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

		_, err = writer.Write([]byte(ide.codeEditor.Text()))
		if err != nil {
			dialog.ShowError(err, ide.window)
			return
		}

		ide.currentFile = writer.URI().Path()
	}, ide.window)
}

func (ide *IDE) runCode() {
	code := ide.codeEditor.Text()
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
