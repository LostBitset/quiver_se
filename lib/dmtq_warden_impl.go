package qse

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

func (warden_config DMTQWardenConfig[N, ATOM, AUG]) Start() {
	go func() {
		var wg sync.WaitGroup
		for update_augmented := range warden_config.in_updates {
			log.Info("[dmtq_warden/go1] Received (augmented) quiver update. ")
			update := update_augmented.Value
			augment := update_augmented.Augment
			out_walks_specific := make(chan QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]])
			log.Info("[dmtq_warden/go1] Applying update to quiver and emitting new walks. ")
			warden_config.dmtq.ApplyUpdateAndEmitWalks(
				out_walks_specific,
				update,
				warden_config.dmtq.ParameterizeIndex(
					warden_config.walk_src,
				),
				warden_config.walk_dst,
			)
			wg.Add(1)
			go func() {
				defer wg.Done()
				log.Info("[dmtq_warden/go1/go1] Listening for (raw) quiver walks. ")
				for walk := range out_walks_specific {
					warden_config.out_walks <- Augmented[
						QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]],
						AUG,
					]{
						walk, augment,
					}
				}
				log.Info("[dmtq_warden/go1/go1] Sent all (augmented) quiver walks. ")
			}()
		}
		go func() {
			defer close(warden_config.out_walks)
			wg.Wait()
		}()
	}()
}

func NewAugmentedSimple[A any](value A) (aug Augmented[A, struct{}]) {
	aug = Augmented[A, struct{}]{
		value, struct{}{},
	}
	return
}
