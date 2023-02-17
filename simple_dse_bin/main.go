package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("[simple_dse] Started orchestration process.")
	fmt.Println("[simple_dse] Performing simple concolic execution.")
	// Setup
	if len(os.Args) != 2 {
		panic("ERR! Need one arguments: message prefix. ")
	}
	msg_prefix := os.Args[1]
	// Handle messages
	msgdir := `../.eidin-run/PathCondition`
	var wg sync.WaitGroup
	for {
		entries, err := os.ReadDir(msgdir)
		if err != nil {
			panic(err)
		}
	currentPCMsgsLoop:
		for _, entry := range entries {
			if entry.IsDir() {
				continue currentPCMsgsLoop
			}
			filename := entry.Name()
			if !strings.HasPrefix(filename, msg_prefix) {
				continue currentPCMsgsLoop
			}
			fmt.Println(filename)
			contents, errf := os.ReadFile(filename)
			if err != nil {
				panic(errf)
			}
			msg := &eidin.PathCondition{}
			erru := proto.Unmarshal(contents, msg)
			if err != nil {
				panic(erru)
			}
			fmt.Println("[simple_dse] Successfully deserialized PathCondition message. ")
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer fmt.Println("[simple_dse] Deleted message, done processing. ")
				defer os.Remove(msgdir + "/" + filename)
				HandlePathCondition(*msg, msg_prefix)
			}()
		}
		timer := time.After(200 * time.Millisecond)
		wg.Wait()
		<-timer
	}
}
