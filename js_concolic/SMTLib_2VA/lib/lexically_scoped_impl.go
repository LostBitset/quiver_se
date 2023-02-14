package smtlib2va

// lvbls = lo vanbi be lo snicne = the environment of the variables

func NewLexicallyScoped() (lvbls LexicallyScoped) {
	lvbls = LexicallyScoped{
		NewSliceStack[[]Var](),
		make(map[string]LexicallyScopedIndex),
	}
	return
}
