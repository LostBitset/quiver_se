package qse

import (
	"fmt"

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

func (pm PHashMap[K, V]) ToStdlibMap() (m map[K]V) {
	m = make(map[K]V, pm.length)
	for itr := pm.inner.Iterator(); itr.HasElem(); itr.Next() {
		key_any, val_any := itr.Elem()
		key := key_any.(K)
		val := val_any.(V)
		m[key] = val
	}
	return
}

func StdlibMapToPHashMap[K hashable, V any](m map[K]V) (pm PHashMap[K, V]) {
	pm = NewPHashMap[K, V]()
	for k, v := range m {
		pm = pm.Assoc(k, v)
	}
	return
}

func (a PHashMap[K, V]) Equal(b PHashMap[K, V]) (eq bool) {
	fmt.Println("actually ran Equal")
	a_copy := PHashMap[K, V]{
		a.inner,
		a.length,
		PhantomData[K]{},
		PhantomData[V]{},
	}
	for itr := b.inner.Iterator(); itr.HasElem(); itr.Next() {
		b_key_any, b_val_any := itr.Elem()
		b_key := b_key_any.(K)
		b_val := b_val_any.(V)
		a_val, ok := a_copy.Index(b_key)
		if !ok || ComparableEqualFunc(a_val, b_val) {
			eq = false
			return
		}
		a_copy = a_copy.Dissoc(b_key)
	}
	if a_copy.length == 0 {
		eq = false
		return
	}
	eq = true
	return
}

