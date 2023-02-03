package qse

import (
	"fmt"
	"sync"
)

func (warden_config DMTQWardenConfig[N, ATOM, AUG]) Start() {
	go func() {
		fmt.Println("<bgn> dmtq warden")
		var wg sync.WaitGroup
		for update_augmented := range warden_config.in_updates {
			update := update_augmented.value
			augment := update_augmented.augment
			out_walks_specific := make(chan QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]])
			fmt.Println("dmtq warden emit bgn")
			warden_config.dmtq.ApplyUpdateAndEmitWalks(
				out_walks_specific,
				update,
				warden_config.dmtq.ParameterizeIndex(
					warden_config.walk_src,
				),
				warden_config.walk_dst,
			)
			fmt.Println("dmtq warden emit end")
			wg.Add(1)
			go func() {
				defer wg.Done()
				for walk := range out_walks_specific {
					warden_config.out_walks <- Augmented[
						QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]],
						AUG,
					]{
						walk, augment,
					}
				}
			}()
		}
		go func() {
			defer fmt.Println("<end> dmtq warden")
			defer close(warden_config.out_walks)
			wg.Wait()
		}()
		fmt.Println("for dmtq warden, in_updates channel closed")
	}()
}

func NewAugmentedSimple[A any](value A) (aug Augmented[A, struct{}]) {
	aug = Augmented[A, struct{}]{
		value, struct{}{},
	}
	return
}
