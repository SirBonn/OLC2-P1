package value

import (
	"strconv"
	"strings"
)

type StringValue struct {
	BaseValue     // Embed BaseValue for default behaviors
	InternalValue string
}

// Constructor
func NewStringValue(value string) *StringValue {
	return &StringValue{
		BaseValue:     BaseValue{valueType: IVOR_STRING},
		InternalValue: value,
	}
}

// Basic methods
func (s *StringValue) Value() interface{} {
	return s.InternalValue
}

func (s *StringValue) Copy() IVOR {
	return NewStringValue(s.InternalValue)
}

// String operations
func (s *StringValue) Add(other IVOR) (IVOR, bool, string) {
	// String concatenation
	switch o := other.(type) {
	case *StringValue:
		return NewStringValue(s.InternalValue + o.InternalValue), true, ""
	case *IntValue:
		return NewStringValue(s.InternalValue + strconv.Itoa(int(o.InternalValue))), true, ""
	case *FloatValue:
		return NewStringValue(s.InternalValue + strconv.FormatFloat(o.InternalValue, 'f', -1, 64)), true, ""
	case *BoolValue:
		return NewStringValue(s.InternalValue + strconv.FormatBool(o.InternalValue)), true, ""
	default:
		return nil, false, "cannot concatenate string with " + other.Type()
	}
}

// Comparison operations
func (s *StringValue) Equals(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*StringValue); ok {
		return NewBoolValue(s.InternalValue == o.InternalValue), true, ""
	}
	return NewBoolValue(false), true, ""
}

func (s *StringValue) NotEquals(other IVOR) (IVOR, bool, string) {
	res, ok, msg := s.Equals(other)
	if !ok {
		return nil, false, msg
	}
	return NewBoolValue(!res.(*BoolValue).InternalValue), true, ""
}

func (s *StringValue) LessThan(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*StringValue); ok {
		return NewBoolValue(s.InternalValue < o.InternalValue), true, ""
	}
	return nil, false, "cannot compare string with " + other.Type()
}

func (s *StringValue) LessOrEqual(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*StringValue); ok {
		return NewBoolValue(s.InternalValue <= o.InternalValue), true, ""
	}
	return nil, false, "cannot compare string with " + other.Type()
}

func (s *StringValue) GreaterThan(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*StringValue); ok {
		return NewBoolValue(s.InternalValue > o.InternalValue), true, ""
	}
	return nil, false, "cannot compare string with " + other.Type()
}

func (s *StringValue) GreaterOrEqual(other IVOR) (IVOR, bool, string) {
	if o, ok := other.(*StringValue); ok {
		return NewBoolValue(s.InternalValue >= o.InternalValue), true, ""
	}
	return nil, false, "cannot compare string with " + other.Type()
}

// Unsupported operations for strings
func (s *StringValue) Subtract(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '-' not supported for strings"
}

func (s *StringValue) Multiply(other IVOR) (IVOR, bool, string) {
	// Optional: Could implement string repetition if desired
	return nil, false, "operation '*' not supported for strings"
}

func (s *StringValue) Divide(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '/' not supported for strings"
}

func (s *StringValue) Mod(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '%' not supported for strings"
}

// Logical operations
func (s *StringValue) And(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '&&' not supported for strings"
}

func (s *StringValue) Or(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '||' not supported for strings"
}

func (s *StringValue) Not() (IVOR, bool, string) {
	return NewBoolValue(len(s.InternalValue) == 0), true, ""
}

// Compound assignment operations
func (s *StringValue) AddAssign(other IVOR) (IVOR, bool, string) {
	res, ok, msg := s.Add(other)
	if !ok {
		return nil, false, msg
	}
	if str, ok := res.(*StringValue); ok {
		s.InternalValue = str.InternalValue
		return s, true, ""
	}
	return nil, false, "invalid concatenation result"
}

func (s *StringValue) SubAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '-=' not supported for strings"
}

func (s *StringValue) MulAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '*=' not supported for strings"
}

func (s *StringValue) DivAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '/=' not supported for strings"
}

func (s *StringValue) ModAssign(other IVOR) (IVOR, bool, string) {
	return nil, false, "operation '%=' not supported for strings"
}

// Additional string-specific methods
// func (s *StringValue) Length() IVOR {
// 	return NewIntValue(int64(len(s.InternalValue)))
// }

func (s *StringValue) ToUpper() IVOR {
	return NewStringValue(strings.ToUpper(s.InternalValue))
}

func (s *StringValue) ToLower() IVOR {
	return NewStringValue(strings.ToLower(s.InternalValue))
}
