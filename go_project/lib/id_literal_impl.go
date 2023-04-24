package qse

func (idsrc *IdSource) Gen() (id NumericId) {
	id = idsrc.next_id
	idsrc.next_id++
	return
}

func (wi WithId_H[T]) Hash() (digest digest_t) {
	digest = uint32_H{wi.Id}.Hash()
	return
}

func (wi WithId_H[T]) Hash32() (fixed_digest uint32) {
	fixed_digest = uint32_H{wi.Id}.Hash32()
	return
}

func (wi WithId_H[T]) GeneralDeref() (val T) {
	val = wi.Value
	return
}

func (lit IdLiteral[ATOM]) Hash() (digest digest_t) {
	digest = lit.Value.Hash()
	if !lit.Eq {
		digest = WrapInvert(digest)
	}
	return
}

func (lit IdLiteral[ATOM]) Hash32() (fixed_digest uint32) {
	fixed_digest = lit.Value.Hash32()
	if !lit.Eq {
		fixed_digest = WrapInvert32(fixed_digest)
	}
	return
}
