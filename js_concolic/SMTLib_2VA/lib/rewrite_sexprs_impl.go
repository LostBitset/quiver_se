package smtlib2va

import (
	"strings"
)

func RewriteSexprs(orig []byte, out strings.Builder, rewrites map[string](func() string)) {
	// TODO
}
