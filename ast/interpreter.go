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

// Environment representa el entorno de ejecución
type Environment struct {
	parent    *Environment
	variables map[string]Value
	functions map[string]*FuncDecl
}

// NewEnvironment crea un nuevo entorno
func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		parent:    parent,
		variables: make(map[string]Value),
		functions: make(map[string]*FuncDecl),
	}
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
	// Primero, registrar todas las funciones
	for _, stmt := range program.Statements {
		if funcDecl, ok := stmt.(*FuncDecl); ok {
			i.env.SetFunction(funcDecl.Name, funcDecl)
		}
	}

	// Luego ejecutar los statements
	for _, stmt := range program.Statements {
		if _, ok := stmt.(*FuncDecl); ok {
			continue // Ya procesamos las funciones
		}

		err := i.executeStatement(stmt)
		if err != nil {
			return i.output.String(), err
		}

		if i.shouldExit {
			break
		}
	}

	return i.output.String(), nil
}

// executeStatement ejecuta un statement
func (i *Interpreter) executeStatement(stmt Statement) error {
	if i.shouldExit || i.shouldBreak || i.shouldContinue {
		return nil
	}

	switch s := stmt.(type) {
	case *PrintStmt:
		return i.executePrintStmt(s)
	case *VarDecl:
		return i.executeVarDecl(s)
	case *Assignment:
		return i.executeAssignment(s)
	case *PlusAssign:
		return i.executePlusAssign(s)
	case *MinusAssign: // AGREGAR
		return i.executeMinusAssign(s)
	case *MulAssign: // AGREGAR
		return i.executeMulAssign(s)
	case *DivAssign: // AGREGAR
		return i.executeDivAssign(s)
	case *ModAssign: // AGREGAR
		return i.executeModAssign(s)
	case *IfStmt:
		return i.executeIfStmt(s)
	case *WhileStmt:
		return i.executeWhileStmt(s)
	case *ForStmt:
		return i.executeForStmt(s)
	case *FuncDecl:
		// Ya procesado en Interpret
		return nil
	case *Return:
		return i.executeReturn(s)
	case *Break:
		i.shouldBreak = true
		return nil
	case *Continue:
		i.shouldContinue = true
		return nil
	default:
		return fmt.Errorf("unhandled statement type: %T", s)
	}
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
	condition, err := i.evaluateExpression(stmt.Condition)
	if err != nil {
		return err
	}

	if i.isTruthy(condition) {
		for _, s := range stmt.ThenBranch {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}
			if i.shouldExit || i.shouldBreak || i.shouldContinue {
				break
			}
		}
	} else if len(stmt.ElseBranch) > 0 {
		for _, s := range stmt.ElseBranch {
			err := i.executeStatement(s)
			if err != nil {
				return err
			}
			if i.shouldExit || i.shouldBreak || i.shouldContinue {
				break
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
	fn, exists := i.env.GetFunction(expr.Name)
	if !exists {
		return Value{}, fmt.Errorf("undefined function: %s", expr.Name)
	}

	// Evaluar argumentos
	args := make([]Value, len(expr.Arguments))
	for idx, arg := range expr.Arguments {
		value, err := i.evaluateExpression(arg)
		if err != nil {
			return Value{}, err
		}
		args[idx] = value
	}

	// Verificar número de parámetros
	if len(args) != len(fn.Parameters) {
		return Value{}, fmt.Errorf("function %s expects %d arguments, got %d",
			expr.Name, len(fn.Parameters), len(args))
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
		// Conversión automática para comparación
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
		return 1 // Tipos diferentes, left > right arbitrariamente
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
