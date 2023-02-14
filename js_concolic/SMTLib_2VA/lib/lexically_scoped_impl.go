package smtlib2va

// lvbls = lo vanbi be lo snicne = the environment of the variables

func NewLexicallyScoped() (lvbls LexicallyScoped) {
	lvbls = LexicallyScoped{
		NewSliceStack[*[]Var](),
		make(map[string]*SliceStack[LexicallyScopedIndex]),
	}
	return
}

func (lvbls *LexicallyScoped) EnterScope() {
	backing := make([]Var, 0)
	lvbls.stack.Push(&backing)
}

func (lvbls *LexicallyScoped) LeaveScope() {
	vars := lvbls.stack.Peek()
	for _, v := range *vars {
		if lvbls.names[v.name].Length() < 2 {
			delete(lvbls.names, v.name)
		} else {
			lvbls.names[v.name].SilentPop()
		}
	}
	lvbls.stack.SilentPop()
}

func (lvlbs *LexicallyScoped) SetVar(name string, val *string) {
	scope_ref := lvlbs.stack.Peek()
	stack_index := lvlbs.stack.Length() - 1
	frame_index := len(*scope_ref)
	if val == nil {
		*scope_ref = append(*scope_ref, Var{name, NewVarSlot()})
	} else {
		slot := NewVarSlot()
		slot.Write(*val)
		*scope_ref = append(*scope_ref, Var{name, slot})
	}
	if stack_ref, ok := lvlbs.names[name]; ok {
		stack_ref.Push(
			LexicallyScopedIndex{
				stack_index,
				frame_index,
			},
		)
	}
}
