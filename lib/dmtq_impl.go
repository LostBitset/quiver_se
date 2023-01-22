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
