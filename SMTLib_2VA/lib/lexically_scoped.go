package smtlib2va

type LexicallyScoped struct {
	stack SliceStack[*[]Var]
	names map[string]*SliceStack[LexicallyScopedIndex]
}

type LexicallyScopedIndex struct {
	stack_index int
	frame_index int
}

type Var struct {
	name string
	slot VarSlot
}
