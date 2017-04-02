package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceUint16(t *testing.T) {
	T := reflect.TypeOf(SliceUint16Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint16 SliceUint16")
	var sliceUint16 = new(SliceUint16Value)
	Var(sliceUint16, "sliceUint16", "use SliceUint16")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceUint16.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceUint16.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceUint16.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceUint16.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceUint16, *sliceUint16)
	// Parse()
	// Usage()
}

