package smtlib2va

import (
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
	src_v2 = src_2va // just testing
	return
}
