package qse

func (idsrc *IdSource) Gen() (id NumericId) {
	id = idsrc.next_id
	idsrc.next_id++
	return
}

func (wi WithId_H[T]) Hash() (digest digest_t) {
	digest = uint32_H{wi.id}.Hash()
	return
}

func (wi WithId_H[T]) Hash32() (fixed_digest uint32) {
	fixed_digest = uint32_H{wi.id}.Hash32()
	return
}

func (wi WithId_H[T]) GeneralDeref() (val T) {
	val = wi.value
	return
}
