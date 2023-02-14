package smtlib2va

// lvbls = lo vanbi be lo snicne = the environment of the variables

func NewLexicallyScoped() (lvbls LexicallyScoped) {
	lvbls = LexicallyScoped{
		NewSliceStack[[]Var](),
		make(map[string]*SliceStack[LexicallyScopedIndex]),
	}
	return
}

func (lvbls *LexicallyScoped) EnterScope() {
	lvbls.stack.Push(make([]Var, 0))
}

func (lvbls *LexicallyScoped) LeaveScope() {
	vars := lvbls.stack.Peek()
	for _, v := range vars {
		if lvbls.names[v.name].Length() < 2 {
			delete(lvbls.names, v.name)
		} else {
			lvbls.names[v.name].SilentPop()
		}
	}
	lvbls.stack.SilentPop()
}
