package main

func Pto[T any](x T) (p *T) {
	y := x
	p = &y
	return
}
