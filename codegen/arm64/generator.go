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

	// Campos para manejar bucles anidados
	currentLoopStart        string
	currentLoopEnd          string
	currentLoopContinue     string
	currentLoopBreakUsed    bool
	currentLoopContinueUsed bool           // etiqueta para continue
	currentVars             map[string]int // Mapa de nombres de variables a offsets en el stack
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
		currentVars:    make(map[string]int),
	}
}

// Generate es el punto de entrada principal
func (g *ARM64Generator) Generate(node ast.Node) (string, error) {
	g.Reset()
	fmt.Println("Iniciando generación de código ARM64...")
	// El nodo debería ser un *ast.Program
	program, ok := node.(*ast.Program)
	if !ok {
		return "", fmt.Errorf("expected *ast.Program, got %T", node)
	}

	// Visitar el programa para generar código
	program.Accept(g)
	fmt.Printf("Generación completada, errores: %v\n", g.HasErrors()) // Debug print
	// Si hay errores, retornarlos
	if g.HasErrors() {
		return "", g.GetErrors()
	}
	output := g.buildFinalOutput()
	fmt.Println("Código generado:\n", output) // Debug print del código generado
	// Construir el output final
	return g.buildFinalOutput(), nil
}

// buildFinalOutput construye el código final con secciones apropiadas
func (g *ARM64Generator) buildFinalOutput() string {
	var output strings.Builder

	// Sección de datos
	output.WriteString(".data\n")
	output.WriteString(".align 4\n")
	output.WriteString("print_fmt: .asciz \"%d\"\n")     // Sin \n
	output.WriteString("print_str_fmt: .asciz \"%s\"\n") // Sin \n
	output.WriteString("print_space_fmt: .asciz \"%c\"\n")

	// String literals
	for str, label := range g.stringLiterals {
		output.WriteString(fmt.Sprintf("%s:\n", label))
		output.WriteString(fmt.Sprintf("\t.asciz \"%s\"\n", escapeString(str)))
	}
	output.WriteString("\n")

	// Resto del código como estaba...
	output.WriteString(".text\n")
	output.WriteString(".align 4\n")

	if g.hasMainFunction() {
		output.WriteString(".global main\n")
	} else {
		output.WriteString(".global _start\n")
	}
	output.WriteString("\n")

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
	// CORRECCION: Usar registros callee-saved para evitar conflictos
	leftReg := "x19"  // Registro callee-saved
	rightReg := "x20" // Registro callee-saved

	// Evaluar operando izquierdo
	node.Left.Accept(g)
	g.Emit("\tmov %s, x0 // Guardar operando izquierdo", leftReg)

	// Evaluar operando derecho
	node.Right.Accept(g)
	g.Emit("\tmov %s, x0 // Guardar operando derecho", rightReg)

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
		g.Emit("\tsdiv x21, %s, %s", leftReg, rightReg)
		g.Emit("\tmsub x0, x21, %s, %s", rightReg, leftReg)
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
	if offset, exists := g.currentVars[node.Name]; exists {
		// CORRECCION: Cargar usando offset positivo
		g.Emit("\tldr x0, [x29, #%d] // Load %s", offset, node.Name)
	} else {
		g.AddError(fmt.Errorf("variable no declarada: %s", node.Name))
		g.Emit("\tmov x0, #0 // variable no encontrada")
	}
	return nil
}

// VisitPrintStmt genera código para print/println
func (g *ARM64Generator) VisitPrintStmt(node *ast.PrintStmt) interface{} {
	// Guardar registros que usaremos
	g.Emit("\tstp x19, x20, [sp, #-16]!")

	// Procesar todos los argumentos primero
	for i, arg := range node.Arguments {
		arg.Accept(g) // El valor a imprimir queda en x0

		// Determinar el formato adecuado
		switch arg.(type) {
		case *ast.Literal:
			if lit, ok := arg.(*ast.Literal); ok && lit.Type == "string" {
				g.Emit("\tadrp x19, print_str_fmt")
				g.Emit("\tadd x19, x19, :lo12:print_str_fmt")
			} else {
				g.Emit("\tadrp x19, print_fmt")
				g.Emit("\tadd x19, x19, :lo12:print_fmt")
			}
		default:
			g.Emit("\tadrp x19, print_fmt")
			g.Emit("\tadd x19, x19, :lo12:print_fmt")
		}

		g.Emit("\tmov x20, x0") // Guardar el valor
		g.Emit("\tmov x0, x19") // Formato
		g.Emit("\tmov x1, x20") // Valor

		// Para todos menos el último, usar espacio como separador
		if i < len(node.Arguments)-1 {
			g.Emit("\tmov x2, #' '") // Separador de espacio
			g.Emit("\tbl printf")
		} else {
			g.Emit("\tbl printf")
		}
	}

	// Solo agregar newline si es println (node.NewLine)
	if node.NewLine {
		g.Emit("\tmov x0, #10")
		g.Emit("\tbl putchar")
	}

	// Restaurar registros
	g.Emit("\tldp x19, x20, [sp], #16")
	return nil
}

