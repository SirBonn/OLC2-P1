package ast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Value representa un valor en el intérprete
type Value struct {
	Value interface{}
	Type  string
}

type BuiltinFunc func(args []Value) (Value, error)

// Environment representa el entorno de ejecución
type Environment struct {
	parent    *Environment
	variables map[string]Value
	functions map[string]*FuncDecl
	builtins  map[string]BuiltinFunc
}

// NewEnvironment crea un nuevo entorno
func NewEnvironment(parent *Environment) *Environment {
	env := &Environment{
		parent:    parent,
		variables: make(map[string]Value),
		functions: make(map[string]*FuncDecl),
		builtins:  make(map[string]BuiltinFunc),
	}

	// Solo registrar builtins en el ambiente global
	if parent == nil {
		env.registerBuiltins()
	}

	return env
}

// Set establece una variable en el entorno
func (env *Environment) Set(name string, value Value) {
	env.variables[name] = value
}

// Get obtiene una variable del entorno
func (env *Environment) Get(name string) (Value, bool) {
	if value, exists := env.variables[name]; exists {
		return value, true
	}
	if env.parent != nil {
		return env.parent.Get(name)
	}
	return Value{}, false
}

// SetFunction establece una función en el entorno
func (env *Environment) SetFunction(name string, fn *FuncDecl) {
	env.functions[name] = fn
}

// GetFunction obtiene una función del entorno
func (env *Environment) GetFunction(name string) (*FuncDecl, bool) {
	if fn, exists := env.functions[name]; exists {
		return fn, true
	}
	if env.parent != nil {
		return env.parent.GetFunction(name)
	}
	return nil, false
}

// Interpreter es el intérprete del AST
type Interpreter struct {
	env            *Environment
	output         strings.Builder
	shouldExit     bool
	returnVal      Value
	shouldBreak    bool
	shouldContinue bool
}

// NewInterpreter crea un nuevo intérprete
func NewInterpreter() *Interpreter {
	return &Interpreter{
		env: NewEnvironment(nil),
	}
}

