// visitor.go
package repl

import (
	compiler "compiler/parser"
	parser "compiler/parser"
	"compiler/value"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor personalizado para recorrer el árbol de sintaxis

// Constructor del visitor
type ReplVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	CallStack   *CallStack
	Console     *Console
	ErrorTable  *ErrorTable
	StructNames []string
}

// Recibe DclVisitor e inicializa el ReplVisitor
func NewVisitor(dclVisitor *DclVisitor) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  dclVisitor.ScopeTrace,
		ErrorTable:  dclVisitor.ErrorTable,
		StructNames: dclVisitor.StructNames,
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

// tambien podemos crear un ReplVisitor sin DclVisitor
// Esto es util para hacer tests de una vez
func NewReplVisitor(errorTable *ErrorTable) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  NewScopeTrace(),
		ErrorTable:  errorTable,
		StructNames: []string{},
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

// GetReplContext devuelve el contexto del REPL
func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ScopeTrace: v.ScopeTrace,
		CallStack:  v.CallStack,
		ErrorTable: v.ErrorTable,
	}
}

type BreakResult struct {
	Label string
}

type ContinueResult struct {
	Label string
}

type FallthroughResult struct{}

func (v *ReplVisitor) ValidType(_type string) bool {
	return v.ScopeTrace.GlobalScope.ValidType(_type)
}

func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		fmt.Println("Árbol nulo recibido.")
		return nil
	}

	fmt.Println("Enrutado------------------")
	fmt.Printf("DEBUG: Visitando tipo: %T\n", tree)
	// fmt.Printf("DEBUG: Texto del nodo: '%s'\n", tree.GetText())

	switch node := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(node.GetText())
	case *parser.ProgramaContext:
		return v.VisitPrograma(node)
	case *parser.PrintlnStmtContext:
		return v.VisitPrintlnStmt(node)
	case *parser.PrintStmtContext:
		return v.VisitPrintStmt(node)
	case *parser.ValorexpresionContext:
		return v.VisitValorexpresion(node)
	case *compiler.DeclAssignContext:
		return v.VisitValueDeclAssign(node)
	case *parser.ValorEnteroContext:
		return v.VisitValorEntero(node)
	case *parser.ValorFloatContext:
		return v.VisitValorFloat(node)
	case *parser.ValorDecimalContext:
		return v.VisitValorDecimal(node)
	case *parser.ValorCadenaContext:
		return v.VisitValorCadena(node)
	case *parser.ValorBooleanoContext:
		return v.VisitValorBooleano(node)
	case *parser.ValorCaracterContext:
		return v.VisitValorCaracter(node)
	case *parser.IdContext:
		return v.VisitId(node)
	case *parser.BinaryExpContext:
		return v.VisitBinaryExp(node)
	case *parser.StmtContext:
		return v.VisitStmt(node)
	case *parser.PlusAssignContext:
		return v.VisitPlusAssign(node)
	case *parser.MinusAssignContext:
		return v.VisitMinusAssign(node)
	case *parser.MulAssignContext:
		return v.VisitMulAssign(node)
	case *parser.DivAssignContext:
		return v.VisitDivAssign(node)
	case *parser.ModAssignContext:
		return v.VisitModAssign(node)
	case *parser.ForConditionContext:
		return v.VisitForCondition(node)
	case *parser.ForClassicContext:
		return v.VisitForClassic(node)
	case *parser.ForIndexValueContext:
		return v.VisitForIndexValue(node)
	case *parser.ForInfiniteContext:
		return v.VisitForInfinite(node)
	case *parser.ForRangeContext:
		return v.VisitForRange(node)
	case *parser.SwitchStmtContext:
		return v.VisitSwitchStmt(node)
	case *parser.SwitchCaseContext:
		return v.VisitSwitchCase(node)
	case *parser.DefaultCaseContext:
		return v.VisitDefaultCase(node)
	case *parser.FallthroughStmtContext:
		return v.VisitFallthroughStmt(node)
	default:
		fmt.Printf("Tipo inesperado en Visit(): %T\n", tree)
		return tree.Accept(v) // fallback por si acaso
	}

	return nil
}

