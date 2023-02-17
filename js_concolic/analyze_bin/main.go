package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"os"
	"strings"

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
	msgdir := `../.eidin-run/Analyze`
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
			contents, errf := os.ReadFile(filename)
			if err != nil {
				panic(errf)
			}
			msg := &eidin.Analyze{}
			erru := proto.Unmarshal(contents, msg)
			if err != nil {
				panic(erru)
			}
			fmt.Println("[js_concolic:AnalyzerProcess] Successfully deserialized Analyze message. ")
			go HandleAnalyze(*msg, target_filename)
		}
	}
}
