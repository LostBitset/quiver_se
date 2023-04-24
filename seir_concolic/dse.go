package main

type PathConditionResult struct {
	pc    []string
	fails bool
}

func SeirTopEventState() (top SeirEventState) {
	top = SeirEventState{"__top__reserved"}
	return
}

type SeirEventState struct {
	name string
}

type SeirPrgm struct {
	source string
}