// VisitPrograma es el metodo que se llama para visitar el nodo Programa
// Este nodo es el nodo raiz del arbol de sintaxis
// En este metodo recorremos todos los statements del programa
func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	if ctx == nil {
		fmt.Println("ProgramaContext nulo")
		return nil
	}

	statements := ctx.AllStmt()
	if len(statements) == 0 {
		fmt.Println("ℹ️ Programa vacío - no contiene statements")
		return nil
	}

	fmt.Printf("🔍 Iniciando visita del programa (%d statements)\n", len(statements))

	var lastResult any
	for i, stmt := range statements {
		fmt.Printf("\n📜 Statement %d/%d: %s\n", i+1, len(statements), stmt.GetText())

		result := v.Visit(stmt)
		if result != nil {
			lastResult = result
		}
		fmt.Printf("🔍 Resultado del statement %d: %v\n", i+1, result)
	}

	fmt.Println("🏁 Finalizada la ejecución del programa")
	return lastResult
}

/*

Ahora recorremos todos los stmts que tenemos en el programa

*/

func (v *ReplVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	// fmt.Println("Visitando statement:", ctx.GetText())

	if ctx == nil || ctx.GetChildCount() == 0 {
		fmt.Println("Stmt vacío o nulo.")
		return nil
	}

	node := ctx.GetChild(0)

	switch stmt := node.(type) {
	case *parser.PrintlnStmtContext:
		fmt.Println("Visitando nodo println")
		return v.VisitPrintlnStmt(stmt)

	case *parser.PrintStmtContext:
		fmt.Println("Visitando nodo print")
		return v.VisitPrintStmt(stmt)

	case *parser.DeclAssignContext:
		fmt.Println("Visitando nodo declAssign")
		return v.VisitValueDeclAssign(stmt)

	case *parser.DirectAssignContext:
		fmt.Println("Visitando nodo directAssign")
		return v.VisitDirectAssign(stmt)

	case *parser.PlusAssignContext:
		fmt.Println("Visitando nodo plusAssign (+=)")
		return v.VisitPlusAssign(stmt)

	case *parser.MinusAssignContext:
		fmt.Println("Visitando nodo minusAssign (-=)")
		return v.VisitMinusAssign(stmt)

	case *parser.MulAssignContext:
		fmt.Println("Visitando nodo mulAssign (*=)")
		return v.VisitMulAssign(stmt)

	case *parser.DivAssignContext:
		fmt.Println("Visitando nodo divAssign (/=)")
		return v.VisitDivAssign(stmt)

	case *parser.ModAssignContext:
		fmt.Println("Visitando nodo modAssign (%=)")
		return v.VisitModAssign(stmt)

	case *parser.IfStmtContext:
		fmt.Println("Visitando nodo ifStmt")
		return v.VisitIfStmt(stmt)

	case *parser.WhileStmtContext:
		fmt.Println("Visitando nodo whileStmt")
		return v.VisitWhileStmt(stmt)

	case *parser.ForStmtContext:
		fmt.Println("Visitando nodo forStmt")
		return v.VisitForStmt(stmt)

	case *parser.FuncCallContext:
		fmt.Println("Visitando nodo funcCall")
		return v.VisitFuncCall(stmt)

	case *parser.FuncDeclContext:
		fmt.Println("Visitando nodo funcDecl")
		return v.VisitFuncDecl(stmt)

	case *parser.StructDeclContext:
		fmt.Println("Visitando nodo structDecl")
		return v.VisitStructDecl(stmt)

	case *parser.StructInstanciaContext:
		fmt.Println("Visitando nodo structInstancia")
		return v.VisitStructInstancia(stmt)

	case *parser.ReturnStmtContext:
		fmt.Println("Visitando nodo returnStmt")
		return v.VisitReturnStmt(stmt)

	case *parser.BreakStmtContext:
		fmt.Println("Visitando nodo breakStmt")
		return v.VisitBreakStmt(stmt)

	case *parser.ContinueStmtContext:
		fmt.Println("Visitando nodo continueStmt")
		return v.VisitContinueStmt(stmt)

	case *parser.IncredecreContext:
		fmt.Println("Visitando nodo incredecre")
		// return v.VisitIncredecre(stmt)

	case *parser.ExpresionContext:
		fmt.Println("Visitando expresión como statement")
		// return v.VisitExpresion(stmt)
	case *parser.ForConditionContext:
		fmt.Println("Visitando nodo forCondition")
		return v.VisitForCondition(stmt)

	case *parser.ForClassicContext:
		fmt.Println("Visitando nodo forClassic")
		return v.VisitForClassic(stmt)

	case *parser.ForIndexValueContext:
		fmt.Println("Visitando nodo forIndexValue")
		return v.VisitForIndexValue(stmt)

	case *parser.ForInfiniteContext:
		fmt.Println("Visitando nodo forInfinite")
		return v.VisitForInfinite(stmt)

	case *parser.ForRangeContext:
		fmt.Println("Visitando nodo forRange")
		return v.VisitForRange(stmt)

	case *parser.SwitchStmtContext:
		fmt.Println("Visitando nodo switchStmt")
		return v.VisitSwitchStmt(stmt)

	case *parser.FallthroughStmtContext:
		fmt.Println("Visitando nodo fallthrough")
		return v.VisitFallthroughStmt(stmt)

	default:
		fmt.Printf("Tipo no reconocido dentro de stmt->: %T\n", node)
		if node != nil {
			// return node.Accept(v)
		}
		return nil
	}

	return nil

}

