package envflagstructconfig

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	trace "github.com/davidwalter0/tracer"
)

var tracer *trace.Tracer = trace.New()

func init() {
	// disable trace of function call path, comment or set to enable
	// traces
	tracer.Disable()
}

// Tracer enable users of package envflagstructconfig to access tracer instance
func Tracer() *trace.Tracer {
	return tracer
}

// SetValue from source text
func (info *Info) SetValue(text string) {
	defer tracer.ScopedTrace()()
	if len(text) > 0 {
		T := info.StructField
		st := (StructField)(info.StructField)
		switch T.Type.Kind() {
		case reflect.String:
			info.AbstractValue = text
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if T.Type.Kind() == reflect.Int64 &&
				T.Type.PkgPath() == "time" &&
				T.Type.Name() == "Duration" {
				lhs, _ := time.ParseDuration(text)
				info.AbstractValue = lhs
			} else {
				lhs, _ := strconv.ParseInt(text, 0, T.Type.Bits())
				info.AbstractValue = lhs
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			lhs, _ := strconv.ParseUint(text, 0, T.Type.Bits())
			info.AbstractValue = lhs
		case reflect.Bool:
			lhs, _ := strconv.ParseBool(text)
			info.AbstractValue = lhs
		case reflect.Float32, reflect.Float64:
			lhs, _ := strconv.ParseFloat(text, T.Type.Bits())
			info.AbstractValue = lhs
		case reflect.Slice:
			var slice = reflect.ValueOf(info.Address)
			slice = slice.Elem()
			// if slice.Kind() == reflect.Ptr {... slice = slice.Elem()
			for _, value := range strings.Split(text, ",") {
				info.AppendSlice(slice, st.CastTextToInterface(value))
			}
		case reflect.Map:
			var m = reflect.MakeMap(st.Type)
			var key = reflect.New(st.Type.Key()).Elem()
			var value = reflect.New(st.Type.Elem()).Elem()
			for _, kv := range strings.Split(text, ",") {
				pair := strings.Split(kv, ":")
				if len(pair) == 2 {
					k, v := pair[0], pair[1]
					if len(k) != 0 && len(v) != 0 {
						key.Set(TypeFromTextToReflectValue(st.Type.Key(), k))
						value.Set(TypeFromTextToReflectValue(st.Type.Elem(), v))
						m.SetMapIndex(key, value)
					}
				}
			}
			info.Field.Set(m)
		default:
			panic(info.Name)
		}
		info.Cast()
	}
}

// AppendSlice reflected append based on type
func (info *Info) AppendSlice(slice reflect.Value, v interface{}) {
	defer tracer.ScopedTrace()()
	value := reflect.New(info.Field.Type().Elem()).Elem()
	switch v.(type) {
	case time.Duration:
		value.SetInt(v.(int64))
		slice.Set(reflect.Append(slice, value))
	case bool:
		value.SetBool(v.(bool))
		slice.Set(reflect.Append(slice, value))
	case float32:
		value.SetFloat(v.(float64))
		slice.Set(reflect.Append(slice, value))
	case float64:
		value.SetFloat(v.(float64))
		slice.Set(reflect.Append(slice, value))
	case int:
		value.SetInt(v.(int64))
		slice.Set(reflect.Append(slice, value))
	case uint:
		value.SetUint(v.(uint64))
		slice.Set(reflect.Append(slice, value))
	case int64:
		value.SetInt(v.(int64))
		slice.Set(reflect.Append(slice, value))
	case uint64:
		value.SetUint(v.(uint64))
		slice.Set(reflect.Append(slice, value))
	case int8:
		value.SetInt(v.(int64))
		slice.Set(reflect.Append(slice, value))
	case uint8:
		value.SetUint(v.(uint64))
		slice.Set(reflect.Append(slice, value))
	case int16:
		value.SetInt(v.(int64))
		slice.Set(reflect.Append(slice, value))
	case uint16:
		value.SetUint(v.(uint64))
		slice.Set(reflect.Append(slice, value))
	case int32:
		value.SetInt(v.(int64))
		slice.Set(reflect.Append(slice, value))
	case uint32:
		value.SetUint(v.(uint64))
		slice.Set(reflect.Append(slice, value))
	case string:
		value.SetString(v.(string))
		slice.Set(reflect.Append(slice, value))
	}
}

// Cast item into reflected field via type
func (info *Info) Cast() {
	defer tracer.ScopedTrace()()
	switch ptr := info.AbstractValue.(type) {
	case time.Duration:
		info.Field.SetInt((int64)(info.AbstractValue.(time.Duration)))
	case bool:
		info.Field.SetBool(info.AbstractValue.(bool))
	case float32:
		info.Field.SetFloat(info.AbstractValue.(float64))
	case float64:
		info.Field.SetFloat(info.AbstractValue.(float64))
	case int:
		info.Field.SetInt(info.AbstractValue.(int64))
	case uint:
		info.Field.SetUint(info.AbstractValue.(uint64))
	case int64:
		info.Field.SetInt(info.AbstractValue.(int64))
	case uint64:
		info.Field.SetUint(info.AbstractValue.(uint64))
	case int8:
		info.Field.SetInt(info.AbstractValue.(int64))
	case uint8:
		info.Field.SetUint(info.AbstractValue.(uint64))
	case int16:
		info.Field.SetInt(info.AbstractValue.(int64))
	case uint16:
		info.Field.SetUint(info.AbstractValue.(uint64))
	case int32:
		info.Field.SetInt(info.AbstractValue.(int64))
	case uint32:
		info.Field.SetUint(info.AbstractValue.(uint64))
	case string:
		info.Field.SetString(info.AbstractValue.(string))
	default:
		if ptr == nil {
			break
		}
		log.Printf("unknown/default type:%v %T %p\n", ptr, ptr, ptr)
		panic("what")
	}
}

// Setkey for environment
func (info *Info) Setkey() {
	defer tracer.ScopedTrace()()
	info.UnderScoreCamelCaseWords()
	// log.Println(info.Name, info.Key)
}

// SetFlagName for environment
func (info *Info) SetFlagName() {
	defer tracer.ScopedTrace()()
	info.HyphenateCamelCaseWords()
}

// Initialize from environment or tag default: key
func (info *Info) Initialize() {
	defer tracer.ScopedTrace()()
	var text string
	info.Setkey()
	info.SetFlagName()
	text = os.Getenv(info.Key)
	if len(text) == 0 {
		text = info.Tag.Get("default")
	}
	info.SetValue(text)
}
