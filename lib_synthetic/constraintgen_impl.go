package libsynthetic

import (
	qse "github.com/LostBitset/quiver_se/lib"
	"math"
	"math/rand"

	"strconv"
	"strings"
)

func GetStandardItemLists() (ops []Op, vals []Val) {
	ops = []Op{
		{"+", []Sort{RealSort, RealSort}, RealSort},
		{"-", []Sort{RealSort, RealSort}, RealSort},
		{"and", []Sort{BoolSort, BoolSort}, BoolSort},
		{"not", []Sort{BoolSort}, BoolSort},
		{"=", []Sort{RealSort, RealSort}, BoolSort},
		{"<", []Sort{RealSort, RealSort}, BoolSort},
		{">", []Sort{RealSort, RealSort}, BoolSort},
		{"ite", []Sort{BoolSort, RealSort, RealSort}, RealSort}, // Generic but it doesn't matter
	}
	vals = make([]Val, 0)
	for i := -2; i <= 3; i++ {
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

func (g ConstraintGenerator) Verify() {
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

func (sort Sort) String() (s string) {
	switch sort {
	case RealSort:
		s = "Real"
		return
	case BoolSort:
		s = "Bool"
		return
	default:
		panic("Unknown sort in Sort.String: #" + strconv.Itoa(int(uint8(sort))))
	}
}

const VARIABLE_PREFIX = "var_"

func (g ConstraintGenerator) AddVariables(n int, sort_distr DDistr[Sort], p_var float64) {
	prev := len(g.vals)
	// p_var = rep*n / (prev + rep*n)
	// p_var*prev + p_var*rep*n = rep*n
	// p_var*prev = (1 - p_var)*rep*n
	// p_var*prev / (1 - p_var)*n = rep
	rep := int(math.Round(
		(p_var * float64(prev)) / ((1 - p_var) * float64(n)),
	))
	for i := 0; i < n; i++ {
		g.AddVariable(sort_distr, rep)
	}
}

func (g ConstraintGenerator) AddVariable(sort_distr DDistr[Sort], rep int) {
	sort := sort_distr.Sample()
	id := *g.next_var_id
	*g.next_var_id++
	val := Val{VARIABLE_PREFIX + strconv.Itoa(id), sort}
	for i := 0; i < rep; i++ {
		if _, ok := g.vals[sort]; !ok {
			g.vals[sort] = make([]Val, 1)
			g.vals[sort][0] = val
		} else {
			g.vals[sort] = append(g.vals[sort], val)
		}
	}
}

func (g ConstraintGenerator) Variables() (vars []Val) {
	vars = make([]Val, 0)
	seen_vars := make(map[string]struct{})
	for _, val_subset := range g.vals {
		for _, val := range val_subset {
			if strings.HasPrefix(val.name, VARIABLE_PREFIX) {
				if _, ok := seen_vars[val.name]; !ok {
					seen_vars[val.name] = struct{}{}
					vars = append(vars, val)
				}
			}
		}
	}
	return
}

func (g ConstraintGenerator) SMTFreeFuns() (smt_free_funs []qse.SMTFreeFun[string, string]) {
	vars := g.Variables()
	smt_free_funs = make([]qse.SMTFreeFun[string, string], len(vars))
	for i, val := range vars {
		smt_free_funs[i] = qse.SMTFreeFun[string, string]{
			Name: val.name,
			Args: []string{},
			Ret:  val.sort.String(),
		}
	}
	return
}

func (g ConstraintGenerator) Generate(sort Sort) (expr string) {
	expr = g.GenerateAtDepth(
		sort,
		int(math.Round(math.Max(
			0,
			(rand.NormFloat64()*g.n_depth_stddev)+g.n_depth_mean,
		))),
	)
	return
}

func (g ConstraintGenerator) GenerateAtDepth(sort Sort, depth int) (expr string) {
	if depth == 0 {
		index := rand.Intn(len(g.vals[sort]))
		return g.vals[sort][index].name
	}
	index := rand.Intn(len(g.ops[sort]))
	head := g.ops[sort][index]
	var sb strings.Builder
	sb.WriteRune('(')
	sb.WriteString(head.name)
	for _, sub_sort := range head.args {
		sb.WriteRune(' ')
		sb.WriteString(
			g.GenerateAtDepth(sub_sort, depth-1),
		)
	}
	sb.WriteRune(')')
	expr = sb.String()
	return
}
