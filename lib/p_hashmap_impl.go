package qse

import (
	"src.elv.sh/pkg/persistent/hashmap"
)


// Actually requires k1, k2 comparable
func ComparableEqualFunc(k1, k2 any) (eq bool) {
	eq = k1 == k2
	return
}

// Actually requires k hashable
func HashableHashFunc(k any) (fixed_digest uint32) {
	fixed_digest = k.(unsafe_relaxed_hashable).Hash32()
	return
}

func NewPHashMap[K hashable, V any]() (pm PHashMap[K, V]) {
	pm = PHashMap[K, V]{
		hashmap.New(ComparableEqualFunc, HashableHashFunc),
		PhantomData[K]{},
		PhantomData[V]{},
	}
	return
}

