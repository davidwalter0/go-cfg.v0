package envflagstructconfig

import (
	"github.com/davidwalter0/envflagstructconfig/flag"
	"log"
	"reflect"
)

// AppEnvVarPrefixOverrideName environment variable application lookup
// prefix override, default prefix is the struct name
var AppEnvVarPrefixOverrideName = "APP_OVERRIDE_PREFIX"

// Process bootstrap the configuration from environment and flags to
// struct with env var name override to replace the prefix of the
// object name. The prefix argument will replace/override the struct
// type name, to default to the use of the struct name call the Parse
// method directly
func Process(prefix string, sptr interface{}) {
	var sti *StructInfo = &StructInfo{
		StructPtr:    sptr,
		EnvVarPrefix: prefix,
	}

	if err := sti.Parse(); err != nil { // parse tags, environment, flags
		log.Fatalf("%v\n", err)
	}

}

// Parse bootstrap the configuration from environment and flags
// to struct with env var name override to replace the prefix of the
// object name
func Parse(sptr interface{}) (err error) {
	var sti *StructInfo = &StructInfo{
		StructPtr: sptr,
	}

	if err = sti.Parse(); err != nil { // parse tags, environment, flags
		log.Printf("%v\n", err)
	}
	return err
}

// Parse the struct and flags initializing configuration from tags,
// environment, and flags in order, additional flags must be defined
// prior to Parse call as it calls flag.Parse
func (structInfo *StructInfo) Parse() (err error) {
	if !structInfo.Processed {
		structInfo.Processed = true
		if err = structInfo.process(); err != nil {
			return
		}
		flag.Parse()
	} else {
		panic("Parse called more than once")
	}
	return
}

// process initialize with *struct (struct pointer) and use the struct
// name as the app name for the env prefix
// func (structInfo *StructInfo) Init(structPtr interface{}) (err error) {
func (structInfo *StructInfo) process() (err error) {
	if reflect.TypeOf(structInfo.StructPtr).Kind() != reflect.Ptr {
		err = ErrInvalidArgPointerRequired
	} else {
		var depth = 0
		var ok bool
		m := reflect.ValueOf(structInfo.StructPtr).Elem()
		switch m.Kind() {
		case reflect.Struct:
			element := reflect.ValueOf(structInfo.StructPtr).Elem()
			elementType := element.Type()
			AppName := elementType.Name()
			if len(structInfo.EnvVarPrefix) == 0 {
				if structInfo.EnvVarPrefix, ok = lookupEnv(AppEnvVarPrefixOverrideName); !ok {
					structInfo.EnvVarPrefix = elementType.Name()
				} else {
					AppName = structInfo.EnvVarPrefix
				}
			}
			prefix := structInfo.EnvVarPrefix

			for i := 0; i < elementType.NumField(); i++ {
				structField := elementType.Field(i)
				ptr := element.Field(i).Addr().Interface()
				var member = &MemberType{
					AppName:      AppName,
					EnvVarPrefix: prefix,
				}
				if err = member.Parse(prefix, structField, ptr, depth); err != nil {
					return
				}
			}
		}
	}
	return
}
