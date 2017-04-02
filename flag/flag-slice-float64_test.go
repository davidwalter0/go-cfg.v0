package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceFloat64(t *testing.T) {
	T := reflect.TypeOf(SliceFloat64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for float64 SliceFloat64")
	var sliceFloat64 = new(SliceFloat64Value)
	Var(sliceFloat64, "sliceFloat64", "use SliceFloat64")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceFloat64.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceFloat64.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceFloat64.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceFloat64.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceFloat64, *sliceFloat64)
	// Parse()
	// Usage()
}

