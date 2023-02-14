package smtlib2va

import (
	"regexp"
	"strings"
)

// lvbls = lo vanbi be lo snicne = the environment of the variables

func TranspileV2From2VA(src_2va string) (src_v2 string) {
	sb := strings.Builder{}
	strs_re := regexp.MustCompile(`\"(?:[^\\\"]|\\.)*\"`)
	strs_re.FindAllStringIndex(src_2va, -1)
	lvbls := NewLexicallyScoped()
	src_v2 = sb.String()
	return
}
