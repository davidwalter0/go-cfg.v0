package envflagstructconfig

import (
	"github.com/davidwalter0/envflagstructconfig/flag"
	"reflect"
)

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
		flag.Usage()
	} else {
		flag.Usage()
		panic("Process called more than once")
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
			if structInfo.EnvVarPrefix, ok = lookupEnv(AppEnvVarPrefixOverrideName); !ok {
				structInfo.EnvVarPrefix = elementType.Name()
			}
			prefix := structInfo.EnvVarPrefix
			for i := 0; i < elementType.NumField(); i++ {
				structField := elementType.Field(i)
				ptr := element.Field(i).Addr().Interface()
				var info = &InfoType{
					AppName:      elementType.Name(),
					EnvVarPrefix: prefix,
				}

				if err = info.Process(prefix, structField, ptr, depth); err != nil {
					return
				}
			}
		}
	}
	return
}
