package value

type IVORArray struct {
	Elements []IVOR
}

func (a *IVORArray) Value() interface{} {
	return a.Elements
}

func (a *IVORArray) Type() string {
	return "array"
}

// Agregar el método Copy()
func (a *IVORArray) Copy() IVOR {
	// Crear una copia profunda del array
	newElements := make([]IVOR, len(a.Elements))
	for i, elem := range a.Elements {
		newElements[i] = elem.Copy()
	}
	return &IVORArray{Elements: newElements}
}

func NewArray(elements []IVOR) IVOR {
	return &IVORArray{Elements: elements}
}
