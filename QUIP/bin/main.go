package main

import (
	quip "LostBitset/quiver_se/QUIP/lib"
	qse "LostBitset/quiver_se/lib"
)

func main() {
	var idsrc qse.IdSource
	sys := qse.SMTLib2VAStringSystem{idsrc}
	in_updates
	qse.StartSiMReQ(updates, out_models, sys)
	quip.StartQUIP(updates, top_node, fail_node, target, msg_prefix)
}
