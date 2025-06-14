package errors

import (
	"github.com/antlr4-go/antlr/v4"
)

// LexicalErrorListener para errores léxicos
type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *ErrorTable
}

func NewLexicalErrorListener() *LexicalErrorListener {
	return &LexicalErrorListener{
		ErrorTable: NewErrorTable(),
	}
}

func (l *LexicalErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := CompilerError{
		Type:    LexicalError,
		Message: msg,
		Line:    line,
		Column:  column,
	}
	l.ErrorTable.AddError(err)
}

// SyntaxErrorListener para errores sintácticos
type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorTable *ErrorTable
}

func NewSyntaxErrorListener(lexicalErrors *ErrorTable) *SyntaxErrorListener {
	return &SyntaxErrorListener{
		ErrorTable: lexicalErrors, // Reutilizar la misma tabla
	}
}

func (s *SyntaxErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := CompilerError{
		Type:    SyntaxError,
		Message: msg,
		Line:    line,
		Column:  column,
	}
	s.ErrorTable.AddError(err)
}
