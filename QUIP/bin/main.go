package main

import (
	quip "LostBitset/quiver_se/QUIP/lib"
	"fmt"
	"os"
)

func main() {
	fmt.Println("[QUIP:(bin/)main.go] Started QUIP.")
	if len(os.Args) != 2 {
		panic("ERR! Need two arguments, filename and message prefix. ")
	}
	target := os.Args[1]
	msg_prefix := quip.GetMessagePrefix(target)
	fmt.Println("[QUIP:(bin/)main.go] Performing initial instrumentation...")
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	quip.InstrumentFunctionInfo(cwd + "/" + target)
	fmt.Println("[QUIP:(bin/)main.go] Initial instrumentation complete.")
	fmt.Println("[QUIP:(bin/)main.go] Running all QUIP components...")
	quip.StartQUIP(target, msg_prefix)
}
