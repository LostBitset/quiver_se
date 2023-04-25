package main

import (
	"encoding/json"
	"fmt"
)

func (sp SeirPrgm) PerformQuery(assignment []AssignedSMTValue) (
	spc []PathCondSegment,
	fails bool,
) {
	respjson := sp.PerformQueryJson(assignment)
	spc, fails = DecodeSeirQueryResp(respjson)
	return
}

func DecodeSeirQueryResp(respjson []byte) (spc []PathCondSegment, fails bool) {
	var deser DseOutSpcObject
	err := json.Unmarshal(respjson, &deser)
	if err != nil {
		fmt.Println("JSON UNMARSHAL RETURNED ERROR, dump below")
		fmt.Println(string(respjson))
		panic(err)
	}
	if deser.Langs.Smt != "smtlib_2va" {
		panic("Only smtlib_2va is supported as returned smt language, got " + deser.Langs.Smt)
	}
	spc = deser.Spc
	fails = false
	for _, segm := range spc {
		if segm.IsAtFailure() {
			fails = true
		}
	}
	return
}

const SEIR_RESERVED_FAILURE_EVENT = "__seirevr_FAIL"

func (segm PathCondSegment) IsAtFailure() (at bool) {
	if segm.Event == nil {
		at = false
		return
	}
	at = *segm.Event == SEIR_RESERVED_FAILURE_EVENT
	return
}
