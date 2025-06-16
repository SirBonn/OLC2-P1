package value

type CharacterValue struct {
	BaseValue
	InternalValue string
}

func NewCharacterValue(value string) *CharacterValue {
	return &CharacterValue{
		BaseValue:     BaseValue{valueType: IVOR_CHARACTER},
		InternalValue: value,
	}
}

func (c *CharacterValue) Value() interface{} {
	return c.InternalValue
}

func (c *CharacterValue) Copy() IVOR {
	return NewCharacterValue(c.InternalValue)
}

func (c *CharacterValue) Equals(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue == o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "operación '==' solo válida entre caracteres"
}
func (c *CharacterValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := c.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

func (c *CharacterValue) Add(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '+' no soportada para caracteres"
}
func (c *CharacterValue) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '-' no soportada para caracteres"
}
func (c *CharacterValue) Multiply(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '*' no soportada para caracteres"
}

func (c *CharacterValue) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '/' no soportada para caracteres"
}

func (c *CharacterValue) Modulo(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '%' no soportada para caracteres"
}

func (c *CharacterValue) Power(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '^' no soportada para caracteres"
}

func (c *CharacterValue) LessThan(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue < o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "operación '<' solo válida entre caracteres"
}

func (c *CharacterValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue > o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "operación '>' solo válida entre caracteres"
}

func (c *CharacterValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue <= o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "operación '<=' solo válida entre caracteres"
}

func (c *CharacterValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue >= o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "operación '>=' solo válida entre caracteres"
}

func (c *CharacterValue) And(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue == o.InternalValue), true, ""
	}
	return nil, false, "operación '&&' solo válida entre caracteres"
}

func (c *CharacterValue) Or(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*CharacterValue); ok {
		return NewBoolValue(c.InternalValue == o.InternalValue), true, ""
	}
	return nil, false, "operación '||' solo válida entre caracteres"
}

func (c *CharacterValue) Not() (IVOR, bool, string) {
	return nil, false, "operación '!' no soportada para caracteres"
}

func (c *CharacterValue) Type() string {
	return c.valueType
}

func (c *CharacterValue) String() string {
	return c.InternalValue
}

func (c *CharacterValue) AddAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '+=' no soportada para caracteres"
}

func (c *CharacterValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '-=' no soportada para caracteres"
}

func (c *CharacterValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '*=' no soportada para caracteres"
}

func (c *CharacterValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '/=' no soportada para caracteres"
}

func (c *CharacterValue) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '%=' no soportada para caracteres"
}

func (c *CharacterValue) PowerAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación '^=' no soportada para caracteres"
}
