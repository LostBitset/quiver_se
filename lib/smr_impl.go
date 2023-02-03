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
	unfinished = &backing_nocopy
	return
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) Start() {
	// TODO
	go func() {}()
}
