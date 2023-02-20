package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	q "LostBitset/quiver_se/lib"
	"fmt"
	"os"
	"regexp"
	"strings"
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
	InstrumentFunctionInfo(location)
	pc_chan := make(chan eidin.PathCondition)
	go PerformDse(location, GetMessagePrefix(location), pc_chan)
	for pc := range pc_chan {
		segment := pc.GetSegmentedPc()[0]
		segment_chan <- GeneralizePartialDseSegment(*segment, pc.GetFreeFuns())
	}
}

func ExtractCallback(cb eidin.CallbackId, target string) (extracted string) {
	start := int64(cb.GetBytesStart())
	end := int64(cb.GetBytesEnd())
	f, err := os.Open(target)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Seek(start, 0)
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, end-start)
	_, err = f.Read(buffer)
	if err != nil {
		panic(err)
	}
	extracted = string(buffer)
	return
}

func PreparePartialDse(cb eidin.CallbackId, function_source string) (full_source string) {
	var sb strings.Builder
	sb.WriteString(GenerateSymbolicPrelude(cb))
	sb.WriteRune('\n')
	sb.WriteString(ExtractJSFunctionBody(function_source))
	sb.WriteRune('\n')
	full_source = sb.String()
	return
}

func ExtractJSFunctionBody(function_original string) (body string) {
	function := strings.TrimSpace(function_original)
	prefix_re := regexp.MustCompile(`function(\s*|\s*\w+\s*)\([^\)]*\)\s*{`)
	leftmost_prefix_loc := prefix_re.FindStringIndex(function)
	body = function[leftmost_prefix_loc[1] : len(function)-1]
	return
}

func GenerateSymbolicPrelude(cb eidin.CallbackId) (prelude string) {
	var sb strings.Builder
	for _, free_fun_ref := range cb.GetUsedFreeFuns() {
		free_fun_ref := free_fun_ref
		free_fun := *free_fun_ref
		name := free_fun.GetName()
		sort := free_fun.GetRetSort()
		sb.WriteString(
			fmt.Sprintf(
				"var sym__outer__%s = \"outer__%s:%s\";\n",
				name, name, sort,
			),
		)
	}
	prelude = sb.String()
	return
}

func GeneralizePartialDseSegment(
	segment eidin.PathConditionSegment,
	free_funs []*eidin.SMTFreeFun,
) (
	general_segment q.Augmented[eidin.PathConditionSegment, []q.SMTFreeFun[string, string]],
) {
	// TODO
}
