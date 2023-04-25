package main

func (sp SeirPrgm) PerformQuery(smt []AssignedSMTValue) (jsonresp []byte) {
	query := sp.MakeQueryJson(smt)
	jsonresp = PerformSeirQuery(query)
	return
}

func PerformSeirQuery(query []byte) (resp []byte) {
	panic("TODO TODO TODO")
}
