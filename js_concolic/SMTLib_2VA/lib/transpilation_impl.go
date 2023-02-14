package smtlib2va

import (
	"strings"
)

// lvbls = lo vanbi be lo snicne = the environment of the variables

func TranspileV2From2VA(src_2va []byte) (src_v2 string) {
	sb := strings.Builder{}
	lvbls := NewLexicallyScoped()
	RewriteSexprs(src_2va, sb, map[string](func() string){
		"*/enter-scope/*": func() (rewritten string) {
			lvbls.EnterScope()
			rewritten = ""
			return
		},
		"*/leave-scope/*": func() (rewritten string) {
			lvbls.LeaveScope()
			rewritten = ""
			return
		},
		// TODO (more)
	})
	src_v2 = sb.String()
	return
}
