package grep

import (
	"regexp"
)

func Grep(ss []string, pat string) (ret []string) {
	r := regexp.MustCompile(pat)
	for _, s := range ss {
		if s != "" && r.MatchString(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Grepnot(ss []string, pat string) (ret []string) {
	r := regexp.MustCompile(pat)
	for _, s := range ss {
		if s == "" || !r.MatchString(s) {
			ret = append(ret, s)
		}
	}
	return
}
