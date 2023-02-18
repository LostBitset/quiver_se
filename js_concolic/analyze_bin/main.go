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

// Takes the target program filename as an argument
// Generates the hash (in binary, alg MD5, out base64, of filename)
// For each Analyze message targeting that program it sees,
//   Convert the model into assignments
//   Call the Jalangi2 analysis code (which writes the PathCondition response)
//   Delete the Analyze message file

func main() {
	fmt.Println("[js_concolic:AnalyzerProcess] Started Analyzer Process.")
	// Setup
	if len(os.Args) != 3 {
		panic("ERR! Need two arguments, filename and message prefix. ")
	}
	target_filename := os.Args[1]
	msg_prefix := os.Args[2]
	// Handle messages
	seen_names := make(map[string]struct{})
	msgdir := `../.eidin-run/Analyze`
	var wg sync.WaitGroup
	fmt.Println("[js_concolic:AnalyzerProcess] Ready.")
mainLoop:
	for {
		entries, err := os.ReadDir(msgdir)
		if err != nil {
			panic(err)
		}
	currentAnalyzeMsgsLoop:
		for _, entry := range entries {
			if entry.IsDir() {
				continue currentAnalyzeMsgsLoop
			}
			filename := entry.Name()
			if !strings.HasPrefix(filename, msg_prefix) {
				continue currentAnalyzeMsgsLoop
			}
			if strings.HasSuffix(filename, "__EIDIN-SIGNAL-STOP") {
				break mainLoop
			}
			if _, ok := seen_names[filename]; ok {
				fmt.Printf("Name \"%s\" already seen.\n", filename)
				continue currentAnalyzeMsgsLoop
			}
			seen_names[filename] = struct{}{}
			fmt.Printf("Added \"%s\" to seen names.\n", filename)
			contents, errf := os.ReadFile(msgdir + "/" + filename)
			if err != nil {
				panic(errf)
			}
			msg := &eidin.Analyze{}
			erru := proto.Unmarshal(contents, msg)
			if err != nil {
				panic(erru)
			}
			fmt.Println("[js_concolic:AnalyzerProcess] Successfully deserialized Analyze message. ")
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer fmt.Println("[js_concolic:AnalyzerProcess] Deleted message, done processing. ")
				defer os.Remove(msgdir + "/" + filename)
				HandleAnalyze(*msg, target_filename)
			}()
		}
		timer := time.After(200 * time.Millisecond)
		wg.Wait()
		<-timer
	}
}
