package smtlib2va

func NewSliceStack[A any]() (stack SliceStack[A]) {
	return SliceStack[A]{
		make([]A, 0),
	}
}

func (stack *SliceStack[A]) Push(val A) {
	stack.backing = append(stack.backing, val)
}

func (stack *SliceStack[A]) SilentPop() {
	var zero A
	stack.backing[len(stack.backing)] = zero // Reclaim memory
	stack.backing = stack.backing[:len(stack.backing)-1]
}

func (stack SliceStack[A]) Peek() (val A) {
	val = stack.backing[len(stack.backing)-1]
	return
}
