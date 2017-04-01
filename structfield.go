package envflagstructconfig

import (
	"reflect"
	"strconv"
	"time"
)

// StructField local definition for receivers
type StructField reflect.StructField

// CastTextToInterface from local StructField definition
func (st StructField) CastTextToInterface(text string) interface{} {
	defer tracer.ScopedTrace()()
	if len(text) != 0 {
		switch st.Type.Elem().Kind() {
		case reflect.String:
			return text
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if st.Type.Kind() == reflect.Int64 &&
				st.Type.PkgPath() == "time" &&
				st.Type.Name() == "Duration" {
				lhs, _ := time.ParseDuration(text)
				return (int64)(lhs)
			}
			lhs, _ := strconv.ParseInt(text, 0, st.Type.Elem().Bits())
			return lhs
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			lhs, _ := strconv.ParseUint(text, 0, st.Type.Elem().Bits())
			return lhs
		case reflect.Bool:
			lhs, _ := strconv.ParseBool(text)
			return lhs
		case reflect.Float32, reflect.Float64:
			lhs, _ := strconv.ParseFloat(text, st.Type.Elem().Bits())
			return lhs
		}
	}
	return nil
}

// SetValue from local StructField definition
func (st StructField) SetValue(text string) interface{} {
	defer tracer.ScopedTrace()()
	if len(text) != 0 {
		if len(text) > 0 {
			switch st.Type.Kind() {
			case reflect.String:
				reflect.ValueOf(st).SetString(text)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if st.Type.Kind() == reflect.Int64 &&
					st.Type.PkgPath() == "time" &&
					st.Type.Name() == "Duration" {
					lhs, _ := time.ParseDuration(text)
					reflect.ValueOf(st).SetInt((int64)(lhs))
				} else {
					lhs, _ := strconv.ParseInt(text, 0, st.Type.Bits())
					reflect.ValueOf(st).SetInt(lhs)
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				lhs, _ := strconv.ParseUint(text, 0, st.Type.Bits())
				reflect.ValueOf(st).SetUint(lhs)
			case reflect.Bool:
				lhs, _ := strconv.ParseBool(text)
				reflect.ValueOf(st).SetBool(lhs)
			case reflect.Float32, reflect.Float64:
				lhs, _ := strconv.ParseFloat(text, st.Type.Bits())
				reflect.ValueOf(st).SetFloat(lhs)
			}
		}
	}
	return reflect.ValueOf(st)
}
