package reports

import (
	// "compiler/errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

// Generar reporte HTML simple del AST
func GenerateASTReport(astRoot interface{}) string {
	var html strings.Builder

	html.WriteString(`<!DOCTYPE html>
<html>
<head>
	<title>Reporte AST - V-Lang Cherry</title>
	<style>
		body { 
			font-family: Arial, sans-serif; 
			margin: 20px;
			background-color: #f5f5f5;
		}
		h1 { 
			color: #2196F3;
			text-align: center;
		}
		.ast-container {
			background: white;
			border-radius: 8px;
			padding: 20px;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}
		.node {
			margin: 10px;
			padding: 10px;
			background: #e3f2fd;
			border-radius: 4px;
			border-left: 4px solid #2196F3;
		}
		.node-type {
			font-weight: bold;
			color: #1976d2;
		}
		pre {
			background: #f5f5f5;
			padding: 10px;
			border-radius: 4px;
			overflow-x: auto;
		}
	</style>
</head>
<body>
	<h1>🌳 Árbol de Sintaxis Abstracta (AST)</h1>
	<div class="ast-container">
		<div class="node">
			<span class="node-type">Program</span>
			<p>Nodo raíz del programa</p>
		</div>
		<p style="text-align: center; color: #666;">
			<em>La visualización completa del AST estará disponible cuando se complete el análisis semántico.</em>
		</p>
	</div>
</body>
</html>`)

	return html.String()
}

// Generar reporte de tabla de símbolos
func GenerateSymbolTableReport(symbolTable interface{}) string {
	var html strings.Builder

	html.WriteString(`<!DOCTYPE html>
<html>
<head>
	<title>Tabla de Símbolos - V-Lang Cherry</title>
	<style>
		body { 
			font-family: Arial, sans-serif; 
			margin: 20px;
			background-color: #f5f5f5;
		}
		h1 { 
			color: #4CAF50;
			text-align: center;
		}
		.table-container {
			background: white;
			border-radius: 8px;
			padding: 20px;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}
		table { 
			border-collapse: collapse; 
			width: 100%; 
			margin-top: 20px; 
		}
		th { 
			background-color: #4CAF50; 
			color: white; 
			padding: 12px; 
			text-align: left; 
		}
		td { 
			border: 1px solid #ddd; 
			padding: 8px; 
		}
		tr:nth-child(even) { 
			background-color: #f2f2f2; 
		}
		.empty-state {
			text-align: center;
			color: #666;
			padding: 40px;
		}
	</style>
</head>
<body>
	<h1>📊 Tabla de Símbolos</h1>
	<div class="table-container">
		<div class="empty-state">
			<p><em>La tabla de símbolos se generará cuando se complete el análisis semántico.</em></p>
		</div>
	</div>
</body>
</html>`)

	return html.String()
}

// Guardar y abrir reporte en el navegador
func SaveAndOpenReport(html string, filename string) error {
	// Guardar archivo HTML
	err := ioutil.WriteFile(filename, []byte(html), 0644)
	if err != nil {
		return err
	}

	// Abrir en navegador según el sistema operativo
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	case "darwin":
		cmd = exec.Command("open", filename)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filename)
	default:
		return fmt.Errorf("plataforma no soportada")
	}

	if cmd != nil {
		return cmd.Start()
	}

	return nil
}

// Funciones placeholder para cuando tengas el AST completo
func GenerateASTDot(astRoot interface{}) string {
	// Por ahora retorna un DOT básico
	return `digraph AST {
		node [shape=box];
		Program [label="Program\n(Root)"];
		Placeholder [label="...\n(En desarrollo)"];
		Program -> Placeholder;
	}`
}

func GenerateASTImage(dot string, filename string) error {
	// Por ahora solo retorna nil
	// Cuando implementes, usarás Graphviz aquí
	return nil
}