// Interpret ejecuta un programa AST
func (i *Interpreter) Interpret(program *Program) (string, error) {
	fmt.Printf("INICIANDO Interpret\n")
	fmt.Printf("Programa recibido: %+v\n", program)

	if program == nil {
		fmt.Printf("ERROR: program es nil\n")
		return "", fmt.Errorf("program is nil")
	}

	fmt.Printf("Programa tiene %d statements\n", len(program.Statements))

	// Mostrar TODOS los statements antes de procesarlos
	for idx, stmt := range program.Statements {
		fmt.Printf("Statement [%d]: %T", idx, stmt)
		if funcDecl, ok := stmt.(*FuncDecl); ok {
			fmt.Printf(" (función: %s)", funcDecl.Name)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("FASE 1: Registrando funciones...\n")

	// Primero, registrar todas las funciones
	functionCount := 0
	for idx, stmt := range program.Statements {
		if funcDecl, ok := stmt.(*FuncDecl); ok {
			fmt.Printf("Registrando función [%d]: %s con %d parámetros\n",
				idx, funcDecl.Name, len(funcDecl.Parameters))
			i.env.SetFunction(funcDecl.Name, funcDecl)
			functionCount++
		}
	}

	fmt.Printf("FASE 1 COMPLETADA: %d funciones registradas\n", functionCount)

	// Listar todas las funciones registradas
	fmt.Printf("Funciones disponibles en el entorno:\n")
	for name, _ := range i.env.functions {
		fmt.Printf("   - %s\n", name)
	}
	fmt.Printf("FASE 2: Buscando función main...\n")
	if mainFunc, exists := i.env.GetFunction("main"); exists {
		fmt.Printf("✅ Función main encontrada, ejecutando...\n")

		// Crear nuevo entorno para la función main
		mainEnv := NewEnvironment(i.env)
		oldEnv := i.env
		i.env = mainEnv
		defer func() { i.env = oldEnv }()

		// Ejecutar cuerpo de la función main
		for _, stmt := range mainFunc.Body {
			err := i.executeStatement(stmt)
			if err != nil {
				return i.output.String(), err
			}
			if i.shouldExit {
				break
			}
		}
	} else {
		fmt.Printf("⚠️ No se encontró función main\n")
	}
	fmt.Printf("FASE 3: Ejecutando statements no-función...\n")

	// Luego ejecutar los statements que no son declaraciones de función
	statementCount := 0
	for idx, stmt := range program.Statements {
		if _, ok := stmt.(*FuncDecl); ok {
			fmt.Printf("⏭Saltando declaración de función [%d]\n", idx)
			continue // Ya procesamos las funciones
		}

		fmt.Printf("Ejecutando statement [%d]: %T\n", idx, stmt)

		err := i.executeStatement(stmt)
		if err != nil {
			fmt.Printf("Error en statement [%d]: %v\n", idx, err)
			return i.output.String(), err
		}

		statementCount++
		fmt.Printf("Statement [%d] ejecutado exitosamente\n", idx)
		fmt.Printf("Output acumulado: '%s'\n", i.output.String())

		if i.shouldExit {
			fmt.Printf("shouldExit=true, terminando en statement [%d]\n", idx)
			break
		}
	}

	fmt.Printf("RESUMEN FINAL:\n")
	fmt.Printf("   - Total statements: %d\n", len(program.Statements))
	fmt.Printf("   - Funciones registradas: %d\n", functionCount)
	fmt.Printf("   - Statements ejecutados: %d\n", statementCount)
	fmt.Printf("   - Output final: '%s'\n", i.output.String())
	fmt.Printf("INTERPRET COMPLETADO\n")

	return i.output.String(), nil
}

// executeStatement ejecuta un statement
func (i *Interpreter) executeStatement(stmt Statement) error {
	fmt.Printf("executeStatement llamado con: %T\n", stmt)

	if i.shouldExit || i.shouldBreak || i.shouldContinue {
		fmt.Printf("Terminando early: exit=%v, break=%v, continue=%v\n",
			i.shouldExit, i.shouldBreak, i.shouldContinue)
		return nil
	}

	switch s := stmt.(type) {
	case *PrintStmt:
		fmt.Printf("Ejecutando PrintStmt con %d argumentos\n", len(s.Arguments))
		return i.executePrintStmt(s)
	case *VarDecl:
		fmt.Printf("Ejecutando VarDecl: %s\n", s.Name)
		return i.executeVarDecl(s)
	case *Assignment:
		fmt.Printf("Ejecutando Assignment\n")
		return i.executeAssignment(s)
	case *PlusAssign:
		fmt.Printf("Ejecutando PlusAssign\n")
		return i.executePlusAssign(s)
	case *MinusAssign:
		fmt.Printf("Ejecutando MinusAssign\n")
		return i.executeMinusAssign(s)
	case *MulAssign:
		fmt.Printf("Ejecutando MulAssign\n")
		return i.executeMulAssign(s)
	case *DivAssign:
		fmt.Printf("Ejecutando DivAssign\n")
		return i.executeDivAssign(s)
	case *ModAssign:
		fmt.Printf("Ejecutando ModAssign\n")
		return i.executeModAssign(s)
	case *IfStmt:
		fmt.Printf("Ejecutando IfStmt con condición: %T\n", s.Condition)
		return i.executeIfStmt(s)
	case *WhileStmt:
		fmt.Printf("Ejecutando WhileStmt con condición: %T\n", s.Condition)
		return i.executeWhileStmt(s)
	case *ForStmt:
		fmt.Printf("Ejecutando ForStmt con iterable: %T\n", s.Iterable)
		return i.executeForStmt(s)
	case *FuncDecl:
		// Las funciones ya se registraron al inicio del programa
		fmt.Printf("Función declarada: %s (ya procesada)\n", s.Name)
		return nil
	case *Return:
		fmt.Printf("Ejecutando Return con valor: %T\n", s.Value)
		return i.executeReturn(s)
	case *ForCondition:
		fmt.Printf("Ejecutando ForCondition\n")
		return i.executeForCondition(s)
	case *ForClassic:
		fmt.Printf("Ejecutando ForClassic\n")
		return i.executeForClassic(s)
	case *ForIndexValue:
		fmt.Printf("Ejecutando ForIndexValue\n")
		return i.executeForIndexValue(s)
	case *ForInfinite:
		fmt.Printf("Ejecutando ForInfinite\n")
		return i.executeForInfinite(s)
	case *ForRange:
		fmt.Printf("Ejecutando ForRange\n")
		return i.executeForRange(s)
	case *SwitchStmt:
		fmt.Printf("Ejecutando SwitchStmt\n")
		return i.executeSwitchStmt(s)
	case *Fallthrough:
		fmt.Printf("Ejecutando Fallthrough\n")
		return i.executeFallthrough(s)
	case *Break:
		fmt.Println("Ejecutando Break")
		i.shouldBreak = true
		return nil
	case *Continue:
		fmt.Println("Ejecutando Continue")
		i.shouldContinue = true
		return nil
	case *ExpressionStatement: // NUEVO CASO AGREGADO
		fmt.Printf("Ejecutando ExpressionStatement con expresión: %T\n", s.Expression)
		// Las llamadas a funciones llegarán aquí como ExpressionStatement
		_, err := i.evaluateExpression(s.Expression)
		return err
	default:
		fmt.Printf("Unhandled statement type: %T\n", s)
		return fmt.Errorf("unhandled statement type: %T", s)
	}
}

func (i *Interpreter) processStringInterpolation(str string) string {
	// Por ahora, retornar el string tal cual
	// TODO: Implementar interpolación real si es necesario
	return str
}

func (i *Interpreter) executeStructDecl(stmt *StructDecl) error {
	// Por ahora, solo registrar que se declaró un struct
	fmt.Printf("Struct declarado: %s\n", stmt.Name)
	// TODO: Implementar registro de tipos de struct si es necesario
	return nil
}

// executePrintStmt ejecuta un statement de print
func (i *Interpreter) executePrintStmt(stmt *PrintStmt) error {
	for j, arg := range stmt.Arguments {
		if j > 0 {
			i.output.WriteString(" ")
		}

		value, err := i.evaluateExpression(arg)
		if err != nil {
			return err
		}

		i.output.WriteString(i.valueToString(value))
	}

	if stmt.NewLine {
		i.output.WriteString("\n")
	}

	return nil
}

// executeVarDecl ejecuta una declaración de variable
func (i *Interpreter) executeVarDecl(stmt *VarDecl) error {
	value, err := i.evaluateExpression(stmt.Value)
	if err != nil {
		return err
	}

	i.env.Set(stmt.Name, value)
	return nil
}

// executeAssignment ejecuta una asignación
func (i *Interpreter) executeAssignment(stmt *Assignment) error {
	value, err := i.evaluateExpression(stmt.Value)
	if err != nil {
		return err
	}

	// Asumimos que Target es un Identifier por simplicidad
	if id, ok := stmt.Target.(*Identifier); ok {
		i.env.Set(id.Name, value)
		return nil
	}

	return fmt.Errorf("unsupported assignment target: %T", stmt.Target)
}

// executeIfStmt ejecuta un statement if
func (i *Interpreter) executeIfStmt(stmt *IfStmt) error {
	fmt.Printf("Evaluando condición del if en línea %d\n", stmt.Line)
	condition, err := i.evaluateExpression(stmt.Condition)
	if err != nil {
		return err
	}
	if i.isTruthy(condition) {
		for _, s := range stmt.ThenBranch {
			if err := i.executeStatement(s); err != nil {
				return err
			}
		}
	} else if stmt.ElseIf != nil {
		fmt.Println("Ejecutando ELSE IF...")
		return i.executeIfStmt(stmt.ElseIf)
	} else if len(stmt.ElseBranch) > 0 {
		for _, s := range stmt.ElseBranch {
			if err := i.executeStatement(s); err != nil {
				return err
			}
		}
	}
	return nil
}

// executeWhileStmt ejecuta un loop while
func (i *Interpreter) executeWhileStmt(stmt *WhileStmt) error {
	for {
		condition, err := i.evaluateExpression(stmt.Condition)
		if err != nil {
			return err
		}

		if !i.isTruthy(condition) {
			break
		}

		i.shouldBreak = false
		i.shouldContinue = false

		for _, s := range stmt.Body {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}

			if i.shouldBreak {
				i.shouldBreak = false
				return nil
			}

			if i.shouldContinue {
				i.shouldContinue = false
				break
			}

			if i.shouldExit {
				return nil
			}
		}
	}

	return nil
}

// executeForStmt ejecuta un loop for
func (i *Interpreter) executeForStmt(stmt *ForStmt) error {
	// Evaluar el iterable
	iterable, err := i.evaluateExpression(stmt.Iterable)
	if err != nil {
		return err
	}

	// Crear nuevo entorno para el scope del for
	forEnv := NewEnvironment(i.env)
	oldEnv := i.env
	i.env = forEnv
	defer func() { i.env = oldEnv }()

	// Iterar sobre el valor (por simplicidad, solo manejamos arrays/slices)
	switch val := iterable.Value.(type) {
	case []interface{}:
		for _, item := range val {
			i.env.Set(stmt.Variable, Value{Value: item, Type: "auto"})

			i.shouldBreak = false
			i.shouldContinue = false

			for _, s := range stmt.Body {
				err := i.executeStatement(s)
				if err != nil {
					return err
				}

				if i.shouldBreak {
					i.shouldBreak = false
					return nil
				}

				if i.shouldContinue {
					i.shouldContinue = false
					break
				}

				if i.shouldExit {
					return nil
				}
			}
		}
	default:
		return fmt.Errorf("cannot iterate over type: %T", val)
	}

	return nil
}

// executeReturn ejecuta un return
func (i *Interpreter) executeReturn(stmt *Return) error {
	if stmt.Value != nil {
		value, err := i.evaluateExpression(stmt.Value)
		if err != nil {
			return err
		}
		i.returnVal = value
	} else {
		i.returnVal = Value{Value: nil, Type: "void"}
	}

	i.shouldExit = true
	return nil
}

// evaluateExpression evalúa una expresión
func (i *Interpreter) evaluateExpression(expr Expression) (Value, error) {
	switch e := expr.(type) {
	case *Literal:
		return Value{Value: e.Value, Type: e.Type}, nil
	case *Identifier:
		value, exists := i.env.Get(e.Name)
		if !exists {
			return Value{}, fmt.Errorf("undefined variable: %s", e.Name)
		}
		return value, nil
	case *BinaryExpr:
		return i.evaluateBinaryExpr(e)
	case *UnaryExpr:
		return i.evaluateUnaryExpr(e)
	case *FuncCall:
		return i.evaluateFuncCall(e)
	case *ArrayLiteral: // ESTE CASO DEBE ESTAR
		return i.evaluateArrayLiteral(e)
	default:
		return Value{}, fmt.Errorf("unhandled expression type: %T", e)
	}
}

// evaluateBinaryExpr evalúa una expresión binaria
func (i *Interpreter) evaluateBinaryExpr(expr *BinaryExpr) (Value, error) {
	left, err := i.evaluateExpression(expr.Left)
	if err != nil {
		return Value{}, err
	}

	right, err := i.evaluateExpression(expr.Right)
	if err != nil {
		return Value{}, err
	}

	switch expr.Operator {
	case "+":
		return i.addValues(left, right)
	case "-":
		return i.subtractValues(left, right)
	case "*":
		return i.multiplyValues(left, right)
	case "/":
		return i.divideValues(left, right)
	case "%":
		return i.moduloValues(left, right)
	case "==":
		return Value{Value: i.compareValues(left, right) == 0, Type: "bool"}, nil
	case "!=":
		return Value{Value: i.compareValues(left, right) != 0, Type: "bool"}, nil
	case "<":
		return Value{Value: i.compareValues(left, right) < 0, Type: "bool"}, nil
	case "<=":
		return Value{Value: i.compareValues(left, right) <= 0, Type: "bool"}, nil
	case ">":
		return Value{Value: i.compareValues(left, right) > 0, Type: "bool"}, nil
	case ">=":
		return Value{Value: i.compareValues(left, right) >= 0, Type: "bool"}, nil
	case "&&":
		return Value{Value: i.isTruthy(left) && i.isTruthy(right), Type: "bool"}, nil
	case "||":
		return Value{Value: i.isTruthy(left) || i.isTruthy(right), Type: "bool"}, nil
	default:
		return Value{}, fmt.Errorf("unsupported binary operator: %s", expr.Operator)
	}
}

// evaluateUnaryExpr evalúa una expresión unaria
func (i *Interpreter) evaluateUnaryExpr(expr *UnaryExpr) (Value, error) {
	operand, err := i.evaluateExpression(expr.Operand)
	if err != nil {
		return Value{}, err
	}

	switch expr.Operator {
	case "-":
		return i.negateValue(operand)
	case "!":
		return Value{Value: !i.isTruthy(operand), Type: "bool"}, nil
	default:
		return Value{}, fmt.Errorf("unsupported unary operator: %s", expr.Operator)
	}
}

// evaluateFuncCall evalúa una llamada a función
func (i *Interpreter) evaluateFuncCall(expr *FuncCall) (Value, error) {
	// Primero buscar en funciones nativas
	if builtin, exists := i.env.GetBuiltin(expr.Name); exists {
		fmt.Printf("📞 Llamando función nativa: %s\n", expr.Name)

		// Evaluar argumentos
		args := make([]Value, len(expr.Arguments))
		for idx, arg := range expr.Arguments {
			value, err := i.evaluateExpression(arg)
			if err != nil {
				return Value{}, err
			}
			args[idx] = value
		}

		// Llamar a la función nativa
		return builtin(args)
	}

	// Si no es builtin, buscar función definida por el usuario
	fn, exists := i.env.GetFunction(expr.Name)
	if !exists {
		return Value{}, fmt.Errorf("undefined function: %s at line %d, column %d",
			expr.Name, expr.Line, expr.Column)
	}

	// ... resto del código existente para funciones de usuario ...

	// Evaluar argumentos
	args := make([]Value, len(expr.Arguments))
	for idx, arg := range expr.Arguments {
		value, err := i.evaluateExpression(arg)
		if err != nil {
			return Value{}, fmt.Errorf("error evaluating argument %d: %v", idx, err)
		}
		args[idx] = value
	}

	// Verificar número de parámetros
	if len(args) != len(fn.Parameters) {
		return Value{}, fmt.Errorf("function %s expects %d arguments, got %d at line %d",
			expr.Name, len(fn.Parameters), len(args), expr.Line)
	}

	// Crear nuevo entorno para la función
	funcEnv := NewEnvironment(i.env)
	oldEnv := i.env
	i.env = funcEnv
	defer func() { i.env = oldEnv }()

	// Establecer parámetros
	for idx, param := range fn.Parameters {
		i.env.Set(param.Name, args[idx])
	}

	// Ejecutar cuerpo de la función
	oldShouldExit := i.shouldExit
	i.shouldExit = false

	for _, stmt := range fn.Body {
		err := i.executeStatement(stmt)
		if err != nil {
			return Value{}, err
		}
		if i.shouldExit {
			break
		}
	}

	result := i.returnVal
	i.shouldExit = oldShouldExit
	i.returnVal = Value{}

	return result, nil
}

// Funciones auxiliares para operaciones aritméticas
func (i *Interpreter) addValues(left, right Value) (Value, error) {
	switch l := left.Value.(type) {
	case int:
		if r, ok := right.Value.(int); ok {
			return Value{Value: l + r, Type: "int"}, nil
		}
		if r, ok := right.Value.(float64); ok {
			return Value{Value: float64(l) + r, Type: "float64"}, nil
		}
	case float64:
		if r, ok := right.Value.(float64); ok {
			return Value{Value: l + r, Type: "float64"}, nil
		}
		if r, ok := right.Value.(int); ok {
			return Value{Value: l + float64(r), Type: "float64"}, nil
		}
	case string:
		if r, ok := right.Value.(string); ok {
			return Value{Value: l + r, Type: "string"}, nil
		}
	}
	return Value{}, fmt.Errorf("cannot add %T and %T", left.Value, right.Value)
}

func (i *Interpreter) subtractValues(left, right Value) (Value, error) {
	switch l := left.Value.(type) {
	case int:
		if r, ok := right.Value.(int); ok {
			return Value{Value: l - r, Type: "int"}, nil
		}
		if r, ok := right.Value.(float64); ok {
			return Value{Value: float64(l) - r, Type: "float64"}, nil
		}
	case float64:
		if r, ok := right.Value.(float64); ok {
			return Value{Value: l - r, Type: "float64"}, nil
		}
		if r, ok := right.Value.(int); ok {
			return Value{Value: l - float64(r), Type: "float64"}, nil
		}
	}
	return Value{}, fmt.Errorf("cannot subtract %T and %T", left.Value, right.Value)
}

func (i *Interpreter) multiplyValues(left, right Value) (Value, error) {
	switch l := left.Value.(type) {
	case int:
		if r, ok := right.Value.(int); ok {
			return Value{Value: l * r, Type: "int"}, nil
		}
		if r, ok := right.Value.(float64); ok {
			return Value{Value: float64(l) * r, Type: "float64"}, nil
		}
	case float64:
		if r, ok := right.Value.(float64); ok {
			return Value{Value: l * r, Type: "float64"}, nil
		}
		if r, ok := right.Value.(int); ok {
			return Value{Value: l * float64(r), Type: "float64"}, nil
		}
	}
	return Value{}, fmt.Errorf("cannot multiply %T and %T", left.Value, right.Value)
}

func (i *Interpreter) divideValues(left, right Value) (Value, error) {
	switch l := left.Value.(type) {
	case int:
		if r, ok := right.Value.(int); ok {
			if r == 0 {
				return Value{}, fmt.Errorf("division by zero")
			}
			return Value{Value: l / r, Type: "int"}, nil
		}
		if r, ok := right.Value.(float64); ok {
			if r == 0 {
				return Value{}, fmt.Errorf("division by zero")
			}
			return Value{Value: float64(l) / r, Type: "float64"}, nil
		}
	case float64:
		if r, ok := right.Value.(float64); ok {
			if r == 0 {
				return Value{}, fmt.Errorf("division by zero")
			}
			return Value{Value: l / r, Type: "float64"}, nil
		}
		if r, ok := right.Value.(int); ok {
			if r == 0 {
				return Value{}, fmt.Errorf("division by zero")
			}
			return Value{Value: l / float64(r), Type: "float64"}, nil
		}
	}
	return Value{}, fmt.Errorf("cannot divide %T and %T", left.Value, right.Value)
}

func (i *Interpreter) moduloValues(left, right Value) (Value, error) {
	if l, ok := left.Value.(int); ok {
		if r, ok := right.Value.(int); ok {
			if r == 0 {
				return Value{}, fmt.Errorf("modulo by zero")
			}
			return Value{Value: l % r, Type: "int"}, nil
		}
	}
	return Value{}, fmt.Errorf("modulo operation requires integers")
}

func (i *Interpreter) negateValue(value Value) (Value, error) {
	switch v := value.Value.(type) {
	case int:
		return Value{Value: -v, Type: "int"}, nil
	case float64:
		return Value{Value: -v, Type: "float64"}, nil
	default:
		return Value{}, fmt.Errorf("cannot negate %T", v)
	}
}

func (i *Interpreter) compareValues(left, right Value) int {
	if reflect.TypeOf(left.Value) != reflect.TypeOf(right.Value) {
		if l, ok := left.Value.(int); ok {
			if r, ok := right.Value.(float64); ok {
				return i.compareFloat64(float64(l), r)
			}
		}
		if l, ok := left.Value.(float64); ok {
			if r, ok := right.Value.(int); ok {
				return i.compareFloat64(l, float64(r))
			}
		}
		return 1
	}

	switch l := left.Value.(type) {
	case int:
		r := right.Value.(int)
		if l < r {
			return -1
		} else if l > r {
			return 1
		}
		return 0
	case float64:
		r := right.Value.(float64)
		return i.compareFloat64(l, r)
	case string:
		r := right.Value.(string)
		if l < r {
			return -1
		} else if l > r {
			return 1
		}
		return 0
	case bool:
		r := right.Value.(bool)
		if l == r {
			return 0
		}
		if l {
			return 1
		}
		return -1
	default:
		return 0
	}
}

func (i *Interpreter) compareFloat64(l, r float64) int {
	if l < r {
		return -1
	} else if l > r {
		return 1
	}
	return 0
}

func (i *Interpreter) isTruthy(value Value) bool {
	switch v := value.Value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0.0
	case string:
		return v != ""
	case nil:
		return false
	default:
		return true
	}
}

