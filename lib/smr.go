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
	in_canidates chan SMTQueryDNFClause[ATOM, IDENT, SORT]
	out_models   chan MODEL
	sys          SYS
	unfinished   SMRUnfinishedArray[ATOM, IDENT, SORT]
}

type SMTQueryDNFClause[
	ATOM comparable,
	IDENT any,
	SORT any,
] struct {
	conjunction []IdLiteral[ATOM]
	free_funs   []SMTFreeFun[IDENT, SORT]
}

type SMRUnfinishedArray[
	ATOM comparable,
	IDENT any,
	SORT any,
] struct {
	*TrustingNoCopySMRUnfinishedArray[ATOM, IDENT, SORT]
}

type TrustingNoCopySMRUnfinishedArray[
	ATOM comparable,
	IDENT any,
	SORT any,
] struct {
	arr []SMTQueryDNFClause[ATOM, IDENT, SORT]
	mu  sync.Mutex
}

type SMRIsSleeping struct {
	*TrustingNoCopySMRIsSleeping
}

type TrustingNoCopySMRIsSleeping struct {
	is bool
	mu sync.Mutex
}