func (v *ReplVisitor) VisitDirectAssign(ctx *parser.DirectAssignContext) interface{} {
	// Vamos a visitar la asignacion
	// Primero vamos a ver si es una asignacion de variable
	// lo que queremos hacer para una asignacion
	// es ver si tenemos el ID
	varName := v.Visit(ctx.Id_pattern()).(string)
	varValue := v.Visit(ctx.Expresion()).(value.IVOR)
	// le pedimos la variable al ScopeTrace
	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		ok, msg := variable.Assign(varValue, true)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
	}

	return nil

}

/*

Con este metodo subimos el valor de la expresion

*/

func (v *ReplVisitor) VisitValorexpresion(ctx *parser.ValorexpresionContext) interface{} {
	// solo vamos a subir al hijo
	fmt.Println("🔍 Visitando Valorexpresion:", ctx.GetText())

	/*
		Todo esta bien, pero, tenemos que verificar el tipo de objeto que
		es el ctx y mandarlo al visit de nuevo para que suba pero
		con una forma de expresion con valores atomicos

		volvemos a subir el valor a la expresion

	*/
	valor := ctx.Valor()

	return v.Visit(valor)
}

// print normal
func (v *ReplVisitor) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	fmt.Println("🔍 Visitando PrintStmt:", ctx.GetText())

	// Total de expresiones
	exprs := ctx.AllExpresion()

	// Lista de resultados para imprimir
	resultados := []string{}

	for _, expr := range exprs {
		result := v.Visit(expr)
		if result == nil {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión vacía dentro de print")
			continue
		}

		val, ok := result.(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "La expresión no devuelve un valor válido")
			continue
		}

		// Convertir el valor a string y agregar a resultados
		text := fmt.Sprintf("%v", val.Value())
		resultados = append(resultados, text)
	}

	// Unir con espacios y enviar a la consola
	finalOutput := strings.Join(resultados, " ")
	v.Console.Println(finalOutput)

	return nil
}

func (v *ReplVisitor) VisitPrintlnStmt(ctx *parser.PrintlnStmtContext) interface{} {
	fmt.Println("🔍 Visitando PrintlnStmt:", ctx.GetText())

	expresiones := ctx.AllExpresion()
	if len(expresiones) == 0 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Se requiere al menos una expresión en println")
		return nil
	}

	var resultados []string

	for _, expr := range expresiones {
		result := v.Visit(expr)
		if result == nil {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión vacía dentro de Println")
			continue
		}

		val, ok := result.(value.IVOR)
		if !ok {
			fmt.Println("🔍 Resultado no es IVOR:", result)
			v.ErrorTable.NewSemanticError(expr.GetStart(), "La expresión no devuelve un valor válido")
			continue
		}

		valInterno := val.Value()
		valTipo := val.Type()
		fmt.Printf("Imprimiendo valor: %v (tipo: %s)\n", valInterno, valTipo)

		resultados = append(resultados, fmt.Sprintf("%v", valInterno))
	}

	if len(resultados) > 0 {
		v.Console.Println(strings.Join(resultados, " "))
	}

	return nil
}

func (v *ReplVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {

	runChain := true

	for _, ifStmt := range ctx.AllIf_chain() {

		runChain = !v.Visit(ifStmt).(bool)
		// Si la condicion del if es verdadera, ejecutamos el bloque
		if !runChain {
			break
		}
	}

	if runChain && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *ReplVisitor) VisitIfChain(ctx *parser.IfChainContext) interface{} {

	condition := v.Visit(ctx.Expresion()).(value.IVOR)

	if condition.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condicion del if debe ser un booleano")
		return false

	}

	if condition.(*value.BoolValue).InternalValue {

		// Push scope
		v.ScopeTrace.PushScope("if")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.ScopeTrace.PopScope()

		return true
	}

	return false
}

