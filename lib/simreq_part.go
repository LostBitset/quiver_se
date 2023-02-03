package qse

type DSEWithCallbacksUpdate[
	QNODE any,
	ATOM comparable,
	IDENT any,
	SORT any,
] struct {
	WithSMTFreeFuns[
		IDENT, SORT,
		QuiverUpdate[
			QNODE, PHashMap[Literal[WithId_H[ATOM]], struct{}], *DMT[WithId_H[ATOM], QuiverIndex],
		],
	]
}

type WithSMTFreeFuns[T any, IDENT any, SORT any] struct {
	value     T
	free_funs []SMTFreeFun[IDENT, SORT]
}
