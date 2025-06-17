package ast

import (
	"fmt"
)

type SemanticAnalyzer struct {
	symbolTable *SymbolTable
}

var builtinFunctions = map[string]struct {
	Parameters int
	ReturnType string
}{
	"int":         {Parameters: 1, ReturnType: "int"},
	"float":       {Parameters: 1, ReturnType: "f64"},
	"string":      {Parameters: 1, ReturnType: "string"},
	"atoi":        {Parameters: 1, ReturnType: "int"},
	"parse_float": {Parameters: 1, ReturnType: "f64"},
	"TypeOf":      {Parameters: 1, ReturnType: "string"},
	"len":         {Parameters: 1, ReturnType: "int"},
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{
		symbolTable: NewSymbolTable(),
	}
}

func (sa *SemanticAnalyzer) Analyze(program *Program) error {
	for _, stmt := range program.Statements {
		switch s := stmt.(type) {
		case *FuncDecl:
			sa.registerFunction(s)
		case *StructDecl:
			sa.registerStruct(s)
		}
	}

	for _, stmt := range program.Statements {
		sa.analyzeStatement(stmt)
	}

	if len(sa.symbolTable.GetErrors()) > 0 {
		return fmt.Errorf("semantic analysis failed with %d errors", len(sa.symbolTable.GetErrors()))
	}

	return nil
}

func (sa *SemanticAnalyzer) registerFunction(funcDecl *FuncDecl) {
	symbol := &Symbol{
		Name:       funcDecl.Name,
		Type:       "function",
		SymbolType: SymbolFunction,
		Line:       funcDecl.Line,
		Column:     funcDecl.Column,
		Parameters: funcDecl.Parameters,
		ReturnType: funcDecl.ReturnType,
	}

	sa.symbolTable.Define(symbol)
}

func (sa *SemanticAnalyzer) registerStruct(structDecl *StructDecl) {
	symbol := &Symbol{
		Name:       structDecl.Name,
		Type:       "struct",
		SymbolType: SymbolStruct,
		Line:       structDecl.Line,
		Column:     structDecl.Column,
		Fields:     structDecl.Fields,
	}

	sa.symbolTable.Define(symbol)
}

func (sa *SemanticAnalyzer) analyzeStatement(stmt Statement) {
	switch s := stmt.(type) {
	case *VarDecl:
		sa.analyzeVarDecl(s)
	case *Assignment:
		sa.analyzeAssignment(s)
	case *PlusAssign, *MinusAssign, *MulAssign, *DivAssign, *ModAssign:
		sa.analyzeCompoundAssignment(stmt)
	case *PrintStmt:
		sa.analyzePrintStmt(s)
	case *IfStmt:
		sa.analyzeIfStmt(s)
	case *WhileStmt:
		sa.analyzeWhileStmt(s)
	case *ForStmt:
		sa.analyzeForStmt(s)
	case *FuncDecl:
		sa.analyzeFuncDecl(s)
	case *Return:
		sa.analyzeReturn(s)
	case *ExpressionStatement:
		sa.analyzeExpression(s.Expression)
	}
}

func (sa *SemanticAnalyzer) analyzeVarDecl(varDecl *VarDecl) {
	valueType := sa.analyzeExpression(varDecl.Value)

	varType := varDecl.Type
	if varType == "" {
		varType = valueType
	} else {
		if !sa.areTypesCompatible(varType, valueType) {
			sa.symbolTable.AddError(fmt.Sprintf("type mismatch: cannot assign %s to %s at line %d",
				valueType, varType, varDecl.Line))
		}
	}

	symbol := &Symbol{
		Name:       varDecl.Name,
		Type:       varType,
		SymbolType: SymbolVariable,
		Line:       varDecl.Line,
		Column:     varDecl.Column,
		IsMutable:  varDecl.IsMutable,
	}

	sa.symbolTable.Define(symbol)
}

