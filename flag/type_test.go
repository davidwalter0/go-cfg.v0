package flag

import (
	"fmt"
	"os"
	"testing"
	"reflect"
)

func TestParseSliceDuration(t *testing.T) {
	T := reflect.TypeOf(sliceDurationValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for duration SliceDuration")
	var sliceDuration = new(sliceDurationValue)
	Var(sliceDuration, "sliceDuration", "use SliceDuration")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceDuration.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceDuration.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceDuration.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceDuration.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceDuration, *sliceDuration)
	// Parse()
	// Usage()
}

func TestParseSliceInt(t *testing.T) {
	T := reflect.TypeOf(sliceIntValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int SliceInt")
	var sliceInt = new(sliceIntValue)
	Var(sliceInt, "sliceInt", "use SliceInt")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt, *sliceInt)
	// Parse()
	// Usage()
}

func TestParseSliceInt8(t *testing.T) {
	T := reflect.TypeOf(sliceInt8Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int8 SliceInt8")
	var sliceInt8 = new(sliceInt8Value)
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

func TestParseSliceInt16(t *testing.T) {
	T := reflect.TypeOf(sliceInt16Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int16 SliceInt16")
	var sliceInt16 = new(sliceInt16Value)
	Var(sliceInt16, "sliceInt16", "use SliceInt16")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt16.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt16.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt16.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt16.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt16, *sliceInt16)
	// Parse()
	// Usage()
}

func TestParseSliceInt32(t *testing.T) {
	T := reflect.TypeOf(sliceInt32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int32 SliceInt32")
	var sliceInt32 = new(sliceInt32Value)
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

func TestParseSliceInt64(t *testing.T) {
	T := reflect.TypeOf(sliceInt64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int64 SliceInt64")
	var sliceInt64 = new(sliceInt64Value)
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

func TestParseSliceUint(t *testing.T) {
	T := reflect.TypeOf(sliceUintValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint SliceUint")
	var sliceUint = new(sliceUintValue)
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

func TestParseSliceUint8(t *testing.T) {
	T := reflect.TypeOf(sliceUint8Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint8 SliceUint8")
	var sliceUint8 = new(sliceUint8Value)
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

func TestParseSliceUint16(t *testing.T) {
	T := reflect.TypeOf(sliceUint16Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint16 SliceUint16")
	var sliceUint16 = new(sliceUint16Value)
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

func TestParseSliceUint32(t *testing.T) {
	T := reflect.TypeOf(sliceUint32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint32 SliceUint32")
	var sliceUint32 = new(sliceUint32Value)
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

func TestParseSliceUint64(t *testing.T) {
	T := reflect.TypeOf(sliceUint64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for uint64 SliceUint64")
	var sliceUint64 = new(sliceUint64Value)
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

func TestParseSliceFloat64(t *testing.T) {
	T := reflect.TypeOf(sliceFloat64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for float64 SliceFloat64")
	var sliceFloat64 = new(sliceFloat64Value)
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

func TestParseSliceFloat32(t *testing.T) {
	T := reflect.TypeOf(sliceFloat32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for float32 SliceFloat32")
	var sliceFloat32 = new(sliceFloat32Value)
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

func TestParseSliceBool(t *testing.T) {
	T := reflect.TypeOf(sliceBoolValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for bool SliceBool")
	var sliceBool = new(sliceBoolValue)
	Var(sliceBool, "sliceBool", "use SliceBool")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceBool.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceBool.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceBool.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceBool.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceBool, *sliceBool)
	// Parse()
	// Usage()
}

func TestParseSliceString(t *testing.T) {
	T := reflect.TypeOf(sliceStringValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for string SliceString")
	var sliceString = new(sliceStringValue)
	Var(sliceString, "sliceString", "use SliceString")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceString.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceString.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceString.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceString.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceString, *sliceString)
	// Parse()
	// Usage()
}

