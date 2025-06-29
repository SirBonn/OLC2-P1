// generator.go
package arm64

import (
	"compiler/ast"
	"compiler/codegen"
	"fmt"
	"strings"
)

// VariableInfo contiene información sobre una variable
type VariableInfo struct {
	Name       string   // nombre de la variable
	Type       string   // tipo de la variable
	Offset     int      // offset en el stack frame (negativo desde x29)
	Size       int      // tamaño en bytes
	IsParam    bool     // si es un parámetro de función
	Register   Register // registro asignado (si aplica)
	InRegister bool     // si está actualmente en un registro
}

type ControlFlowContext struct {
	Type       string // "switch", "while", "for"
	StartLabel string
	EndLabel   string
	BreakLabel string
}

// getSize retorna el tamaño en bytes según el tipo
func (vi *VariableInfo) getSize() int {
	switch vi.Type {
	case "int", "bool":
		return 8 // 64 bits
	case "string":
		return 8 // puntero
	case "f64", "float64":
		return 8 // 64 bits
	case "char":
		return 1 // 8 bits
	default:
		return 8 // por defecto
	}
}

// ARM64Generator genera código ensamblador ARM64
type ARM64Generator struct {
	*codegen.BaseGenerator
	currentFunction string
	stackOffset     int
	labelCount      int
	registers       *RegisterAllocator
	stringLiterals  map[string]string // mapa de literales string a etiquetas
	dataSection     []string          // sección .data

	// Sistema de variables
	symbolTable        map[string]*VariableInfo   // tabla de símbolos
	scopeStack         []map[string]*VariableInfo // stack de scopes
	currentStackOffset int                        // offset actual del stack frame
	controlFlowStack   []*ControlFlowContext
}

// NewARM64Generator crea un nuevo generador ARM64
func NewARM64Generator() *ARM64Generator {
	return &ARM64Generator{
		BaseGenerator:  codegen.NewBaseGenerator(),
		stackOffset:    0,
		labelCount:     0,
		registers:      NewRegisterAllocator(),
		stringLiterals: make(map[string]string),
		dataSection:    make([]string, 0),

		// Sistema de variables
		symbolTable:        make(map[string]*VariableInfo),
		scopeStack:         make([]map[string]*VariableInfo, 0),
		currentStackOffset: 0,
	}
}

// Generate es el punto de entrada principal
func (g *ARM64Generator) Generate(node ast.Node) (string, error) {
	g.Reset()

	// El nodo debería ser un *ast.Program
	program, ok := node.(*ast.Program)
	if !ok {
		return "", fmt.Errorf("expected *ast.Program, got %T", node)
	}

	// Visitar el programa para generar código
	program.Accept(g)

	// Si hay errores, retornarlos
	if g.HasErrors() {
		return "", g.GetErrors()
	}

	// Construir el output final
	return g.buildFinalOutput(), nil
}

func (g *ARM64Generator) buildFinalOutput() string {
	var output strings.Builder

	// Sección de datos mejorada
	output.WriteString(".data\n")
	output.WriteString("print_fmt: .asciz \"%ld\\n\"\n")    // Formato para números (long)
	output.WriteString("print_fmt_no_nl: .asciz \"%ld\"\n") // Sin newline

	// Agregar string literals
	for str, label := range g.stringLiterals {
		output.WriteString(fmt.Sprintf("%s:\n", label))
		output.WriteString(fmt.Sprintf("\t.asciz \"%s\"\n", escapeString(str)))
	}
	output.WriteString("\n")

	// Sección de texto
	output.WriteString(".text\n")
	output.WriteString(".global main\n\n")

	// Agregar el código generado
	output.WriteString(g.GetOutput())

	return output.String()
}

// hasMainFunction verifica si existe una función main
func (g *ARM64Generator) hasMainFunction() bool {
	return strings.Contains(g.GetOutput(), "main:")
}

// reset limpia el estado del generador
func (g *ARM64Generator) reset() {
	g.BaseGenerator.Reset()
	g.currentFunction = ""
	g.stackOffset = 0
	g.labelCount = 0
	g.registers = NewRegisterAllocator()
	g.stringLiterals = make(map[string]string)
	g.dataSection = make([]string, 0)

	// Sistema de variables
	g.symbolTable = make(map[string]*VariableInfo)
	g.scopeStack = make([]map[string]*VariableInfo, 0)
	g.currentStackOffset = 0
}

// newLabel genera una nueva etiqueta única
func (g *ARM64Generator) newLabel(prefix string) string {
	label := fmt.Sprintf(".L%s_%d", prefix, g.labelCount)
	g.labelCount++
	return label
}

// getStringLabel obtiene o crea una etiqueta para un string literal
func (g *ARM64Generator) getStringLabel(str string) string {
	if label, exists := g.stringLiterals[str]; exists {
		return label
	}

	label := g.newLabel("str")
	g.stringLiterals[str] = label
	return label
}

// === MÉTODOS PARA MANEJO DE VARIABLES ===

// enterScope crea un nuevo scope
func (g *ARM64Generator) enterScope() {
	newScope := make(map[string]*VariableInfo)
	g.scopeStack = append(g.scopeStack, newScope)
}

// exitScope sale del scope actual
func (g *ARM64Generator) exitScope() {
	if len(g.scopeStack) > 0 {
		g.scopeStack = g.scopeStack[:len(g.scopeStack)-1]
	}
}

