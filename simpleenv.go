package cfg

// Boolish string interpreted as boolean
import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Boolish string

// On return true if Boolish string text value is in truth set, false
// if in false set panic otherwise
func (e Boolish) On() bool {
	switch e {
	case "1", "t", "T", "true", "TRUE", "True":
		return true
	case "0", "f", "F", "false", "FALSE", "False":
		return false
	default:
		//		panic("func (e Boolish) On() passed invalid bool string " + e)
	}
	return false
}

func IsBoolish(boolish string) bool {
	return Boolish(boolish).On()
}

func EnvTextIsBoolish(name string) bool {
	return IsBoolish(os.Getenv(name))
}

// DumpObjectHelper MarshalIndent output of object
func DumpObjectHelper(o interface{}) string {
	var err error
	return fmt.Sprintln(func() (text string) {
		var rv []byte
		if rv, err = json.MarshalIndent(o, "", "  "); err != nil {
			fmt.Println(err)
		}
		text = string(rv)
		return text
	}())
}

/*
// DumpObjectHelperInternal MarshalIndent output of object append DUMP
// & HELP
func (sti *StructInfo) DumpObjectHelperInternal(o interface{}) string {
	var err error
	return fmt.Sprintln(func() (text string) {
		var rv []byte
		if rv, err = json.MarshalIndent(o, "", "  "); err != nil {
			fmt.Println(err)
		}

		var translate = map[string]interface{}{}
		if err = json.Unmarshal(rv, &translate); err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%+v\n", translate)
		// Copy the self referential StructPtr name
		translate[sti.EnvVarPrefix] = translate["StructPtr"]
		// Remove the self referential StructPtr name
		delete(translate, "StructPtr")
		if rv, err = json.MarshalIndent(translate, "", "  "); err != nil {
			fmt.Println(err)
		}

		text = string(rv)

		return text
	}())
}
*/

// DumpStructInfo return the string from json.MarshalIndent
func (sti *StructInfo) DumpStructInfo() string {
	return DumpObjectHelper(sti.StructPtr)
	///////// var err error
	///////// return fmt.Sprintln(func() (text string) {
	///////// 	var rv []byte
	///////// 	if rv, err = json.MarshalIndent(sti.StructPtr, "", "  "); err != nil {
	///////// 		fmt.Println(err)
	///////// 	}
	///////// 	/*
	///////// 		var translate = map[string]interface{}{}
	///////// 		if err = json.Unmarshal(rv, &translate); err != nil {
	///////// 			fmt.Println(err)
	///////// 		}

	///////// 		fmt.Printf("%+v\n", translate)
	///////// 			// Copy the self referential StructPtr name
	///////// 			translate[sti.EnvVarPrefix] = translate["StructPtr"]
	///////// 			// Remove the self referential StructPtr name
	///////// 			delete(translate, "StructPtr")
	///////// 			if rv, err = json.MarshalIndent(translate, "", "  "); err != nil {
	///////// 				fmt.Println(err)
	///////// 			}

	///////// 	*/
	///////// 	text = string(rv)
	///////// 	return text
	///////// }())
}

// EnvDumpSetWithKV when the EnvVarPrefix + _DUMP are set to true
func (sti *StructInfo) EnvDumpSetWithKV() (string, string, bool) {
	var key string = strings.ToUpper(sti.EnvVarPrefix + "_DUMP")
	var value = os.Getenv(key)
	if Boolish(value).On() {
		return key, value, true
	}
	return key, value, true

}

// EnvHelpSet is true when the EnvVarPrefix + _HELP are set to true
func (sti *StructInfo) EnvHelpSet() bool {
	// Special case test for faux Env var UpperCase(EnvVarPrefix) + _HELP
	var key string = strings.ToUpper(sti.EnvVarPrefix + "_HELP")
	var value = os.Getenv(key)
	if Boolish(value).On() {
		return true
	}
	return false
}

// EnvDumpSet when the EnvVarPrefix + _DUMP are set to true
func (sti *StructInfo) EnvDumpSet() bool {
	var key string = strings.ToUpper(sti.EnvVarPrefix + "_DUMP")
	var value = os.Getenv(key)
	if Boolish(value).On() {
		return true
	}
	return false
}

// LoadCfgSet when the EnvVarPrefix + _DUMP are set to true
func (sti *StructInfo) LoadCfgSet() bool {
	var key string = strings.ToUpper(sti.EnvVarPrefix + "_LOAD_CFG")
	var value = os.Getenv(key)
	if Boolish(value).On() {
		return true
	}
	return false
}

// SaveCfgSet when the EnvVarPrefix + _DUMP are set to true
func (sti *StructInfo) SaveCfgSet() bool {
	var key string = strings.ToUpper(sti.EnvVarPrefix + "_SAVE_CFG")
	var value = os.Getenv(key)
	if Boolish(value).On() {
		return true
	}
	return false
}

// EnvHelpSet is true when the EnvVarPrefix + _HELP are set to true
func (sti *StructInfo) EnvHelpSetWithKV() (string, string, bool) {
	// Special case test for faux Env var UpperCase(EnvVarPrefix) + _HELP
	var key string = strings.ToUpper(sti.EnvVarPrefix + "_HELP")
	var value = os.Getenv(key)
	if Boolish(value).On() {
		return key, value, true
	}
	return key, value, false
}

// Env bootstrap the configuration from environment without flags
// to struct with env var name override to replace the prefix of the
// object name
func Env(sptr interface{}) (err error) {
	var sti *StructInfo = &StructInfo{
		StructPtr: sptr,
	}

	if err = sti.Env(); err != nil { // parse tags, environment, flags
		log.Printf("%v\n", err)
	}

	if sti.EnvDumpSet() {
		fmt.Println(sti.DumpStructInfo())
	}

	// if sti.SaveCfgSet() {
	// 	ioutil.WriteFile(filename, []byte(sti.DumpStructInfo()), 0664)
	// }

	if sti.EnvHelpSet() {
		Usage()
		os.Exit(1)
	}

	return err
}

// Env the struct and flags initializing configuration from tags,
// environment, and flags in order, additional flags must be defined
// prior to Env call as it calls flag.Env
func (structInfo *StructInfo) Env() (err error) {
	if !structInfo.Processed {
		structInfo.Processed = true
		if err = structInfo.process(); err != nil {
			return
		}
	} else {
		panic("Env called more than once")
	}
	return
}
