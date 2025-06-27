package arm64

import (
	"compiler/ast"
	"compiler/codegen"
	"fmt"
	"strings"
)

// ARM64Generator genera código ensamblador ARM64
type ARM64Generator struct {
	*codegen.BaseGenerator
	currentFunction string
	stackOffset     int
	labelCount      int
	registers       *RegisterAllocator
	stringLiterals  map[string]string // mapa de literales string a etiquetas
	dataSection     []string          // sección .data
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

// buildFinalOutput construye el código final con secciones apropiadas
func (g *ARM64Generator) buildFinalOutput() string {
	var output strings.Builder

	// Sección de datos (strings, constantes, etc.)
	if len(g.dataSection) > 0 || len(g.stringLiterals) > 0 {
		output.WriteString(".data\n")

		// Agregar string literals
		for str, label := range g.stringLiterals {
			output.WriteString(fmt.Sprintf("%s:\n", label))
			output.WriteString(fmt.Sprintf("\t.asciz \"%s\"\n", escapeString(str)))
		}

		// Agregar otras entradas de la sección de datos
		for _, data := range g.dataSection {
			output.WriteString(data + "\n")
		}

		output.WriteString("\n")
	}

	// Sección de texto (código)
	output.WriteString(".text\n")
	output.WriteString(".global _start\n\n")

	// Si no hay función main, crear un _start mínimo
	if !g.hasMainFunction() {
		output.WriteString("_start:\n")
		output.WriteString("\t// No main function found, executing top-level code\n")
	}

	// Agregar el código generado
	output.WriteString(g.GetOutput())

	// Si no terminamos con exit, agregarlo
	if !strings.Contains(g.GetOutput(), "svc #0x80") {
		output.WriteString("\n\t// Exit program\n")
		output.WriteString("\tmov x0, #0\n")
		output.WriteString("\tmov x16, #1\n")
		output.WriteString("\tsvc #0x80\n")
	}

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

// === IMPLEMENTACIÓN DEL VISITOR PATTERN ===

// VisitProgram genera código para el programa completo
func (g *ARM64Generator) VisitProgram(node *ast.Program) interface{} {
	// Primero, buscar y procesar funciones
	hasMain := false
	for _, stmt := range node.Statements {
		if funcDecl, ok := stmt.(*ast.FuncDecl); ok {
			if funcDecl.Name == "main" {
				hasMain = true
			}
			stmt.Accept(g)
		}
	}

	// Si no hay main, crear _start para código top-level
	if !hasMain {
		g.Emit("_start:")
		g.Emit("\t// Setup stack frame")
		g.Emit("\tstp x29, x30, [sp, #-16]!")
		g.Emit("\tmov x29, sp")
		g.Emit("")

		// Procesar statements que no son funciones
		for _, stmt := range node.Statements {
			if _, ok := stmt.(*ast.FuncDecl); !ok {
				stmt.Accept(g)
			}
		}
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
	// Por ahora, solo un placeholder
	g.Emit("\t// TODO: Load variable %s", node.Name)
	g.Emit("\tmov x0, #0 // placeholder")
	return nil
}

// VisitPrintStmt genera código para print/println
func (g *ARM64Generator) VisitPrintStmt(node *ast.PrintStmt) interface{} {
	g.Emit("\t// Print statement")

	for i, arg := range node.Arguments {
		// Evaluar el argumento
		arg.Accept(g)

		// Por ahora, solo imprimimos enteros
		// TODO: Manejar diferentes tipos
		g.Emit("\t// Print integer value")
		g.Emit("\tmov x1, x0")        // valor a imprimir
		g.Emit("\tadr x0, print_fmt") // formato (necesitaríamos agregarlo a .data)
		g.Emit("\tbl printf")         // llamar a printf

		// Agregar espacio entre argumentos (excepto el último)
		if i < len(node.Arguments)-1 {
			g.Emit("\t// Print space")
			g.Emit("\tmov x0, #32") // ASCII space
			g.Emit("\tbl putchar")
		}
	}

	// Agregar newline si es println
	if node.NewLine {
		g.Emit("\t// Print newline")
		g.Emit("\tmov x0, #10") // ASCII newline
		g.Emit("\tbl putchar")
	}

	return nil
}

// VisitVarDecl genera código para declaraciones de variables
func (g *ARM64Generator) VisitVarDecl(node *ast.VarDecl) interface{} {
	g.Emit("\t// Variable declaration: %s", node.Name)

	// Evaluar el valor inicial
	if node.Value != nil {
		node.Value.Accept(g)
	} else {
		// Valor por defecto
		g.Emit("\tmov x0, #0")
	}

	// TODO: Almacenar en el stack o en registros según el contexto
	g.Emit("\t// TODO: Store variable %s", node.Name)

	return nil
}

// VisitAssignment genera código para asignaciones
func (g *ARM64Generator) VisitAssignment(node *ast.Assignment) interface{} {
	g.Emit("\t// Assignment")

	// Evaluar el valor
	node.Value.Accept(g)

	// TODO: Almacenar en la variable objetivo
	if id, ok := node.Target.(*ast.Identifier); ok {
		g.Emit("\t// TODO: Store to variable %s", id.Name)
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

// VisitFuncDecl genera código para declaraciones de funciones
func (g *ARM64Generator) VisitFuncDecl(node *ast.FuncDecl) interface{} {
	g.currentFunction = node.Name
	g.stackOffset = 0

	// Si es main, también crear _start
	if node.Name == "main" {
		g.Emit("_start:")
		g.Emit("\tbl main")
		g.Emit("\t// Exit after main")
		g.Emit("\tmov x0, #0")
		g.Emit("\tmov x16, #1")
		g.Emit("\tsvc #0x80")
		g.Emit("")
	}

	// Etiqueta de la función
	g.Emit("%s:", node.Name)

	// Prólogo
	g.Emit("\t// Function prologue")
	g.Emit("\tstp x29, x30, [sp, #-16]!")
	g.Emit("\tmov x29, sp")

	// TODO: Configurar parámetros

	// Generar código del cuerpo
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Epílogo (si no hay return explícito)
	g.Emit("\t// Function epilogue")
	g.Emit("\tldp x29, x30, [sp], #16")
	g.Emit("\tret")
	g.Emit("")

	g.currentFunction = ""

	return nil
}

// VisitReturn genera código para return
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

	// Epílogo y retorno
	g.Emit("\tldp x29, x30, [sp], #16")
	g.Emit("\tret")

	return nil
}

// === MÉTODOS STUB PARA OTROS NODOS ===

func (g *ARM64Generator) VisitFuncCall(node *ast.FuncCall) interface{} {
	g.Emit("\t// Function call: %s", node.Name)

	// Evaluar argumentos
	for i, arg := range node.Arguments {
		arg.Accept(g)
		// Mover a registro de argumento apropiado
		if i < 8 {
			g.Emit("\tmov x%d, x0", i)
		} else {
			// Argumentos adicionales van en el stack
			g.Emit("\tstr x0, [sp, #%d]", (i-8)*8)
		}
	}

	// Llamar a la función
	g.Emit("\tbl %s", node.Name)

	return nil
}

func (g *ARM64Generator) VisitBreak(node *ast.Break) interface{} {
	g.Emit("\t// TODO: Break statement")
	return nil
}

func (g *ARM64Generator) VisitContinue(node *ast.Continue) interface{} {
	g.Emit("\t// TODO: Continue statement")
	return nil
}

func (g *ARM64Generator) VisitForStmt(node *ast.ForStmt) interface{} {
	g.Emit("\t// TODO: For statement")
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

// Implementar los nuevos tipos de for
func (g *ARM64Generator) VisitForCondition(node *ast.ForCondition) interface{} {
	g.Emit("\t// TODO: For with condition")
	return nil
}

func (g *ARM64Generator) VisitForClassic(node *ast.ForClassic) interface{} {
	g.Emit("\t// TODO: Classic for loop")
	return nil
}

func (g *ARM64Generator) VisitForIndexValue(node *ast.ForIndexValue) interface{} {
	g.Emit("\t// TODO: For with index,value")
	return nil
}

func (g *ARM64Generator) VisitForInfinite(node *ast.ForInfinite) interface{} {
	g.Emit("\t// TODO: Infinite for loop")
	return nil
}

func (g *ARM64Generator) VisitForRange(node *ast.ForRange) interface{} {
	g.Emit("\t// TODO: For range")
	return nil
}

func (g *ARM64Generator) VisitSwitchStmt(node *ast.SwitchStmt) interface{} {
	g.Emit("\t// TODO: Switch statement")
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

// Implementar compound assignments
func (g *ARM64Generator) VisitPlusAssign(node *ast.PlusAssign) interface{} {
	g.Emit("\t// TODO: Plus assign (+=)")
	return nil
}

func (g *ARM64Generator) VisitMinusAssign(node *ast.MinusAssign) interface{} {
	g.Emit("\t// TODO: Minus assign (-=)")
	return nil
}

func (g *ARM64Generator) VisitMulAssign(node *ast.MulAssign) interface{} {
	g.Emit("\t// TODO: Multiply assign (*=)")
	return nil
}

func (g *ARM64Generator) VisitDivAssign(node *ast.DivAssign) interface{} {
	g.Emit("\t// TODO: Divide assign (/=)")
	return nil
}

func (g *ARM64Generator) VisitModAssign(node *ast.ModAssign) interface{} {
	g.Emit("\t// TODO: Modulo assign (%=)")
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