// addVariable agrega una variable al scope actual
func (g *ARM64Generator) addVariable(name, varType string) *VariableInfo {
	// Calcular tamaño y offset
	size := g.getTypeSize(varType)
	g.currentStackOffset -= size

	// Alinear a 8 bytes
	if g.currentStackOffset%8 != 0 {
		g.currentStackOffset -= (8 - (-g.currentStackOffset % 8))
	}

	varInfo := &VariableInfo{
		Name:       name,
		Type:       varType,
		Offset:     g.currentStackOffset,
		Size:       size,
		IsParam:    false,
		InRegister: false,
	}

	// Agregar al scope actual y tabla global
	if len(g.scopeStack) > 0 {
		g.scopeStack[len(g.scopeStack)-1][name] = varInfo
	}
	g.symbolTable[name] = varInfo

	return varInfo
}

// findVariable busca una variable en los scopes
func (g *ARM64Generator) findVariable(name string) *VariableInfo {
	// Buscar en scopes desde el más reciente al más antiguo
	for i := len(g.scopeStack) - 1; i >= 0; i-- {
		if varInfo, exists := g.scopeStack[i][name]; exists {
			return varInfo
		}
	}

	// Buscar en tabla global como fallback
	if varInfo, exists := g.symbolTable[name]; exists {
		return varInfo
	}

	return nil
}

// getTypeSize retorna el tamaño en bytes de un tipo
func (g *ARM64Generator) getTypeSize(varType string) int {
	switch varType {
	case "int", "bool":
		return 8
	case "string":
		return 8 // puntero
	case "f64", "float64":
		return 8
	case "char":
		return 1
	default:
		return 8
	}
}

// storeVariable almacena una variable en el stack
func (g *ARM64Generator) storeVariable(varInfo *VariableInfo) {
	g.Emit("\t// Store variable %s at offset %d", varInfo.Name, varInfo.Offset)
	g.Emit("\tstr x0, [x29, #%d]", varInfo.Offset)
}

// loadVariable carga una variable desde el stack
func (g *ARM64Generator) loadVariable(varInfo *VariableInfo) {
	g.Emit("\t// Load variable %s from offset %d", varInfo.Name, varInfo.Offset)
	g.Emit("\tldr x0, [x29, #%d]", varInfo.Offset)
}

// allocateStackSpace reserva espacio en el stack para variables locales
func (g *ARM64Generator) allocateStackSpace() {
	if g.currentStackOffset < 0 {
		// Alinear a 16 bytes (requerimiento ARM64)
		totalSize := -g.currentStackOffset
		if totalSize%16 != 0 {
			totalSize += 16 - (totalSize % 16)
		}

		g.Emit("\t// Allocate stack space for local variables")
		g.Emit("\tsub sp, sp, #%d", totalSize)
	}
}

// === IMPLEMENTACIÓN DEL VISITOR PATTERN ===

func (g *ARM64Generator) VisitProgram(node *ast.Program) interface{} {
	g.Emit("\t// Function declarations")
	for _, stmt := range node.Statements {
		if funcDecl, ok := stmt.(*ast.FuncDecl); ok {
			_ = funcDecl // Solo procesar declaraciones de funciones
			stmt.Accept(g)
		}
	}

	// Segunda pasada: Verificar si hay main
	hasMain := false
	for _, stmt := range node.Statements {
		if funcDecl, ok := stmt.(*ast.FuncDecl); ok {
			if funcDecl.Name == "main" {
				hasMain = true
			}
		}
	}

	// Si no hay main, crear uno para código top-level
	if !hasMain {
		g.Emit("main:")
		g.Emit("\t// Setup stack frame")
		g.Emit("\tstp x29, x30, [sp, #-16]!")
		g.Emit("\tmov x29, sp")
		g.Emit("")

		// Crear scope para main generado
		g.enterScope()

		// Procesar statements que no son funciones
		for _, stmt := range node.Statements {
			if _, ok := stmt.(*ast.FuncDecl); !ok {
				stmt.Accept(g)
			}
		}

		g.exitScope()

		// Retornar 0
		g.Emit("\tmov w0, #0")
		g.Emit("\tldp x29, x30, [sp], #16")
		g.Emit("\tret")
	}

	return nil
}

// VisitBinaryExpr genera código para expresiones binarias
func (g *ARM64Generator) VisitBinaryExpr(node *ast.BinaryExpr) interface{} {
	// Evaluar operando izquierdo
	leftReg := g.allocateRegister()
	node.Left.Accept(g)
	g.Emit("\tmov %s, x0", leftReg) // Asumimos que el resultado está en x0

	// Evaluar operando derecho
	rightReg := g.allocateRegister()
	node.Right.Accept(g)
	g.Emit("\tmov %s, x0", rightReg)

	// Realizar la operación
	switch node.Operator {
	case "+":
		g.Emit("\tadd x0, %s, %s", leftReg, rightReg)
	case "-":
		g.Emit("\tsub x0, %s, %s", leftReg, rightReg)
	case "*":
		g.Emit("\tmul x0, %s, %s", leftReg, rightReg)
	case "/":
		g.Emit("\tsdiv x0, %s, %s", leftReg, rightReg)
	case "%":
		// ARM64 no tiene instrucción de módulo directa
		g.Emit("\tsdiv x2, %s, %s", leftReg, rightReg)
		g.Emit("\tmsub x0, x2, %s, %s", rightReg, leftReg)
	case "==":
		g.Emit("\tcmp %s, %s", leftReg, rightReg)
		g.Emit("\tcset x0, eq")
	case "!=":
		g.Emit("\tcmp %s, %s", leftReg, rightReg)
		g.Emit("\tcset x0, ne")
	case "<":
		g.Emit("\tcmp %s, %s", leftReg, rightReg)
		g.Emit("\tcset x0, lt")
	case "<=":
		g.Emit("\tcmp %s, %s", leftReg, rightReg)
		g.Emit("\tcset x0, le")
	case ">":
		g.Emit("\tcmp %s, %s", leftReg, rightReg)
		g.Emit("\tcset x0, gt")
	case ">=":
		g.Emit("\tcmp %s, %s", leftReg, rightReg)
		g.Emit("\tcset x0, ge")
	case "&&":
		g.Emit("\tand x0, %s, %s", leftReg, rightReg)
	case "||":
		g.Emit("\torr x0, %s, %s", leftReg, rightReg)
	default:
		g.AddError(fmt.Errorf("unsupported binary operator: %s", node.Operator))
	}

	// Liberar registros
	g.freeRegister(leftReg)
	g.freeRegister(rightReg)

	return nil
}