func (v *ReplVisitor) VisitElseStmt(ctx *parser.ElseStmtContext) interface{} {

	// Push scope {   }
	v.ScopeTrace.PushScope("else")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.ScopeTrace.PopScope()

	return nil
}

func (v *ReplVisitor) VisitValueDeclAssign(ctx *parser.DeclAssignContext) interface{} {

	/*
		Cuando declaramos una variable tenemos que

		2. Saber el nombre de la variable
		3. Saber el valor de la variable (Depende, si solo la declaramos)
		mut

	*/
	// verificamos si es una constante

	// isConst := isDeclConst(ctx.Var_type().GetText())
	fmt.Println("🔍 Visitando DeclAssign:", ctx.GetText())
	varName := ctx.ID().GetText()
	varValue := v.Visit(ctx.Expresion()).(value.IVOR)
	varType := varValue.Type()

	if varType == "[]" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede inferir el tipo de un vector vacio '"+varName+"'")
		return nil
	}

	// copy object
	// if obj, ok := varValue.(*ObjectValue); ok {
	// 	varValue = obj.Copy()
	// }

	// if IsVectorType(varValue.Type()) {
	// 	varValue = varValue.Copy()
	// }

	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, false, false, ctx.GetStart())
	/*

	   fn main() {
	   // entorno padre

	   x = 6
	   while (true){
	    x = 5
	   }
	   while (true){
	   x = 6
	   }
	   }

	*/
	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}
	return nil
}

/*


Devolvemos el valor de la expresion en este caso es el tope,
es un entero asi que solamente tomaremos lo que nos da el usuario
y lo convertiremos a un entero

*/

// el hijo de valorexpresioncontext -> valorEnter | valorBooleano |

func (v *ReplVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {

	fmt.Println("🔍 Visitando ValorEntero:", ctx.GetText())

	intVal, _ := strconv.Atoi(ctx.GetText())

	return &value.IntValue{
		InternalValue: intVal,
	}

}

// Manejar Floats
func (v *ReplVisitor) VisitValorFloat(ctx *parser.ValorFloatContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}
func (v *ReplVisitor) VisitIdPattern(ctx *parser.IdContext) interface{} {
	// hay que ir a buscar el Id a el entorno
	return ctx.GetText()
}

func (v *ReplVisitor) visitArrayExpresion(ctx *parser.ArrayexpreContext) interface{} {
	fmt.Println("🔍 Visitando arrayexpre:", ctx.GetText())

	var elementos []value.IVOR
	for _, exprCtx := range ctx.AllExpresion() {
		result := v.Visit(exprCtx)

		if result == nil {
			v.ErrorTable.NewSemanticError(exprCtx.GetStart(), "Elemento inválido en el arreglo")
			continue
		}

		val, ok := result.(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(exprCtx.GetStart(), "Expresión no válida en el arreglo")
			continue
		}

		elementos = append(elementos, val)
	}

	// Retornar una representación del array, por ejemplo un wrapper IVORArray
	return value.NewArray(elementos)
}

// Suma Re

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	// TODO: buscar variable en entorno
	fmt.Println("Accediendo a variable:", id)
	// debemos buscar la variable en el ScopeTrace
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	} else {
		// Si la variable existe, devolvemos su valor
		fmt.Println("Variable encontrada:", id, "con valor:", variable.Value)
		return variable
	}
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	return v.Visit(ctx.Incredecre())
}

func (v *ReplVisitor) VisitExpdotexp1(ctx *parser.Expdotexp1Context) interface{} {
	left := ctx.ID(0).GetText()
	right := ctx.ID(1).GetText()
	// TODO: acceder a campo estructurado: ID.ID
	fmt.Println("Accediendo a campo estructurado:", left, ".", right)
	return nil
}

func (v *ReplVisitor) VisitExpdotexp(ctx *parser.ExpdotexpContext) interface{} {
	id := ctx.ID().GetText()
	expr := v.Visit(ctx.Expresion())
	// TODO: ID.EXPR, posiblemente array o campo dinámico
	fmt.Println("Accediendo a campo dinámico:", id, "con expresión:", expr)
	return nil
}

func (v *ReplVisitor) VisitAsignacionfor(ctx *parser.AsignacionforContext) interface{} {
	id := ctx.ID().GetText()
	expr := v.Visit(ctx.Expresion())
	fmt.Println("Asignando valor a variable:", id, "con expresión:", expr)
	// TODO: asignar expr a id
	return nil
}