func (i *Interpreter) valueToString(value Value) string {
	switch v := value.Value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (i *Interpreter) executePlusAssign(stmt *PlusAssign) error {
	if id, ok := stmt.Target.(*Identifier); ok {
		currentValue, exists := i.env.Get(id.Name)
		if !exists {
			return fmt.Errorf("undefined variable: %s", id.Name)
		}

		addValue, err := i.evaluateExpression(stmt.Value)
		if err != nil {
			return err
		}

		result, err := i.addValues(currentValue, addValue)
		if err != nil {
			return err
		}

		i.env.Set(id.Name, result)
		return nil
	}

	return fmt.Errorf("unsupported plus assign target: %T", stmt.Target)
}

func (i *Interpreter) executeMinusAssign(stmt *MinusAssign) error {
	if id, ok := stmt.Target.(*Identifier); ok {
		currentValue, exists := i.env.Get(id.Name)
		if !exists {
			return fmt.Errorf("undefined variable: %s", id.Name)
		}

		subValue, err := i.evaluateExpression(stmt.Value)
		if err != nil {
			return err
		}

		result, err := i.subtractValues(currentValue, subValue)
		if err != nil {
			return err
		}

		i.env.Set(id.Name, result)
		return nil
	}

	return fmt.Errorf("unsupported minus assign target: %T", stmt.Target)
}

func (i *Interpreter) executeMulAssign(stmt *MulAssign) error {
	if id, ok := stmt.Target.(*Identifier); ok {
		currentValue, exists := i.env.Get(id.Name)
		if !exists {
			return fmt.Errorf("undefined variable: %s", id.Name)
		}

		mulValue, err := i.evaluateExpression(stmt.Value)
		if err != nil {
			return err
		}

		result, err := i.multiplyValues(currentValue, mulValue)
		if err != nil {
			return err
		}

		i.env.Set(id.Name, result)
		return nil
	}

	return fmt.Errorf("unsupported multiply assign target: %T", stmt.Target)
}

func (i *Interpreter) executeDivAssign(stmt *DivAssign) error {
	if id, ok := stmt.Target.(*Identifier); ok {
		currentValue, exists := i.env.Get(id.Name)
		if !exists {
			return fmt.Errorf("undefined variable: %s", id.Name)
		}

		divValue, err := i.evaluateExpression(stmt.Value)
		if err != nil {
			return err
		}

		result, err := i.divideValues(currentValue, divValue)
		if err != nil {
			return err
		}

		i.env.Set(id.Name, result)
		return nil
	}

	return fmt.Errorf("unsupported divide assign target: %T", stmt.Target)
}

func (i *Interpreter) executeModAssign(stmt *ModAssign) error {
	if id, ok := stmt.Target.(*Identifier); ok {
		currentValue, exists := i.env.Get(id.Name)
		if !exists {
			return fmt.Errorf("undefined variable: %s", id.Name)
		}

		modValue, err := i.evaluateExpression(stmt.Value)
		if err != nil {
			return err
		}

		result, err := i.moduloValues(currentValue, modValue)
		if err != nil {
			return err
		}

		i.env.Set(id.Name, result)
		return nil
	}

	return fmt.Errorf("unsupported modulo assign target: %T", stmt.Target)
}

func (i *Interpreter) evaluateArrayLiteral(expr *ArrayLiteral) (Value, error) {
	elements := make([]interface{}, len(expr.Elements))

	for idx, elementExpr := range expr.Elements {
		value, err := i.evaluateExpression(elementExpr)
		if err != nil {
			return Value{}, err
		}
		elements[idx] = value.Value
	}

	return Value{Value: elements, Type: "array"}, nil
}

func (env *Environment) registerBuiltins() {
	// Conversión de tipos
	env.builtins["int"] = builtinInt
	env.builtins["float"] = builtinFloat
	env.builtins["string"] = builtinString

	// Funciones
	env.builtins["atoi"] = builtinAtoi
	env.builtins["parse_float"] = builtinParseFloat
	env.builtins["TypeOf"] = builtinTypeOf
	env.builtins["len"] = builtinLen
}

func builtinInt(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("int() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch v := arg.Value.(type) {
	case string:
		// Intentar parsear como float primero para manejar decimales
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return Value{}, fmt.Errorf("cannot convert '%s' to int", v)
		}
		return Value{Value: int(f), Type: "int"}, nil
	case float64:
		return Value{Value: int(v), Type: "int"}, nil
	case int:
		return arg, nil
	default:
		return Value{}, fmt.Errorf("int() cannot convert %T to int", v)
	}
}

