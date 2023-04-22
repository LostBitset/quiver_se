package main

import (
	"os"
)

func main() {
	uprgm := GenerateEvaluationMicroprogram()
	dot := MicroprogramQuiverDot(uprgm)
	f, errC := os.Create("reprdigraph.dot")
	if errC != nil {
		panic(errC)
	}
	defer f.Close()
	_, errW := f.Write([]byte(dot.String()))
	if errW != nil {
		panic(errW)
	}
}
