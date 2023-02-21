package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
	"os"
	"os/exec"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

const SUBROUTINE_DSE_TIMEOUT_MILLIS = 10000
const SUBROUTINE_DSE_CYCLE_WAIT_TIME_MILLIS = 100

// Should generally be called as a goroutine
func PerformDse(
	location string,
	msg_prefix string,
	single_callback_mode bool,
	pc_chan chan eidin.PathCondition,
) {
	go RunSimpleDSE(
		msg_prefix,
		SUBROUTINE_DSE_CYCLE_WAIT_TIME_MILLIS,
		single_callback_mode,
		true,
	)
	go RunAnalyzer(location, msg_prefix)
	go UsePathConditionChannel(msg_prefix, pc_chan)
	defer close(pc_chan)
	<-time.After(time.Duration(SUBROUTINE_DSE_TIMEOUT_MILLIS) * time.Millisecond)
	msgdir := `../../js_concolic/.eidin-run/PathCondition`
	f, err := os.Create(msgdir + "/" + msg_prefix + "__EIDIN-SIGNAL-STOP")
	if err != nil {
		panic(err)
	}
	defer os.Remove(msgdir + "/" + f.Name())
	defer f.Close()
	msgdir = `../../js_concolic/.eidin-run/Analyze`
	f, err = os.Create(msgdir + "/" + msg_prefix + "__EIDIN-SIGNAL-STOP")
	if err != nil {
		panic(err)
	}
	defer os.Remove(msgdir + "/" + f.Name())
	defer f.Close()
}

func KickstartDse(msg_prefix string) {
	cmd := exec.Command(
		"cp",
		"../../js_concolic/analyze_bin/empty_Analyze.eidin.bin",
		"../../js_concolic/.eidin-run/Analyze/"+msg_prefix+"_spec-empty.eidin.bin",
	)
	fmt.Println("[QUIP:perform_dse.go] Ran \"" + cmd.String() + "\".")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Println("[QUIP:perform_dse.go] Kickstarted DSE with empty Analyze message.")
}

func UsePathConditionChannel(msg_prefix string, pc_chan chan eidin.PathCondition) {
	seen_pc_hashes := make(map[uint32]struct{})
	// Handle messages
	msgdir := `../../js_concolic/.eidin-run/PathCondition`
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
			fmt.Println("[QUIP:perform_dse.go] Successfully deserialized PathCondition message. ")
			pc_chan <- *msg
		}
		<-time.After(
			time.Duration(SUBROUTINE_DSE_CYCLE_WAIT_TIME_MILLIS+50) * time.Millisecond,
		)
	}
}