// VisitUnaryExpr genera código para expresiones unarias
func (g *ARM64Generator) VisitUnaryExpr(node *ast.UnaryExpr) interface{} {
	// Evaluar el operando
	node.Operand.Accept(g)

	switch node.Operator {
	case "-":
		g.Emit("\tneg x0, x0")
	case "!":
		g.Emit("\tcmp x0, #0")
		g.Emit("\tcset x0, eq")
	default:
		g.AddError(fmt.Errorf("unsupported unary operator: %s", node.Operator))
	}

	return nil
}

// VisitLiteral genera código para literales
func (g *ARM64Generator) VisitLiteral(node *ast.Literal) interface{} {
	switch node.Type {
	case "int":
		if val, ok := node.Value.(int); ok {
			g.Emit("\tmov x0, #%d", val)
		}
	case "bool":
		if val, ok := node.Value.(bool); ok {
			if val {
				g.Emit("\tmov x0, #1")
			} else {
				g.Emit("\tmov x0, #0")
			}
		}
	case "string":
		if val, ok := node.Value.(string); ok {
			label := g.getStringLabel(val)
			g.Emit("\tadr x0, %s", label)
		}
	case "f64", "float64":
		// Para floats, necesitaríamos usar registros de punto flotante
		g.AddError(fmt.Errorf("float literals not yet implemented"))
	default:
		g.AddError(fmt.Errorf("unsupported literal type: %s", node.Type))
	}

	return nil
}

// VisitIdentifier genera código para identificadores
func (g *ARM64Generator) VisitIdentifier(node *ast.Identifier) interface{} {
	// Buscar la variable en la tabla de símbolos
	varInfo := g.findVariable(node.Name)
	if varInfo == nil {
		g.AddError(fmt.Errorf("undefined variable: %s", node.Name))
		g.Emit("\tmov x0, #0 // undefined variable")
		return nil
	}

	// Cargar la variable
	g.loadVariable(varInfo)
	return nil
}

func (g *ARM64Generator) VisitPrintStmt(node *ast.PrintStmt) interface{} {
	g.Emit("\t// Print statement")

	for i, arg := range node.Arguments {
		arg.Accept(g)

		isStringLiteral := false
		if strLit, ok := arg.(*ast.Literal); ok && strLit.Type == "string" {
			isStringLiteral = true
		}

		if isStringLiteral {
			g.Emit("\t// Print string value")
			g.Emit("\tbl puts") // puts maneja strings directamente
		} else {
			g.Emit("\t// Print integer value")
			g.Emit("\tmov x1, x0")        // valor a imprimir
			g.Emit("\tadr x0, print_fmt") // formato
			g.Emit("\tbl printf")         // llamar a printf
		}

		if i < len(node.Arguments)-1 {
			g.Emit("\t// Print space")
			g.Emit("\tmov x0, #32") // ASCII space
			g.Emit("\tbl putchar")
		}
	}

	if node.NewLine {
		needsNewline := true
		if len(node.Arguments) == 1 {
			if strLit, ok := node.Arguments[0].(*ast.Literal); ok && strLit.Type == "string" {
				needsNewline = false // puts ya agrega newline
			}
		}

		if needsNewline {
			g.Emit("\t// Print newline")
			g.Emit("\tmov x0, #10") // ASCII newline
			g.Emit("\tbl putchar")
		}
	}

	return nil
}

// VisitVarDecl genera código para declaraciones de variables
func (g *ARM64Generator) VisitVarDecl(node *ast.VarDecl) interface{} {
	g.Emit("\t// Variable declaration: %s", node.Name)

	varType := "int" // tipo por defecto
	if node.Type != "" {
		varType = node.Type
	}

	varInfo := g.addVariable(node.Name, varType)

	if node.Value != nil {
		node.Value.Accept(g)
	} else {
		switch varType {
		case "int":
			g.Emit("\tmov x0, #0")
		case "bool":
			g.Emit("\tmov x0, #0") // false
		case "string":
			g.Emit("\tmov x0, #0") // null string
		default:
			g.Emit("\tmov x0, #0")
		}
	}

	g.storeVariable(varInfo)

	return nil
}

