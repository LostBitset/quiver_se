package qse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie[int, int]()
	trie.Insert([]int{0, 0, 7}, 44)
	fmt.Printf("trie: %v\n", trie)
	assert.Equal(
		t,
		44,
		trie.Lookup([]int{0, 0, 7}),
	)
}
