package libsynthetic

func pto[T any](x T) (p *T) {
	y := x
	p = &y
	return
}