func (v *ReplVisitor) VisitValorDecimal(ctx *parser.ValorDecimalContext) interface{} {
	text := ctx.DECIMAL().GetText()
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		// Manejo de error
		// mandamos el error a ErrorTable
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión vacía dentro de paréntesis")
		log.Printf("Error al convertir '%s' a decimal: %v", text, err)
		return nil
	}
	return &value.FloatValue{
		InternalValue: val,
	}
}

func (v *ReplVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	stringVal = strings.ReplaceAll(stringVal, "\\\"", "\"")
	stringVal = strings.ReplaceAll(stringVal, "\\\\", "\\")
	stringVal = strings.ReplaceAll(stringVal, "\\n", "\n")
	stringVal = strings.ReplaceAll(stringVal, "\\r", "\r")

	// si es un caracter -> solamente un char
	if len(stringVal) == 1 {
		return &value.CharacterValue{
			InternalValue: stringVal,
		}
	}

	return &value.StringValue{
		InternalValue: stringVal,
	}
}

func (v *ReplVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
	text := ctx.BOOLEANO().GetText()
	// vemos si es true o false y devolvemos eso
	// si es true entonces devolvemos un valor booleano
	if text == "true" {
		return &value.BoolValue{
			InternalValue: true,
		}
	}
	if text == "false" {
		return &value.BoolValue{
			InternalValue: false,
		}
	}
	return nil
}

func (v *ReplVisitor) VisitValorCaracter(ctx *parser.ValorCaracterContext) interface{} {
	text := ctx.CARACTER().GetText()
	// Suponiendo que vienen entre comillas simples: 'a'
	if len(text) >= 3 {
		char := text[1 : len(text)-1]
		return &value.CharacterValue{
			InternalValue: char,
		}
	}
	return nil
}

/*
Ahora haremos la seccion de las expresiones binarias completa
en esta seccion haremos uso de un strategy, el cual
explicare en la clase del dia LUNES
*/
func (v *ReplVisitor) VisitBinaryExp(ctx *parser.BinaryExpContext) interface{} {

	op := ctx.GetOp().GetText()
	left := v.Visit(ctx.GetLeft()).(value.IVOR)
	// Verificamos si hay una estrategia de retorno anticipado para este operador
	earlyCheck, ok := EarlyReturnStrats[op]

	if ok {
		ok, _, result := earlyCheck.Validate(left)

		if ok {
			return result
		}
	}

	right := v.Visit(ctx.GetRight()).(value.IVOR)

	/*
			Aqui es donde sucede la magia, buscamos el operador
			y haciendo uso de un strategy pattern, validamos
			y ejecutamos la operacion de manera dinamica
		Si el operador no existe, lanzamos un error
		Si el operador existe, validamos los tipos de datos
	*/
	strat, ok := BinaryStrats[op]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(left, right)

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetOp(), msg)
		return value.DefaultNilValue
	}

	return result
}

//llamadas a funciones

func (v *ReplVisitor) VisitFuncCall(ctx *compiler.FuncCallContext) interface{} {

	// find if its a func or constructor of a struct

	canditateName := v.Visit(ctx.Id_pattern()).(string)
	funcObj, msg1 := v.ScopeTrace.GetFunction(canditateName)
	structObj, msg2 := v.ScopeTrace.GlobalScope.GetStruct(canditateName)

	if funcObj == nil && structObj == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg1+msg2)
		return value.DefaultNilValue
	}

	args := make([]*Argument, 0)
	if ctx.Parametros() != nil {
		args = v.Visit(ctx.Parametros()).([]*Argument)
	}

	// struct has priority over func
	if structObj != nil {

		switch funcObj := funcObj.(type) {
		case *BuiltInFunction:
			returnValue, ok, msg := funcObj.Exec(v.GetReplContext(), args)

			if !ok {

				if msg != "" {
					v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
				}

				return value.DefaultNilValue

			}

			return returnValue

		case *Function:
			funcObj.Exec(v, args, ctx.GetStart())
			return funcObj.ReturnValue

		// case *ObjectBuiltInFunction:
		// 	funcObj.Exec(v, args, ctx.GetStart())
		// 	return funcObj.ReturnValue

		default:
			log.Fatal("Function type not found")
		}

		return value.DefaultNilValue
	}
	return nil
}

