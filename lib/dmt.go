package qse

type hashable interface {
	Hash() (digest int)
	comparable
}

type Literal[NODE hashable] struct {
	value NODE
	eq bool
}

func BufferingLiteral[NODE hashable](value NODE) (lit Literal[NODE]) {
	lit = Literal[NODE]{value, true}
	return
}

func InvertingLiteral[NODE hashable](value NODE) (lit Literal[NODE]) {
	lit = Literal[NODE]{value, false}
	return
}

type DMT[NODE hashable, LEAF comparable] struct {
	trie Trie[Literal[NODE], LEAF]
}

func NewDMT[NODE hashable, LEAF comparable]() (t DMT[NODE, LEAF]) {
	t = DMT[NODE, LEAF]{
		NewTrie[Literal[NODE], LEAF](),
	}
	return
}