// VisitAssignment genera código para asignaciones
func (g *ARM64Generator) VisitAssignment(node *ast.Assignment) interface{} {
	g.Emit("\t// Assignment")

	// Evaluar el valor
	node.Value.Accept(g)

	// Obtener el target (debe ser un identificador)
	if id, ok := node.Target.(*ast.Identifier); ok {
		varInfo := g.findVariable(id.Name)
		if varInfo == nil {
			g.AddError(fmt.Errorf("undefined variable: %s", id.Name))
			return nil
		}

		// Almacenar el valor
		g.storeVariable(varInfo)
	} else {
		g.AddError(fmt.Errorf("assignment target must be an identifier"))
	}

	return nil
}

// VisitIfStmt genera código para declaraciones if
func (g *ARM64Generator) VisitIfStmt(node *ast.IfStmt) interface{} {
	elseLabel := g.newLabel("else")
	endLabel := g.newLabel("endif")

	g.Emit("\t// If statement")

	// Evaluar condición
	node.Condition.Accept(g)

	// Saltar a else si es falso
	g.Emit("\tcbz x0, %s", elseLabel)

	// Generar código del then
	for _, stmt := range node.ThenBranch {
		stmt.Accept(g)
	}

	// Saltar al final
	g.Emit("\tb %s", endLabel)

	// Etiqueta else
	g.Emit("%s:", elseLabel)

	// Generar código del else (si existe)
	if node.ElseIf != nil {
		node.ElseIf.Accept(g)
	} else if len(node.ElseBranch) > 0 {
		for _, stmt := range node.ElseBranch {
			stmt.Accept(g)
		}
	}

	// Etiqueta final
	g.Emit("%s:", endLabel)

	return nil
}

// VisitWhileStmt genera código para loops while
func (g *ARM64Generator) VisitWhileStmt(node *ast.WhileStmt) interface{} {
	startLabel := g.newLabel("while_start")
	endLabel := g.newLabel("while_end")

	// Registrar contexto
	g.pushControlFlow("while", startLabel, endLabel)
	defer g.popControlFlow()

	g.Emit("\t// While loop")
	g.Emit("%s:", startLabel)

	// Evaluar condición
	node.Condition.Accept(g)

	// Salir si es falso
	g.Emit("\tcbz x0, %s", endLabel)

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Volver al inicio
	g.Emit("\tb %s", startLabel)

	// Etiqueta final
	g.Emit("%s:", endLabel)

	return nil
}

func (g *ARM64Generator) VisitFuncDecl(node *ast.FuncDecl) interface{} {
	oldFunction := g.currentFunction
	oldStackOffset := g.currentStackOffset

	g.currentFunction = node.Name
	g.currentStackOffset = 0

	// Generar etiqueta de función
	if node.Name == "main" {
		g.Emit(".global %s", node.Name)
	}
	g.Emit("%s:", node.Name)

	// Prólogo estándar
	g.Emit("\t// Function prologue")
	g.Emit("\tstp x29, x30, [sp, #-16]!")
	g.Emit("\tmov x29, sp")

	// Crear scope para la función
	g.enterScope()

	// Procesar parámetros si existen
	if len(node.Parameters) > 0 {
		g.Emit("\t// Store function parameters")
		for i, param := range node.Parameters {
			// Crear variable para el parámetro
			paramVar := g.addVariable(param.Name, param.Type)

			if i < 8 {
				// Parámetros en registros x0-x7
				g.Emit("\t// Parameter %s from x%d", param.Name, i)
				g.Emit("\tstr x%d, [x29, #%d]", i, paramVar.Offset)
			} else {
				// Parámetros adicionales desde caller stack
				g.Emit("\t// Parameter %s from caller stack", param.Name)
				g.Emit("\tldr x0, [x29, #%d]", 16+(i-8)*8)
				g.Emit("\tstr x0, [x29, #%d]", paramVar.Offset)
			}
		}
	}

	// Reservar espacio para variables locales después de procesar parámetros
	// pero antes del cuerpo de la función
	stackSpaceNeeded := false
	savedPosition := len(strings.Split(g.GetOutput(), "\n"))

	_ = savedPosition    // Guardar posición para insertar stack space
	_ = stackSpaceNeeded // Variable para saber si necesitamos reservar espacio

	// Generar código del cuerpo
	hasExplicitReturn := false
	for _, stmt := range node.Body {
		stmt.Accept(g)
		if _, ok := stmt.(*ast.Return); ok {
			hasExplicitReturn = true
		}
	}

	// Calcular y reservar espacio para variables locales
	if g.currentStackOffset < 0 {
		stackSize := -g.currentStackOffset
		// Alinear a 16 bytes
		if stackSize%16 != 0 {
			stackSize += 16 - (stackSize % 16)
		}

		// Insertar reserva de stack después del prólogo
		g.Emit("\t// Note: Stack space of %d bytes needed for local variables", stackSize)
		// TODO: En una implementación real, insertaríamos esto después del prólogo
		stackSpaceNeeded = true
	}

	// Salir del scope
	g.exitScope()

	// Epílogo si no hay return explícito
	if !hasExplicitReturn {
		g.Emit("\t// Function epilogue (implicit return)")

		// Valor de retorno por defecto
		if node.ReturnType != "" && node.ReturnType != "void" {
			g.Emit("\tmov x0, #0 // Default return value")
		}

		// Restaurar stack y retornar
		g.Emit("\tldp x29, x30, [sp], #16")
		g.Emit("\tret")
	}

	// Restaurar estado
	g.currentFunction = oldFunction
	g.currentStackOffset = oldStackOffset

	return nil
}

