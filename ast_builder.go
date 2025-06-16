package main

import (
	ast "compiler/ast"
	parser "compiler/parser"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// ASTBuilder es el visitor que construye el AST
type ASTBuilder struct {
	parser.BaseVlangVisitor
	errors []string
}

// NewASTBuilder crea un nuevo constructor del AST
func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		errors: make([]string, 0),
	}
}

// Build construye el AST a partir del árbol de análisis sintáctico
func (b *ASTBuilder) Build(tree antlr.ParseTree) (*ast.Program, error) {
	result := b.Visit(tree)
	if len(b.errors) > 0 {
		return nil, fmt.Errorf("errors building AST: %v", b.errors)
	}

	program, ok := result.(*ast.Program)
	if !ok {
		return nil, fmt.Errorf("unexpected root node type: %T", result)
	}

	return program, nil
}

// Visit dispatches to the correct visit method
func (b *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(b)
}

// VisitPrograma construye el nodo Program del AST
func (b *ASTBuilder) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	program := &ast.Program{
		Statements: make([]ast.Statement, 0),
	}

	// Visitar todos los statements
	for _, stmtCtx := range ctx.AllStmt() {
		if stmt := b.visitStmt(stmtCtx.(*parser.StmtContext)); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}

func (b *ASTBuilder) visitStmt(ctx *parser.StmtContext) ast.Statement {
	if ctx == nil || ctx.GetChildCount() == 0 {
		return nil
	}

	child := ctx.GetChild(0)

	switch node := child.(type) {
	case *parser.PrintlnStmtContext:
		return b.visitPrintlnStmt(node)
	case *parser.PrintStmtContext:
		return b.visitPrintStmt(node)
	case *parser.DeclAssignContext:
		return b.visitDeclAssign(node)
	case *parser.DirectAssignContext:
		return b.visitDirectAssign(node)
	case *parser.PlusAssignContext:
		return b.visitPlusAssign(node)
	case *parser.MinusAssignContext:
		return b.visitMinusAssign(node)
	case *parser.MulAssignContext:
		return b.visitMulAssign(node)
	case *parser.DivAssignContext:
		return b.visitDivAssign(node)
	case *parser.ModAssignContext:
		return b.visitModAssign(node)
	case *parser.IfStmtContext:
		return b.visitIfStmt(node)
	case *parser.WhileStmtContext:
		return b.visitWhileStmt(node)
	case *parser.ForStmtContext:
		return b.visitForStmt(node)
	case *parser.FuncDeclContext:
		return b.visitFuncDecl(node)
	case *parser.FuncCallContext:
		// IMPORTANTE: Crear un ExpressionStatement para las llamadas a funciones
		expr := b.visitFuncCall(node)
		return &ast.ExpressionStatement{
			Expression: expr,
			Line:       node.GetStart().GetLine(),
			Column:     node.GetStart().GetColumn(),
		}
	case *parser.StructDeclContext:
		return b.visitStructDecl(node)
	case *parser.ReturnStmtContext:
		return b.visitReturnStmt(node)
	case *parser.BreakStmtContext:
		return b.visitBreakStmt(node)
	case *parser.ContinueStmtContext:
		return b.visitContinueStmt(node)
	default:
		b.addError(fmt.Sprintf("unhandled statement type: %T", node))
		return nil
	}
}

// AGREGAMOS LOS MÉTODOS FALTANTES:

// VisitIf_chain procesa una cadena if
func (b *ASTBuilder) VisitIf_chain(ctx *parser.IfChainContext) interface{} {
	condition := b.visitExpresion(ctx.Expresion())
	statements := make([]ast.Statement, 0)

	for _, stmtCtx := range ctx.AllStmt() {
		if stmt := b.visitStmt(stmtCtx.(*parser.StmtContext)); stmt != nil {
			statements = append(statements, stmt)
		}
	}

	return &IfChainNode{
		Condition:  condition,
		Statements: statements,
	}
}