// VisitPlusAssign maneja la asignación de suma
func (v *ReplVisitor) VisitPlusAssign(ctx *parser.PlusAssignContext) interface{} {
	id := ctx.Id_pattern().GetText()
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	}

	// Obtenemos el valor actual de la variable
	currentValue := variable.Value.Copy()

	// Obtenemos el valor a sumar
	valueToAdd := v.Visit(ctx.Expresion()).(value.IVOR)
	if valueToAdd == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida para suma")
		return nil
	}

	// Verificamos si el tipo soporta la operación Add usando type assertion
	var newValue value.IVOR
	var ok bool
	var msg string

	switch cv := currentValue.(type) {
	case *value.IntValue:
		newValue, ok, msg = cv.Add(valueToAdd)
	case *value.FloatValue:
		newValue, ok, msg = cv.Add(valueToAdd)
	case *value.StringValue:
		newValue, ok, msg = cv.Add(valueToAdd)
	case *value.CharacterValue:
		newValue, ok, msg = cv.Add(valueToAdd)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación '+=' no soportada para "+currentValue.Type())
		return nil
	}

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	// Asignamos el nuevo valor a la variable
	variable.Assign(newValue, true)
	return nil
}

// VisitMinusAssign maneja la asignación de resta
func (v *ReplVisitor) VisitMinusAssign(ctx *parser.MinusAssignContext) interface{} {
	id := ctx.Id_pattern().GetText()
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	}

	// Obtenemos el valor actual de la variable
	currentValue := variable.Value.Copy()

	// Obtenemos el valor a restar
	valueToSubtract := v.Visit(ctx.Expresion()).(value.IVOR)
	if valueToSubtract == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida para resta")
		return nil
	}

	// Verificamos si el tipo soporta la operación Subtract usando type assertion
	var newValue value.IVOR
	var ok bool
	var msg string

	switch cv := currentValue.(type) {
	case *value.IntValue:
		newValue, ok, msg = cv.Subtract(valueToSubtract)
	case *value.FloatValue:
		newValue, ok, msg = cv.Subtract(valueToSubtract)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación '-=' no soportada para "+currentValue.Type())
		return nil
	}

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	variable.Assign(newValue, true)
	return nil
}

func (v *ReplVisitor) VisitMulAssign(ctx *parser.MulAssignContext) interface{} {
	id := ctx.Id_pattern().GetText()
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	}

	currentValue := variable.Value.Copy()

	valueToMultiply := v.Visit(ctx.Expresion()).(value.IVOR)
	if valueToMultiply == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida para multiplicación")
		return nil
	}

	var newValue value.IVOR
	var ok bool
	var msg string

	switch cv := currentValue.(type) {
	case *value.IntValue:
		newValue, ok, msg = cv.Multiply(valueToMultiply)
	case *value.FloatValue:
		newValue, ok, msg = cv.Multiply(valueToMultiply)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación '*=' no soportada para "+currentValue.Type())
		return nil
	}

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	variable.Assign(newValue, true)
	return nil
}

func (v *ReplVisitor) VisitDivAssign(ctx *parser.DivAssignContext) interface{} {
	id := ctx.Id_pattern().GetText()
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	}

	currentValue := variable.Value.Copy()

	valueToDivide := v.Visit(ctx.Expresion()).(value.IVOR)
	if valueToDivide == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida para división")
		return nil
	}

	var newValue value.IVOR
	var ok bool
	var msg string

	switch cv := currentValue.(type) {
	case *value.IntValue:
		newValue, ok, msg = cv.Divide(valueToDivide)
	case *value.FloatValue:
		newValue, ok, msg = cv.Divide(valueToDivide)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación '/=' no soportada para "+currentValue.Type())
		return nil
	}

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	variable.Assign(newValue, true)
	return nil
}

func (v *ReplVisitor) VisitModAssign(ctx *parser.ModAssignContext) interface{} {
	id := ctx.Id_pattern().GetText()
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	}

	currentValue := variable.Value.Copy()

	valueToMod := v.Visit(ctx.Expresion()).(value.IVOR)
	if valueToMod == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida para módulo")
		return nil
	}

	var newValue value.IVOR
	var ok bool
	var msg string

	switch cv := currentValue.(type) {
	case *value.IntValue:
		newValue, ok, msg = cv.Mod(valueToMod)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación '%=' no soportada para "+currentValue.Type())
		return nil
	}

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	variable.Assign(newValue, true)
	return nil
}

