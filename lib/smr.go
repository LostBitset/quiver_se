package qse

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
	in_canidates chan []IdLiteral[ATOM]
	out_models   chan []IdLiteral[ATOM]
	sys          SYS
}
