package smtlib2va

type LexicallyScoped struct {
	stack SliceStack[[]VarSlot]
	names map[string]LexicallyScopedIndex
}

type LexicallyScopedIndex struct {
	stack_index int
	frame_index int
}
