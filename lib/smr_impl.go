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
	in_canidates chan SMTQueryDNFClause[ATOM, IDENT, SORT],
	out_models chan MODEL,
	sys SYS,
) (
	smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS],
) {
	smr_config = SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]{
		in_canidates: in_canidates,
		out_models:   out_models,
		sys:          sys,
		unfinished:   NewSMRUnfinishedArray[ATOM, IDENT, SORT](),
	}
	return
}

func NewSMRUnfinishedArray[
	ATOM comparable,
	IDENT any,
	SORT any,
]() (unfinished SMRUnfinishedArray[ATOM, IDENT, SORT]) {
	backing_nocopy := TrustingNoCopySMRUnfinishedArray[ATOM, IDENT, SORT]{
		arr: make([]SMTQueryDNFClause[ATOM, IDENT, SORT], 0),
	}
	unfinished = SMRUnfinishedArray[ATOM, IDENT, SORT]{
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

func (unfinished *TrustingNoCopySMRUnfinishedArray[ATOM, IDENT, SORT]) Append(
	elems ...SMTQueryDNFClause[ATOM, IDENT, SORT],
) {
	unfinished.mu.Lock()
	defer unfinished.mu.Unlock()
	unfinished.arr = append(unfinished.arr, elems...)
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) RunSMR() (done bool) {
	smr_config.unfinished.mu.Lock()
	defer smr_config.unfinished.mu.Unlock()
	done = (len(smr_config.unfinished.arr) == 0)
	if done {
		return
	}
	smr_config.SMRIterationUnfinishedUnlocked()
	return
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) SMRIterationUnfinishedUnlocked() {
	finished := make([]int, 0)
	for i := range smr_config.unfinished.arr {
		elem := smr_config.unfinished.arr[i]
		sctx := smr_config.sys.CheckSat(
			elem.conjunction,
			elem.free_funs,
		)
		is_sat_ptr := sctx.IsSat()
		if is_sat_ptr == nil {
			finished = append(finished, i)
			continue
		}
		if *is_sat_ptr {
			smr_config.out_models <- *sctx.GetModel()
		} else {
			// TODO unsat case
		}
	}
	for _, index := range finished {
		SpliceOutReclaim(
			&smr_config.unfinished.arr,
			index,
		)
	}
}