// Función float() - convierte a float
func builtinFloat(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("float() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch v := arg.Value.(type) {
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return Value{}, fmt.Errorf("cannot convert '%s' to float", v)
		}
		return Value{Value: f, Type: "float64"}, nil
	case int:
		return Value{Value: float64(v), Type: "float64"}, nil
	case float64:
		return arg, nil
	default:
		return Value{}, fmt.Errorf("float() cannot convert %T to float", v)
	}
}

// Función string() - convierte a string
func builtinString(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("string() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch v := arg.Value.(type) {
	case string:
		return arg, nil
	case int:
		return Value{Value: strconv.Itoa(v), Type: "string"}, nil
	case float64:
		return Value{Value: strconv.FormatFloat(v, 'g', -1, 64), Type: "string"}, nil
	case bool:
		if v {
			return Value{Value: "true", Type: "string"}, nil
		}
		return Value{Value: "false", Type: "string"}, nil
	default:
		return Value{Value: fmt.Sprintf("%v", v), Type: "string"}, nil
	}
}

// Función atoi() - string a int (compatible con strconv.Atoi)
func builtinAtoi(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("atoi() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	if str, ok := arg.Value.(string); ok {
		val, err := strconv.Atoi(str)
		if err != nil {
			return Value{}, fmt.Errorf("atoi: %v", err)
		}
		return Value{Value: val, Type: "int"}, nil
	}

	return Value{}, fmt.Errorf("atoi() expects string argument, got %T", arg.Value)
}

