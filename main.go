package main

import (
	"compiler/ast"
	"compiler/cst"
	"compiler/errors"
	parser "compiler/parser"
	"compiler/reports"
	"fmt"
	"io/ioutil"
	"strings"

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

func NewLineNumberedEditor() *LineNumberedEditor {
	editor := &LineNumberedEditor{}
	editor.ExtendBaseWidget(editor)

	editor.entry = widget.NewMultiLineEntry()
	editor.entry.Wrapping = fyne.TextWrapOff
	editor.entry.SetPlaceHolder("// Escribe tu código V-Lang Cherry aquí...")

	editor.lineNumbers = widget.NewLabel("1")
	editor.lineNumbers.Alignment = fyne.TextAlignTrailing

	editor.entry.OnChanged = func(text string) {
		editor.updateLineNumbers(text)
	}

	editor.lineNumbers.Resize(fyne.NewSize(60, 0))

	editor.container = container.NewBorder(
		nil, nil,
		editor.lineNumbers, nil,
		editor.entry,
	)

	editor.updateLineNumbers("")

	return editor
}

func (e *LineNumberedEditor) updateLineNumbers(text string) {
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		lines = []string{""}
	}

	var lineNumbers []string
	for i := 1; i <= len(lines); i++ {
		lineNumbers = append(lineNumbers, fmt.Sprintf("%3d", i))
	}

	numbersText := strings.Join(lineNumbers, "\n")
	e.lineNumbers.SetText(numbersText)
}

func (e *LineNumberedEditor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.container)
}

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

func (e *LineNumberedEditor) OnChanged(callback func(string)) {
	originalCallback := e.entry.OnChanged
	e.entry.OnChanged = func(text string) {
		originalCallback(text)
		if callback != nil {
			callback(text)
		}
	}
}

type IDE struct {
	window           fyne.Window
	codeEditor       *LineNumberedEditor
	outputEntry      *widget.Entry
	currentFile      string
	app              fyne.App
	errorTable       *errors.ErrorTable
	symbolTable      *ast.SymbolTable
	semanticAnalyzer *ast.SemanticAnalyzer
}

func main() {
	a := app.NewWithID("com.vlang.ide")
	w := a.NewWindow("V-Lang Cherry IDE")

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
	ide.codeEditor = NewLineNumberedEditor()

	ide.outputEntry = widget.NewMultiLineEntry()
	ide.outputEntry.Wrapping = fyne.TextWrapWord
	ide.outputEntry.Disable()

	tabs := container.NewAppTabs(
		container.NewTabItem("Editor", container.NewVScroll(ide.codeEditor)),
		container.NewTabItem("Consola", container.NewVScroll(ide.outputEntry)),
	)

	toolbar := ide.createToolbar()

	return container.NewBorder(toolbar, nil, nil, nil, tabs)
}

