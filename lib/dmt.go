package qse

// A trie representing mappings from clauses of a disjunctive normal form (DNF), with literals of
// type LIT, to values of type LEAF. The Merkle component comes from the fact that the hashes of
// child subtries are kept and recomputed as necessary in order to make it fast to simplify clause
// pairs of the form (A && B) || (A && !B) to just A.
type DnfMerkleTrie[LIT comparable, LEAF comparable] struct {
	backing_trie MerkleTrie[LIT, LEAF]
}
