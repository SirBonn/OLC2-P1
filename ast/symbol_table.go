package ast

import (
	"fmt"
)

type SymbolType int

const (
	SymbolVariable SymbolType = iota
	SymbolFunction
	SymbolStruct
	SymbolParameter
)

type Symbol struct {
	Name       string
	Type       string
	SymbolType SymbolType
	Scope      int
	Line       int
	Column     int
	IsMutable  bool
	Parameters []Parameter
	ReturnType string
	Fields     []Field
}

type Scope struct {
	Parent  *Scope
	Symbols map[string]*Symbol
	Level   int
}

type SymbolTable struct {
	CurrentScope *Scope
	Errors       []string
}

func NewSymbolTable() *SymbolTable {
	globalScope := &Scope{
		Parent:  nil,
		Symbols: make(map[string]*Symbol),
		Level:   0,
	}

	return &SymbolTable{
		CurrentScope: globalScope,
		Errors:       make([]string, 0),
	}
}

func (st *SymbolTable) EnterScope() {
	newScope := &Scope{
		Parent:  st.CurrentScope,
		Symbols: make(map[string]*Symbol),
		Level:   st.CurrentScope.Level + 1,
	}
	st.CurrentScope = newScope
}

func (st *SymbolTable) ExitScope() {
	if st.CurrentScope.Parent != nil {
		st.CurrentScope = st.CurrentScope.Parent
	}
}

func (st *SymbolTable) Define(symbol *Symbol) error {
	symbol.Scope = st.CurrentScope.Level

	if _, exists := st.CurrentScope.Symbols[symbol.Name]; exists {
		err := fmt.Errorf("symbol '%s' already defined in this scope at line %d",
			symbol.Name, symbol.Line)
		st.Errors = append(st.Errors, err.Error())
		return err
	}

	st.CurrentScope.Symbols[symbol.Name] = symbol
	return nil
}

func (st *SymbolTable) Lookup(name string) (*Symbol, bool) {
	scope := st.CurrentScope

	for scope != nil {
		if symbol, exists := scope.Symbols[name]; exists {
			return symbol, true
		}
		scope = scope.Parent
	}

	return nil, false
}

func (st *SymbolTable) LookupInCurrentScope(name string) (*Symbol, bool) {
	symbol, exists := st.CurrentScope.Symbols[name]
	return symbol, exists
}

func (st *SymbolTable) GetErrors() []string {
	return st.Errors
}

func (st *SymbolTable) AddError(err string) {
	st.Errors = append(st.Errors, err)
}

func (st *SymbolTable) PrintTable() {
	fmt.Println("=== Symbol Table ===")
	st.printScope(st.CurrentScope, "")
}

func (st *SymbolTable) printScope(scope *Scope, indent string) {
	if scope == nil {
		return
	}

	if scope.Parent != nil {
		st.printScope(scope.Parent, indent)
	}

	fmt.Printf("%sScope Level %d:\n", indent, scope.Level)
	for name, symbol := range scope.Symbols {
		fmt.Printf("%s  %s: %s\n", indent, name, st.symbolToString(symbol))
	}
}

func (st *SymbolTable) symbolToString(symbol *Symbol) string {
	switch symbol.SymbolType {
	case SymbolVariable:
		mut := ""
		if symbol.IsMutable {
			mut = "mut "
		}
		return fmt.Sprintf("%s%s %s (line %d)", mut, symbol.Type, "variable", symbol.Line)
	case SymbolFunction:
		params := ""
		for i, p := range symbol.Parameters {
			if i > 0 {
				params += ", "
			}
			params += p.Name + " " + p.Type
		}
		return fmt.Sprintf("function(%s) %s (line %d)", params, symbol.ReturnType, symbol.Line)
	case SymbolStruct:
		return fmt.Sprintf("struct with %d fields (line %d)", len(symbol.Fields), symbol.Line)
	case SymbolParameter:
		return fmt.Sprintf("%s parameter (line %d)", symbol.Type, symbol.Line)
	default:
		return "unknown"
	}
}
