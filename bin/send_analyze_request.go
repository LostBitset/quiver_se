package main

import (
	eidin "LostBitset/quiver_se/EIDIN/proto_lib"
	"fmt"
	"hash/fnv"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func SendAnalyzeRequest(msg_prefix string, model string) {
	log.Info("[bin/send_analyze_request.go] Sending Analyze message. ")
	message := eidin.Analyze{
		ForbidCaching:  false,
		Model:          &model,
		SingleCallback: false,
	}
	msg_raw, proto_enc_err := proto.Marshal(&message)
	if proto_enc_err != nil {
		panic(proto_enc_err)
	}
	cwd, cwd_err := os.Getwd()
	if cwd_err != nil {
		panic(cwd_err)
	}
	hasher := fnv.New64a()
	hasher.Write(msg_raw)
	f, file_err := os.Create(
		fmt.Sprintf(
			"%s/../js_concolic/.eidin-run/Analyze/%s_%d_fromqse.eidin.bin",
			cwd,
			msg_prefix,
			hasher.Sum64(),
		),
	)
	if file_err != nil {
		panic(file_err)
	}
	defer f.Close()
	f.Write(msg_raw)
}
