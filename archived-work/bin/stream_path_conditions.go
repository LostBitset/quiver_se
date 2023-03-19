package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

func StreamPathConditions(msg_prefix_original string, pc_chan chan eidin.PathCondition) {
	msgdir := `../js_concolic/.eidin-run/PathCondition`
	var msg_prefix string
	if os.Args[len(os.Args)-1] == "--watch-persisted" {
		msg_prefix = "persist_" + msg_prefix_original
	} else {
		msg_prefix = msg_prefix_original
	}
	seen_pc_hashes := make(map[uint32]struct{})
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
			fmt.Println("CONTENTS OF PATHCONDITION MESSAGE BELOW:")
			fmt.Println(contents)
			msg := &eidin.PathCondition{}
			erru := proto.Unmarshal(contents, msg)
			msg_value := *msg
			fmt.Println(msg_value)
			if err != nil {
				panic(erru)
			}
			fmt.Println(
				"[bin/stream_path_conditions.go] Successfully deserialized PathCondition message. ",
			)
			pc_chan <- msg_value
		}
		timer := time.After(200 * time.Millisecond)
		<-timer
	}
}