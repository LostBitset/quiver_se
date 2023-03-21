package main

import "fmt"

func main() {
	uprgm_gen := BuildTestingMicroprogramGenerator()
	demo := uprgm_gen.RandomMicroprogram()
	fmt.Println(demo)
}