func (g *ARM64Generator) VisitReturn(node *ast.Return) interface{} {
	g.Emit("\t// Return statement")

	if node.Value != nil {
		// Evaluar el valor de retorno
		node.Value.Accept(g)
		// El resultado ya está en x0
	} else {
		// Return sin valor
		g.Emit("\tmov x0, #0")
	}

	// Restaurar stack si había variables locales
	if g.currentStackOffset < 0 {
		stackSize := -g.currentStackOffset
		if stackSize%16 != 0 {
			stackSize += 16 - (stackSize % 16)
		}
		g.Emit("\t// Note: Should restore %d bytes of stack space", stackSize)
		// TODO: En implementación real, restaurar stack aquí
	}

	// Epílogo y retorno
	g.Emit("\tldp x29, x30, [sp], #16")
	g.Emit("\tret")

	return nil
}

// === MÉTODOS STUB PARA OTROS NODOS ===

func (g *ARM64Generator) VisitFuncCall(node *ast.FuncCall) interface{} {
	g.Emit("\t// Function call: %s", node.Name)

	// Evaluar argumentos y colocarlos en registros/stack
	for i, arg := range node.Arguments {
		arg.Accept(g) // Resultado en x0

		if i < 8 {
			// Primeros 8 argumentos van en registros x0-x7
			if i > 0 {
				g.Emit("\tmov x%d, x0", i)
			}
			// El primer argumento ya está en x0
		} else {
			// Argumentos adicionales van en el stack
			g.Emit("\tstr x0, [sp, #%d]", (i-8)*8)
		}
	}

	// Llamar a la función
	g.Emit("\tbl %s", node.Name)

	// El resultado ya está en x0, no necesitamos hacer nada más

	return nil
}

func (g *ARM64Generator) VisitBreak(node *ast.Break) interface{} {
	g.Emit("\t// Break statement")

	ctx := g.currentControlFlow()
	if ctx == nil {
		g.AddError(fmt.Errorf("break statement outside of loop or switch"))
		return nil
	}

	g.Emit("\tb %s // Break to end of %s", ctx.BreakLabel, ctx.Type)
	return nil
}

func (g *ARM64Generator) VisitContinue(node *ast.Continue) interface{} {
	g.Emit("\t// Continue statement")

	ctx := g.currentControlFlow()
	if ctx == nil {
		g.AddError(fmt.Errorf("continue statement outside of loop"))
		return nil
	}

	if ctx.Type != "for" && ctx.Type != "while" {
		g.AddError(fmt.Errorf("continue statement only valid in loops"))
		return nil
	}

	g.Emit("\tb %s // Continue to start of %s", ctx.StartLabel, ctx.Type)
	return nil
}

func (g *ARM64Generator) VisitForStmt(node *ast.ForStmt) interface{} {
	startLabel := g.newLabel("for_in_start")
	endLabel := g.newLabel("for_in_end")
	continueLabel := g.newLabel("for_in_continue")

	g.Emit("\t// For-in loop: %s", node.Variable)

	// Registrar contexto para break/continue
	ctx := &ControlFlowContext{
		Type:       "for",
		StartLabel: continueLabel,
		EndLabel:   endLabel,
		BreakLabel: endLabel,
	}
	g.controlFlowStack = append(g.controlFlowStack, ctx)
	defer g.popControlFlow()

	// Crear scope para el loop
	g.enterScope()
	defer g.exitScope()

	// TODO: Por ahora implementación básica
	// En una implementación completa necesitaríamos:
	// 1. Evaluar la expresión/range
	// 2. Determinar si es un rango numérico o una colección
	// 3. Crear variable de iteración
	// 4. Generar loop apropiado

	g.Emit("\t// TODO: Implement for-in loop properly")
	g.Emit("\t// Variable: %s", node.Variable)

	// Por ahora, generar placeholder que no hace nada
	g.Emit("%s:", startLabel)
	g.Emit("\t// Evaluate iterable expression")
	node.Iterable.Accept(g)

	g.Emit("\t// TODO: Check if more elements")
	g.Emit("\tcbz x0, %s", endLabel)

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	g.Emit("%s:", continueLabel)
	g.Emit("\t// TODO: Advance iterator")
	g.Emit("\tb %s", startLabel)

	g.Emit("%s:", endLabel)

	return nil
}

func (g *ARM64Generator) VisitStructDecl(node *ast.StructDecl) interface{} {
	g.Emit("\t// TODO: Struct declaration: %s", node.Name)
	return nil
}

func (g *ARM64Generator) VisitStructInstance(node *ast.StructInstance) interface{} {
	g.Emit("\t// TODO: Struct instance")
	return nil
}

func (g *ARM64Generator) VisitArrayLiteral(node *ast.ArrayLiteral) interface{} {
	g.Emit("\t// TODO: Array literal")
	return nil
}

func (g *ARM64Generator) VisitExpressionStatement(node *ast.ExpressionStatement) interface{} {
	// Simplemente evaluar la expresión
	node.Expression.Accept(g)
	return nil
}

