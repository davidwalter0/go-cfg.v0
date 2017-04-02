package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceUint(t *testing.T) {
	T := reflect.TypeOf(SliceUintValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint SliceUint")
	var sliceUint = new(SliceUintValue)
	Var(sliceUint, "sliceUint", "use SliceUint")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceUint.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceUint.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceUint.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceUint.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceUint, *sliceUint)
	// Parse()
	// Usage()
}

