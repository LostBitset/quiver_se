package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
)

const SUBROUTINE_DSE_TIMEOUT_MILLIS = 10000
const SUBROUTINE_DSE_CYCLE_WAIT_TIME_MILLIS = 100

// Should generally be called as a goroutine
func PerformDse(
	location string,
	msg_prefix string,
	single_callback_mode bool,
	pc_chan chan eidin.PathCondition,
) {
	go RunSimpleDSE(msg_prefix, SUBROUTINE_DSE_CYCLE_WAIT_TIME_MILLIS)
	go RunAnalyzer(location, msg_prefix)
	UsePathConditionChannel(msg_prefix, pc_chan)
}

func UsePathConditionChannel(msg_prefix string, pc_chan chan eidin.PathCondition) {
	// TODO
}
