package reports

import (
	"compiler/ast"
	"compiler/errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func GenerateErrorReport(errorTable *errors.ErrorTable) string {
	var html strings.Builder

	html.WriteString(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Reporte de Errores - V-Lang Cherry</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background-color: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #d32f2f;
            text-align: center;
            margin-bottom: 30px;
        }
        .summary {
            background-color: #ffebee;
            padding: 15px;
            border-radius: 5px;
            margin-bottom: 30px;
            border-left: 4px solid #d32f2f;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 20px;
        }
        th {
            background-color: #d32f2f;
            color: white;
            padding: 12px;
            text-align: left;
            font-weight: 600;
        }
        td {
            padding: 12px;
            border-bottom: 1px solid #ddd;
        }
        tr:hover {
            background-color: #f5f5f5;
        }
        .error-type {
            display: inline-block;
            padding: 4px 8px;
            border-radius: 3px;
            font-size: 12px;
            font-weight: bold;
        }
        .lexical {
            background-color: #ff5722;
            color: white;
        }
        .syntax {
            background-color: #f44336;
            color: white;
        }
        .semantic {
            background-color: #e91e63;
            color: white;
        }
        .timestamp {
            text-align: right;
            color: #666;
            font-size: 14px;
            margin-top: 30px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Reporte de Errores</h1>
        <div class="summary">
            <h3>Resumen</h3>
            <p><strong>Total de errores:</strong> %d</p>
        </div>
`)

	if len(errorTable.Errors) > 0 {
		html.WriteString(`
        <table>
            <thead>
                <tr>
                    <th>#</th>
                    <th>Tipo</th>
                    <th>Línea</th>
                    <th>Columna</th>
                    <th>Mensaje</th>
                    <th>Descripción</th>
                </tr>
            </thead>
            <tbody>
`)

		for i, err := range errorTable.Errors {
			errorType := "semantic"
			if err.Type == errors.LexicalError {
				errorType = "lexical"
			} else if err.Type == errors.SyntaxError {
				errorType = "syntax"
			}

			html.WriteString(fmt.Sprintf(`
                <tr>
                    <td>%d</td>
                    <td><span class="error-type %s">%s</span></td>
                    <td>%d</td>
                    <td>%d</td>
                    <td>%s</td>
                    <td>%s</td>
                </tr>
`, i+1, errorType, strings.ToUpper(errorType), err.Line, err.Column, err.Message, err.Description))
		}

		html.WriteString(`
            </tbody>
        </table>
`)
	} else {
		html.WriteString(`
        <p style="text-align: center; color: #4caf50; font-size: 18px;">
            No se encontraron errores en el código.
        </p>
`)
	}

	html.WriteString(fmt.Sprintf(`
        <div class="timestamp">
            Generado: %s
        </div>
    </div>
</body>
</html>
`, time.Now().Format("02/01/2006 15:04:05")))

	return html.String()
}

func GenerateSymbolTableReport(symbolTable *ast.SymbolTable) string {
	var html strings.Builder

	html.WriteString(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Tabla de Símbolos - V-Lang Cherry</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background-color: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #1976d2;
            text-align: center;
            margin-bottom: 30px;
        }
        .scope {
            margin-bottom: 30px;
            border: 1px solid #ddd;
            border-radius: 5px;
            overflow: hidden;
        }
        .scope-header {
            background-color: #e3f2fd;
            padding: 10px 15px;
            font-weight: bold;
            color: #1976d2;
        }
        table {
            width: 100%;
            border-collapse: collapse;
        }
        th {
            background-color: #1976d2;
            color: white;
            padding: 12px;
            text-align: left;
            font-weight: 600;
        }
        td {
            padding: 12px;
            border-bottom: 1px solid #eee;
        }
        tr:hover {
            background-color: #f5f5f5;
        }
        .symbol-type {
            display: inline-block;
            padding: 4px 8px;
            border-radius: 3px;
            font-size: 12px;
            font-weight: bold;
        }
        .variable {
            background-color: #4caf50;
            color: white;
        }
        .function {
            background-color: #2196f3;
            color: white;
        }
        .struct {
            background-color: #ff9800;
            color: white;
        }
        .parameter {
            background-color: #9c27b0;
            color: white;
        }
        .mutable {
            color: #f44336;
            font-weight: bold;
        }
        .timestamp {
            text-align: right;
            color: #666;
            font-size: 14px;
            margin-top: 30px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Tabla de Símbolos</h1>
`)

	generateScopeHTML(&html, symbolTable.CurrentScope, symbolTable)

	html.WriteString(fmt.Sprintf(`
        <div class="timestamp">
            Generado: %s
        </div>
    </div>
</body>
</html>
`, time.Now().Format("02/01/2006 15:04:05")))

	return html.String()
}

func generateScopeHTML(html *strings.Builder, scope *ast.Scope, symbolTable *ast.SymbolTable) {
	if scope == nil {
		return
	}

	if scope.Parent != nil {
		generateScopeHTML(html, scope.Parent, symbolTable)
	}

	html.WriteString(fmt.Sprintf(`
        <div class="scope">
            <div class="scope-header">Alcance Nivel %d</div>
            <table>
                <thead>
                    <tr>
                        <th>Nombre</th>
                        <th>Tipo</th>
                        <th>Categoría</th>
                        <th>Línea</th>
                        <th>Columna</th>
                        <th>Detalles</th>
                    </tr>
                </thead>
                <tbody>
`, scope.Level))

	for name, symbol := range scope.Symbols {
		symbolTypeClass := "variable"
		symbolTypeStr := "Variable"
		details := ""

		switch symbol.SymbolType {
		case ast.SymbolVariable:
			symbolTypeClass = "variable"
			symbolTypeStr = "Variable"
			if symbol.IsMutable {
				details = `<span class="mutable">mut</span>`
			}
		case ast.SymbolFunction:
			symbolTypeClass = "function"
			symbolTypeStr = "Función"
			params := []string{}
			for _, p := range symbol.Parameters {
				params = append(params, fmt.Sprintf("%s: %s", p.Name, p.Type))
			}
			details = fmt.Sprintf("(%s) → %s", strings.Join(params, ", "), symbol.ReturnType)
		case ast.SymbolStruct:
			symbolTypeClass = "struct"
			symbolTypeStr = "Struct"
			details = fmt.Sprintf("%d campos", len(symbol.Fields))
		case ast.SymbolParameter:
			symbolTypeClass = "parameter"
			symbolTypeStr = "Parámetro"
		}

		html.WriteString(fmt.Sprintf(`
                    <tr>
                        <td><strong>%s</strong></td>
                        <td>%s</td>
                        <td><span class="symbol-type %s">%s</span></td>
                        <td>%d</td>
                        <td>%d</td>
                        <td>%s</td>
                    </tr>
`, name, symbol.Type, symbolTypeClass, symbolTypeStr, symbol.Line, symbol.Column, details))
	}

	html.WriteString(`
                </tbody>
            </table>
        </div>
`)
}

func SaveAndOpenReport(htmlContent string, filename string) error {
	err := os.WriteFile(filename, []byte(htmlContent), 0644)
	if err != nil {
		return fmt.Errorf("error al guardar el reporte: %v", err)
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("error al obtener ruta absoluta: %v", err)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", absPath)
	case "darwin":
		cmd = exec.Command("open", absPath)
	case "linux":
		cmd = exec.Command("xdg-open", absPath)
	default:
		return fmt.Errorf("sistema operativo no soportado")
	}

	return cmd.Start()
}
