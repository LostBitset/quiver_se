package qse

func (warden_config DMTQWardenConfig[N, ATOM]) Start() {
	go func() {
		defer close(warden_config.out_walks)
		for update := range warden_config.in_updates {
			warden_config.dmtq.ApplyUpdateAndEmitWalks(
				warden_config.out_walks,
				update,
				warden_config.dmtq.ParameterizeIndex(
					warden_config.walk_src,
				),
			)
		}
	}()
}
