package envflagstructconfig

import (
	"log"
	"regexp"
	"strings"
)

var regExpr = regexp.MustCompile("([^A-Z]+|[A-Z][^A-Z]+|[A-Z]+)")

func init() {
	// defer tracer.ScopedTrace()()
	if false {
		log.Println("")
	}
}

// UnderScoreCamelCaseWords split on CamelCase words
func (info *InfoType) UnderScoreCamelCaseWords() {
	// defer tracer.ScopedTrace()()
	words := regExpr.FindAllStringSubmatch(info.Name, -1)
	if len(words) > 0 {
		var envSepName []string
		for _, words := range words {
			envSepName = append(envSepName, words[0])
		}
		info.KeyName = strings.Join(envSepName, "_")
		if len(info.EnvVarPrefix) > 0 {
			info.KeyName = strings.ToUpper(info.EnvVarPrefix + "_" + info.KeyName)
		}
	}
}

// HyphenateCamelCaseWords converts camel case name string and
// hyphenates words for flags between words
func (info *InfoType) HyphenateCamelCaseWords() {
	info.FlagName = strings.Replace(info.FlagName, "_", "", -1)
	// defer tracer.ScopedTrace()()
	words := regExpr.FindAllStringSubmatch(info.FlagName, -1)
	if len(words) > 0 {
		var name []string
		for _, words := range words {
			name = append(name, strings.ToLower(words[0]))
		}
		info.FlagName = strings.Join(name, "-")
	}
}
