package main

import (
	quip "LostBitset/quiver_se/QUIP/lib"
	"fmt"
	"os"
)

func main() {
	fmt.Println("[QUIP:(bin/)main.go] Started QUIP.")
	if len(os.Args) != 3 {
		panic("ERR! Need two arguments, filename and message prefix. ")
	}
	target := os.Args[1]
	msg_prefix := os.Args[2]
	quip.StartQUIP(target, msg_prefix)
}
