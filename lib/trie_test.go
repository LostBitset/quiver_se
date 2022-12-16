package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie[int, int]()
	entries := []TrieEntry[int, int]{
		{
			map[int]struct{}{0: {}, 1: {}, 7: {}},
			44,
		},
		{
			map[int]struct{}{0: {}, 1: {}, 9: {}},
			55,
		},
		{
			map[int]struct{}{0: {}, 1: {}, 2: {}, 9: {}},
			12,
		},
		{
			map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}},
			12,
		},
		{
			map[int]struct{}{9: {}},
			99,
		},
		{
			map[int]struct{}{0: {}, 5: {}, 6: {}, 7: {}, 4: {}},
			31,
		},
		{
			map[int]struct{}{0: {}, 5: {}, 6: {}, 8: {}, 4: {}},
			32,
		},
	}
	for _, entry := range entries {
		trie.Insert(entry.key, entry.value)
	}
	assert.ElementsMatch(t, trie.EntryList(), entries)
}
