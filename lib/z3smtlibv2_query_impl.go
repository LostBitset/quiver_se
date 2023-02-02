package qse

func NewZ3SMTLibv2Query(query_str string) (query Z3SMTLibv2Query) {
	query = Z3SMTLibv2Query{query_str}
	return
}

func (query Z3SMTLibv2Query) Run() (output string) {
	// TODO
	return
}
