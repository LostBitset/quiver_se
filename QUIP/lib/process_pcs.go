package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	q "LostBitset/quiver_se/lib"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

// This function does the following:
// - Runs partial DSE on new callbacks
// - Listens for path conditions, updating the quiver accordingly
func ProcessPathConditions(
	out_updates chan q.Augmented[
		q.QuiverUpdate[
			int,
			q.PHashMap[q.Literal[q.WithId_H[string]], struct{}],
			*q.DMT[q.WithId_H[string], q.QuiverIndex],
		],
		[]q.SMTFreeFun[string, string],
	],
	top_node q.QuiverIndex,
	fail_node q.QuiverIndex,
	target string,
	msg_prefix string,
) {
	segment_chan := make(chan q.Augmented[eidin.PathConditionSegment, []q.SMTFreeFun[string, string]])
	seen_callbacks := make(map[uint64]struct{})
	go InterpretPathConditionSegments(segment_chan, msg_prefix)
	for segment_augmented := range segment_chan {
		segment := segment_augmented.Value
		cb_this := segment.GetThisCallbackId()
		cb_next := segment.GetNextCallbackId()
		if cb_this.GetBytesStart() != cb_this.GetBytesEnd() {
			if _, ok := seen_callbacks[cb_this.GetBytesStart()]; !ok {
				PerformPartialDse(*cb_this, target, msg_prefix)
			}
		}
		if cb_next.GetBytesStart() != cb_next.GetBytesEnd() {
			if _, ok := seen_callbacks[cb_next.GetBytesStart()]; !ok {
				PerformPartialDse(*cb_next, target, msg_prefix)
			}
		}
		out_updates <- SegmentToQuiverUpdate(segment, segment_augmented.Augment, top_node, fail_node)
	}
	close(out_updates)
}

func InterpretPathConditionSegments(
	segment_chan chan q.Augmented[eidin.PathConditionSegment, []q.SMTFreeFun[string, string]],
	msg_prefix string,
) {
	InterpretPathConditionSegmentsInner(segment_chan, "persist_"+msg_prefix)
}

func InterpretPathConditionSegmentsInner(
	segment_chan chan q.Augmented[eidin.PathConditionSegment, []q.SMTFreeFun[string, string]],
	msg_prefix string,
) {
	cycle_time_millis := 200
	seen_pc_hashes := make(map[uint32]struct{})
	// Handle messages
	msgdir := `../../js_concolic/.eidin-run/PathCondition`
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
				fmt.Println("[QUIP:process_pcs.go] Ignored already-seen path condition.")
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
			fmt.Println("[QUIP:process_pcs.go] Successfully deserialized PathCondition message. ")
			wg.Add(1)
			func() {
				defer wg.Done()
				defer func() {
					os.Remove(msgdir + "/" + filename)
					fmt.Println("[QUIP:process_pcs.go] Deleted message, done processing. ")
				}()
				smt_free_funs := make([]q.SMTFreeFun[string, string], 0)
				for _, free_fun := range msg.GetFreeFuns() {
					smt_free_funs = append(
						smt_free_funs,
						q.SMTFreeFun[string, string]{
							Name: free_fun.GetName(),
							Args: free_fun.GetArgSorts(),
							Ret:  free_fun.GetRetSort(),
						},
					)
				}
				for _, segment := range msg.GetSegmentedPc() {
					segment := segment
					segment_item := *segment
					segment_chan <- q.Augmented[
						eidin.PathConditionSegment,
						[]q.SMTFreeFun[string, string],
					]{
						Value:   segment_item,
						Augment: smt_free_funs,
					}
				}
			}()
		}
		timer := time.After(time.Duration(cycle_time_millis) * time.Millisecond)
		wg.Wait()
		<-timer
	}
}
