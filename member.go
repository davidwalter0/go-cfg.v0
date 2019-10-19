package cfg // import "github.com/davidwalter0/go-cfg"

import (
	"strconv"

	"github.com/davidwalter0/go-flag"

	"fmt"
	"log"
	"reflect"
	"strings"
)

var announceDuplicates bool

func ignore() {
	log.Println()
}

var allFlagNames = make(map[string]bool)

// TagInit initialize the MemberType struct from the struct tags
func (member *MemberType) TagInit(tag reflect.StructTag) {
	var jsonTag string

	defer func() {
		if err := recover(); err != nil {
		}
	}()

	member.Default = tag.Get("default")
	member.Value = member.Default
	member.Name = tag.Get("name")

	if jsonTag = tag.Get("json"); len(jsonTag) > 0 && len(member.Name) == 0 {
		member.Name = jsonTag
	}

	member.Short = tag.Get("short")
	var text = tag.Get("usage")
	if text = tag.Get("usage"); len(text) > 0 {
		member.Usage = text
	}
	if text = tag.Get("doc"); len(text) > 0 {
		member.Usage = text
	}
	if text = tag.Get("help"); len(text) > 0 {
		member.Usage = text
	}

	// ignore fields with ,omit in json tag
	if strings.Index(jsonTag, ",omitempty") == -1 && strings.Index(jsonTag, ",omit") > 0 {
		member.Ignore = true
	} else if jsonTag == "-" {
		// ignore fields with ,omit in json tag
		member.Ignore = true
	}
	// ignore fields with 'ignore' in json tag
	text = tag.Get("ignore")
	if v, err := strconv.ParseBool(text); err == nil {
		member.Ignore = v
	}

	required := tag.Get("required")
	if v, err := strconv.ParseBool(required); err == nil {
		member.Required = v
	}
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
	member.FlagName = member.Name
	member.UnderScoreCamelCaseWords()
	member.HyphenateCamelCaseWords()
	member.EnvInit()
	// Env variable names historically didn't allow hyphenation
	member.KeyName = strings.Replace(member.KeyName, "-", "_", -1)

	switch structField.Type.Kind() {
	case reflect.Struct:
		depth = depth + 1
		element := reflect.ValueOf(ptr).Elem()
		elementType := element.Type()
		if len(prefix) > 0 {
			prefix = prefix + "_" + Capitalize(member.Name)
		} else {
			prefix = Capitalize(member.Name)
		}
		for i := 0; i < elementType.NumField(); i++ {
			structField := elementType.Field(i)
			ptr := element.Field(i).Addr().Interface()
			if err = member.Parse(prefix, structField, ptr, depth); err != nil {
				return
			}
		}
	default:
		if !member.Ignore {
			var usage string
			if len(member.Usage) > 0 {
				usage = "usage: " + member.Usage
			}
			if _, ok := allFlagNames[member.FlagName]; !ok {

				if member.Required && len(member.Value) == 0 {
					return fmt.Errorf("Required field: Name [%s] FlagName [%s] Env var name [%s] has no default and is unset", member.Name, member.FlagName, member.KeyName)
				}
				flag.MakeVar(ptr, member.FlagName, member.Default,
					usage+fmt.Sprintf(" Env %-32s : (%v)",
						member.KeyName, structField.Type), member.Value)
				allFlagNames[member.FlagName] = true
			} else {
				if !announceDuplicates {
					fmt.Printf("Duplicate flag(s)/env vars found\n")
					fmt.Println(strings.ToUpper(fmt.Sprintf("%-20s %-20s", "flag", "env vars")))
					fmt.Println("-----------------------------------------")
					announceDuplicates = true
				}
				fmt.Printf("%-20s %-20s\n", member.FlagName, member.KeyName)
			}
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
