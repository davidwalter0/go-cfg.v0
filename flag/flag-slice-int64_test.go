package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceInt64(t *testing.T) {
	T := reflect.TypeOf(SliceInt64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int64 SliceInt64")
	var sliceInt64 = new(SliceInt64Value)
	Var(sliceInt64, "sliceInt64", "use SliceInt64")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt64.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt64.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt64.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt64.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt64, *sliceInt64)
	// Parse()
	// Usage()
}

