package envflagstructconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidSpecification = errors.New("Specification must be a struct pointer")
var ErrInvalidArgPointerRequired = errors.New("Argument must be a pointer")
var ErrInvalidArgMapParseSpec = errors.New("Map argument requires pairs")

/*
// TextTypedValMap interface formatted text to map
func TextTypedValMap(m interface{}, text string) error {
	if reflect.TypeOf(m).Kind() != reflect.Ptr {
		return ErrInvalidArgPointerRequired
	}
	pairs := strings.Split(text, ",")
	if len(pairs) > 0 {
		mp := reflect.MakeMap(reflect.TypeOf(m).Elem())
		for _, pair := range pairs {
			kvpair := strings.Split(pair, ":")
			v := reflect.ValueOf(m).Type().Elem().Elem()
			k := reflect.ValueOf(m).Type().Elem().Key()

			key := reflect.New(k).Elem()
			{ // scoped assignment
				lhs := kvpair[0]
				TextTypedVal(key.Addr().Interface(), lhs)
			}
			value := reflect.New(v).Elem()
			{ // scoped assignment
				lhs := kvpair[1]
				TextTypedVal(value.Addr().Interface(), lhs)
			}
			// mp.SetMapIndex(key, value)
			mp.SetMapIndex(key, value)
		}
		reflect.ValueOf(m).Elem().Set(mp)
	}
	return nil
}

// TextTypeSlice interface text to slice
func TextTypeSlice(m interface{}, text string) error {
	if reflect.TypeOf(m).Kind() != reflect.Ptr {
		return ErrInvalidArgPointerRequired
	}

	values := strings.Split(text, ",")
	if len(values) > 0 {
		T := reflect.TypeOf(m).Elem().Elem()
		var slice = reflect.ValueOf(m).Elem()
		for _, v := range values {
			svalue := reflect.New(T)
			value := svalue.Elem().Addr().Interface()
			if err := TextTypedVal(value, v); err != nil {
				return err
			}
			slice.Set(reflect.Append(slice, svalue.Elem()))
		}
	}
	return nil
}
*/
// TextTypedVal interface text to value
func TextTypedVal(arg interface{}, textArg ...string) (err error) {
	var text = ""
	if reflect.TypeOf(arg).Kind() != reflect.Ptr {
		err = ErrInvalidArgPointerRequired
	} else {
		if len(textArg) > 0 {
			text = textArg[0]
		}
		m := reflect.ValueOf(arg).Elem()
		if len(text) != 0 {
			switch m.Kind() {
			case reflect.String:
				m.SetString(text)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var lhs int64
				if m.Kind() == reflect.Int64 &&
					m.Type().PkgPath() == "time" &&
					m.Type().Name() == "Duration" {
					var t time.Duration
					t, err = time.ParseDuration(text)
					if err != nil {
						return
					}
					lhs = (int64)(t)
				} else {
					lhs, err = strconv.ParseInt(text, 10, m.Type().Bits())
					if err != nil {
						return
					}
				}
				m.SetInt(lhs)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				var lhs uint64
				lhs, err = strconv.ParseUint(text, 0, m.Type().Bits())
				if err != nil {
					return
				}
				m.SetUint(lhs)
			case reflect.Bool:
				var lhs bool
				lhs, err = strconv.ParseBool(text)
				if err != nil {
					return
				}
				m.SetBool(lhs)
			case reflect.Float32, reflect.Float64:
				var lhs float64
				lhs, err = strconv.ParseFloat(text, m.Type().Bits())
				if err != nil {
					return
				}
				m.SetFloat(lhs)
			case reflect.Slice:
				values := strings.Split(text, ",")
				if len(values) > 0 {
					T := reflect.TypeOf(arg).Elem().Elem()
					var slice = reflect.ValueOf(arg).Elem()
					for _, v := range values {
						svalue := reflect.New(T)
						value := svalue.Elem().Addr().Interface()
						if err := TextTypedVal(value, v); err != nil {
							return err
						}
						slice.Set(reflect.Append(slice, svalue.Elem()))
					}
				}
			case reflect.Map:
				pairs := strings.Split(text, ",")
				if len(pairs) > 0 {
					mp := reflect.MakeMap(reflect.TypeOf(arg).Elem())
					for _, pair := range pairs {
						kvpair := strings.Split(pair, ":")
						if len(kvpair) != 2 {
							return fmt.Errorf("Map argument requires pairs text source: %s", pair)
						}
						v := reflect.TypeOf(arg).Elem().Elem()
						k := reflect.TypeOf(arg).Elem().Key()

						key := reflect.New(k).Elem()
						{ // scoped assignment
							lhs := kvpair[0]
							TextTypedVal(key.Addr().Interface(), lhs)
						}
						value := reflect.New(v).Elem()
						{ // scoped assignment
							lhs := kvpair[1]
							TextTypedVal(value.Addr().Interface(), lhs)
						}
						// mp.SetMapIndex(key, value)
						mp.SetMapIndex(key, value)
					}
					reflect.ValueOf(arg).Elem().Set(mp)
				}
			case reflect.Struct:
				return ParseStructTypedVal(arg)
			}
		}
	}
	return
}

// ParseStructTypedVal take structure tag info and assign to items from tags
func ParseStructTypedVal(arg interface{}, text ...func() string) (err error) {
	if reflect.TypeOf(arg).Kind() != reflect.Ptr {
		err = ErrInvalidArgPointerRequired
	} else {
		m := reflect.ValueOf(arg).Elem()
		switch m.Kind() {
		case reflect.Struct:
			element := reflect.ValueOf(arg).Elem()
			elementType := element.Type()
			for i := 0; i < elementType.NumField(); i++ {
				structField := elementType.Field(i)
				// Get tags for this structField
				text := structField.Tag.Get("default")
				name := structField.Tag.Get("name")
				if len(name) == 0 {
					name = structField.Name
				}
				if structField.Type.Kind() == reflect.Struct {
					if err = ParseStructTypedVal(element.Field(i).Addr().Interface()); err != nil {
						return
					}
				} else {
					if len(text) != 0 {
						if err = TextTypedVal(element.Field(i).Addr().Interface(), text); err != nil {
							return
						}
					}
				}
			}
		}
	}
	return
}
