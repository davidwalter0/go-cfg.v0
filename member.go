package envflagstructconfig

import (
	"github.com/davidwalter0/envflagstructconfig/flag"

	"fmt"
	"log"
	"reflect"
	"strings"
)

func ignore() {
	log.Println()
}

// TagInit initialize the MemberType struct from the struct tags
func (member *MemberType) TagInit(tag reflect.StructTag) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	member.Default = tag.Get("default")
	member.Value = member.Default
	member.Name = tag.Get("name")
	member.Short = tag.Get("short")
	member.Usage = tag.Get("usage")
	member.Ignore = tag.Get("ignore")

}

// Capitalize text
func Capitalize(text string) string {
	var capitalized = text
	if len(text) > 0 {
		capitalized = strings.ToUpper(text[0:1]) + text[1:]
	}
	return capitalized
}

// Parse the MemberType structure from the current structure member
// assign the environment variable name, the flag name set the value
// from the default or parsed parse the default text
func (member *MemberType) Parse(prefix string,
	structField reflect.StructField,
	ptr interface{}, depth int) (err error) {
	member.TagInit(structField.Tag)
	if len(member.Name) == 0 {
		member.Name = structField.Name
	}

	member.EnvVarPrefix = prefix
	member.KeyName = member.Name
	member.FlagName = member.Name
	member.UnderScoreCamelCaseWords()
	member.HyphenateCamelCaseWords()
	member.EnvInit()

	switch structField.Type.Kind() {
	case reflect.Struct:
		depth = depth + 1
		element := reflect.ValueOf(ptr).Elem()
		elementType := element.Type()
		prefix = prefix + "_" + Capitalize(member.Name)
		for i := 0; i < elementType.NumField(); i++ {
			structField := elementType.Field(i)
			ptr := element.Field(i).Addr().Interface()
			if err = member.Parse(prefix, structField, ptr, depth); err != nil {
				return
			}
		}
	default:
		if member.Ignore != "true" {
			var usage string
			if len(member.Usage) > 0 {
				usage = "usage: " + member.Usage
			}
			flag.MakeVar(ptr, member.FlagName, member.Default,
				usage+fmt.Sprintf(" env var name(%s) : (%v)",
					member.KeyName, structField.Type), member.Value)
		}
	}
	return
}

// EnvInit uses the value from the environment for this structure
// member replacing the default tag value
func (member *MemberType) EnvInit() {
	if len(member.KeyName) > 0 {
		member.EnvText, _ = lookupEnv(member.KeyName)
	}
	if len(member.EnvText) != 0 {
		member.Value = member.EnvText
	}
}
