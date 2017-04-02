package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceBool(t *testing.T) {
	T := reflect.TypeOf(SliceBoolValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for bool SliceBool")
	var sliceBool = new(SliceBoolValue)
	Var(sliceBool, "sliceBool", "use SliceBool")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceBool.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceBool.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceBool.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceBool.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceBool, *sliceBool)
	// Parse()
	// Usage()
}

