package qse

import (
	"encoding/binary"
	"hash/fnv"
)

func (i QuiverIndex) Hash() (digest digest_t) {
	fixed_size := uint64(i)
	pseudo_digest := binary.LittleEndian.AppendUint64([]byte{}, fixed_size)
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{0xCF, 0x64, 0x6A})
	digest = hasher.Sum([]byte{})
	return
}

func (i QuiverIndex) Hash32() (fixed_digest uint32) {
	fixed_size := uint64(i)
	pseudo_digest := binary.LittleEndian.AppendUint64([]byte{}, fixed_size)
	hasher := fnv.New32a()
	hasher.Write(pseudo_digest)
	hasher.Write([]byte{0xCF, 0x64, 0x6A})
	fixed_digest = hasher.Sum32()
	return
}

