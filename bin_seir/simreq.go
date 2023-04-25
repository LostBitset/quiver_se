package main

import (
	qse "github.com/LostBitset/quiver_se/lib"
)

type SiMReQConstrainedTrxn struct {
	src_event   string
	dst_event   string
	constraints []qse.IdLiteral[string]
}
