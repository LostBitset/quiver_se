package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateExampleDMT() (dmt DMT[uint32_H, uint32_H], entries []TrieEntry[Literal[uint32_H], uint32_H]) {
	dmt = NewDMT[uint32_H, uint32_H]()
	entries = []TrieEntry[Literal[uint32_H], uint32_H]{
		{
			map[Literal[uint32_H]]struct{}{
				{uint32_H{14}, true}: {},
				{uint32_H{15}, true}: {},
				{uint32_H{16}, false}: {},
				{uint32_H{17}, false}: {},
			},
			uint32_H{1},
		},
		{
			map[Literal[uint32_H]]struct{}{
				{uint32_H{14}, true}: {},
				{uint32_H{15}, false}: {},
				{uint32_H{16}, false}: {},
				{uint32_H{17}, false}: {},
			},
			uint32_H{1},
		},
		{
			map[Literal[uint32_H]]struct{}{
				{uint32_H{14}, false}: {},
				{uint32_H{18}, true}: {},
				{uint32_H{19}, true}: {},
			},
			uint32_H{2},
		},
	}
	for _, entry := range entries {
		dmt.Insert(entry.key, entry.value)
	}
	return
}

func TestDMT(t *testing.T) {
	dmt, entries := CreateExampleDMT()
	assert.ElementsMatch(t, dmt.EntryList(), entries)
}
