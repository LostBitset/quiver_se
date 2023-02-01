package qse

import "sync"

func (warden_config DMTQWardenConfig[N, ATOM]) Start() {
	go func() {
		var wg sync.WaitGroup
		for update := range warden_config.in_updates {
			out_walks_specific := make(chan QuiverWalk[N, PHashMap[Literal[ATOM], struct{}]])
			warden_config.dmtq.ApplyUpdateAndEmitWalks(
				out_walks_specific,
				update,
				warden_config.dmtq.ParameterizeIndex(
					warden_config.walk_src,
				),
			)
			wg.Add(1)
			go func() {
				defer wg.Done()
				for walk := range out_walks_specific {
					warden_config.out_walks <- walk
				}
			}()
		}
		go func() {
			defer close(warden_config.out_walks)
			wg.Wait()
		}()
	}()
}
