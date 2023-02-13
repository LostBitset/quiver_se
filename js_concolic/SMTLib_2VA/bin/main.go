package main

import (
	smtlib2va "LostBitset/quiver_se/SMTLib_2VA/lib"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const SMTLIB2VA_EXTENSION = ".smt2va"

func main() {
	args := os.Args[1:]
	if len(args) == 2 && args[1] == "log-info" {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		if len(args) == 1 {
			logrus.SetLevel(logrus.WarnLevel)
		} else {
			panic("Requires 1 or 2 arguments: {binary} {2va_file} [log-info]")
		}
	}
	if !strings.HasSuffix(args[0], SMTLIB2VA_EXTENSION) {
		panic("The filename (first argument) must end in  \".smt2va\"")
	}
	b, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	transpiled := smtlib2va.TranspileV2From2VA(b)
	transpiled_bytes := []byte(transpiled)
	filename_root := args[0][:len(args[0])-len(SMTLIB2VA_EXTENSION)]
	output_filename := filename_root + ".TRANSPILED-orig_smt2va.smt"
}
