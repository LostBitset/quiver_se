package main

import "strconv"

func GetStandardItems() (ops []Op, vals []Val) {
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

func (g Generator) AddVariables(n int, sort_distr DDistr[Sort]) {
	for i := 0; i < n; i++ {
		g.AddVariable(sort_distr)
	}
}

func (g Generator) AddVariable(sort_distr DDistr[Sort]) {

}