func (v *ReplVisitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	for {
		condition := v.Visit(ctx.Expresion()).(value.IVOR)
		if condition.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del while debe ser booleana")
			return nil
		}

		if !condition.(*value.BoolValue).InternalValue {
			break
		}

		v.ScopeTrace.PushScope("while")

		shouldBreak := false
		shouldContinue := false

		for _, stmt := range ctx.AllStmt() {
			result := v.Visit(stmt)

			if br, ok := result.(*BreakResult); ok {
				shouldBreak = true
				if br.Label != "" {
					v.ErrorTable.NewSemanticError(ctx.GetStart(), "Break con label no implementado")
				}
				break
			}

			if cr, ok := result.(*ContinueResult); ok {
				shouldContinue = true
				if cr.Label != "" {
					v.ErrorTable.NewSemanticError(ctx.GetStart(), "Continue con label no implementado")
				}
				break
			}
		}

		v.ScopeTrace.PopScope()

		if shouldBreak {
			break
		}

		if shouldContinue {
			continue
		}
	}

	return nil
}

func (v *ReplVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	v.ScopeTrace.PushScope("switch")
	defer v.ScopeTrace.PopScope()

	var switchValue value.IVOR
	if ctx.Expresion() != nil {
		switchValue = v.Visit(ctx.Expresion()).(value.IVOR)
	}

	shouldFallthrough := false

	for _, caseCtx := range ctx.AllSwitch_case() {
		matched := false

		if !shouldFallthrough {
			caseValuesCtx := caseCtx.(*parser.SwitchCaseContext).Case_values()
			caseValues := v.Visit(caseValuesCtx).([]value.IVOR)
			for _, caseVal := range caseValues {
				if switchValue != nil {
					strat, ok := BinaryStrats["=="]
					if ok {
						ok, _, result := strat.Validate(switchValue, caseVal)
						if ok {
							if boolResult, isBool := result.(*value.BoolValue); isBool && boolResult.InternalValue {
								matched = true
								break
							}
						}
					}
				} else {
					if boolVal, ok := caseVal.(*value.BoolValue); ok && boolVal.InternalValue {
						matched = true
						break
					}
				}
			}
		}

		_ = matched

		// if matched || shouldFallthrough {
		// 	for _, stmt := range caseCtx.AllStmt() {
		// 		result := v.Visit(stmt)
		// 		if _, isFallthrough := result.(*FallthroughResult); isFallthrough {
		// 			shouldFallthrough = true
		// 			break
		// 		}
		// 	}

		// 	if !shouldFallthrough {
		// 		break
		// 	}
		// }

	}

	if !shouldFallthrough && ctx.Default_case() != nil {
		v.Visit(ctx.Default_case())
	}

	return nil
}

func (v *ReplVisitor) VisitCaseValueList(ctx *parser.CaseValueListContext) interface{} {
	var values []value.IVOR
	for _, expr := range ctx.AllExpresion() {
		val := v.Visit(expr).(value.IVOR)
		values = append(values, val)
	}
	return values
}

func (v *ReplVisitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *ReplVisitor) VisitFallthroughStmt(ctx *parser.FallthroughStmtContext) interface{} {
	return &FallthroughResult{}
}

func (v *ReplVisitor) VisitForCondition(ctx *parser.ForConditionContext) interface{} {

	for {
		condition := v.Visit(ctx.Expresion()).(value.IVOR)
		if condition.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser booleana")
			return nil
		}

		if !condition.(*value.BoolValue).InternalValue {
			break
		}

		v.ScopeTrace.PushScope("for")
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}
		v.ScopeTrace.PopScope()
	}
	return nil
}

func (v *ReplVisitor) VisitForClassic(ctx *parser.ForClassicContext) interface{} {
	v.ScopeTrace.PushScope("for")
	defer v.ScopeTrace.PopScope()

	if ctx.For_init() != nil {
		v.Visit(ctx.For_init())
	}

	for {
		if ctx.Expresion() != nil {
			condition := v.Visit(ctx.Expresion()).(value.IVOR)
			if condition.Type() != value.IVOR_BOOL {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser booleana")
				return nil
			}
			if !condition.(*value.BoolValue).InternalValue {
				break
			}
		}

		v.ScopeTrace.PushScope("for-body")
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}
		v.ScopeTrace.PopScope()

		if ctx.For_update() != nil {
			v.Visit(ctx.For_update())
		}
	}

	return nil
}