func (g *ARM64Generator) VisitForCondition(node *ast.ForCondition) interface{} {
	startLabel := g.newLabel("for_cond_start")
	endLabel := g.newLabel("for_cond_end")
	continueLabel := g.newLabel("for_cond_continue")

	g.Emit("\t// For loop with condition")

	// Registrar contexto para break/continue
	ctx := &ControlFlowContext{
		Type:       "for",
		StartLabel: continueLabel, // continue va a la evaluación de condición
		EndLabel:   endLabel,
		BreakLabel: endLabel,
	}
	g.controlFlowStack = append(g.controlFlowStack, ctx)
	defer g.popControlFlow()

	// Crear scope para el loop
	g.enterScope()
	defer g.exitScope()

	// Etiqueta de inicio y continue
	g.Emit("%s:", startLabel)
	g.Emit("%s:", continueLabel)

	// Evaluar condición
	node.Condition.Accept(g)

	// Salir si es falso
	g.Emit("\tcbz x0, %s", endLabel)

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Volver a evaluar condición
	g.Emit("\tb %s", continueLabel)

	// Etiqueta final
	g.Emit("%s:", endLabel)

	return nil
}

func (g *ARM64Generator) VisitForClassic(node *ast.ForClassic) interface{} {
	startLabel := g.newLabel("for_classic_start")
	endLabel := g.newLabel("for_classic_end")
	continueLabel := g.newLabel("for_classic_continue")
	conditionLabel := g.newLabel("for_classic_condition")

	g.Emit("\t// Classic for loop")

	// Registrar contexto para break/continue
	ctx := &ControlFlowContext{
		Type:       "for",
		StartLabel: continueLabel, // continue va al update
		EndLabel:   endLabel,
		BreakLabel: endLabel,
	}
	g.controlFlowStack = append(g.controlFlowStack, ctx)
	defer g.popControlFlow()

	// Crear scope para el loop (incluye la variable de inicialización)
	g.enterScope()
	defer g.exitScope()

	// Inicialización (si existe)
	if node.Init != nil {
		node.Init.Accept(g)
	}

	// Saltar a la evaluación de condición
	g.Emit("\tb %s", conditionLabel)

	// Etiqueta de inicio del cuerpo
	g.Emit("%s:", startLabel)

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Etiqueta continue (ejecutar update)
	g.Emit("%s:", continueLabel)

	// Update (si existe)
	if node.Update != nil {
		node.Update.Accept(g)
	}

	// Etiqueta para evaluación de condición
	g.Emit("%s:", conditionLabel)

	// Evaluar condición (si existe)
	if node.Condition != nil {
		node.Condition.Accept(g)
		// Continuar si es verdadero
		g.Emit("\tcbnz x0, %s", startLabel)
	} else {
		// Sin condición = loop infinito
		g.Emit("\tb %s", startLabel)
	}

	// Etiqueta final
	g.Emit("%s:", endLabel)

	return nil
}

func (g *ARM64Generator) VisitForIndexValue(node *ast.ForIndexValue) interface{} {
	startLabel := g.newLabel("for_idx_val_start")
	endLabel := g.newLabel("for_idx_val_end")
	continueLabel := g.newLabel("for_idx_val_continue")

	g.Emit("\t// For index,value loop: %s, %s", node.Index, node.Value)

	// Registrar contexto para break/continue
	ctx := &ControlFlowContext{
		Type:       "for",
		StartLabel: continueLabel,
		EndLabel:   endLabel,
		BreakLabel: endLabel,
	}
	g.controlFlowStack = append(g.controlFlowStack, ctx)
	defer g.popControlFlow()

	// Crear scope para el loop
	g.enterScope()
	defer g.exitScope()

	// Crear variables para index y value
	indexVar := g.addVariable(node.Index, "int")
	valueVar := g.addVariable(node.Value, "int") // Asumimos int por ahora

	// Inicializar index a 0
	g.Emit("\tmov x0, #0")
	g.storeVariable(indexVar)

	g.Emit("%s:", startLabel)

	// TODO: Evaluar la expresión iterable y obtener elemento en índice actual
	g.Emit("\t// TODO: Get element at current index")
	node.Iterable.Accept(g)

	// Por ahora, simular que obtenemos un valor
	g.Emit("\t// TODO: Check if index is valid")
	g.loadVariable(indexVar)
	g.Emit("\tcmp x0, #10") // Simular límite de 10 elementos
	g.Emit("\tbge %s", endLabel)

	// Almacenar valor (placeholder)
	g.Emit("\t// TODO: Store actual value")
	g.loadVariable(indexVar)
	g.storeVariable(valueVar)

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	g.Emit("%s:", continueLabel)

	// Incrementar index
	g.loadVariable(indexVar)
	g.Emit("\tadd x0, x0, #1")
	g.storeVariable(indexVar)

	g.Emit("\tb %s", startLabel)
	g.Emit("%s:", endLabel)

	return nil
}

func (g *ARM64Generator) VisitForInfinite(node *ast.ForInfinite) interface{} {
	startLabel := g.newLabel("for_infinite_start")
	endLabel := g.newLabel("for_infinite_end")

	g.Emit("\t// Infinite for loop")

	// Registrar contexto para break/continue
	g.pushControlFlow("for", startLabel, endLabel)
	defer g.popControlFlow()

	// Crear scope para el loop
	g.enterScope()
	defer g.exitScope()

	// Etiqueta de inicio
	g.Emit("%s:", startLabel)

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Volver al inicio (loop infinito)
	g.Emit("\tb %s", startLabel)

	// Etiqueta final (solo alcanzable con break)
	g.Emit("%s:", endLabel)

	return nil
}

