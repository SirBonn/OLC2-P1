// base.go
package codegen

import (
	"compiler/ast"
	"fmt"
	"strings"
)

// CodeGenerator define la interfaz para generar código
type CodeGenerator interface {
	Generate(node ast.Node) (string, error)
}

// BaseGenerator proporciona funcionalidad común para todos los generadores
type BaseGenerator struct {
	output      []string
	indentLevel int
	errors      []error
}

// NewBaseGenerator crea un nuevo generador base
func NewBaseGenerator() *BaseGenerator {
	return &BaseGenerator{
		output:      make([]string, 0),
		indentLevel: 0,
		errors:      make([]error, 0),
	}
}

// Emit agrega una línea al output
func (g *BaseGenerator) Emit(format string, args ...interface{}) {
	g.output = append(g.output, fmt.Sprintf(format, args...))
}

// EmitWithIndent agrega una línea con indentación
func (g *BaseGenerator) EmitWithIndent(format string, args ...interface{}) {
	indent := strings.Repeat("\t", g.indentLevel)
	g.output = append(g.output, indent+fmt.Sprintf(format, args...))
}

// GetOutput retorna el código generado como string
func (g *BaseGenerator) GetOutput() string {
	return strings.Join(g.output, "\n")
}

// AddError agrega un error a la lista
func (g *BaseGenerator) AddError(err error) {
	g.errors = append(g.errors, err)
}

// HasErrors retorna true si hay errores
func (g *BaseGenerator) HasErrors() bool {
	return len(g.errors) > 0
}

// GetErrors retorna todos los errores acumulados
func (g *BaseGenerator) GetErrors() error {
	if !g.HasErrors() {
		return nil
	}

	errMsgs := make([]string, len(g.errors))
	for i, err := range g.errors {
		errMsgs[i] = err.Error()
	}

	return fmt.Errorf("code generation errors:\n%s", strings.Join(errMsgs, "\n"))
}

// PushIndent incrementa el nivel de indentación
func (g *BaseGenerator) PushIndent() {
	g.indentLevel++
}

// PopIndent decrementa el nivel de indentación
func (g *BaseGenerator) PopIndent() {
	if g.indentLevel > 0 {
		g.indentLevel--
	}
}

// Reset limpia el estado del generador
func (g *BaseGenerator) Reset() {
	g.output = make([]string, 0)
	g.errors = make([]error, 0)
	g.indentLevel = 0
}

// TargetPlatform representa la plataforma objetivo
type TargetPlatform string

const (
	ARM64  TargetPlatform = "arm64"
	X86_64 TargetPlatform = "x86_64"
	WASM   TargetPlatform = "wasm"
)

// CompilerOptions contiene las opciones de compilación
type CompilerOptions struct {
	Platform     TargetPlatform
	Optimization int  // 0 = no optimization, 1 = basic, 2 = aggressive
	Debug        bool // incluir información de debug
	OutputFile   string
}

// DefaultOptions retorna las opciones por defecto
func DefaultOptions() *CompilerOptions {
	return &CompilerOptions{
		Platform:     ARM64,
		Optimization: 0,
		Debug:        false,
		OutputFile:   "output.s",
	}
}
