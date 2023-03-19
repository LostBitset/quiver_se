package smtlib2va

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// lvbls = lo vanbi be lo snicne = the environment of the variables

func TranspileV2From2VA(src_2va string) (src_v2 string) {
	comments_re := regexp.MustCompile(`;.*`)
	comments := make([]string, 0)
	wo_comments := comments_re.ReplaceAllStringFunc(
		src_2va,
		func(orig string) (repl string) {
			repl = "<<tmp:cmnt>>@" + strconv.Itoa(len(comments))
			comments = append(comments, orig)
			return
		},
	)
	strs_re := regexp.MustCompile(`\"(?:[^\\\"]|\\.)*\"`)
	string_lits := make([]string, 0)
	wo_strings := strs_re.ReplaceAllStringFunc(
		wo_comments,
		func(orig string) (repl string) {
			repl = "<<tmp:string>>@" + strconv.Itoa(len(string_lits))
			string_lits = append(string_lits, orig)
			return
		},
	)
	lvbls := NewLexicallyScoped()
	wo_strings_transpiled := TranspileV2From2VANoStrings(wo_strings, &lvbls)
	cut_strs_re := regexp.MustCompile(`<<tmp\:string>>@\d+`)
	with_strs := cut_strs_re.ReplaceAllStringFunc(
		wo_strings_transpiled,
		func(orig string) (repl string) {
			_, index_str, _ := strings.Cut(orig, "@")
			index, _ := strconv.Atoi(index_str)
			repl = string_lits[index]
			return
		},
	)
	cut_cmnts_re := regexp.MustCompile(`<<tmp\:cmnt>>@\d+`)
	with_strs_and_cmnts := cut_cmnts_re.ReplaceAllStringFunc(
		with_strs,
		func(orig string) (repl string) {
			_, index_str, _ := strings.Cut(orig, "@")
			index, _ := strconv.Atoi(index_str)
			repl = comments[index]
			return
		},
	)
	src_v2 = with_strs_and_cmnts
	return
}

func TranspileV2From2VANoStrings(src_2va string, lvbls *LexicallyScoped) (src_v2 string) {
	sexpr_2va_re := regexp.MustCompile(
		`\(\*\/[a-z\-]+\/\*(\s\*\*[\w\-]+(\s\*{{.*}}\*)?)?\)`,
	)
	src_v2 = sexpr_2va_re.ReplaceAllStringFunc(
		src_2va,
		func(orig string) (repl string) {
			head_section_raw, _, _ := strings.Cut(orig, "/*")
			head := head_section_raw[3:]
			logrus.Infof(
				"Replacing SMTLib_2VA sexpr (head \"*/%s/*\").",
				head,
			)
			switch head {
			case "enter-scope":
				lvbls.EnterScope()
				repl = ""
			case "leave-scope":
				lvbls.LeaveScope()
				repl = ""
			case "decl-var":
				_, name_section_raw, _ := strings.Cut(orig, "**")
				name := name_section_raw[:len(name_section_raw)-1]
				lvbls.DeclVar(name)
				repl = ""
			case "write-var":
				_, after_section_raw, _ := strings.Cut(orig, "**")
				name_section_raw, _, _ := strings.Cut(after_section_raw, "*{{")
				name := strings.TrimSpace(name_section_raw)
				capture_start := strings.Index(orig, "*{{") + 3
				capture_end := strings.Index(orig, "}}*")
				inner := orig[capture_start:capture_end]
				inner_rec := TranspileV2From2VANoStrings(inner, lvbls)
				lvbls.WriteVar(name, inner_rec)
				repl = ""
			case "read-var":
				_, name_section_raw, _ := strings.Cut(orig, "**")
				name := name_section_raw[:len(name_section_raw)-1]
				repl = lvbls.ReadVar(name)
			case "is-defined?":
				_, name_section_raw, _ := strings.Cut(orig, "**")
				name := name_section_raw[:len(name_section_raw)-1]
				defined := lvbls.IsDefined(name)
				if defined {
					repl = "true"
				} else {
					repl = "false"
				}
			default:
				panic("Unrecognized SMTLib_2VA head (one enclosed by slanted earmuffs)")
			}
			return
		},
	)
	return
}
