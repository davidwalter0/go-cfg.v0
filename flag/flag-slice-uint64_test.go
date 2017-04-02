package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceUint64(t *testing.T) {
	T := reflect.TypeOf(SliceUint64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint64 SliceUint64")
	var sliceUint64 = new(SliceUint64Value)
	Var(sliceUint64, "sliceUint64", "use SliceUint64")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceUint64.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceUint64.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceUint64.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceUint64.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceUint64, *sliceUint64)
	// Parse()
	// Usage()
}

