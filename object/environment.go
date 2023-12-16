// Package object environment implement symbol table
// This implement the symbol table we will use to store the bindings of identifiers to values.
package object

// NewEnclosedEnvironment add a new environment with outer environment
// when we evaluate a function, we create a new environment that is enclosed by the environment that was active when the function was created
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment create a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment struct
type Environment struct {
	store map[string]Object //symbol table
	outer *Environment      //outer environment
}

// Get symbol from current environment, if not found, get from outer environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		// recursively get symbol from outer environment
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set add symbol to current environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
