package smtlib2va

// lvbls = lo vanbi be lo snicne = the environment of the variables

func NewLexicallyScoped() (lvbls LexicallyScoped) {
	lvbls = LexicallyScoped{
		NewSliceStack[[]Var](),
		make(map[string][]LexicallyScopedIndex),
	}
	return
}

func (lvbls LexicallyScoped) EnterScope() {
	lvbls.stack.Push(make([]Var, 0))
}

func (lvbls LexicallyScoped) LeaveScope() {
	vars := lvbls.stack.Peek()
	for _, v := range vars {
		delete(lvbls.names, v.name)
	}
}
