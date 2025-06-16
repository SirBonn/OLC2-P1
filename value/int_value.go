package value

type IntValue struct {
	BaseValue
	InternalValue int
}

func NewIntValue(value int) *IntValue {
	return &IntValue{
		BaseValue:     BaseValue{valueType: IVOR_INT},
		InternalValue: value,
	}
}

func (i *IntValue) Value() interface{} {
	return i.InternalValue
}

func (i *IntValue) Copy() IVOR {
	return NewIntValue(i.InternalValue)
}

func (i *IntValue) Add(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewIntValue(i.InternalValue + o.InternalValue), true, ""
	case *FloatValue:
		return NewFloatValue(float64(i.InternalValue) + o.InternalValue), true, ""
	default:
		return nil, false, "no se puede sumar entero con " + other.Type()
	}
}

func (i *IntValue) Subtract(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewIntValue(i.InternalValue - o.InternalValue), true, ""
	case *FloatValue:
		return NewFloatValue(float64(i.InternalValue) - o.InternalValue), true, ""
	default:
		return nil, false, "no se puede restar entero con " + other.Type()
	}
}

func (i *IntValue) Multiply(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewIntValue(i.InternalValue * o.InternalValue), true, ""
	case *FloatValue:
		return NewFloatValue(float64(i.InternalValue) * o.InternalValue), true, ""
	default:
		return nil, false, "no se puede multiplicar entero con " + other.Type()
	}
}

func (i *IntValue) Divide(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		if o.InternalValue == 0 {
			return nil, false, "división por cero"
		}
		return NewIntValue(i.InternalValue / o.InternalValue), true, ""
	case *FloatValue:
		if o.InternalValue == 0 {
			return nil, false, "división por cero"
		}
		return NewFloatValue(float64(i.InternalValue) / o.InternalValue), true, ""
	default:
		return nil, false, "no se puede dividir entero con " + other.Type()
	}
}

func (i *IntValue) Mod(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*IntValue); ok {
		if o.InternalValue == 0 {
			return nil, false, "módulo por cero"
		}
		return NewIntValue(i.InternalValue % o.InternalValue), true, ""
	}
	return nil, false, "operación módulo solo válida entre enteros"
}

// Operaciones de comparación
func (i *IntValue) Equals(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(i.InternalValue == o.InternalValue), true, ""
	case *FloatValue:
		return NewBoolValue(float64(i.InternalValue) == o.InternalValue), true, ""
	default:
		return NewBoolValue(false), true, ""
	}
}

func (i *IntValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := i.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

// Operaciones de asignación compuesta
func (i *IntValue) AddAssign(other IVOR) (IVOR, bool, string) {
	return i.Add(other) // Para enteros, es lo mismo que Add
}

func (i *IntValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return i.Subtract(other)
}

func (i *IntValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return i.Multiply(other)
}
func (i *IntValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return i.Divide(other)
}
func (i *IntValue) ModAssign(other IVOR) (IVOR, bool, string) {
	res, ok, msg := i.Mod(other)
	if !ok {
		return nil, false, msg
	}
	return res, true, ""
}

// Funciones de comparación
func (i *IntValue) LessThan(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(i.InternalValue < o.InternalValue), true, ""
	case *FloatValue:
		return NewBoolValue(float64(i.InternalValue) < o.InternalValue), true, ""
	default:
		return NewBoolValue(false), true, "operación '<' no soportada entre enteros y " + other.Type()
	}
}
func (i *IntValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(i.InternalValue <= o.InternalValue), true, ""
	case *FloatValue:
		return NewBoolValue(float64(i.InternalValue) <= o.InternalValue), true, ""
	default:
		return NewBoolValue(false), true, "operación '<=' no soportada entre enteros y " + other.Type()
	}
}

func (i *IntValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(i.InternalValue > o.InternalValue), true, ""
	case *FloatValue:
		return NewBoolValue(float64(i.InternalValue) > o.InternalValue), true, ""
	default:
		return NewBoolValue(false), true, "operación '>' no soportada entre enteros y " + other.Type()
	}
}
func (i *IntValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(i.InternalValue >= o.InternalValue), true, ""
	case *FloatValue:
		return NewBoolValue(float64(i.InternalValue) >= o.InternalValue), true, ""
	default:
		return NewBoolValue(false), true, "operación '>=' no soportada entre enteros y " + other.Type()
	}
}

// Operaciones lógicas (solo para compatibilidad)
func (i *IntValue) And(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación lógica AND no soportada para enteros"
}

func (i *IntValue) Or(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación lógica OR no soportada para enteros"
}

func (i *IntValue) Not() (IVOR, bool, string) {
	return nil, false, "operación lógica NOT no soportada para enteros"
}
