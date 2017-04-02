package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceInt32(t *testing.T) {
	T := reflect.TypeOf(SliceInt32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int32 SliceInt32")
	var sliceInt32 = new(SliceInt32Value)
	Var(sliceInt32, "sliceInt32", "use SliceInt32")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt32.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt32.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt32.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt32.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt32, *sliceInt32)
	// Parse()
	// Usage()
}

