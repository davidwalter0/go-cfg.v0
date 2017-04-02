package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceUint8(t *testing.T) {
	T := reflect.TypeOf(SliceUint8Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint8 SliceUint8")
	var sliceUint8 = new(SliceUint8Value)
	Var(sliceUint8, "sliceUint8", "use SliceUint8")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceUint8.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceUint8.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceUint8.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceUint8.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceUint8, *sliceUint8)
	// Parse()
	// Usage()
}

