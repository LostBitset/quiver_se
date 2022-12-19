package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateExampleTrie() (trie Trie[int, int], entries []TrieEntry[int, int]) {
	trie = NewTrie[int, int]()
	entries = []TrieEntry[int, int]{
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
			21,
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
	return
}

func TestTrie(t *testing.T) {
	trie, entries := CreateExampleTrie()
	assert.ElementsMatch(t, trie.EntryList(), entries)
}

func TestTrieLookup(t *testing.T) {
	trie, entries := CreateExampleTrie()
	for _, entry := range entries {
		assert.Equal(
			t,
			entry.value,
			*trie.Lookup(entry.key),
		)
	}
	assert.Nil(t, trie.Lookup(make(map[int]struct{})))
	assert.Nil(t, trie.Lookup(map[int]struct{}{
		0: {}, 1: {}, 443: {},
	}))
}

func TestTrieLookupLeaf(t *testing.T) {
	trie, entries := CreateExampleTrie()
	for _, entry := range entries {
		assert.Equal(
			t,
			[]map[int]struct{}{
				entry.key,
			},
			trie.LookupLeaf(entry.value),
		)
	}
}

func TestTrieLookupInvalid(t *testing.T) {
	trie, _ := CreateExampleTrie()
	assert.Nil(
		t,
		trie.Lookup(map[int]struct{}{
			0: {}, 1777: {},
		}),
	)
	assert.Nil(
		t,
		trie.LookupLeaf(444444),
	)
}

func TestTrieLookupLeafDuplicates(t *testing.T) {
	trie := NewTrie[int, int]()
	keys := []map[int]struct{}{
		{0: {}, 1: {}, 7: {}},
		{0: {}, 1: {}, 9: {}},
	}
	for _, key := range keys {
		trie.Insert(key, 77)
	}
	assert.Equal(
		t,
		keys,
		trie.LookupLeaf(77),
	)
}
