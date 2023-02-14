package smtlib2va

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// lvbls = lo vanbi be lo snicne = the environment of the variables

func TranspileV2From2VA(src_2va string) (src_v2 string) {
	comments_re := regexp.MustCompile(`;.*`)
	wo_comments := comments_re.ReplaceAllString(src_2va, "")
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
	wo_strings_transpiled := TranspileV2From2VANoStrings(wo_strings)
	cut_strs_re := regexp.MustCompile(`<<tmp\:string>>@\d+`)
	src_v2 = cut_strs_re.ReplaceAllStringFunc(
		wo_strings_transpiled,
		func(orig string) (repl string) {
			_, index_str, _ := strings.Cut(orig, "@")
			index, _ := strconv.Atoi(index_str)
			repl = string_lits[index]
			return
		},
	)
	return
}

func TranspileV2From2VANoStrings(src_2va string) (src_v2 string) {
	lvbls := NewLexicallyScoped()
	sexpr_2va_re := regexp.MustCompile(
		`\(\*\/[a-z\-]+\/\*(\s\*\*[\w\-]+(\s\*{{.*}}\*)?)?\)`,
	)
	src_v2 = sexpr_2va_re.ReplaceAllStringFunc(
		src_2va,
		func(orig string) (repl string) {
			head_section_raw, _, _ := strings.Cut(orig, "/*")
			head := head_section_raw[3:]
			fmt.Println(head)
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
				lvbls.WriteVar(name, inner)
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
