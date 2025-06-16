package value

type BoolValue struct {
	BaseValue     // Embed BaseValue para comportamientos básicos
	InternalValue bool
}

// Constructor mejorado
func NewBoolValue(value bool) *BoolValue {
	return &BoolValue{
		BaseValue:     BaseValue{valueType: IVOR_BOOL},
		InternalValue: value,
	}
}

// Métodos básicos
func (b *BoolValue) Value() interface{} {
	return b.InternalValue
}

func (b *BoolValue) Copy() IVOR {
	return NewBoolValue(b.InternalValue)
}

// Operaciones lógicas (implementaciones propias de booleanos)
func (b *BoolValue) And(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*BoolValue); ok {
		return NewBoolValue(b.InternalValue && o.InternalValue), true, ""
	}
	return nil, false, "operación AND solo válida entre booleanos"
}

func (b *BoolValue) Or(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*BoolValue); ok {
		return NewBoolValue(b.InternalValue || o.InternalValue), true, ""
	}
	return nil, false, "operación OR solo válida entre booleanos"
}

func (b *BoolValue) Not() (IVOR, bool, string) {
	return NewBoolValue(!b.InternalValue), true, ""
}

// Operaciones de comparación
func (b *BoolValue) Equals(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*BoolValue); ok {
		return NewBoolValue(b.InternalValue == o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, ""
}

func (b *BoolValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := b.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

// Operaciones no soportadas para booleanos
func (b *BoolValue) Add(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '+' no soportada para booleanos"
}

func (b *BoolValue) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '-' no soportada para booleanos"
}

func (b *BoolValue) Multiply(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '*' no soportada para booleanos"
}

func (b *BoolValue) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '/' no soportada para booleanos"
}

func (b *BoolValue) Mod(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '%' no soportada para booleanos"
}

func (b *BoolValue) LessThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '<' no soportada para booleanos"
}

func (b *BoolValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '<=' no soportada para booleanos"
}

func (b *BoolValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '>' no soportada para booleanos"
}

func (b *BoolValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '>=' no soportada para booleanos"
}

// Operaciones de asignación compuesta
func (b *BoolValue) AddAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '+=' no soportada para booleanos"
}

func (b *BoolValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '-=' no soportada para booleanos"
}

func (b *BoolValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '*=' no soportada para booleanos"
}

func (b *BoolValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '/=' no soportada para booleanos"
}

func (b *BoolValue) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '%=' no soportada para booleanos"
}
