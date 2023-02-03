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

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) Start() {
	go func() {
		for canidate := range smr_config.in_canidates {
			smr_config.unfinished.Append(canidate)
		}
	}()
	go func() {
		defer close(smr_config.out_models)
		smr_config.RunSMR()
	}()
}

func (unfinished *TrustingNoCopySMRUnfinishedArray[ATOM]) Append(elems ...map[NumericId]IdLiteral[ATOM]) {
	unfinished.mu.Lock()
	defer unfinished.mu.Unlock()
	unfinished.arr = append(unfinished.arr, elems...)
}
