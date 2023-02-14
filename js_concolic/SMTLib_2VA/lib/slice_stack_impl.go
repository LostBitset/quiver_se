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
	if len(stack.backing) == 0 {
		panic("Stack underflow")
	}
	var zero A
	stack.backing[len(stack.backing)-1] = zero // Reclaim memory
	stack.backing = stack.backing[:len(stack.backing)-1]
}

func (stack SliceStack[A]) Peek() (val A) {
	val = stack.backing[len(stack.backing)-1]
	return
}

func (stack SliceStack[A]) Length() (length int) {
	length = len(stack.backing)
	return
}

func (stack SliceStack[A]) Index(index int) (val A) {
	val = stack.backing[index]
	return
}