// VisitElse_stmt procesa un else
func (b *ASTBuilder) VisitElse_stmt(ctx *parser.ElseStmtContext) interface{} {
	statements := make([]ast.Statement, 0)

	for _, stmtCtx := range ctx.AllStmt() {
		if stmt := b.visitStmt(stmtCtx.(*parser.StmtContext)); stmt != nil {
			statements = append(statements, stmt)
		}
	}

	return &ElseNode{
		Statements: statements,
	}
}

func (b *ASTBuilder) visitIfStmt(ctx *parser.IfStmtContext) ast.Statement {
	// Obtener la primera cadena if con type assertion
	firstChain := ctx.If_chain(0)

	// Convertir de interface a tipo concreto
	var condition ast.Expression
	var thenBranch []ast.Statement

	// Usar el visitor genérico para if_chain
	chainResult := b.Visit(firstChain)
	if ifChainNode, ok := chainResult.(*IfChainNode); ok {
		condition = ifChainNode.Condition
		thenBranch = ifChainNode.Statements
	}

	// Procesar else si existe
	elseBranch := make([]ast.Statement, 0)
	if ctx.Else_stmt() != nil {
		elseResult := b.Visit(ctx.Else_stmt())
		if elseNode, ok := elseResult.(*ElseNode); ok {
			elseBranch = elseNode.Statements
		}
	}

	return &ast.IfStmt{
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
		Line:       ctx.GetStart().GetLine(),
		Column:     ctx.GetStart().GetColumn(),
	}
}

