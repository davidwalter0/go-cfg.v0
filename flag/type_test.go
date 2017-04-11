package flag

import (
        "fmt"
        "testing"
        "reflect"
)

func init(){
   if false {
      fmt.Println("")
   }
}


func TestParseSliceDuration(t *testing.T) {
	T := reflect.TypeOf(sliceDurationValue{}).Elem()

	var sliceDuration = new(sliceDurationValue)
	Var(sliceDuration, "sliceDuration", "use SliceDuration")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceDuration.Set("1h2m3s,720h4m3s")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceDuration.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceDuration.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceDuration.Set("false,true")
  case reflect.String:
    sliceDuration.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("duration SliceDuration %v\n", *sliceDuration)
}

func TestParseSliceInt(t *testing.T) {
	T := reflect.TypeOf(sliceIntValue{}).Elem()

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
  case reflect.String:
    sliceInt.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("int SliceInt %v\n", *sliceInt)
}

func TestParseSliceInt8(t *testing.T) {
	T := reflect.TypeOf(sliceInt8Value{}).Elem()

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
  case reflect.String:
    sliceInt8.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("int8 SliceInt8 %v\n", *sliceInt8)
}

func TestParseSliceInt16(t *testing.T) {
	T := reflect.TypeOf(sliceInt16Value{}).Elem()

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
  case reflect.String:
    sliceInt16.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("int16 SliceInt16 %v\n", *sliceInt16)
}

func TestParseSliceInt32(t *testing.T) {
	T := reflect.TypeOf(sliceInt32Value{}).Elem()

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
  case reflect.String:
    sliceInt32.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("int32 SliceInt32 %v\n", *sliceInt32)
}

func TestParseSliceInt64(t *testing.T) {
	T := reflect.TypeOf(sliceInt64Value{}).Elem()

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
  case reflect.String:
    sliceInt64.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("int64 SliceInt64 %v\n", *sliceInt64)
}

func TestParseSliceUint(t *testing.T) {
	T := reflect.TypeOf(sliceUintValue{}).Elem()

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
  case reflect.String:
    sliceUint.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("uint SliceUint %v\n", *sliceUint)
}

func TestParseSliceUint8(t *testing.T) {
	T := reflect.TypeOf(sliceUint8Value{}).Elem()

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
  case reflect.String:
    sliceUint8.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("uint8 SliceUint8 %v\n", *sliceUint8)
}

func TestParseSliceUint16(t *testing.T) {
	T := reflect.TypeOf(sliceUint16Value{}).Elem()

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
  case reflect.String:
    sliceUint16.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("uint16 SliceUint16 %v\n", *sliceUint16)
}

func TestParseSliceUint32(t *testing.T) {
	T := reflect.TypeOf(sliceUint32Value{}).Elem()

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
  case reflect.String:
    sliceUint32.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("uint32 SliceUint32 %v\n", *sliceUint32)
}

func TestParseSliceUint64(t *testing.T) {
	T := reflect.TypeOf(sliceUint64Value{}).Elem()

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
  case reflect.String:
    sliceUint64.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("uint64 SliceUint64 %v\n", *sliceUint64)
}

func TestParseSliceFloat64(t *testing.T) {
	T := reflect.TypeOf(sliceFloat64Value{}).Elem()

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
  case reflect.String:
    sliceFloat64.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("float64 SliceFloat64 %v\n", *sliceFloat64)
}

func TestParseSliceFloat32(t *testing.T) {
	T := reflect.TypeOf(sliceFloat32Value{}).Elem()

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
  case reflect.String:
    sliceFloat32.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("float32 SliceFloat32 %v\n", *sliceFloat32)
}

func TestParseSliceBool(t *testing.T) {
	T := reflect.TypeOf(sliceBoolValue{}).Elem()

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
  case reflect.String:
    sliceBool.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("bool SliceBool %v\n", *sliceBool)
}

func TestParseSliceString(t *testing.T) {
	T := reflect.TypeOf(sliceStringValue{}).Elem()

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
  case reflect.String:
    sliceString.Set("a,b,c,false,true,another string value")
  }
	// fmt.Printf("string SliceString %v\n", *sliceString)
}

