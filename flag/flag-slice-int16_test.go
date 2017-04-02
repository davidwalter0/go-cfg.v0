package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceInt16(t *testing.T) {
	T := reflect.TypeOf(SliceInt16Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int16 SliceInt16")
	var sliceInt16 = new(SliceInt16Value)
	Var(sliceInt16, "sliceInt16", "use SliceInt16")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt16.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt16.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt16.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt16.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt16, *sliceInt16)
	// Parse()
	// Usage()
}

