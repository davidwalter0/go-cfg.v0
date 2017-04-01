package envflagstructconfig

import (
	"github.com/davidwalter0/envflagstructconfig/flag"

	"log"
	"reflect"
	"time"
)

// MakeFlag create a flag address of object initialize if not set from
// initial
func MakeFlag(name, usage string, initial interface{}, addr interface{}) {
	switch ptr := addr.(type) {
	case *time.Duration:
		var v time.Duration
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(time.Duration)
			}
		}
		flag.DurationVar(ptr, name, v, usage)
	case *bool:
		var v bool
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(bool)
			}
		}
		flag.BoolVar(ptr, name, v, usage)
	case *float32:
		var v float32
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(float32)
			}
		}
		flag.Float32Var(ptr, name, v, usage)
	case *float64:
		var v float64
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(float64)
			}
		}
		flag.Float64Var(ptr, name, v, usage)
	case *int:
		var v int
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = (int)(initial.(int64))
			}
		}
		flag.IntVar(ptr, name, v, usage)
	case *uint:
		var v uint
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = (uint)(initial.(uint64))
			}
		}
		flag.UintVar(ptr, name, v, usage)
	case *int64:
		var v int64
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(int64)
			}
		}
		flag.Int64Var(ptr, name, v, usage)
	case *uint64:
		var v uint64
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(uint64)
			}
		}
		flag.Uint64Var(ptr, name, v, usage)
	case *int8:
		var v int8
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(int8)
			}
		}
		flag.Int8Var(ptr, name, v, usage)
	case *uint8:
		var v uint8
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(uint8)
			}
		}
		flag.Uint8Var(ptr, name, v, usage)
	case *int16:
		var v int16
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(int16)
			}
		}
		flag.Int16Var(ptr, name, v, usage)
	case *uint16:
		var v uint16
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(uint16)
			}
		}
		flag.Uint16Var(ptr, name, v, usage)
	case *int32:
		var v int32
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(int32)
			}
		}
		flag.Int32Var(ptr, name, v, usage)
	case *uint32:
		var v uint32
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(uint32)
			}
		}
		flag.Uint32Var(ptr, name, v, usage)
	case *[]string:
		var v flag.SliceValue
		if initial != nil {
			sliceValue := initial.(flag.SliceValue)
			if len(sliceValue) > 0 {
				for _, u := range sliceValue {
					v = append(v, u)
				}
			}
			v = initial.(flag.SliceValue)
		}
		flag.SliceVar((*flag.SliceValue)(ptr), name, v, usage)
	case *string:
		var v string
		if *ptr != v {
			v = *ptr
		} else {
			if initial != nil {
				v = initial.(string)
			}
		}
		flag.StringVar(ptr, name, v, usage)
	default:
		log.Printf("unknown/default type:%v %T %p\n", ptr, ptr, ptr)
	}
}

// Initialize global initialization strategy
func Initialize(prefix string, structure interface{}) {
	if reflect.TypeOf(structure).Kind() != reflect.Ptr {
		log.Fatalln(ErrInvalidSpecification)
	}

	element := reflect.ValueOf(structure).Elem()
	elementType := element.Type()
	for i := 0; i < elementType.NumField(); i++ {
		structField := elementType.Field(i)
		// Get tags for this structField
		name := structField.Tag.Get("name")
		if len(name) == 0 {
			name = structField.Name
		}

		info := &Info{
			Num:         i,
			Prefix:      prefix,
			Name:        name,
			Address:     element.Field(i).Addr().Interface(),
			Field:       element.Field(i),
			StructField: structField,
			Tag:         structField.Tag,
			Short:       structField.Tag.Get("short"),
			Usage:       structField.Tag.Get("usage"),
		}
		switch info.Field.Kind() {
		case reflect.Map:
			info.Kind = info.Field.Type().Kind()
			info.KeyKind = info.StructField.Type.Key().Kind()
			info.ValueKind = info.Field.Type().Elem().Kind()
		case reflect.Slice:
			info.Kind = info.Field.Kind()
			info.ValueKind = info.Field.Type().Elem().Kind()
		}

		info.Initialize()

		names := []string{}
		if len(info.FlagName) > 0 {
			names = append(names, info.FlagName)
		}
		if len(info.Short) > 0 {
			names = append(names, info.Short)
		}

		addr := element.Field(i).Addr().Interface()
		usage := info.Usage
		var initial interface{}
		for _, name := range names {
			MakeFlag(name, usage, initial, addr)
		}
	}
	flag.Parse()
}