func (sa *SemanticAnalyzer) analyzeAssignment(assignment *Assignment) {
	if id, ok := assignment.Target.(*Identifier); ok {
		symbol, exists := sa.symbolTable.Lookup(id.Name)
		if !exists {
			sa.symbolTable.AddError(fmt.Sprintf("undefined variable '%s' at line %d",
				id.Name, assignment.Line))
			return
		}

		if symbol.SymbolType != SymbolVariable || !symbol.IsMutable {
			sa.symbolTable.AddError(fmt.Sprintf("cannot assign to immutable variable '%s' at line %d",
				id.Name, assignment.Line))
			return
		}

		valueType := sa.analyzeExpression(assignment.Value)
		if !sa.areTypesCompatible(symbol.Type, valueType) {
			sa.symbolTable.AddError(fmt.Sprintf("type mismatch: cannot assign %s to %s at line %d",
				valueType, symbol.Type, assignment.Line))
		}
	}
}

func (sa *SemanticAnalyzer) analyzeCompoundAssignment(stmt Statement) {
	var target Expression
	var value Expression
	var line int

	switch s := stmt.(type) {
	case *PlusAssign:
		target, value, line = s.Target, s.Value, s.Line
	case *MinusAssign:
		target, value, line = s.Target, s.Value, s.Line
	case *MulAssign:
		target, value, line = s.Target, s.Value, s.Line
	case *DivAssign:
		target, value, line = s.Target, s.Value, s.Line
	case *ModAssign:
		target, value, line = s.Target, s.Value, s.Line
	}

	if id, ok := target.(*Identifier); ok {
		symbol, exists := sa.symbolTable.Lookup(id.Name)
		if !exists {
			sa.symbolTable.AddError(fmt.Sprintf("undefined variable '%s' at line %d", id.Name, line))
			return
		}

		if !symbol.IsMutable {
			sa.symbolTable.AddError(fmt.Sprintf("cannot modify immutable variable '%s' at line %d", id.Name, line))
			return
		}

		valueType := sa.analyzeExpression(value)
		if !sa.isNumericType(symbol.Type) || !sa.isNumericType(valueType) {
			sa.symbolTable.AddError(fmt.Sprintf("compound assignment requires numeric types at line %d", line))
		}
	}
}

func (sa *SemanticAnalyzer) analyzePrintStmt(printStmt *PrintStmt) {
	for _, arg := range printStmt.Arguments {
		sa.analyzeExpression(arg)
	}
}

func (sa *SemanticAnalyzer) analyzeIfStmt(ifStmt *IfStmt) {
	condType := sa.analyzeExpression(ifStmt.Condition)
	if condType != "bool" && condType != "unknown" {
		sa.symbolTable.AddError(fmt.Sprintf("if condition must be boolean, got %s at line %d",
			condType, ifStmt.Line))
	}

	sa.symbolTable.EnterScope()
	for _, stmt := range ifStmt.ThenBranch {
		sa.analyzeStatement(stmt)
	}
	sa.symbolTable.ExitScope()

	if len(ifStmt.ElseBranch) > 0 {
		sa.symbolTable.EnterScope()
		for _, stmt := range ifStmt.ElseBranch {
			sa.analyzeStatement(stmt)
		}
		sa.symbolTable.ExitScope()
	}
}

func (sa *SemanticAnalyzer) analyzeWhileStmt(whileStmt *WhileStmt) {
	condType := sa.analyzeExpression(whileStmt.Condition)
	if condType != "bool" && condType != "unknown" {
		sa.symbolTable.AddError(fmt.Sprintf("while condition must be boolean, got %s at line %d",
			condType, whileStmt.Line))
	}

	sa.symbolTable.EnterScope()
	for _, stmt := range whileStmt.Body {
		sa.analyzeStatement(stmt)
	}
	sa.symbolTable.ExitScope()
}

