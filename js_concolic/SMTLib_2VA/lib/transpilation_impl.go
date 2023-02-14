package smtlib2va

import (
	"strings"
)

// lvbls = lo vanbi be lo snicne = the environment of the variables

func TranspileV2From2VA(src_2va []byte) (src_v2 string) {
	sb := strings.Builder{}
	lvbls := NewLexicallyScoped()
	src_v2 = sb.String()
	return
}
