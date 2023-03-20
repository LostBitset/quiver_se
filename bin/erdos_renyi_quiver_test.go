package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestErdosRenyiQuiverBasic(t *testing.T) {
	for i := 0; i < 500; i++ {
		n := rand.Intn(20) + 3
		p := (rand.Float64() * 0.8) + 0.1
		avg_ne := rand.Intn(20) + 3
		adj_list := ErdosRenyiQuiverGivenEdges(n, p, avg_ne)
		fmt.Printf("adj_list: %#+v\n", adj_list)
	}
}
