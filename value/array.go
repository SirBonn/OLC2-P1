package value

type IVORArray struct {
	BaseValue // Embed BaseValue for default behaviors
	Elements  []IVOR
}

// Constructor
func NewArray(elements []IVOR) *IVORArray {
	return &IVORArray{
		BaseValue: BaseValue{valueType: IVOR_VECTOR},
		Elements:  elements,
	}
}

// Basic methods
func (a *IVORArray) Value() interface{} {
	return a.Elements
}

func (a *IVORArray) Copy() IVOR {
	newElements := make([]IVOR, len(a.Elements))
	for i, elem := range a.Elements {
		if elem != nil {
			newElements[i] = elem.Copy()
		}
	}
	return NewArray(newElements)
}

// // Array-specific operations
// func (a *IVORArray) Length() IVOR {
//     return NewIntValue(int64(len(a.Elements)))
// }

func (a *IVORArray) Get(index int) (IVOR, bool) {
	if index < 0 || index >= len(a.Elements) {
		return nil, false
	}
	return a.Elements[index], true
}

func (a *IVORArray) Set(index int, value IVOR) bool {
	if index < 0 || index >= len(a.Elements) {
		return false
	}
	a.Elements[index] = value
	return true
}

func (a *IVORArray) Append(value IVOR) {
	a.Elements = append(a.Elements, value)
}

// Operation implementations
func (a *IVORArray) Add(other IVOR) (IVOR, bool, string) {
	// Array concatenation
	if otherArr, ok := other.(*IVORArray); ok {
		newElements := make([]IVOR, len(a.Elements)+len(otherArr.Elements))
		copy(newElements, a.Elements)
		copy(newElements[len(a.Elements):], otherArr.Elements)
		return NewArray(newElements), true, ""
	}
	return nil, false, "cannot concatenate array with " + other.Type()
}

// func (a *IVORArray) Equals(other IVOR) (IVOR, bool, string) {
//     if otherArr, ok := other.(*IVORArray); ok {
//         if len(a.Elements) != len(otherArr.Elements) {
//             return NewBoolValue(false), true, ""
//         }
//         for i := range a.Elements {
//             eq, ok, _ := a.Elements[i].Equals(otherArr.Elements[i])
//             if !ok || !eq.(*BoolValue).InternalValue {
//                 return NewBoolValue(false), true, ""
//             }
//         }
//         return NewBoolValue(true), true, ""
//     }
//     return NewBoolValue(false), true, ""
// }

// func (a *IVORArray) NotEquals(other IVOR) (IVOR, bool, string) {
//     res, ok, msg := a.Equals(other)
//     if !ok {
//         return nil, false, msg
//     }
//     return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
// }

// Unsupported operations for arrays
func (a *IVORArray) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '-' not supported for arrays"
}

func (a *IVORArray) Multiply(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '*' not supported for arrays"
}

func (a *IVORArray) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '/' not supported for arrays"
}

func (a *IVORArray) Mod(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '%' not supported for arrays"
}

func (a *IVORArray) LessThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '<' not supported for arrays"
}

func (a *IVORArray) LessOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '<=' not supported for arrays"
}

func (a *IVORArray) GreaterThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '>' not supported for arrays"
}

func (a *IVORArray) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '>=' not supported for arrays"
}

func (a *IVORArray) And(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '&&' not supported for arrays"
}

func (a *IVORArray) Or(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '||' not supported for arrays"
}

func (a *IVORArray) Not() (IVOR, bool, string) {
	return NewBoolValue(len(a.Elements) == 0), true, ""
}

func (a *IVORArray) AddAssign(other IVOR) (IVOR, bool, string) {
	if otherArr, ok := other.(*IVORArray); ok {
		a.Elements = append(a.Elements, otherArr.Elements...)
		return a, true, ""
	}
	return nil, false, "cannot concatenate array with " + other.Type()
}

func (a *IVORArray) SubAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '-=' not supported for arrays"
}

func (a *IVORArray) MulAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '*=' not supported for arrays"
}

func (a *IVORArray) DivAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '/=' not supported for arrays"
}

func (a *IVORArray) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '%=' not supported for arrays"
}
