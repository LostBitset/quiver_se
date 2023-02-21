package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
)

func GetMessagePrefix(location string) (pfx string) {
	hasher_md5 := md5.New()
	hasher_md5.Write([]byte(location))
	hash_md5 := hasher_md5.Sum([]byte{})
	hash_md5_b64 := base64.StdEncoding.EncodeToString(hash_md5)
	pfx = fmt.Sprintf("m_%s_", hash_md5_b64)
	pfx = strings.ReplaceAll(pfx, "/", "^")
	return
}
