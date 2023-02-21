package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

// The time to wait between cycles, in milliseconds.
const CYCLE_WAIT_TIME_OPTION_PREFIX = "--cycle-wait-time="

func main() {
	fmt.Println("[simple_dse] Started orchestration process.")
	fmt.Println("[simple_dse] Performing simple concolic execution.")
	// Setup
	if len(os.Args) < 2 {
		panic("ERR! Need one arguments: message prefix (--cycle-wait-time=<int> optional). ")
	}
	msg_prefix := os.Args[1]
	cycle_time_millis := 200
	for _, arg := range os.Args {
		// The time to wait between cycles, in milliseconds.
		if strings.HasPrefix(arg, CYCLE_WAIT_TIME_OPTION_PREFIX) {
			value_str, _ := strings.CutPrefix(arg, CYCLE_WAIT_TIME_OPTION_PREFIX)
			arg_cycle_time_millis, err := strconv.Atoi(value_str)
			if err != nil {
				panic(err)
			}
			cycle_time_millis = arg_cycle_time_millis
		}
	}
	fmt.Printf("[simple_dse] cycle_time_millis=%d\n", cycle_time_millis)
	seen_pc_hashes := make(map[uint32]struct{})
	seen_analyze_hashes := make(map[uint32]struct{})
	// Handle messages
	msgdir := `../js_concolic/.eidin-run/PathCondition`
	var wg sync.WaitGroup
mainLoop:
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
				fmt.Println("[simple_dse] Skipping unwatched. DBG DBG DBG")
				fmt.Println(filename)
				fmt.Println("prefix:")
				fmt.Println(msg_prefix)
				continue currentPCMsgsLoop
			}
			if strings.HasSuffix(filename, "__EIDIN-SIGNAL-STOP") {
				break mainLoop
			}
			contents, errf := os.ReadFile(msgdir + "/" + filename)
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
			fmt.Println(contents)
			msg := &eidin.PathCondition{}
			erru := proto.Unmarshal(contents, msg)
			fmt.Println(*msg)
			if err != nil {
				panic(erru)
			}
			fmt.Println("[simple_dse] Successfully deserialized PathCondition message. ")
			wg.Add(1)
			func() {
				defer wg.Done()
				defer func() {
					if SliceContains(os.Args, "--rename-persist-path-conditions") {
						os.Rename(msgdir+"/"+filename, msgdir+"/persist_"+filename)
						fmt.Println("[simple_dse] Renamed message (with different prefix), done processing.")
					} else {
						os.Remove(msgdir + "/" + filename)
						fmt.Println("[simple_dse] Deleted message, done processing. ")
					}
				}()
				reqs := PathConditionToAnalyzeMessages(*msg)
				fmt.Printf("[simple_dse] Generated %v possible Analyze messages.\n", len(reqs))
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
		timer := time.After(time.Duration(cycle_time_millis) * time.Millisecond)
		wg.Wait()
		<-timer
	}
}

func SendAnalyzeMessage(amsg []byte, msg_prefix string) {
	hasher := fnv.New64a()
	hasher.Write(amsg)
	hash := hasher.Sum64()
	hash_s := strconv.Itoa(int(hash))
	filename := `../js_concolic/.eidin-run/Analyze/` + msg_prefix + hash_s + ".eidin.bin"
	fmt.Println(filename)
	err := os.WriteFile(filename, amsg, 0644)
	if err != nil {
		panic(err)
	}
}
