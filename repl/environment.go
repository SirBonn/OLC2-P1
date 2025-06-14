package repl

type Environment struct {
	values map[string]interface{}
	parent *Environment
}

func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		values: make(map[string]interface{}),
		parent: parent,
	}
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) Get(name string) (interface{}, bool) {
	value, exists := e.values[name]
	if !exists && e.parent != nil {
		return e.parent.Get(name)
	}
	return value, exists
}

func (e *Environment) Set(name string, value interface{}) bool {
	if _, exists := e.values[name]; exists {
		e.values[name] = value
		return true
	}

	if e.parent != nil {
		return e.parent.Set(name, value)
	}

	return false
}
