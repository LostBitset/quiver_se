package qse

import "sync"

type SMRConfig[
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
] struct {
	in_canidates chan map[NumericId]IdLiteral[ATOM]
	out_models   chan map[NumericId]IdLiteral[ATOM]
	sys          SYS
	unfinished   SMRUnfinishedArray[ATOM]
}

type SMRUnfinishedArray[ATOM comparable] *TrustingNoCopySMRUnfinishedArray[ATOM]

type TrustingNoCopySMRUnfinishedArray[ATOM comparable] struct {
	arr []map[NumericId]IdLiteral[ATOM]
	mu  sync.Mutex
}
