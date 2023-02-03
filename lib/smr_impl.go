package qse

func NewSMRConfig[
	ATOM comparable,
	IDENT any,
	SORT any,
	MODEL any,
	SCTX SMTSolvedContext[MODEL],
	SYS SMTSystem[
		IdLiteral[ATOM],
		IDENT,
		SORT,
		MODEL,
		SCTX,
	],
](
	in_canidates chan map[NumericId]IdLiteral[ATOM],
	out_models chan map[NumericId]IdLiteral[ATOM],
	sys SYS,
) (
	smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS],
) {
	smr_config = SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]{
		in_canidates: in_canidates,
		out_models:   out_models,
		sys:          sys,
		unfinished:   NewSMRUnfinishedArray[ATOM](),
	}
	return
}

func NewSMRUnfinishedArray[ATOM comparable]() (unfinished SMRUnfinishedArray[ATOM]) {
	backing_nocopy := TrustingNoCopySMRUnfinishedArray[ATOM]{
		arr: make([]map[uint32]IdLiteral[ATOM], 0),
	}
	unfinished = SMRUnfinishedArray[ATOM]{
		&backing_nocopy,
	}
	return
}

func NewSMRIsSleeping() (is_sleeping SMRIsSleeping) {
	backing_nocopy := TrustingNoCopySMRIsSleeping{
		is: false,
	}
	is_sleeping = SMRIsSleeping{
		&backing_nocopy,
	}
	return
}

func (is_sleeping SMRIsSleeping) Sleep() (was bool) {
	is_sleeping.mu.Lock()
	defer is_sleeping.mu.Unlock()
	was = is_sleeping.is
	is_sleeping.is = true
	return
}

func (is_sleeping SMRIsSleeping) Wake() (was bool) {
	is_sleeping.mu.Lock()
	defer is_sleeping.mu.Unlock()
	was = is_sleeping.is
	is_sleeping.is = false
	return
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) Start() {
	wakeup_chan := make(chan struct{})
	is_sleeping := NewSMRIsSleeping()
	go func() {
		for {
			for smr_config.RunSMR() {
			}
			is_sleeping.Sleep()
			<-wakeup_chan
		}
	}()
	go func() {
		defer close(wakeup_chan)
		for canidate := range smr_config.in_canidates {
			smr_config.unfinished.Append(canidate)
			if !is_sleeping.Wake() {
				wakeup_chan <- struct{}{}
			}
		}
	}()
}

func (unfinished *TrustingNoCopySMRUnfinishedArray[ATOM]) Append(
	elems ...map[NumericId]IdLiteral[ATOM],
) {
	unfinished.mu.Lock()
	defer unfinished.mu.Unlock()
	unfinished.arr = append(unfinished.arr, elems...)
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) RunSMR() (done bool) {
	// TODO
	smr_config.unfinished.mu.Lock()
	defer smr_config.unfinished.mu.Unlock()
	done
	return
}
