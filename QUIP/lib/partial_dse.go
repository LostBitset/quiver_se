package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	q "LostBitset/quiver_se/lib"
	"os"
)

const PARTIAL_DSE_TEMP_PATTERN = "tmp_QSE_quip-partial-dse_*.js"

func PerformPartialDse(
	cb eidin.CallbackId,
	target string,
	segment_chan chan q.Augmented[eidin.PathConditionSegment, []q.SMTFreeFun[string, string]],
) {
	original := ExtractCallback(cb, target)
	complete := PreparePartialDse(cb, original)
	f, err_temp := os.CreateTemp("/tmp", PARTIAL_DSE_TEMP_PATTERN)
	if err_temp != nil {
		panic(err_temp)
	}
	f.WriteString(complete)
	location := "/tmp/" + f.Name()
	f.Close()
	defer os.Remove(location)
	pc_chan := make(chan eidin.PathCondition)
	go PerformDse(location, GetMessagePrefix(location), pc_chan)
	for pc := range pc_chan {
		segment := pc.GetSegmentedPc()[0]
		segment_chan <- GeneralizePartialDseSegment(*segment, cb)
	}
}
