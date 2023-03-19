package main

import (
	"strconv"
)

func GetStandardItemLists() (ops []Op, vals []Val) {
	ops = []Op{
		{"+", []Sort{RealSort, RealSort}, RealSort},
		{"-", []Sort{RealSort, RealSort}, RealSort},
		{"and", []Sort{BoolSort, BoolSort}, BoolSort},
		{"or", []Sort{BoolSort, BoolSort}, BoolSort},
		{"not", []Sort{BoolSort}, BoolSort},
		{"=", []Sort{RealSort, RealSort}, BoolSort},
		{"<", []Sort{RealSort, RealSort}, BoolSort},
		{">", []Sort{RealSort, RealSort}, BoolSort},
		{"ite", []Sort{BoolSort, RealSort, RealSort}, RealSort}, // Generic but it doesn't matter
	}
	vals = make([]Val, 0)
	for i := -10; i <= 10; i++ {
		vals = append(vals, Val{strconv.Itoa(i), RealSort})
	}
	vals = append(vals, []Val{
		{"true", BoolSort},
		{"false", BoolSort},
	}...)
	return
}

func IndexOps(ops_list []Op) (ops map[Sort][]Op) {
	ops = make(map[Sort][]Op)
	for _, op := range ops_list {
		if _, ok := ops[op.ret]; !ok {
			ops[op.ret] = make([]Op, 1)
			ops[op.ret][0] = op
		} else {
			ops[op.ret] = append(ops[op.ret], op)
		}
	}
	return
}

func IndexVals(vals_list []Val) (vals map[Sort][]Val) {
	vals = make(map[Sort][]Val)
	for _, op := range vals_list {
		if _, ok := vals[op.sort]; !ok {
			vals[op.sort] = make([]Val, 1)
			vals[op.sort][0] = op
		} else {
			vals[op.sort] = append(vals[op.sort], op)
		}
	}
	return
}

func GetStandardItems() (ops map[Sort][]Op, vals map[Sort][]Val) {
	ops_list, vals_list := GetStandardItemLists()
	ops = IndexOps(ops_list)
	vals = IndexVals(vals_list)
	return
}

func (g Generator) Verify() {
	for k, v := range g.ops {
		for _, i := range v {
			if k != i.ret {
				panic("Incorrect sort used as key for generator operator.")
			}
		}
	}
	for k, v := range g.vals {
		for _, i := range v {
			if k != i.sort {
				panic("Incorrect sort used as key for generator value.")
			}
		}
	}
}

func (g Generator) AddVariables(n int, sort_distr DDistr[Sort]) {
	for i := 0; i < n; i++ {
		g.AddVariable(sort_distr)
	}
}

func (g Generator) AddVariable(sort_distr DDistr[Sort]) {
	sort := sort_distr.Sample()
	id := g.next_var_id
	g.next_var_id++
	val := Val{"var_" + strconv.Itoa(id), sort}
	if _, ok := g.vals[sort]; !ok {
		g.vals[sort] = make([]Val, 1)
		g.vals[sort][0] = val
	} else {
		g.vals[sort] = append(g.vals[sort], val)
	}
}
