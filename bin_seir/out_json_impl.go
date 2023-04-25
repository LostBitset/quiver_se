package main

import (
	"encoding/json"
	"fmt"
)

func (sp SeirPrgm) PerformQuery(assignment []AssignedSMTValue) (spc []PathCondSegment) {
	respjson := sp.PerformQueryJson(assignment)
	spc = DecodeSeirQueryResp(respjson)
	return
}

func DecodeSeirQueryResp(respjson []byte) (spc []PathCondSegment) {
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
	return
}
