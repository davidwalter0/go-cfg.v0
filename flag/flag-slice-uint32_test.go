package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceUint32(t *testing.T) {
	T := reflect.TypeOf(SliceUint32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint32 SliceUint32")
	var sliceUint32 = new(SliceUint32Value)
	Var(sliceUint32, "sliceUint32", "use SliceUint32")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceUint32.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceUint32.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceUint32.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceUint32.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceUint32, *sliceUint32)
	// Parse()
	// Usage()
}

