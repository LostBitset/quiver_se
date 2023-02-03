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
	in_canidates chan PHashMap[IdLiteral[ATOM], struct{}]
	out_models   chan PHashMap[IdLiteral[ATOM], struct{}]
	sys          SYS
}