func (sa *SemanticAnalyzer) analyzeForStmt(forStmt *ForStmt) {
	sa.symbolTable.EnterScope()

	iterableType := sa.analyzeExpression(forStmt.Iterable)
	elementType := sa.getElementType(iterableType)

	symbol := &Symbol{
		Name:       forStmt.Variable,
		Type:       elementType,
		SymbolType: SymbolVariable,
		Line:       forStmt.Line,
		Column:     forStmt.Column,
		IsMutable:  false,
	}
	sa.symbolTable.Define(symbol)

	for _, stmt := range forStmt.Body {
		sa.analyzeStatement(stmt)
	}

	sa.symbolTable.ExitScope()
}

func (sa *SemanticAnalyzer) analyzeFuncDecl(funcDecl *FuncDecl) {
	sa.symbolTable.EnterScope()

	for _, param := range funcDecl.Parameters {
		symbol := &Symbol{
			Name:       param.Name,
			Type:       param.Type,
			SymbolType: SymbolParameter,
			Line:       funcDecl.Line,
			Column:     funcDecl.Column,
			IsMutable:  false,
		}
		sa.symbolTable.Define(symbol)
	}

	for _, stmt := range funcDecl.Body {
		sa.analyzeStatement(stmt)
	}

	sa.symbolTable.ExitScope()
}

func (sa *SemanticAnalyzer) analyzeReturn(returnStmt *Return) {
	if returnStmt.Value != nil {
		sa.analyzeExpression(returnStmt.Value)
	}
}

func (sa *SemanticAnalyzer) analyzeExpression(expr Expression) string {
	if expr == nil {
		return "unknown"
	}

	switch e := expr.(type) {
	case *Literal:
		return e.Type
	case *Identifier:
		symbol, exists := sa.symbolTable.Lookup(e.Name)
		if !exists {
			sa.symbolTable.AddError(fmt.Sprintf("undefined identifier '%s' at line %d",
				e.Name, e.Line))
			return "unknown"
		}
		return symbol.Type
	case *BinaryExpr:
		return sa.analyzeBinaryExpr(e)
	case *UnaryExpr:
		return sa.analyzeUnaryExpr(e)
	case *FuncCall:
		return sa.analyzeFuncCall(e)
	case *ArrayLiteral:
		return sa.analyzeArrayLiteral(e)
	default:
		return "unknown"
	}
}

func (sa *SemanticAnalyzer) analyzeBinaryExpr(expr *BinaryExpr) string {
	leftType := sa.analyzeExpression(expr.Left)
	rightType := sa.analyzeExpression(expr.Right)

	switch expr.Operator {
	case "+", "-", "*", "/", "%":
		if !sa.isNumericType(leftType) || !sa.isNumericType(rightType) {
			sa.symbolTable.AddError(fmt.Sprintf("operator %s requires numeric operands at line %d",
				expr.Operator, expr.Line))
			return "unknown"
		}
		if leftType == "float64" || rightType == "float64" {
			return "float64"
		}
		return "int"
	case "<", "<=", ">", ">=":
		if !sa.isNumericType(leftType) || !sa.isNumericType(rightType) {
			sa.symbolTable.AddError(fmt.Sprintf("comparison operator requires numeric operands at line %d",
				expr.Line))
		}
		return "bool"
	case "==", "!=":
		if !sa.areTypesCompatible(leftType, rightType) {
			sa.symbolTable.AddError(fmt.Sprintf("cannot compare %s with %s at line %d",
				leftType, rightType, expr.Line))
		}
		return "bool"
	case "&&", "||":
		if leftType != "bool" || rightType != "bool" {
			sa.symbolTable.AddError(fmt.Sprintf("logical operator requires boolean operands at line %d",
				expr.Line))
		}
		return "bool"
	default:
		return "unknown"
	}
}

