package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceFloat32(t *testing.T) {
	T := reflect.TypeOf(SliceFloat32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for float32 SliceFloat32")
	var sliceFloat32 = new(SliceFloat32Value)
	Var(sliceFloat32, "sliceFloat32", "use SliceFloat32")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceFloat32.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceFloat32.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceFloat32.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceFloat32.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceFloat32, *sliceFloat32)
	// Parse()
	// Usage()
}

