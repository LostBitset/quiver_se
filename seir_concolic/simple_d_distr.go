package libsynthetic

type SimpleDDistr[T comparable] struct {
	Outcomes map[T]float64
}
