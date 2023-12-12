// environment implement symbol table
package object

// add a new environment with outer environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// create a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Environment struct
type Environment struct {
	store map[string]Object //符号表
	outer *Environment      //外层环境
}

// get symbol from current environment, if not found, get from outer environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		// recursively get symbol from outer environment
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// add symbol to current environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
