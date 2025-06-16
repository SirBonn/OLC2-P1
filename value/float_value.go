package value

type FloatValue struct {
	BaseValue     // Embed BaseValue para comportamientos básicos
	InternalValue float64
}

// Constructor mejorado
func NewFloatValue(value float64) *FloatValue {
	return &FloatValue{
		BaseValue:     BaseValue{valueType: IVOR_FLOAT},
		InternalValue: value,
	}
}

// Métodos básicos
func (f *FloatValue) Value() interface{} {
	return f.InternalValue
}

func (f *FloatValue) Copy() IVOR {
	return NewFloatValue(f.InternalValue)
}

// Operaciones aritméticas
func (f *FloatValue) Add(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewFloatValue(f.InternalValue + float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewFloatValue(f.InternalValue + o.InternalValue), true, ""
	default:
		return nil, false, "no se puede sumar float con " + other.Type()
	}
}

func (f *FloatValue) Subtract(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewFloatValue(f.InternalValue - float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewFloatValue(f.InternalValue - o.InternalValue), true, ""
	default:
		return nil, false, "no se puede restar float con " + other.Type()
	}
}

func (f *FloatValue) Multiply(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewFloatValue(f.InternalValue * float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewFloatValue(f.InternalValue * o.InternalValue), true, ""
	default:
		return nil, false, "no se puede multiplicar float con " + other.Type()
	}
}

func (f *FloatValue) Divide(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		if o.InternalValue == 0 {
			return nil, false, "división por cero"
		}
		return NewFloatValue(f.InternalValue / float64(o.InternalValue)), true, ""
	case *FloatValue:
		if o.InternalValue == 0.0 {
			return nil, false, "división por cero"
		}
		return NewFloatValue(f.InternalValue / o.InternalValue), true, ""
	default:
		return nil, false, "no se puede dividir float con " + other.Type()
	}
}

func (f *FloatValue) Mod(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación módulo no soportada para floats"
}

// Operaciones de comparación
func (f *FloatValue) Equals(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(f.InternalValue == float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewBoolValue(f.InternalValue == o.InternalValue), true, ""
	default:
		return NewBoolValue(false), true, ""
	}
}

func (f *FloatValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := f.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

func (f *FloatValue) LessThan(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(f.InternalValue < float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewBoolValue(f.InternalValue < o.InternalValue), true, ""
	default:
		return nil, false, "no se puede comparar float con " + other.Type()
	}
}

func (f *FloatValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(f.InternalValue <= float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewBoolValue(f.InternalValue <= o.InternalValue), true, ""
	default:
		return nil, false, "no se puede comparar float con " + other.Type()
	}
}

func (f *FloatValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(f.InternalValue > float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewBoolValue(f.InternalValue > o.InternalValue), true, ""
	default:
		return nil, false, "no se puede comparar float con " + other.Type()
	}
}

func (f *FloatValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	switch o := other.(type) {
	case *IntValue:
		return NewBoolValue(f.InternalValue >= float64(o.InternalValue)), true, ""
	case *FloatValue:
		return NewBoolValue(f.InternalValue >= o.InternalValue), true, ""
	default:
		return nil, false, "no se puede comparar float con " + other.Type()
	}
}

// Operaciones de asignación compuesta
func (f *FloatValue) AddAssign(other IVOR) (IVOR, bool, string) {
	return f.Add(other)
}

func (f *FloatValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return f.Subtract(other)
}

func (f *FloatValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return f.Multiply(other)
}

func (f *FloatValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return f.Divide(other)
}

func (f *FloatValue) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación módulo no soportada para floats"
}

// Operaciones lógicas (no aplicables)
func (f *FloatValue) And(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación lógica AND no soportada para floats"
}

func (f *FloatValue) Or(other IVOR) (IVOR, bool, string) {
	return nil, false, "operación lógica OR no soportada para floats"
}

func (f *FloatValue) Not() (IVOR, bool, string) {
	return nil, false, "operación lógica NOT no soportada para floats"
}
