package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieAsReversibleAssoc(t *testing.T) {
	var ra ReversibleAssoc[map[int]struct{}, int]
	trie := NewTrie[int, int]()
	ra = &trie
	test_key := map[int]struct{}{
		7: {}, 1: {},
	}
	ra.Insert(test_key, 42)
	assert.Equal(
		t,
		42,
		*ra.FwdLookup(test_key),
	)
	assert.Equal(
		t,
		[]map[int]struct{}{
			test_key,
		},
		ra.RevLookup(42),
	)
}
