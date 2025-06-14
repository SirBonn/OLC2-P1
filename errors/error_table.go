package errors

import "fmt"

type ErrorType string

const (
	LexicalError  ErrorType = "Léxico"
	SyntaxError   ErrorType = "Sintáctico"
	SemanticError ErrorType = "Semántico"
	RuntimeError  ErrorType = "Runtime"
)

type CompilerError struct {
	Type        ErrorType
	Message     string
	Line        int
	Column      int
	Description string
}

func (e *CompilerError) String() string {
	return fmt.Sprintf("%s error en línea %d:%d - %s",
		e.Type, e.Line, e.Column, e.Message)
}

type ErrorTable struct {
	Errors []CompilerError
}

func NewErrorTable() *ErrorTable {
	return &ErrorTable{
		Errors: make([]CompilerError, 0),
	}
}

func (et *ErrorTable) AddError(err CompilerError) {
	et.Errors = append(et.Errors, err)
}

func (et *ErrorTable) HasErrors() bool {
	return len(et.Errors) > 0
}