// Función parse_float() - string a float
func builtinParseFloat(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("parse_float() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	if str, ok := arg.Value.(string); ok {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return Value{}, fmt.Errorf("parse_float: %v", err)
		}
		return Value{Value: val, Type: "float64"}, nil
	}

	return Value{}, fmt.Errorf("parse_float() expects string argument, got %T", arg.Value)
}

func builtinTypeOf(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("TypeOf() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	typeStr := arg.Type

	switch arg.Value.(type) {
	case []interface{}:
		if arr, ok := arg.Value.([]interface{}); ok && len(arr) > 0 {
			firstElem := Value{Value: arr[0]}
			elemType := getTypeString(firstElem)
			typeStr = "[]" + elemType
		} else {
			typeStr = "[]unknown"
		}
	}

	return Value{Value: typeStr, Type: "string"}, nil
}

// Función len() - longitud de strings o arrays
func builtinLen(args []Value) (Value, error) {
	if len(args) != 1 {
		return Value{}, fmt.Errorf("len() expects exactly 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch v := arg.Value.(type) {
	case string:
		return Value{Value: len(v), Type: "int"}, nil
	case []interface{}:
		return Value{Value: len(v), Type: "int"}, nil
	default:
		return Value{}, fmt.Errorf("len() expects string or array, got %T", v)
	}
}

// Función auxiliar para obtener el tipo como string
func getTypeString(v Value) string {
	switch v.Value.(type) {
	case int:
		return "int"
	case float64:
		return "f64"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func (env *Environment) GetBuiltin(name string) (BuiltinFunc, bool) {
	if fn, exists := env.builtins[name]; exists {
		return fn, true
	}
	if env.parent != nil {
		return env.parent.GetBuiltin(name)
	}
	return nil, false
}

func (i *Interpreter) executeSwitchStmt(stmt *SwitchStmt) error {
	var switchValue Value
	var err error

	// Evaluar expresión del switch (si existe)
	if stmt.Expression != nil {
		switchValue, err = i.evaluateExpression(stmt.Expression)
		if err != nil {
			return err
		}
	}

	executedCase := false
	shouldFallthrough := false

	// Evaluar cada case
	for _, caseClause := range stmt.Cases {
		matched := false

		if !shouldFallthrough {
			// Evaluar si algún valor del case coincide
			for _, caseExpr := range caseClause.Values {
				caseValue, err := i.evaluateExpression(caseExpr)
				if err != nil {
					return err
				}

				if stmt.Expression != nil {
					// Comparar valores
					if i.compareValues(switchValue, caseValue) == 0 {
						matched = true
						break
					}
				} else {
					// Switch sin expresión - evaluar cada case como booleano
					if i.isTruthy(caseValue) {
						matched = true
						break
					}
				}
			}
		}

		if matched || shouldFallthrough {
			executedCase = true
			shouldFallthrough = false

			// Ejecutar statements del case
			for _, s := range caseClause.Statements {
				// Verificar si es fallthrough
				if _, ok := s.(*Fallthrough); ok {
					shouldFallthrough = true
					break
				}

				err := i.executeStatement(s)
				if err != nil {
					return err
				}

				if i.shouldBreak {
					i.shouldBreak = false
					return nil
				}

				if i.shouldExit {
					return nil
				}
			}

			if !shouldFallthrough {
				break
			}
		}
	}

	// Ejecutar default si no se ejecutó ningún case
	if !executedCase && stmt.Default != nil {
		for _, s := range stmt.Default.Statements {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}

			if i.shouldBreak {
				i.shouldBreak = false
				return nil
			}

			if i.shouldExit {
				return nil
			}
		}
	}

	return nil
}

// executeFallthrough ejecuta un fallthrough (usado dentro de switch)
func (i *Interpreter) executeFallthrough(stmt *Fallthrough) error {
	// El manejo real del fallthrough se hace en executeSwitchStmt
	// Este método existe por consistencia con el patrón visitor
	return nil
}

// executeForCondition ejecuta un for con solo condición
func (i *Interpreter) executeForCondition(stmt *ForCondition) error {
	for {
		condition, err := i.evaluateExpression(stmt.Condition)
		if err != nil {
			return err
		}

		if !i.isTruthy(condition) {
			break
		}

		i.shouldBreak = false
		i.shouldContinue = false

		for _, s := range stmt.Body {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}

			if i.shouldBreak {
				i.shouldBreak = false
				return nil
			}

			if i.shouldContinue {
				i.shouldContinue = false
				break
			}

			if i.shouldExit {
				return nil
			}
		}
	}

	return nil
}

// executeForClassic ejecuta un for clásico (init; condition; update)
func (i *Interpreter) executeForClassic(stmt *ForClassic) error {
	// Crear nuevo entorno para el scope del for
	forEnv := NewEnvironment(i.env)
	oldEnv := i.env
	i.env = forEnv
	defer func() { i.env = oldEnv }()

	// Ejecutar inicialización
	if stmt.Init != nil {
		err := i.executeStatement(stmt.Init)
		if err != nil {
			return err
		}
	}

	// Loop principal
	for {
		// Evaluar condición (si existe)
		if stmt.Condition != nil {
			condition, err := i.evaluateExpression(stmt.Condition)
			if err != nil {
				return err
			}

			if !i.isTruthy(condition) {
				break
			}
		}

		i.shouldBreak = false
		i.shouldContinue = false

		// Ejecutar cuerpo
		for _, s := range stmt.Body {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}

			if i.shouldBreak {
				i.shouldBreak = false
				return nil
			}

			if i.shouldContinue {
				i.shouldContinue = false
				break
			}

			if i.shouldExit {
				return nil
			}
		}

		// Ejecutar actualización
		if stmt.Update != nil {
			err := i.executeStatement(stmt.Update)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// executeForIndexValue ejecuta un for con índice y valor
func (i *Interpreter) executeForIndexValue(stmt *ForIndexValue) error {
	// Evaluar el iterable
	iterable, err := i.evaluateExpression(stmt.Iterable)
	if err != nil {
		return err
	}

	// Crear nuevo entorno para el scope del for
	forEnv := NewEnvironment(i.env)
	oldEnv := i.env
	i.env = forEnv
	defer func() { i.env = oldEnv }()

	// Verificar que es un array
	if arr, ok := iterable.Value.([]interface{}); ok {
		for idx, item := range arr {
			// Establecer índice y valor
			i.env.Set(stmt.Index, Value{Value: idx, Type: "int"})
			i.env.Set(stmt.Value, Value{Value: item, Type: "auto"})

			i.shouldBreak = false
			i.shouldContinue = false

			for _, s := range stmt.Body {
				err := i.executeStatement(s)
				if err != nil {
					return err
				}

				if i.shouldBreak {
					i.shouldBreak = false
					return nil
				}

				if i.shouldContinue {
					i.shouldContinue = false
					break
				}

				if i.shouldExit {
					return nil
				}
			}
		}
	} else {
		return fmt.Errorf("cannot iterate with index,value over type: %T", iterable.Value)
	}

	return nil
}

// executeForInfinite ejecuta un for infinito
func (i *Interpreter) executeForInfinite(stmt *ForInfinite) error {
	for {
		i.shouldBreak = false
		i.shouldContinue = false

		for _, s := range stmt.Body {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}

			if i.shouldBreak {
				i.shouldBreak = false
				return nil
			}

			if i.shouldContinue {
				i.shouldContinue = false
				break
			}

			if i.shouldExit {
				return nil
			}
		}
	}
}

// executeForRange ejecuta un for range (similar a ForIndexValue)
func (i *Interpreter) executeForRange(stmt *ForRange) error {
	// Evaluar el iterable
	iterable, err := i.evaluateExpression(stmt.Iterable)
	if err != nil {
		return err
	}

	// Crear nuevo entorno para el scope del for
	forEnv := NewEnvironment(i.env)
	oldEnv := i.env
	i.env = forEnv
	defer func() { i.env = oldEnv }()

	// Verificar que es un array
	if arr, ok := iterable.Value.([]interface{}); ok {
		for idx, item := range arr {
			// Establecer índice y valor
			i.env.Set(stmt.Index, Value{Value: idx, Type: "int"})
			i.env.Set(stmt.Value, Value{Value: item, Type: "auto"})

			i.shouldBreak = false
			i.shouldContinue = false

			for _, s := range stmt.Body {
				err := i.executeStatement(s)
				if err != nil {
					return err
				}

				if i.shouldBreak {
					i.shouldBreak = false
					return nil
				}

				if i.shouldContinue {
					i.shouldContinue = false
					break
				}

				if i.shouldExit {
					return nil
				}
			}
		}
	} else {
		return fmt.Errorf("cannot range over type: %T", iterable.Value)
	}

	return nil
}
