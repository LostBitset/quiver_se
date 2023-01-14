package qse

type Literal[NODE comparable] struct {
	value NODE
	eq bool
}

func BufferingLiteral[NODE comparable](value NODE) (lit Literal[NODE]) {
	lit = Literal[NODE]{value, true}
	return
}

func InvertingLiteral[NODE comparable](value NODE) (lit Literal[NODE]) {
	lit = Literal[NODE]{value, false}
	return
}

func (lit Literal[NODE]) Equal()

type DMT[NODE comparable, LEAF comparable] struct {
	trie Trie[Literal[NODE], LEAF]
}

func NewDMT[NODE comparable, LEAF comparable]() (t DMT[NODE, LEAF]) {
	t = DMT[NODE, LEAF]{
		NewTrie[Literal[NODE], LEAF](),
	}
	return
}

