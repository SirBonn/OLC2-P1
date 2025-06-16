package value

const (
	IVOR_INT              = "Int"
	IVOR_FLOAT            = "Float"
	IVOR_STRING           = "String"
	IVOR_BOOL             = "Bool"
	IVOR_CHARACTER        = "Character"
	IVOR_NIL              = "nil"
	IVOR_BUILTIN_FUNCTION = "builtin_function"
	IVOR_FUNCTION         = "function"
	IVOR_VECTOR           = "vector"
	IVOR_OBJECT           = "object"
	IVOR_ARRAY            = "array"
	IVOR_ANY              = "any"
	IVOR_POINTER          = "pointer"
	IVOR_MATRIX           = "matrix"
	IVOR_SELF             = "self"
	IVOR_UNINITIALIZED    = "uninitialized"
	IVOR_UNDEFINED        = "undefined"
	IVOR_ERROR            = "error"
)

// IVOR stands for Internal Value Object Representation
type IVOR interface {
	Value() interface{}
	Type() string
	Copy() IVOR
}

type BaseValue struct {
	valueType string
}

func (b *BaseValue) Type() string {
	return b.valueType
}

// Implementaciones base que devuelven error por defecto
func (b *BaseValue) Add(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '+' no soportada para " + b.Type()
}

func (b *BaseValue) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '-' no soportada para " + b.Type()
}

func (b *BaseValue) Multiply(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '*' no soportada para " + b.Type()
}

func (b *BaseValue) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '/' no soportada para " + b.Type()
}

func (b *BaseValue) Modulo(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '%' no soportada para " + b.Type()
}

func (b *BaseValue) Power(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '^' no soportada para " + b.Type()
}

func (b *BaseValue) Equal(other IVOR) (bool, string) {
	return false, "operación '==' no soportada para " + b.Type()
}

func (b *BaseValue) Equals(other IVOR) (bool, string) {
	return false, "operación ecuals no soportado para" + b.Type()
}

func (b *BaseValue) NotEqual(other IVOR) (bool, string) {
	return false, "operación '!=' no soportada para " + b.Type()
}

func (b *BaseValue) LessThan(other IVOR) (bool, string) {
	return false, "operación '<' no soportada para " + b.Type()
}

func (b *BaseValue) GreaterThan(other IVOR) (bool, string) {
	return false, "operación '>' no soportada para " + b.Type()
}

func (b *BaseValue) LessThanOrEqual(other IVOR) (bool, string) {
	return false, "operación '<=' no soportada para " + b.Type()
}

func (b *BaseValue) GreaterThanOrEqual(other IVOR) (bool, string) {
	return false, "operación '>=' no soportada para " + b.Type()
}

func (b *BaseValue) And(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '&&' no soportada para " + b.Type()
}

func (b *BaseValue) Or(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '||' no soportada para " + b.Type()
}
