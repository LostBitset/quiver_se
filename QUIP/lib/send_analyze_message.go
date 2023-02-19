package lib

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"

	"google.golang.org/protobuf/proto"
)

func SendAnalyzeMessage(model string, msg_prefix string) {
	SendRawAnalyzeMessage(
		MakeAnalyzeMessage(model, false),
		msg_prefix,
	)
}

func MakeAnalyzeMessage(model string, single_callback_mode bool) (msg_raw []byte) {
	msg := &eidin.Analyze{
		ForbidCaching:  false,
		Model:          &model,
		SingleCallback: single_callback_mode,
	}
	out, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	msg_raw = out
	return
}

func SendRawAnalyzeMessage(amsg []byte, msg_prefix string) {
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
