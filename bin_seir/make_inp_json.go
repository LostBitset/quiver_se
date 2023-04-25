package main

import "encoding/json"

func (sp SeirPrgm) MakeQuery(smt []AssignedSMTValue) (obj any) {
	vars := make([]map[string]any, 0)
	for _, av := range smt {
		if len(av.smt_free_fun.Args) > 0 {
			panic("cannot make queries involving smt functions")
		}
		smt_name := av.smt_free_fun.Name
		new_var := map[string]any{
			"smt_name":       smt_name,
			"source_name":    sp.names_source_symb(smt_name),
			"sort":           av.smt_free_fun.Ret,
			"assigned_value": av.value_repr,
		}
		vars = append(vars, new_var)
	}
	obj = map[string]any{
		"languages": map[string]any{
			"source": "seir",
			"smt":    "smtlib_2va",
		},
		"source": sp.source,
		"vars":   vars,
	}
	return
}

func (sp SeirPrgm) MakeQueryJson(smt []AssignedSMTValue) (jsonstr []byte) {
	json_string, err := json.Marshal(sp.MakeQuery(smt))
	if err != nil {
		panic(err)
	}
	jsonstr = json_string
	return
}
