package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie[int, int]()
	assert.Nil(
		t,
		trie.Lookup([]int{0, 0, 7}),
	)
}
