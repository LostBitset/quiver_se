package qse

// This is just a wrapper around the elves/elvish implementation

import (
	"src.elv.sh/pkg/persistent/hashmap"
)

// Effectively does the same thing as in Rust
// Unused type parameters are *usually* a mistake
// so it's nice to be explicit
type PhantomData[T any] struct {}

// A persistent hash map
// Struct embedding would just make things confusing because
// this is all pre-generics
type PHashMap[K comparable, V any] struct {
	inner         hashmap.Map
	phantom_key   PhantomData[K]
	phantom_value PhantomData[V]
}

