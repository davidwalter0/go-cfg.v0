package envflagstructconfig

import (
	"testing"
)

// TestUnderScoreCamelCaseWords split on CamelCase words
func TestUnderScoreCamelCaseWords(t *testing.T) {
	var info = &Info{}
	member.UnderScoreCamelCaseWords()

}

// TestHyphenateCamelCaseWords converts camel case name string and
// hyphenates words for flags between words
func TestHyphenateCamelCaseWords() {
	var info = &Info{}
	member.HyphenateCamelCaseWords()
}
