package main

import (
	"os"

	synth "github.com/LostBitset/quiver_se/lib_synthetic"
)

func main() {
	uprgm := synth.GenerateEvaluationMicroprogram()
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