// === PRINT STATEMENTS ===
func (b *ASTBuilder) visitPrintlnStmt(ctx *parser.PrintlnStmtContext) ast.Statement {
	var args []ast.Expression
	for _, exprCtx := range ctx.AllExpresion() {
		expr := b.visitExpresion(exprCtx)
		args = append(args, expr)
	}

	return &ast.PrintStmt{
		Arguments: args,
		NewLine:   true,
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitPrintStmt(ctx *parser.PrintStmtContext) ast.Statement {
	args := make([]ast.Expression, 0)
	for _, exprCtx := range ctx.AllExpresion() {
		args = append(args, b.visitExpresion(exprCtx))
	}

	return &ast.PrintStmt{
		Arguments: args,
		NewLine:   false,
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
	}
}

// === VARIABLE DECLARATIONS ===
func (b *ASTBuilder) visitDeclAssign(ctx *parser.DeclAssignContext) ast.Statement {
	name := ctx.ID().GetText()
	value := b.visitExpresion(ctx.Expresion())

	return &ast.VarDecl{
		Name:      name,
		Type:      "", // Tipo inferido
		Value:     value,
		IsMutable: true, // mut keyword
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitDirectAssign(ctx *parser.DirectAssignContext) ast.Statement {
	target := b.visitIdPattern(ctx.Id_pattern())
	value := b.visitExpresion(ctx.Expresion())

	return &ast.Assignment{
		Target: target,
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

// === CONTROL FLOW ===
// Removemos el método VisitIfStmt duplicado y usamos solo visitIfStmt

func (b *ASTBuilder) visitWhileStmt(ctx *parser.WhileStmtContext) ast.Statement {
	condition := b.visitExpresion(ctx.Expresion())
	body := make([]ast.Statement, 0)

	for _, stmtCtx := range ctx.AllStmt() {
		if stmt := b.visitStmt(stmtCtx.(*parser.StmtContext)); stmt != nil {
			body = append(body, stmt)
		}
	}

	return &ast.WhileStmt{
		Condition: condition,
		Body:      body,
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitForStmt(ctx *parser.ForStmtContext) ast.Statement {
	variable := ctx.ID().GetText()

	// El iterable puede ser una expresión o un range
	var iterable ast.Expression
	if exprCtx := ctx.Expresion(); exprCtx != nil {
		iterable = b.visitExpresion(exprCtx)
	} else if rangeCtx := ctx.Range_(); rangeCtx != nil {
		// TODO: Manejar ranges
		b.addError("ranges not yet implemented in AST")
	}

	body := make([]ast.Statement, 0)
	for _, stmtCtx := range ctx.AllStmt() {
		if stmt := b.visitStmt(stmtCtx.(*parser.StmtContext)); stmt != nil {
			body = append(body, stmt)
		}
	}

	return &ast.ForStmt{
		Variable: variable,
		Iterable: iterable,
		Body:     body,
		Line:     ctx.GetStart().GetLine(),
		Column:   ctx.GetStart().GetColumn(),
	}
}

// === FUNCTIONS ===
func (b *ASTBuilder) visitFuncDecl(ctx *parser.FuncDeclContext) ast.Statement {
	name := ctx.ID().GetText()
	params := make([]ast.Parameter, 0)

	if argList := ctx.Arg_list(); argList != nil {
		params = b.visitArgList(argList.(*parser.ArgListContext))
	}

	body := make([]ast.Statement, 0)
	for _, stmtCtx := range ctx.AllStmt() {
		if stmt := b.visitStmt(stmtCtx.(*parser.StmtContext)); stmt != nil {
			body = append(body, stmt)
		}
	}

	return &ast.FuncDecl{
		Name:       name,
		Parameters: params,
		ReturnType: "", // TODO: Manejar tipos de retorno
		Body:       body,
		Line:       ctx.GetStart().GetLine(),
		Column:     ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitArgList(ctx *parser.ArgListContext) []ast.Parameter {
	params := make([]ast.Parameter, 0)

	for _, argCtx := range ctx.AllFunc_arg() {
		param := b.visitFuncArg(argCtx.(*parser.FuncArgContext))
		params = append(params, param)
	}

	return params
}

func (b *ASTBuilder) visitFuncArg(ctx *parser.FuncArgContext) ast.Parameter {
	name := ctx.ID().GetText()
	varType := b.getVarType(ctx.Var_type())

	return ast.Parameter{
		Name: name,
		Type: varType,
	}
}

// === TRANSFER STATEMENTS ===
func (b *ASTBuilder) visitReturnStmt(ctx *parser.ReturnStmtContext) ast.Statement {
	var value ast.Expression
	if exprCtx := ctx.Expresion(); exprCtx != nil {
		value = b.visitExpresion(exprCtx)
	}

	return &ast.Return{
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitBreakStmt(ctx *parser.BreakStmtContext) ast.Statement {
	return &ast.Break{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitContinueStmt(ctx *parser.ContinueStmtContext) ast.Statement {
	return &ast.Continue{
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

// === STRUCTS ===
func (b *ASTBuilder) visitStructDecl(ctx *parser.StructDeclContext) ast.Statement {
	name := ctx.ID().GetText()
	fields := make([]ast.Field, 0)

	// TODO: La gramática muestra assign_stmt dentro del struct,
	// pero según el PDF deberían ser declaraciones de campos
	// Por ahora, lo dejamos vacío

	return &ast.StructDecl{
		Name:   name,
		Fields: fields,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

// === EXPRESSIONS ===
func (b *ASTBuilder) visitExpresion(ctx parser.IExpresionContext) ast.Expression {
	if ctx == nil {
		return nil
	}

	switch expr := ctx.(type) {
	case *parser.ValorexpresionContext:
		return b.visitValorExpresion(expr)
	case *parser.FuncionexpreContext:
		return b.visitFuncCall(expr.GetChild(0).(*parser.FuncCallContext))
	case *parser.ParentesisexpreContext:
		return b.visitExpresion(expr.Expresion())
	case *parser.UnarioContext:
		return b.visitUnaryExpr(expr)
	case *parser.BinaryExpContext:
		return b.visitBinaryExpr(expr)
	case *parser.IdContext:
		return b.visitIdExpr(expr)
	case *parser.ArrayexpreContext:
		return b.visitArrayExpr(expr)
	case *parser.AsignacionExprContext:
		return b.visitAssignmentExpr(expr)
	default:
		b.addError(fmt.Sprintf("unhandled expression type: %T", expr))
		return nil
	}
}

func (b *ASTBuilder) visitValorExpresion(ctx *parser.ValorexpresionContext) ast.Expression {
	valor := ctx.Valor()

	switch v := valor.(type) {
	case *parser.ValorEnteroContext:
		return b.visitValorEntero(v)
	case *parser.ValorFloatContext:
		return b.visitValorFloat(v)
	case *parser.ValorDecimalContext:
		return b.visitValorDecimal(v)
	case *parser.ValorCadenaContext:
		return b.visitValorCadena(v)
	case *parser.ValorBooleanoContext:
		return b.visitValorBooleano(v)
	case *parser.ValorCaracterContext:
		return b.visitValorCaracter(v)
	default:
		b.addError(fmt.Sprintf("unhandled valor type: %T", v))
		return nil
	}
}

func (b *ASTBuilder) visitUnaryExpr(ctx *parser.UnarioContext) ast.Expression {
	operator := ctx.GetOp().GetText()
	operand := b.visitExpresion(ctx.Expresion())

	return &ast.UnaryExpr{
		Operator: operator,
		Operand:  operand,
		Line:     ctx.GetStart().GetLine(),
		Column:   ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitBinaryExpr(ctx *parser.BinaryExpContext) ast.Expression {
	left := b.visitExpresion(ctx.GetLeft())
	operator := ctx.GetOp().GetText()
	right := b.visitExpresion(ctx.GetRight())

	return &ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
		Line:     ctx.GetStart().GetLine(),
		Column:   ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitIdExpr(ctx *parser.IdContext) ast.Expression {
	return &ast.Identifier{
		Name:   ctx.ID().GetText(),
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitFuncCall(ctx *parser.FuncCallContext) ast.Expression {
	// Manejar id_pattern correctamente
	idPattern := ctx.Id_pattern().(*parser.IdPatternContext)
	ids := idPattern.AllID()

	var name string
	if len(ids) > 0 {
		name = ids[0].GetText()
	}

	args := make([]ast.Expression, 0)

	if params := ctx.Parametros(); params != nil {
		paramList := params.(*parser.ParamListContext)
		for _, paramCtx := range paramList.AllFunc_param() {
			funcParam := paramCtx.(*parser.FuncParamContext)
			expr := funcParam.Expresion()
			args = append(args, b.visitExpresion(expr))
		}
	}

	return &ast.FuncCall{
		Name:      name,
		Arguments: args,
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitIdPattern(ctx parser.IId_patternContext) ast.Expression {
	idPattern := ctx.(*parser.IdPatternContext)
	ids := idPattern.AllID()

	if len(ids) == 1 {
		return &ast.Identifier{
			Name:   ids[0].GetText(),
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		}
	}

	// TODO: Manejar acceso a campos (ID.ID.ID)
	// Por ahora, solo retornamos el primer ID
	return &ast.Identifier{
		Name:   ids[0].GetText(),
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

// === LITERALS ===
func (b *ASTBuilder) visitValorEntero(ctx *parser.ValorEnteroContext) ast.Expression {
	val, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		b.addError(fmt.Sprintf("invalid integer: %s", ctx.GetText()))
		val = 0
	}

	return &ast.Literal{
		Value:  val,
		Type:   "int",
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitValorFloat(ctx *parser.ValorFloatContext) ast.Expression {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		b.addError(fmt.Sprintf("invalid float: %s", ctx.GetText()))
		val = 0.0
	}

	return &ast.Literal{
		Value:  val,
		Type:   "float64",
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitValorDecimal(ctx *parser.ValorDecimalContext) ast.Expression {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		b.addError(fmt.Sprintf("invalid decimal: %s", ctx.GetText()))
		val = 0.0
	}

	return &ast.Literal{
		Value:  val,
		Type:   "float64",
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitValorCadena(ctx *parser.ValorCadenaContext) ast.Expression {
	// Remover comillas y procesar escapes
	text := ctx.GetText()
	text = text[1 : len(text)-1] // Remover comillas

	text = strings.ReplaceAll(text, "\\\"", "\"")
	text = strings.ReplaceAll(text, "\\\\", "\\")
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\r", "\r")
	text = strings.ReplaceAll(text, "\\t", "\t")

	return &ast.Literal{
		Value:  text,
		Type:   "string",
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitValorBooleano(ctx *parser.ValorBooleanoContext) ast.Expression {
	val := ctx.GetText() == "true"

	return &ast.Literal{
		Value:  val,
		Type:   "bool",
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitValorCaracter(ctx *parser.ValorCaracterContext) ast.Expression {
	text := ctx.GetText()
	// Remover comillas simples
	if len(text) >= 3 {
		char := text[1 : len(text)-1]
		return &ast.Literal{
			Value:  char,
			Type:   "char",
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		}
	}

	return &ast.Literal{
		Value:  "",
		Type:   "char",
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

// === COMPOUND ASSIGNMENTS ===
func (b *ASTBuilder) visitPlusAssign(ctx *parser.PlusAssignContext) ast.Statement {
	target := b.visitIdPattern(ctx.Id_pattern())
	value := b.visitExpresion(ctx.Expresion())

	return &ast.PlusAssign{
		Target: target,
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitMinusAssign(ctx *parser.MinusAssignContext) ast.Statement {
	target := b.visitIdPattern(ctx.Id_pattern())
	value := b.visitExpresion(ctx.Expresion())

	return &ast.MinusAssign{
		Target: target,
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitMulAssign(ctx *parser.MulAssignContext) ast.Statement {
	target := b.visitIdPattern(ctx.Id_pattern())
	value := b.visitExpresion(ctx.Expresion())

	return &ast.MulAssign{
		Target: target,
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitDivAssign(ctx *parser.DivAssignContext) ast.Statement {
	target := b.visitIdPattern(ctx.Id_pattern())
	value := b.visitExpresion(ctx.Expresion())

	return &ast.DivAssign{
		Target: target,
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitModAssign(ctx *parser.ModAssignContext) ast.Statement {
	target := b.visitIdPattern(ctx.Id_pattern())
	value := b.visitExpresion(ctx.Expresion())

	return &ast.ModAssign{
		Target: target,
		Value:  value,
		Line:   ctx.GetStart().GetLine(),
		Column: ctx.GetStart().GetColumn(),
	}
}

func (b *ASTBuilder) visitArrayExpr(ctx *parser.ArrayexpreContext) ast.Expression {
	elements := make([]ast.Expression, 0)

	// Iterar sobre todos los elementos del array
	for _, exprCtx := range ctx.AllExpresion() {
		element := b.visitExpresion(exprCtx)
		if element != nil {
			elements = append(elements, element)
		}
	}

	return &ast.ArrayLiteral{
		Elements: elements,
		Line:     ctx.GetStart().GetLine(),
		Column:   ctx.GetStart().GetColumn(),
	}
}

// Si tienes asignaciones como expresiones, también agregar esta función
func (b *ASTBuilder) visitAssignmentExpr(ctx *parser.AsignacionExprContext) ast.Expression {
	// Esto dependería de cómo esté estructurada tu gramática para asignaciones como expresiones
	// Por ahora, retornar nil o manejar según tu caso específico
	b.addError("assignment expressions not yet implemented")
	return nil
}

// === HELPER METHODS ===
func (b *ASTBuilder) getVarType(ctx parser.IVar_typeContext) string {
	if ctx == nil {
		return ""
	}

	switch ctx.(type) {
	case *parser.IntTypeContext:
		return "int"
	case *parser.FloatTypeContext:
		return "float64"
	case *parser.StringTypeContext:
		return "string"
	case *parser.BoolTypeContext:
		return "bool"
	case *parser.CharTypeContext:
		return "char"
	case *parser.VoidTypeContext:
		return "void"
	default:
		return ""
	}
}

func (b *ASTBuilder) addError(msg string) {
	b.errors = append(b.errors, msg)
}

// === NODOS AUXILIARES (temporales para el AST builder) ===
type IfChainNode struct {
	Condition  ast.Expression
	Statements []ast.Statement
}

type ElseNode struct {
	Statements []ast.Statement
}

// === EXPRESIONES ADICIONALES ===
type RangeExpr struct {
	Start  ast.Expression
	End    ast.Expression
	Line   int
	Column int
}

func (r *RangeExpr) IsExpression()  {}
func (r *RangeExpr) GetLine() int   { return r.Line }
func (r *RangeExpr) GetColumn() int { return r.Column }
