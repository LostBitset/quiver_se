package main

import "strings"

func (sff StringSMTFreeFun) DefinitionString(rhs string) (stmt string) {
	if len(sff.FreeFun.Args) != 0 {
		panic("Invalid. Cannot generate definition string for parametric SMT fun.")
	}
	var sb strings.Builder
	sb.WriteRune('(')
	sb.WriteString(sff.FreeFun.Name)
	sb.WriteString(" () ")
	sb.WriteString(sff.FreeFun.Ret)
	sb.WriteRune(' ')
	sb.WriteString(rhs)
	sb.WriteRune(')')
	stmt = sb.String()
	return
}