func (ide *IDE) createMenu() *fyne.MainMenu {
	fileMenu := fyne.NewMenu("Archivo",
		fyne.NewMenuItem("Nuevo", func() { ide.newFile() }),
		fyne.NewMenuItem("Abrir...", func() { ide.openFile() }),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Guardar", func() { ide.saveFile() }),
		fyne.NewMenuItem("Guardar como...", func() { ide.saveFileAs() }),
	)

	toolsMenu := fyne.NewMenu("Herramientas",
		fyne.NewMenuItem("Ejecutar", func() { ide.runCode() }),
	)

	reportsMenu := fyne.NewMenu("Reportes",
		fyne.NewMenuItem("Reporte de Errores", func() { ide.showErrorsReport() }),
		fyne.NewMenuItem("Tabla de Símbolos", func() { ide.showSymbolTableReport() }),
		fyne.NewMenuItem("Reporte CST", func() { ide.showCSTReport() }),
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
	ide.outputEntry.SetText("Compilando...\n\n")

	lexicalErrs := errors.NewLexicalErrorListener()
	lexer := parser.NewVlangLexer(antlr.NewInputStream(code))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

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

	ide.outputEntry.SetText("Análisis léxico y sintáctico completado\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "Construyendo AST...\n")

	astBuilder := NewASTBuilder()
	astProgram, err := astBuilder.Build(tree)

	if err != nil {
		ide.outputEntry.SetText(ide.outputEntry.Text + fmt.Sprintf("Error al construir el AST: %v\n", err))
		return
	}

	ide.semanticAnalyzer = ast.NewSemanticAnalyzer()
	err = ide.semanticAnalyzer.Analyze(astProgram)
	if err != nil {
		for _, e := range ide.semanticAnalyzer.GetSymbolTable().GetErrors() {
			fmt.Println("Error semántico:", e)
		}
		ide.outputEntry.SetText(ide.outputEntry.Text + fmt.Sprintf("Error: %v\n", err))
	}

	ide.symbolTable = ide.semanticAnalyzer.GetSymbolTable()

	ide.outputEntry.SetText(ide.outputEntry.Text + "AST construido exitosamente\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "\nEjecutando programa...\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "────────────────────────────\n\n")

	interpreter := ast.NewInterpreter()
	output, err := interpreter.Interpret(astProgram)

	if err != nil {
		errorMsg := fmt.Sprintf("Error durante la ejecución: %v\n", err)
		ide.outputEntry.SetText(ide.outputEntry.Text + errorMsg)
		fmt.Print(errorMsg)
		ide.outputEntry.SetText(ide.outputEntry.Text + "\n────────────────────────────\n")
		return
	}

	ide.outputEntry.SetText(ide.outputEntry.Text + output)
	fmt.Print(output)
	ide.outputEntry.SetText(ide.outputEntry.Text + "\n────────────────────────────\n")
	ide.outputEntry.SetText(ide.outputEntry.Text + "Ejecución completada\n")
}

func (ide *IDE) showErrors() {
	var output strings.Builder
	output.WriteString("Se encontraron errores:\n\n")

	for i, err := range ide.errorTable.Errors {
		output.WriteString(fmt.Sprintf("[%d] %s\n", i+1, err.String()))
	}

	ide.outputEntry.SetText(output.String())
}

func (ide *IDE) showErrorsReport() {
	if ide.errorTable == nil {
		ide.errorTable = errors.NewErrorTable()
	}

	html := reports.GenerateErrorReport(ide.errorTable)
	err := reports.SaveAndOpenReport(html, "error_report.html")

	if err != nil {
		dialog.ShowError(err, ide.window)
		return
	}

	dialog.ShowInformation("Reporte de Errores",
		"Reporte generado exitosamente en error_report.html", ide.window)
}

func (ide *IDE) showSymbolTableReport() {
	if ide.symbolTable == nil {
		dialog.ShowInformation("Tabla de Símbolos",
			"No hay tabla de símbolos disponible. Ejecuta el código primero.", ide.window)
		return
	}

	html := reports.GenerateSymbolTableReport(ide.symbolTable)
	err := reports.SaveAndOpenReport(html, "symbol_table_report.html")

	if err != nil {
		dialog.ShowError(err, ide.window)
		return
	}

	dialog.ShowInformation("Tabla de Símbolos",
		"Reporte generado exitosamente en symbol_table_report.html", ide.window)
}

func (ide *IDE) showCSTReport() {
	code := ide.codeEditor.Text()
	if code == "" {
		dialog.ShowInformation("Reporte CST",
			"No hay código para analizar.", ide.window)
		return
	}

	svgContent := cst.CstReport(code)
	if svgContent == "" {
		dialog.ShowError(fmt.Errorf("No se pudo generar el reporte CST"), ide.window)
		return
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Reporte CST - V-Lang Cherry</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 100%%;
            margin: 0 auto;
            background-color: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #388e3c;
            text-align: center;
            margin-bottom: 30px;
        }
        .tree-container {
            width: 100%%;
            overflow: auto;
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
            background-color: #fafafa;
        }
        svg {
            max-width: 100%%;
            height: auto;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Árbol de Sintaxis Concreta (CST)</h1>
        <div class="tree-container">
            %s
        </div>
    </div>
</body>
</html>`, svgContent)

	err := reports.SaveAndOpenReport(html, "cst_report.html")
	if err != nil {
		dialog.ShowError(err, ide.window)
		return
	}

	dialog.ShowInformation("Reporte CST",
		"Reporte generado exitosamente en cst_report.html", ide.window)
}
