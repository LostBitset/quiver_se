package main

import (
	"bytes"
	"os"
	"os/exec"
)

func (sp SeirPrgm) PerformQueryJson(smt []AssignedSMTValue) (jsonresp []byte) {
	query := sp.MakeQueryJson(smt)
	jsonresp = PerformSeirQuery(query)
	return
}

func PerformSeirQuery(query []byte) (resp []byte) {
	f, errC := os.CreateTemp("/tmp", "qse-bin_seir-TEMP_SEIR_QUERY-*.json")
	if errC != nil {
		panic(errC)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	_, errW := f.Write(query)
	if errW != nil {
		panic(errW)
	}
	cmd := exec.Command("./seirspc.sh", f.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	errE := cmd.Run()
	if errE != nil {
		panic(errE)
	}
	resp = out.Bytes()
	return
}
