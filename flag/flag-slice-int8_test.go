package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceInt8(t *testing.T) {
	T := reflect.TypeOf(SliceInt8Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int8 SliceInt8")
	var sliceInt8 = new(SliceInt8Value)
	Var(sliceInt8, "sliceInt8", "use SliceInt8")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt8.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt8.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt8.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt8.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt8, *sliceInt8)
	// Parse()
	// Usage()
}

