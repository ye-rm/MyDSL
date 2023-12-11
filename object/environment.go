package object

// 创建‘子’环境，将其外层环境作为参数传入
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object //符号表
	outer *Environment      //外层环境
}

// 从当前环境中查找符号,如果没有找到,则递归的向外层查找
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		//递归的向外层查找符号
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// 向当前环境中添加符号
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
