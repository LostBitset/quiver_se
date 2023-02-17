package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
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
	seen_pc_hashes := make(map[uint32]struct{})
	seen_analyze_hashes := make(map[uint32]struct{})
	// Handle messages
	msgdir := `../js_concolic/.eidin-run/PathCondition`
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
			contents, errf := os.ReadFile(filename)
			if err != nil {
				panic(errf)
			}
			contents_hasher := fnv.New32a()
			contents_hasher.Write(contents)
			hash := contents_hasher.Sum32()
			if _, ok := seen_pc_hashes[hash]; ok {
				errr := os.Remove(msgdir + "/" + filename)
				if errr != nil {
					panic(errr)
				}
				fmt.Println("Ignored already-seen path condition.")
				continue currentPCMsgsLoop
			}
			seen_pc_hashes[hash] = struct{}{}
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
				reqs := PathConditionToAnalyzeMessages(*msg)
			sendAnalyzeMsgsLoop:
				for _, amsg := range reqs {
					amsg_hasher := fnv.New32a()
					amsg_hasher.Write(amsg)
					hash := amsg_hasher.Sum32()
					if _, ok := seen_analyze_hashes[hash]; ok {
						continue sendAnalyzeMsgsLoop
					}
					SendAnalyzeMessage(amsg, msg_prefix)
				}
			}()
		}
		timer := time.After(200 * time.Millisecond)
		wg.Wait()
		<-timer
	}
}

func SendAnalyzeMessage(amsg []byte, msg_prefix string) {
	filename := `../js_concolic/.eidin-run/Analyze`
	err := os.WriteFile(filename, amsg, 0644)
	if err != nil {
		panic(err)
	}
}
