package main

import (
	smtlib2va "LostBitset/quiver_se/SMTLib_2VA/lib"
	"fmt"
)

func main() {
	fmt.Println(smtlib2va.TranspileV2From2VA("Hello, world!"))
}