func (g *ARM64Generator) VisitForRange(node *ast.ForRange) interface{} {
	// startLabel := g.newLabel("for_range_start")
	// endLabel := g.newLabel("for_range_end")
	// continueLabel := g.newLabel("for_range_continue")

	// g.Emit("\t// For range loop: %s, %s", node.Index, node.Value)

	// // Registrar contexto para break/continue
	// ctx := &ControlFlowContext{
	// 	Type:       "for",
	// 	StartLabel: continueLabel,
	// 	EndLabel:   endLabel,
	// 	BreakLabel: endLabel,
	// }
	// g.controlFlowStack = append(g.controlFlowStack, ctx)
	// defer g.popControlFlow()

	// // Crear scope para el loop
	// g.enterScope()
	// defer g.exitScope()

	// // Crear variables para index y value
	// indexVar := g.addVariable(node.Index, "int")
	// valueVar := g.addVariable(node.Value, "int")

	// // Evaluar la expresión de rango
	// g.Emit("\t// Evaluate range expression")
	// node.VisitForRange.Accept(g)

	// // TODO: Determinar tipo de rango (array, slice, rango numérico)
	// // Por ahora, asumimos un rango numérico simple

	// // Inicializar index a 0
	// g.Emit("\tmov x0, #0")
	// g.storeVariable(indexVar)

	// g.Emit("%s:", startLabel)

	// // Cargar index actual
	// g.loadVariable(indexVar)

	// // TODO: Verificar límites del rango
	// g.Emit("\t// TODO: Check range bounds")
	// g.Emit("\tcmp x0, #10") // Placeholder: límite de 10
	// g.Emit("\tbge %s", endLabel)

	// // Calcular valor actual (para rango numérico sería start + index)
	// g.Emit("\t// TODO: Calculate current value")
	// g.loadVariable(indexVar)
	// g.storeVariable(valueVar)

	// // Generar código del cuerpo
	// for _, stmt := range node.Body {
	// 	stmt.Accept(g)
	// }

	// g.Emit("%s:", continueLabel)

	// // Incrementar index
	// g.loadVariable(indexVar)
	// g.Emit("\tadd x0, x0, #1")
	// g.storeVariable(indexVar)

	// g.Emit("\tb %s", startLabel)
	// g.Emit("%s:", endLabel)

	// return nil

	g.AddError(fmt.Errorf("for-range loops not yet implemented"))
	return nil
}

func (g *ARM64Generator) VisitSwitchStmt(node *ast.SwitchStmt) interface{} {
	g.Emit("\t// Switch statement")

	// Evaluar la expresión del switch si existe
	if node.Expression != nil {
		node.Expression.Accept(g)
		g.Emit("\tmov x9, x0 // Guardar valor del switch en x9")
	} else {
		g.Emit("\tmov x9, xzr // Switch sin expresión")
	}

	endLabel := g.newLabel("switch_end")
	defaultLabel := g.newLabel("switch_default")
	caseLabels := make([]string, len(node.Cases))

	// Registrar contexto para manejar breaks
	g.pushControlFlow("switch", "", endLabel)
	defer g.popControlFlow()

	// Generar comparaciones para cada caso
	for i, clause := range node.Cases {
		caseLabels[i] = g.newLabel(fmt.Sprintf("case_%d", i))

		for _, valueExpr := range clause.Values {
			valueExpr.Accept(g)
			g.Emit("\tcmp x9, x0 // Comparar con valor del case")
			g.Emit("\tbeq %s // Saltar si igual", caseLabels[i])
		}
	}

	// Saltar al default o al final si no hay match
	if node.Default != nil {
		g.Emit("\tb %s", defaultLabel)
	} else {
		g.Emit("\tb %s", endLabel)
	}

	// Generar código para cada caso
	for i, clause := range node.Cases {
		g.Emit("%s:", caseLabels[i])

		for _, stmt := range clause.Statements {
			stmt.Accept(g)
		}

		// Saltar al final (a menos que haya fallthrough)
		g.Emit("\tb %s", endLabel)
	}

	// Generar código para el default si existe
	if node.Default != nil {
		g.Emit("%s:", defaultLabel)
		for _, stmt := range node.Default.Statements {
			stmt.Accept(g)
		}
	}

	g.Emit("%s:", endLabel)
	return nil
}
func (g *ARM64Generator) VisitCaseClause(node *ast.CaseClause) interface{} {
	g.Emit("\t// TODO: Case clause")
	return nil
}

func (g *ARM64Generator) VisitDefaultClause(node *ast.DefaultClause) interface{} {
	g.Emit("\t// TODO: Default clause")
	return nil
}

func (g *ARM64Generator) VisitFallthrough(node *ast.Fallthrough) interface{} {
	g.Emit("\t// TODO: Fallthrough")
	return nil
}

func (g *ARM64Generator) VisitPlusAssign(node *ast.PlusAssign) interface{} {
	g.Emit("\t// Plus assign (+=)")

	// Obtener el target (debe ser un identificador)
	if id, ok := node.Target.(*ast.Identifier); ok {
		varInfo := g.findVariable(id.Name)
		if varInfo == nil {
			g.AddError(fmt.Errorf("undefined variable: %s", id.Name))
			return nil
		}

		// Cargar la variable actual
		g.loadVariable(varInfo)
		leftReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", leftReg)

		// Evaluar la expresión del lado derecho
		node.Value.Accept(g)
		rightReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", rightReg)

		// Realizar la suma
		g.Emit("\tadd x0, %s, %s", leftReg, rightReg)

		// Almacenar el resultado
		g.storeVariable(varInfo)

		// Liberar registros
		g.freeRegister(leftReg)
		g.freeRegister(rightReg)
	} else {
		g.AddError(fmt.Errorf("compound assignment target must be an identifier"))
	}

	return nil
}