func (v *ReplVisitor) VisitForIndexValue(ctx *parser.ForIndexValueContext) interface{} {
	// indexName := ctx.ID(0).GetText()
	// valueName := ctx.ID(1).GetText()

	// collection := v.Visit(ctx.Expresion()).(value.IVOR)

	// if !IsArrayType(collection.Type()) {
	// 	v.ErrorTable.NewSemanticError(ctx.GetStart(), "For index,value requiere un array")
	// 	return nil
	// }

	// array := collection.(*value.ArrayValue)
	// for i, elem := range array.Elements {
	// 	v.ScopeTrace.PushScope("for")

	// 	// Agregar índice y valor al scope
	// 	v.ScopeTrace.AddVariable(indexName, value.IVOR_INT, &value.IntValue{InternalValue: i}, false, false, ctx.GetStart())
	// 	v.ScopeTrace.AddVariable(valueName, elem.Type(), elem, false, false, ctx.GetStart())

	// 	// Ejecutar statements
	// 	for _, stmt := range ctx.AllStmt() {
	// 		v.Visit(stmt)
	// 	}

	// 	v.ScopeTrace.PopScope()
	// }

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "For index,value no implementado")

	return nil
}

func (v *ReplVisitor) VisitForInfinite(ctx *parser.ForInfiniteContext) interface{} {
	for {
		v.ScopeTrace.PushScope("for")

		shouldBreak := false
		for _, stmt := range ctx.AllStmt() {
			result := v.Visit(stmt)
			// Verificar si hay break
			if _, isBreak := result.(*BreakResult); isBreak {
				shouldBreak = true
				break
			}
		}

		v.ScopeTrace.PopScope()

		if shouldBreak {
			break
		}
	}
	return nil
}

func (v *ReplVisitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	indexName := ctx.ID(0).GetText()
	valueName := ctx.ID(1).GetText()

	collection := v.Visit(ctx.Expresion()).(value.IVOR)

	if !IsArrayType(collection.Type()) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "For range requiere un array")
		return nil
	}

	array := collection.(*value.IVORArray)
	for i, elem := range array.Elements {
		v.ScopeTrace.PushScope("for")

		v.ScopeTrace.AddVariable(indexName, value.IVOR_INT, &value.IntValue{InternalValue: i}, false, false, ctx.GetStart())
		v.ScopeTrace.AddVariable(valueName, elem.Type(), elem, false, false, ctx.GetStart())

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		v.ScopeTrace.PopScope()
	}

	return nil
}

func (v *ReplVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	label := ""
	if ctx.ID() != nil {
		label = ctx.ID().GetText()
	}
	return &BreakResult{Label: label}
}

func (v *ReplVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	label := ""
	if ctx.ID() != nil {
		label = ctx.ID().GetText()
	}
	return &ContinueResult{Label: label}
}

func IsArrayType(t string) bool {
	return strings.HasPrefix(t, "[]")
}

func (v *ReplVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	varName := ctx.ID().GetText()

	var start, end int
	var collection value.IVOR

	if ctx.Range_() != nil {
		rangeCtx := ctx.Range_().(*parser.NumericRangeContext)
		startVal := v.Visit(rangeCtx.Expresion(0)).(value.IVOR)
		endVal := v.Visit(rangeCtx.Expresion(1)).(value.IVOR)

		if startVal.Type() != value.IVOR_INT || endVal.Type() != value.IVOR_INT {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Range debe ser entre enteros")
			return nil
		}

		start = startVal.(*value.IntValue).InternalValue
		end = endVal.(*value.IntValue).InternalValue
	} else if ctx.Expresion() != nil {
		collection = v.Visit(ctx.Expresion()).(value.IVOR)
		if !IsArrayType(collection.Type()) {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "For requiere un array o range")
			return nil
		}
	}

	v.ScopeTrace.PushScope("for")
	defer v.ScopeTrace.PopScope()

	if collection != nil {

		array := collection.(*value.IVORArray)
		for i, elem := range array.Elements {
			v.ScopeTrace.AddVariable(varName, elem.Type(), elem, false, false, ctx.GetStart())
			_ = i // Usar el índice si es necesario

			for _, stmt := range ctx.AllStmt() {
				result := v.Visit(stmt)
				if _, isBreak := result.(*BreakResult); isBreak {
					return nil
				}
				if _, isContinue := result.(*ContinueResult); isContinue {
					break
				}
			}
		}
	} else {
		for i := start; i <= end; i++ {
			v.ScopeTrace.AddVariable(varName, value.IVOR_INT, &value.IntValue{InternalValue: i}, false, false, ctx.GetStart())

			for _, stmt := range ctx.AllStmt() {
				result := v.Visit(stmt)
				if _, isBreak := result.(*BreakResult); isBreak {
					return nil
				}
				if _, isContinue := result.(*ContinueResult); isContinue {
					break
				}
			}
		}
	}

	return nil
}