func (sa *SemanticAnalyzer) analyzeUnaryExpr(expr *UnaryExpr) string {
	operandType := sa.analyzeExpression(expr.Operand)

	switch expr.Operator {
	case "-":
		if !sa.isNumericType(operandType) {
			sa.symbolTable.AddError(fmt.Sprintf("unary minus requires numeric operand at line %d",
				expr.Line))
			return "unknown"
		}
		return operandType
	case "!":
		if operandType != "bool" {
			sa.symbolTable.AddError(fmt.Sprintf("logical not requires boolean operand at line %d",
				expr.Line))
			return "unknown"
		}
		return "bool"
	default:
		return "unknown"
	}
}

func (sa *SemanticAnalyzer) analyzeFuncCall(funcCall *FuncCall) string {

	if builtin, isBuiltin := builtinFunctions[funcCall.Name]; isBuiltin {
		// Verificar número de argumentos
		if len(funcCall.Arguments) != builtin.Parameters {
			sa.symbolTable.AddError(fmt.Sprintf("builtin function '%s' expects %d arguments, got %d at line %d",
				funcCall.Name, builtin.Parameters, len(funcCall.Arguments), funcCall.Line))
		}

		// Analizar argumentos
		for _, arg := range funcCall.Arguments {
			sa.analyzeExpression(arg)
		}

		return builtin.ReturnType
	}

	symbol, exists := sa.symbolTable.Lookup(funcCall.Name)
	if !exists {
		sa.symbolTable.AddError(fmt.Sprintf("undefined function '%s' at line %d",
			funcCall.Name, funcCall.Line))
		return "unknown"
	}

	if symbol.SymbolType != SymbolFunction {
		sa.symbolTable.AddError(fmt.Sprintf("'%s' is not a function at line %d",
			funcCall.Name, funcCall.Line))
		return "unknown"
	}

	if len(funcCall.Arguments) != len(symbol.Parameters) {
		sa.symbolTable.AddError(fmt.Sprintf("function '%s' expects %d arguments, got %d at line %d",
			funcCall.Name, len(symbol.Parameters), len(funcCall.Arguments), funcCall.Line))
		return symbol.ReturnType
	}

	for i, arg := range funcCall.Arguments {
		argType := sa.analyzeExpression(arg)
		expectedType := symbol.Parameters[i].Type
		if !sa.areTypesCompatible(expectedType, argType) {
			sa.symbolTable.AddError(fmt.Sprintf("argument %d of function '%s': expected %s, got %s at line %d",
				i+1, funcCall.Name, expectedType, argType, funcCall.Line))
		}
	}

	return symbol.ReturnType
}

func (sa *SemanticAnalyzer) analyzeArrayLiteral(arrayLit *ArrayLiteral) string {
	if len(arrayLit.Elements) == 0 {
		return "[]unknown"
	}

	firstType := sa.analyzeExpression(arrayLit.Elements[0])

	for i, elem := range arrayLit.Elements[1:] {
		elemType := sa.analyzeExpression(elem)
		if !sa.areTypesCompatible(firstType, elemType) {
			sa.symbolTable.AddError(fmt.Sprintf("array element %d has inconsistent type at line %d",
				i+1, arrayLit.Line))
		}
	}

	return "[]" + firstType
}

func (sa *SemanticAnalyzer) isNumericType(typeName string) bool {
	return typeName == "int" || typeName == "float64" || typeName == "f64"
}

func (sa *SemanticAnalyzer) areTypesCompatible(type1, type2 string) bool {
	if type1 == type2 {
		return true
	}

	if sa.isNumericType(type1) && sa.isNumericType(type2) {
		return true
	}

	if type1 == "unknown" || type2 == "unknown" {
		return true
	}

	return false
}

func (sa *SemanticAnalyzer) getElementType(iterableType string) string {
	if len(iterableType) > 2 && iterableType[:2] == "[]" {
		return iterableType[2:]
	}
	return "unknown"
}

func (sa *SemanticAnalyzer) GetSymbolTable() *SymbolTable {
	return sa.symbolTable
}
