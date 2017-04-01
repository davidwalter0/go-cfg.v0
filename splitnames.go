package envflagstructconfig

import (
	"log"
	"regexp"
	"strings"
)

var regExpr = regexp.MustCompile("([^A-Z]+|[A-Z][^A-Z]+|[A-Z]+)")

func init() {
	defer tracer.ScopedTrace()()
	if false {
		log.Println("")
	}
}

// UnderScoreCamelCaseWords split on CamelCase words
func (info *Info) UnderScoreCamelCaseWords() {
	defer tracer.ScopedTrace()()
	words := regExpr.FindAllStringSubmatch(info.Name, -1)
	if len(words) > 0 {
		var envSepName []string
		for _, words := range words {
			envSepName = append(envSepName, words[0])
		}
		info.Key = strings.Join(envSepName, "_")
		if len(info.Prefix) > 0 {
			info.Key = strings.ToUpper(info.Prefix + "_" + info.Key)
		}
	}
}

// HyphenateCamelCaseWords converts camel case name string and
// hyphenates words for flags between words
func (info *Info) HyphenateCamelCaseWords() {
	defer tracer.ScopedTrace()()
	words := regExpr.FindAllStringSubmatch(info.Name, -1)
	if len(words) > 0 {
		var name []string
		for _, words := range words {
			name = append(name, strings.ToLower(words[0]))
		}
		info.FlagName = strings.Join(name, "-")
	}
}