func (g *ARM64Generator) VisitMinusAssign(node *ast.MinusAssign) interface{} {
	g.Emit("\t// Minus assign (-=)")

	// Obtener el target (debe ser un identificador)
	if id, ok := node.Target.(*ast.Identifier); ok {
		varInfo := g.findVariable(id.Name)
		if varInfo == nil {
			g.AddError(fmt.Errorf("undefined variable: %s", id.Name))
			return nil
		}

		// Cargar la variable actual
		g.loadVariable(varInfo)
		leftReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", leftReg)

		// Evaluar la expresión del lado derecho
		node.Value.Accept(g)
		rightReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", rightReg)

		// Realizar la resta
		g.Emit("\tsub x0, %s, %s", leftReg, rightReg)

		// Almacenar el resultado
		g.storeVariable(varInfo)

		// Liberar registros
		g.freeRegister(leftReg)
		g.freeRegister(rightReg)
	} else {
		g.AddError(fmt.Errorf("compound assignment target must be an identifier"))
	}

	return nil
}

func (g *ARM64Generator) VisitMulAssign(node *ast.MulAssign) interface{} {
	g.Emit("\t// Multiply assign (*=)")

	// Obtener el target (debe ser un identificador)
	if id, ok := node.Target.(*ast.Identifier); ok {
		varInfo := g.findVariable(id.Name)
		if varInfo == nil {
			g.AddError(fmt.Errorf("undefined variable: %s", id.Name))
			return nil
		}

		// Cargar la variable actual
		g.loadVariable(varInfo)
		leftReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", leftReg)

		// Evaluar la expresión del lado derecho
		node.Value.Accept(g)
		rightReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", rightReg)

		// Realizar la multiplicación
		g.Emit("\tmul x0, %s, %s", leftReg, rightReg)

		// Almacenar el resultado
		g.storeVariable(varInfo)

		// Liberar registros
		g.freeRegister(leftReg)
		g.freeRegister(rightReg)
	} else {
		g.AddError(fmt.Errorf("compound assignment target must be an identifier"))
	}

	return nil
}

func (g *ARM64Generator) VisitDivAssign(node *ast.DivAssign) interface{} {
	g.Emit("\t// Divide assign (/=)")

	// Obtener el target (debe ser un identificador)
	if id, ok := node.Target.(*ast.Identifier); ok {
		varInfo := g.findVariable(id.Name)
		if varInfo == nil {
			g.AddError(fmt.Errorf("undefined variable: %s", id.Name))
			return nil
		}

		// Cargar la variable actual
		g.loadVariable(varInfo)
		leftReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", leftReg)

		// Evaluar la expresión del lado derecho
		node.Value.Accept(g)
		rightReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", rightReg)

		// Realizar la división
		g.Emit("\tsdiv x0, %s, %s", leftReg, rightReg)

		// Almacenar el resultado
		g.storeVariable(varInfo)

		// Liberar registros
		g.freeRegister(leftReg)
		g.freeRegister(rightReg)
	} else {
		g.AddError(fmt.Errorf("compound assignment target must be an identifier"))
	}

	return nil
}

func (g *ARM64Generator) VisitModAssign(node *ast.ModAssign) interface{} {
	g.Emit("\t// Modulo assign (%=)")

	// Obtener el target (debe ser un identificador)
	if id, ok := node.Target.(*ast.Identifier); ok {
		varInfo := g.findVariable(id.Name)
		if varInfo == nil {
			g.AddError(fmt.Errorf("undefined variable: %s", id.Name))
			return nil
		}

		// Cargar la variable actual
		g.loadVariable(varInfo)
		leftReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", leftReg)

		// Evaluar la expresión del lado derecho
		node.Value.Accept(g)
		rightReg := g.allocateRegister()
		g.Emit("\tmov %s, x0", rightReg)

		// Realizar el módulo (ARM64 no tiene instrucción directa)
		g.Emit("\tsdiv x2, %s, %s", leftReg, rightReg)
		g.Emit("\tmsub x0, x2, %s, %s", rightReg, leftReg)

		// Almacenar el resultado
		g.storeVariable(varInfo)

		// Liberar registros
		g.freeRegister(leftReg)
		g.freeRegister(rightReg)
	} else {
		g.AddError(fmt.Errorf("compound assignment target must be an identifier"))
	}

	return nil
}

// === HELPERS ===

// allocateRegister obtiene un registro temporal
func (g *ARM64Generator) allocateRegister() string {
	reg := g.registers.Allocate()
	return string(reg)
}

// freeRegister libera un registro
func (g *ARM64Generator) freeRegister(reg string) {
	g.registers.Free(Register(reg))
}

// escapeString escapa caracteres especiales en strings
func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

// switch:
func (g *ARM64Generator) pushControlFlow(ctxType, startLabel, endLabel string) {
	ctx := &ControlFlowContext{
		Type:       ctxType,
		StartLabel: startLabel,
		EndLabel:   endLabel,
		BreakLabel: endLabel,
	}
	g.controlFlowStack = append(g.controlFlowStack, ctx)
}

func (g *ARM64Generator) popControlFlow() {
	if len(g.controlFlowStack) > 0 {
		g.controlFlowStack = g.controlFlowStack[:len(g.controlFlowStack)-1]
	}
}

func (g *ARM64Generator) currentControlFlow() *ControlFlowContext {
	if len(g.controlFlowStack) == 0 {
		return nil
	}
	return g.controlFlowStack[len(g.controlFlowStack)-1]
}
