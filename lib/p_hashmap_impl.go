package qse

import (
	"github.com/google/go-cmp/cmp"
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
	eq = cmp.Equal(a.ToStdlibMap(), b.ToStdlibMap())
	return
}

func (pm PHashMap[K, V]) Clone() (cloned PHashMap[K, V]) {
	cloned = PHashMap[K, V]{
		pm.inner,
		pm.length,
		PhantomData[K]{},
		PhantomData[V]{},
	}
	return
}

