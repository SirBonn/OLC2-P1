package value

type NilValue struct {
	BaseValue // Embed BaseValue for default behaviors
}

func (n *NilValue) Value() interface{} {
	return nil
}

func (n *NilValue) Copy() IVOR {
	return DefaultNilValue
}

// Operation implementations for NilValue
func (n *NilValue) Add(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) Multiply(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) Mod(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) Equals(other IVOR) (IVOR, bool, string) {
	_, isNil := other.(*NilValue)
	_, isUninit := other.(*UnInitializedValue)
	return NewBoolValue(isNil || isUninit), true, ""
}

func (n *NilValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := n.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

func (n *NilValue) LessThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare nil"
}

func (n *NilValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare nil"
}

func (n *NilValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare nil"
}

func (n *NilValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare nil"
}

func (n *NilValue) And(other IVOR) (IVOR, bool, string) {
	return NewBoolValue(false), true, "" // nil is falsy
}

func (n *NilValue) Or(other IVOR) (IVOR, bool, string) {
	if b, ok := other.(*BoolValue); ok {
		return NewBoolValue(b.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "" // nil OR non-bool is false
}

func (n *NilValue) Not() (IVOR, bool, string) {
	return NewBoolValue(true), true, "" // !nil is true
}

func (n *NilValue) AddAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

func (n *NilValue) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on nil"
}

var DefaultNilValue = &NilValue{BaseValue: BaseValue{valueType: IVOR_NIL}}

type UnInitializedValue struct {
	BaseValue // Embed BaseValue for default behaviors
}

func (u *UnInitializedValue) Value() interface{} {
	return nil
}

func (u *UnInitializedValue) Copy() IVOR {
	return DefaultUnInitializedValue
}

// Operation implementations for UnInitializedValue
func (u *UnInitializedValue) Add(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) Multiply(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) Mod(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) Equals(other IVOR) (IVOR, bool, string) {
	_, isUninit := other.(*UnInitializedValue)
	return NewBoolValue(isUninit), true, ""
}

func (u *UnInitializedValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := u.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

func (u *UnInitializedValue) LessThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare uninitialized value"
}

func (u *UnInitializedValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare uninitialized value"
}

func (u *UnInitializedValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare uninitialized value"
}

func (u *UnInitializedValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot compare uninitialized value"
}

func (u *UnInitializedValue) And(other IVOR) (IVOR, bool, string) {
	return NewBoolValue(false), true, "" // uninitialized is falsy
}

func (u *UnInitializedValue) Or(other IVOR) (IVOR, bool, string) {
	if b, ok := other.(*BoolValue); ok {
		return NewBoolValue(b.InternalValue), true, ""
	}
	return NewBoolValue(false), true, "" // uninitialized OR non-bool is false
}

func (u *UnInitializedValue) Not() (IVOR, bool, string) {
	return NewBoolValue(true), true, "" // !uninitialized is true
}

func (u *UnInitializedValue) AddAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

func (u *UnInitializedValue) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "cannot perform operations on uninitialized value"
}

var DefaultUnInitializedValue = &UnInitializedValue{BaseValue: BaseValue{valueType: IVOR_UNINITIALIZED}}
