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

// TagInit initialize the InfoType struct from the struct tags
func (info *InfoType) TagInit(tag reflect.StructTag) {
	info.Default = tag.Get("default")
	info.Initial = info.Default
	info.Name = tag.Get("name")
	info.Short = tag.Get("short")
	info.Usage = tag.Get("usage")
}

// Capitalize text
func Capitalize(text string) string {
	var capitalized = text
	if len(text) > 0 {
		capitalized = strings.ToUpper(text[0:1]) + text[1:]
	}
	return capitalized
}

// Process the info structure from the current structure member
func (info *InfoType) Process(prefix string, structField reflect.StructField, ptr interface{}, depth int) (err error) {
	info.TagInit(structField.Tag)
	if len(info.Name) == 0 {
		info.Name = structField.Name
	}

	info.EnvVarPrefix = prefix
	info.KeyName = info.Name
	info.FlagName = info.Name
	info.UnderScoreCamelCaseWords()
	info.HyphenateCamelCaseWords()
	info.EnvInit()

	switch structField.Type.Kind() {
	case reflect.Struct:
		depth = depth + 1
		element := reflect.ValueOf(ptr).Elem()
		elementType := element.Type()
		prefix = prefix + "_" + Capitalize(info.Name)
		for i := 0; i < elementType.NumField(); i++ {
			structField := elementType.Field(i)
			ptr := element.Field(i).Addr().Interface()
			if err = info.Process(prefix, structField, ptr, depth); err != nil {
				return
			}
		}
	default:
		if len(info.Initial) != 0 {
			if err = TextTypedVal(ptr, info.Initial); err != nil {
				return
			}
			flag.MakeVar(ptr, info.FlagName, info.Default,
				info.Usage+fmt.Sprintf(":Env var name(%s) : (%v)", info.KeyName, structField.Type))

		}
	}
	// flag.Var...
	return
}

// EnvInit capture the value from the environment for this structure
// member
func (info *InfoType) EnvInit() {
	if len(info.KeyName) > 0 {
		info.EnvText, _ = lookupEnv(info.KeyName)
	}
	if len(info.EnvText) != 0 {
		info.Initial = info.EnvText
	}
}