// VisitVarDecl genera código para declaraciones de variables
func (g *ARM64Generator) VisitVarDecl(node *ast.VarDecl) interface{} {
	// El offset ya fue establecido en VisitFuncDecl
	offset, exists := g.currentVars[node.Name]
	if !exists {
		g.AddError(fmt.Errorf("offset no encontrado para variable %s", node.Name))
		return nil
	}

	// Generar código para el valor inicial
	if node.Value != nil {
		node.Value.Accept(g)
	} else {
		g.Emit("\tmov x0, #0")
	}

	// Almacenar en el stack
	g.Emit("\tstr x0, [x29, #%d] // Store %s", offset, node.Name)

	return nil
}

// VisitAssignment genera código para asignaciones
func (g *ARM64Generator) VisitAssignment(node *ast.Assignment) interface{} {
	// Evaluar el valor de la derecha
	node.Value.Accept(g)

	if id, ok := node.Target.(*ast.Identifier); ok {
		if offset, exists := g.currentVars[id.Name]; exists {
			// CORRECCION: Almacenar usando offset positivo
			g.Emit("\tstr x0, [x29, #%d] // Store to %s", offset, id.Name)
		} else {
			g.AddError(fmt.Errorf("variable no declarada: %s", id.Name))
		}
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
	g.currentVars = make(map[string]int)

	// Calcular espacio necesario para variables (8 bytes cada una)
	varCount := 0
	for _, stmt := range node.Body {
		if _, ok := stmt.(*ast.VarDecl); ok {
			varCount++
		}
	}

	// Stack layout:
	// - 16 bytes: x29, x30 (frame pointer, link register)
	// - 16 bytes: x19, x20 (callee-saved)
	// - 8 bytes: x21 (callee-saved)
	// - 8 bytes por variable
	// Total alineado a 16 bytes
	stackSize := ((16 + 16 + 8 + (varCount * 8)) + 15) &^ 15

	// Prologue
	g.Emit("%s:", node.Name)
	g.Emit("\tstp x29, x30, [sp, #-%d]!", stackSize)
	g.Emit("\tmov x29, sp")
	g.Emit("\tstp x19, x20, [sp, #16]")
	g.Emit("\tstr x21, [sp, #32]")

	// Asignar offsets a variables
	varOffset := 40 // Después de los registros guardados
	for _, stmt := range node.Body {
		if decl, ok := stmt.(*ast.VarDecl); ok {
			g.currentVars[decl.Name] = varOffset
			varOffset += 8
		}
	}

	// Generar cuerpo de la función
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Epilogue
	g.Emit("\tmov x0, #0") // Valor de retorno por defecto
	g.Emit("\tldr x21, [sp, #32]")
	g.Emit("\tldp x19, x20, [sp, #16]")
	g.Emit("\tldp x29, x30, [sp], #%d", stackSize)
	g.Emit("\tret")

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
	if g.currentLoopEnd == "" {
		g.AddError(fmt.Errorf("break fuera de bucle en línea %d", node.Line))
		return nil
	}
	g.Emit("\tb %s  // break", g.currentLoopEnd)
	g.currentLoopBreakUsed = true
	return nil
}

func (g *ARM64Generator) VisitContinue(node *ast.Continue) interface{} {
	if g.currentLoopContinue == "" {
		g.AddError(fmt.Errorf("continue fuera de bucle en línea %d", node.Line))
		return nil
	}
	g.Emit("\tb %s  // continue", g.currentLoopContinue)
	g.currentLoopContinueUsed = true
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

func (g *ARM64Generator) VisitForCondition(node *ast.ForCondition) interface{} {
	startLabel := g.newLabel("for_start")
	endLabel := g.newLabel("for_end")

	// Guardar contexto del bucle actual
	prevStart := g.currentLoopStart
	prevEnd := g.currentLoopEnd
	g.currentLoopStart = startLabel
	g.currentLoopEnd = endLabel

	g.Emit("%s:", startLabel)

	// Generar código para la condición
	node.Condition.Accept(g)

	// Saltar al final si la condición es falsa
	g.Emit("\tcbz x0, %s", endLabel)

	// Generar cuerpo del bucle
	for _, stmt := range node.Body {
		stmt.Accept(g)
	}

	// Volver al inicio
	g.Emit("\tb %s", startLabel)

	// Etiqueta de fin
	g.Emit("%s:", endLabel)

	// Restaurar contexto del bucle
	g.currentLoopStart = prevStart
	g.currentLoopEnd = prevEnd

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
	g.Emit("\t// TODO: Modulo assign")
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
