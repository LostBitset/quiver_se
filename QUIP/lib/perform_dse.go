package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
)

const SUBROUTINE_DSE_TIMEOUT_MILLIS = 10000

// Should generally be called as a goroutine
func PerformDse(
	location string,
	msg_prefix string,
	single_callback_mode bool,
	pc_chan chan eidin.PathCondition,
) {
	// TODO
}
