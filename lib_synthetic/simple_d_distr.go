package libsynthetic

type SimpleDDistr[T comparable] struct {
	outcomes map[T]float64
}
