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
		0,
		PhantomData[K]{},
		PhantomData[V]{},
	}
	return
}

func (pm PHashMap[K, V]) Assoc(key K, val V) (updated PHashMap[K, V]) {
	updated = PHashMap[K, V]{
		pm.inner.Assoc(key, val),
		pm.length + 1,
		PhantomData[K]{},
		PhantomData[V]{},
	}
	return
}

func (pm PHashMap[K, V]) Dissoc(key K) (updated PHashMap[K, V]) {
	updated = PHashMap[K, V]{
		pm.inner.Dissoc(key),
		pm.length - 1,
		PhantomData[K]{},
		PhantomData[V]{},
	}
	return
}

func (pm PHashMap[K, V]) HasKey(key K) (has bool) {
	_, has = pm.inner.Index(key)
	return
}

func (pm PHashMap[K, V]) Index(key K) (val V, ok bool) {
	var val_raw any
	val_raw, ok = pm.inner.Index(key)
	val = val_raw.(V)
	return
}

