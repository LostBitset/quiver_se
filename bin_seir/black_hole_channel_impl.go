package main

func BlackHoleChannel[T any]() (bh chan T) {
	bh = make(chan T)
	go func(bh chan T) {
		for range bh {
		}
	}(bh)
	return
}
