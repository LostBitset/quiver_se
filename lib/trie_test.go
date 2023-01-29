package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateExampleTrie() (trie Trie[uint32_H, int, struct{}], entries []TrieEntry[uint32_H, int]) {
	trie = NewTrie[uint32_H, int]()
	entries = []TrieEntry[uint32_H, int]{
		{
			map[uint32_H]struct{}{{0}: {}, {1}: {}, {7}: {}},
			44,
		},
		{
			map[uint32_H]struct{}{{0}: {}, {1}: {}, {9}: {}},
			55,
		},
		{
			map[uint32_H]struct{}{{0}: {}, {1}: {}, {2}: {}, {9}: {}},
			12,
		},
		{
			map[uint32_H]struct{}{{0}: {}, {1}: {}, {2}: {}, {3}: {}},
			21,
		},
		{
			map[uint32_H]struct{}{{9}: {}},
			99,
		},
		{
			map[uint32_H]struct{}{{0}: {}, {5}: {}, {6}: {}, {7}: {}, {4}: {}},
			31,
		},
		{
			map[uint32_H]struct{}{{0}: {}, {5}: {}, {6}: {}, {8}: {}, {4}: {}},
			32,
		},
	}
	for _, entry := range entries {
		trie.Insert(StdlibMapToPHashMap(entry.key), entry.value)
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
			*trie.Lookup(StdlibMapToPHashMap(entry.key)),
		)
	}
	assert.Nil(t, trie.Lookup(NewPHashMap[uint32_H, struct{}]()))
	assert.Nil(t, trie.Lookup(StdlibMapToPHashMap(
		map[uint32_H]struct{}{
			{0}: {}, {1}: {}, {443}: {},
		},
	)))
}

func TestTrieLookupLeaf(t *testing.T) {
	trie, entries := CreateExampleTrie()
	for _, entry := range entries {
		assert.Equal(
			t,
			[]map[uint32_H]struct{}{
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
		trie.Lookup(
			StdlibMapToPHashMap(
				map[uint32_H]struct{}{
					{0}: {}, {1777}: {},
				},
			),
		),
	)
	assert.Nil(
		t,
		trie.LookupLeaf(444444),
	)
}

func TestTrieLookupLeafDuplicates(t *testing.T) {
	trie := NewTrie[uint32_H, int]()
	keys := []map[uint32_H]struct{}{
		{{0}: {}, {1}: {}, {7}: {}},
		{{0}: {}, {1}: {}, {9}: {}},
	}
	for _, key := range keys {
		trie.Insert(StdlibMapToPHashMap(key), 77)
	}
	assert.Equal(
		t,
		keys,
		trie.LookupLeaf(77),
	)
}
