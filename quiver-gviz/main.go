package main

import "fmt"

func main() {
	uprgm := GenerateEvaluationMicroprogram()
	dot := MicroprogramQuiverDot(uprgm)
	fmt.Println(dot.String())
}
