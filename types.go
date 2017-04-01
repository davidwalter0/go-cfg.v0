package envflagstructconfig

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

// ErrInvalidSpecification indicates that a specification is of the wrong type.
var ErrInvalidSpecification = errors.New("Specification must be a struct pointer")

// Info struct to hold parse information
type Info struct {
	Num           int // field number
	Prefix        string
	Name          string
	Key           string // ENV variable name
	Field         reflect.Value
	Tag           reflect.StructTag
	Initial       string
	Short         string // short flag name
	Usage         string
	FlagName      string // hyphenated flag name
	Value         string
	AbstractValue interface{}
	StructField   reflect.StructField
	Address       interface{}
	Kind          reflect.Kind
	KeyKind       reflect.Kind
	ValueKind     reflect.Kind
}

// TypeFromTextToReflectValue string to reflect value
func TypeFromTextToReflectValue(T reflect.Type, text string) reflect.Value {
	defer tracer.ScopedTrace()()
	value := reflect.New(T).Elem()
	if len(text) > 0 {
		switch T.Kind() {
		case reflect.String:
			value.SetString(text)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if T.Kind() == reflect.Int64 &&
				T.PkgPath() == "time" &&
				T.Name() == "Duration" {
				lhs, _ := time.ParseDuration(text)
				value.SetInt((int64)(lhs))
			}
			lhs, _ := strconv.ParseInt(text, 0, T.Bits())
			value.SetInt(lhs)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			lhs, _ := strconv.ParseUint(text, 0, T.Bits())
			value.SetUint(lhs)
		case reflect.Bool:
			lhs, _ := strconv.ParseBool(text)
			value.SetBool(lhs)
		case reflect.Float32, reflect.Float64:
			lhs, _ := strconv.ParseFloat(text, T.Bits())
			value.SetFloat(lhs)
		default:
		}
	}
	return value
}
