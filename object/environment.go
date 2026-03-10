package object

type Environment struct {
	store            map[string]Object
	outerEnvironment *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outerEnvironment: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outerEnvironment != nil {
		obj, ok = e.outerEnvironment.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outerEnvironment *Environment) *Environment {
	env := NewEnvironment()
	env.outerEnvironment = outerEnvironment
	return env
}
