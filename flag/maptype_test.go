package flag

import (
        "fmt"
        "testing"
)

func init(){
   if false {
      fmt.Println("")
   }
}



func TestParseMapDurationDuration(t *testing.T) {
	// T := reflect.TypeOf(mapDurationDurationValue{}).Elem()
	var varMapDurationDuration = new(mapDurationDurationValue)
	Var(varMapDurationDuration, "varMapDurationDuration", "Use mapDurationDuration")
  varMapDurationDuration.Set("1h2m3s:1h2m3s")
  if varMapDurationDuration.String() != "1h2m3s:1h2m3s" {
    t.Fatalf("%v %v %v", varMapDurationDuration,
      varMapDurationDuration.String(), "1h2m3s:1h2m3s")
  }
	// fmt.Printf("duration mapDurationDuration\n%v\n1h2m3s:1h2m3s\n", *varMapDurationDuration)
}


func TestParseMapDurationInt(t *testing.T) {
	// T := reflect.TypeOf(mapDurationIntValue{}).Elem()
	var varMapDurationInt = new(mapDurationIntValue)
	Var(varMapDurationInt, "varMapDurationInt", "Use mapDurationInt")
  varMapDurationInt.Set("1h2m3s:-1")
  if varMapDurationInt.String() != "1h2m3s:-1" {
    t.Fatalf("%v %v %v", varMapDurationInt,
      varMapDurationInt.String(), "1h2m3s:-1")
  }
	// fmt.Printf("int mapDurationInt\n%v\n1h2m3s:-1\n", *varMapDurationInt)
}


func TestParseMapDurationInt8(t *testing.T) {
	// T := reflect.TypeOf(mapDurationInt8Value{}).Elem()
	var varMapDurationInt8 = new(mapDurationInt8Value)
	Var(varMapDurationInt8, "varMapDurationInt8", "Use mapDurationInt8")
  varMapDurationInt8.Set("1h2m3s:-1")
  if varMapDurationInt8.String() != "1h2m3s:-1" {
    t.Fatalf("%v %v %v", varMapDurationInt8,
      varMapDurationInt8.String(), "1h2m3s:-1")
  }
	// fmt.Printf("int8 mapDurationInt8\n%v\n1h2m3s:-1\n", *varMapDurationInt8)
}


func TestParseMapDurationInt16(t *testing.T) {
	// T := reflect.TypeOf(mapDurationInt16Value{}).Elem()
	var varMapDurationInt16 = new(mapDurationInt16Value)
	Var(varMapDurationInt16, "varMapDurationInt16", "Use mapDurationInt16")
  varMapDurationInt16.Set("1h2m3s:-1")
  if varMapDurationInt16.String() != "1h2m3s:-1" {
    t.Fatalf("%v %v %v", varMapDurationInt16,
      varMapDurationInt16.String(), "1h2m3s:-1")
  }
	// fmt.Printf("int16 mapDurationInt16\n%v\n1h2m3s:-1\n", *varMapDurationInt16)
}


func TestParseMapDurationInt32(t *testing.T) {
	// T := reflect.TypeOf(mapDurationInt32Value{}).Elem()
	var varMapDurationInt32 = new(mapDurationInt32Value)
	Var(varMapDurationInt32, "varMapDurationInt32", "Use mapDurationInt32")
  varMapDurationInt32.Set("1h2m3s:-1")
  if varMapDurationInt32.String() != "1h2m3s:-1" {
    t.Fatalf("%v %v %v", varMapDurationInt32,
      varMapDurationInt32.String(), "1h2m3s:-1")
  }
	// fmt.Printf("int32 mapDurationInt32\n%v\n1h2m3s:-1\n", *varMapDurationInt32)
}


func TestParseMapDurationInt64(t *testing.T) {
	// T := reflect.TypeOf(mapDurationInt64Value{}).Elem()
	var varMapDurationInt64 = new(mapDurationInt64Value)
	Var(varMapDurationInt64, "varMapDurationInt64", "Use mapDurationInt64")
  varMapDurationInt64.Set("1h2m3s:-1")
  if varMapDurationInt64.String() != "1h2m3s:-1" {
    t.Fatalf("%v %v %v", varMapDurationInt64,
      varMapDurationInt64.String(), "1h2m3s:-1")
  }
	// fmt.Printf("int64 mapDurationInt64\n%v\n1h2m3s:-1\n", *varMapDurationInt64)
}


func TestParseMapDurationUint(t *testing.T) {
	// T := reflect.TypeOf(mapDurationUintValue{}).Elem()
	var varMapDurationUint = new(mapDurationUintValue)
	Var(varMapDurationUint, "varMapDurationUint", "Use mapDurationUint")
  varMapDurationUint.Set("1h2m3s:2")
  if varMapDurationUint.String() != "1h2m3s:2" {
    t.Fatalf("%v %v %v", varMapDurationUint,
      varMapDurationUint.String(), "1h2m3s:2")
  }
	// fmt.Printf("uint mapDurationUint\n%v\n1h2m3s:2\n", *varMapDurationUint)
}


func TestParseMapDurationUint8(t *testing.T) {
	// T := reflect.TypeOf(mapDurationUint8Value{}).Elem()
	var varMapDurationUint8 = new(mapDurationUint8Value)
	Var(varMapDurationUint8, "varMapDurationUint8", "Use mapDurationUint8")
  varMapDurationUint8.Set("1h2m3s:2")
  if varMapDurationUint8.String() != "1h2m3s:2" {
    t.Fatalf("%v %v %v", varMapDurationUint8,
      varMapDurationUint8.String(), "1h2m3s:2")
  }
	// fmt.Printf("uint8 mapDurationUint8\n%v\n1h2m3s:2\n", *varMapDurationUint8)
}


func TestParseMapDurationUint16(t *testing.T) {
	// T := reflect.TypeOf(mapDurationUint16Value{}).Elem()
	var varMapDurationUint16 = new(mapDurationUint16Value)
	Var(varMapDurationUint16, "varMapDurationUint16", "Use mapDurationUint16")
  varMapDurationUint16.Set("1h2m3s:2")
  if varMapDurationUint16.String() != "1h2m3s:2" {
    t.Fatalf("%v %v %v", varMapDurationUint16,
      varMapDurationUint16.String(), "1h2m3s:2")
  }
	// fmt.Printf("uint16 mapDurationUint16\n%v\n1h2m3s:2\n", *varMapDurationUint16)
}


func TestParseMapDurationUint32(t *testing.T) {
	// T := reflect.TypeOf(mapDurationUint32Value{}).Elem()
	var varMapDurationUint32 = new(mapDurationUint32Value)
	Var(varMapDurationUint32, "varMapDurationUint32", "Use mapDurationUint32")
  varMapDurationUint32.Set("1h2m3s:2")
  if varMapDurationUint32.String() != "1h2m3s:2" {
    t.Fatalf("%v %v %v", varMapDurationUint32,
      varMapDurationUint32.String(), "1h2m3s:2")
  }
	// fmt.Printf("uint32 mapDurationUint32\n%v\n1h2m3s:2\n", *varMapDurationUint32)
}


func TestParseMapDurationUint64(t *testing.T) {
	// T := reflect.TypeOf(mapDurationUint64Value{}).Elem()
	var varMapDurationUint64 = new(mapDurationUint64Value)
	Var(varMapDurationUint64, "varMapDurationUint64", "Use mapDurationUint64")
  varMapDurationUint64.Set("1h2m3s:2")
  if varMapDurationUint64.String() != "1h2m3s:2" {
    t.Fatalf("%v %v %v", varMapDurationUint64,
      varMapDurationUint64.String(), "1h2m3s:2")
  }
	// fmt.Printf("uint64 mapDurationUint64\n%v\n1h2m3s:2\n", *varMapDurationUint64)
}


func TestParseMapDurationFloat64(t *testing.T) {
	// T := reflect.TypeOf(mapDurationFloat64Value{}).Elem()
	var varMapDurationFloat64 = new(mapDurationFloat64Value)
	Var(varMapDurationFloat64, "varMapDurationFloat64", "Use mapDurationFloat64")
  varMapDurationFloat64.Set("1h2m3s:2.71828")
  if varMapDurationFloat64.String() != "1h2m3s:2.71828" {
    t.Fatalf("%v %v %v", varMapDurationFloat64,
      varMapDurationFloat64.String(), "1h2m3s:2.71828")
  }
	// fmt.Printf("float64 mapDurationFloat64\n%v\n1h2m3s:2.71828\n", *varMapDurationFloat64)
}


func TestParseMapDurationFloat32(t *testing.T) {
	// T := reflect.TypeOf(mapDurationFloat32Value{}).Elem()
	var varMapDurationFloat32 = new(mapDurationFloat32Value)
	Var(varMapDurationFloat32, "varMapDurationFloat32", "Use mapDurationFloat32")
  varMapDurationFloat32.Set("1h2m3s:2.71828")
  if varMapDurationFloat32.String() != "1h2m3s:2.71828" {
    t.Fatalf("%v %v %v", varMapDurationFloat32,
      varMapDurationFloat32.String(), "1h2m3s:2.71828")
  }
	// fmt.Printf("float32 mapDurationFloat32\n%v\n1h2m3s:2.71828\n", *varMapDurationFloat32)
}


func TestParseMapDurationBool(t *testing.T) {
	// T := reflect.TypeOf(mapDurationBoolValue{}).Elem()
	var varMapDurationBool = new(mapDurationBoolValue)
	Var(varMapDurationBool, "varMapDurationBool", "Use mapDurationBool")
  varMapDurationBool.Set("1h2m3s:true")
  if varMapDurationBool.String() != "1h2m3s:true" {
    t.Fatalf("%v %v %v", varMapDurationBool,
      varMapDurationBool.String(), "1h2m3s:true")
  }
	// fmt.Printf("bool mapDurationBool\n%v\n1h2m3s:true\n", *varMapDurationBool)
}


func TestParseMapDurationString(t *testing.T) {
	// T := reflect.TypeOf(mapDurationStringValue{}).Elem()
	var varMapDurationString = new(mapDurationStringValue)
	Var(varMapDurationString, "varMapDurationString", "Use mapDurationString")
  varMapDurationString.Set("1h2m3s:one")
  if varMapDurationString.String() != "1h2m3s:one" {
    t.Fatalf("%v %v %v", varMapDurationString,
      varMapDurationString.String(), "1h2m3s:one")
  }
	// fmt.Printf("string mapDurationString\n%v\n1h2m3s:one\n", *varMapDurationString)
}


func TestParseMapIntDuration(t *testing.T) {
	// T := reflect.TypeOf(mapIntDurationValue{}).Elem()
	var varMapIntDuration = new(mapIntDurationValue)
	Var(varMapIntDuration, "varMapIntDuration", "Use mapIntDuration")
  varMapIntDuration.Set("-1:1h2m3s")
  if varMapIntDuration.String() != "-1:1h2m3s" {
    t.Fatalf("%v %v %v", varMapIntDuration,
      varMapIntDuration.String(), "-1:1h2m3s")
  }
	// fmt.Printf("duration mapIntDuration\n%v\n-1:1h2m3s\n", *varMapIntDuration)
}


func TestParseMapIntInt(t *testing.T) {
	// T := reflect.TypeOf(mapIntIntValue{}).Elem()
	var varMapIntInt = new(mapIntIntValue)
	Var(varMapIntInt, "varMapIntInt", "Use mapIntInt")
  varMapIntInt.Set("-1:-1")
  if varMapIntInt.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapIntInt,
      varMapIntInt.String(), "-1:-1")
  }
	// fmt.Printf("int mapIntInt\n%v\n-1:-1\n", *varMapIntInt)
}


func TestParseMapIntInt8(t *testing.T) {
	// T := reflect.TypeOf(mapIntInt8Value{}).Elem()
	var varMapIntInt8 = new(mapIntInt8Value)
	Var(varMapIntInt8, "varMapIntInt8", "Use mapIntInt8")
  varMapIntInt8.Set("-1:-1")
  if varMapIntInt8.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapIntInt8,
      varMapIntInt8.String(), "-1:-1")
  }
	// fmt.Printf("int8 mapIntInt8\n%v\n-1:-1\n", *varMapIntInt8)
}


func TestParseMapIntInt16(t *testing.T) {
	// T := reflect.TypeOf(mapIntInt16Value{}).Elem()
	var varMapIntInt16 = new(mapIntInt16Value)
	Var(varMapIntInt16, "varMapIntInt16", "Use mapIntInt16")
  varMapIntInt16.Set("-1:-1")
  if varMapIntInt16.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapIntInt16,
      varMapIntInt16.String(), "-1:-1")
  }
	// fmt.Printf("int16 mapIntInt16\n%v\n-1:-1\n", *varMapIntInt16)
}


func TestParseMapIntInt32(t *testing.T) {
	// T := reflect.TypeOf(mapIntInt32Value{}).Elem()
	var varMapIntInt32 = new(mapIntInt32Value)
	Var(varMapIntInt32, "varMapIntInt32", "Use mapIntInt32")
  varMapIntInt32.Set("-1:-1")
  if varMapIntInt32.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapIntInt32,
      varMapIntInt32.String(), "-1:-1")
  }
	// fmt.Printf("int32 mapIntInt32\n%v\n-1:-1\n", *varMapIntInt32)
}


func TestParseMapIntInt64(t *testing.T) {
	// T := reflect.TypeOf(mapIntInt64Value{}).Elem()
	var varMapIntInt64 = new(mapIntInt64Value)
	Var(varMapIntInt64, "varMapIntInt64", "Use mapIntInt64")
  varMapIntInt64.Set("-1:-1")
  if varMapIntInt64.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapIntInt64,
      varMapIntInt64.String(), "-1:-1")
  }
	// fmt.Printf("int64 mapIntInt64\n%v\n-1:-1\n", *varMapIntInt64)
}


func TestParseMapIntUint(t *testing.T) {
	// T := reflect.TypeOf(mapIntUintValue{}).Elem()
	var varMapIntUint = new(mapIntUintValue)
	Var(varMapIntUint, "varMapIntUint", "Use mapIntUint")
  varMapIntUint.Set("-1:2")
  if varMapIntUint.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapIntUint,
      varMapIntUint.String(), "-1:2")
  }
	// fmt.Printf("uint mapIntUint\n%v\n-1:2\n", *varMapIntUint)
}


func TestParseMapIntUint8(t *testing.T) {
	// T := reflect.TypeOf(mapIntUint8Value{}).Elem()
	var varMapIntUint8 = new(mapIntUint8Value)
	Var(varMapIntUint8, "varMapIntUint8", "Use mapIntUint8")
  varMapIntUint8.Set("-1:2")
  if varMapIntUint8.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapIntUint8,
      varMapIntUint8.String(), "-1:2")
  }
	// fmt.Printf("uint8 mapIntUint8\n%v\n-1:2\n", *varMapIntUint8)
}


func TestParseMapIntUint16(t *testing.T) {
	// T := reflect.TypeOf(mapIntUint16Value{}).Elem()
	var varMapIntUint16 = new(mapIntUint16Value)
	Var(varMapIntUint16, "varMapIntUint16", "Use mapIntUint16")
  varMapIntUint16.Set("-1:2")
  if varMapIntUint16.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapIntUint16,
      varMapIntUint16.String(), "-1:2")
  }
	// fmt.Printf("uint16 mapIntUint16\n%v\n-1:2\n", *varMapIntUint16)
}


func TestParseMapIntUint32(t *testing.T) {
	// T := reflect.TypeOf(mapIntUint32Value{}).Elem()
	var varMapIntUint32 = new(mapIntUint32Value)
	Var(varMapIntUint32, "varMapIntUint32", "Use mapIntUint32")
  varMapIntUint32.Set("-1:2")
  if varMapIntUint32.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapIntUint32,
      varMapIntUint32.String(), "-1:2")
  }
	// fmt.Printf("uint32 mapIntUint32\n%v\n-1:2\n", *varMapIntUint32)
}


func TestParseMapIntUint64(t *testing.T) {
	// T := reflect.TypeOf(mapIntUint64Value{}).Elem()
	var varMapIntUint64 = new(mapIntUint64Value)
	Var(varMapIntUint64, "varMapIntUint64", "Use mapIntUint64")
  varMapIntUint64.Set("-1:2")
  if varMapIntUint64.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapIntUint64,
      varMapIntUint64.String(), "-1:2")
  }
	// fmt.Printf("uint64 mapIntUint64\n%v\n-1:2\n", *varMapIntUint64)
}


func TestParseMapIntFloat64(t *testing.T) {
	// T := reflect.TypeOf(mapIntFloat64Value{}).Elem()
	var varMapIntFloat64 = new(mapIntFloat64Value)
	Var(varMapIntFloat64, "varMapIntFloat64", "Use mapIntFloat64")
  varMapIntFloat64.Set("-1:2.71828")
  if varMapIntFloat64.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapIntFloat64,
      varMapIntFloat64.String(), "-1:2.71828")
  }
	// fmt.Printf("float64 mapIntFloat64\n%v\n-1:2.71828\n", *varMapIntFloat64)
}


func TestParseMapIntFloat32(t *testing.T) {
	// T := reflect.TypeOf(mapIntFloat32Value{}).Elem()
	var varMapIntFloat32 = new(mapIntFloat32Value)
	Var(varMapIntFloat32, "varMapIntFloat32", "Use mapIntFloat32")
  varMapIntFloat32.Set("-1:2.71828")
  if varMapIntFloat32.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapIntFloat32,
      varMapIntFloat32.String(), "-1:2.71828")
  }
	// fmt.Printf("float32 mapIntFloat32\n%v\n-1:2.71828\n", *varMapIntFloat32)
}


func TestParseMapIntBool(t *testing.T) {
	// T := reflect.TypeOf(mapIntBoolValue{}).Elem()
	var varMapIntBool = new(mapIntBoolValue)
	Var(varMapIntBool, "varMapIntBool", "Use mapIntBool")
  varMapIntBool.Set("-1:true")
  if varMapIntBool.String() != "-1:true" {
    t.Fatalf("%v %v %v", varMapIntBool,
      varMapIntBool.String(), "-1:true")
  }
	// fmt.Printf("bool mapIntBool\n%v\n-1:true\n", *varMapIntBool)
}


func TestParseMapIntString(t *testing.T) {
	// T := reflect.TypeOf(mapIntStringValue{}).Elem()
	var varMapIntString = new(mapIntStringValue)
	Var(varMapIntString, "varMapIntString", "Use mapIntString")
  varMapIntString.Set("-1:one")
  if varMapIntString.String() != "-1:one" {
    t.Fatalf("%v %v %v", varMapIntString,
      varMapIntString.String(), "-1:one")
  }
	// fmt.Printf("string mapIntString\n%v\n-1:one\n", *varMapIntString)
}


func TestParseMapInt8Duration(t *testing.T) {
	// T := reflect.TypeOf(mapInt8DurationValue{}).Elem()
	var varMapInt8Duration = new(mapInt8DurationValue)
	Var(varMapInt8Duration, "varMapInt8Duration", "Use mapInt8Duration")
  varMapInt8Duration.Set("-1:1h2m3s")
  if varMapInt8Duration.String() != "-1:1h2m3s" {
    t.Fatalf("%v %v %v", varMapInt8Duration,
      varMapInt8Duration.String(), "-1:1h2m3s")
  }
	// fmt.Printf("duration mapInt8Duration\n%v\n-1:1h2m3s\n", *varMapInt8Duration)
}


func TestParseMapInt8Int(t *testing.T) {
	// T := reflect.TypeOf(mapInt8IntValue{}).Elem()
	var varMapInt8Int = new(mapInt8IntValue)
	Var(varMapInt8Int, "varMapInt8Int", "Use mapInt8Int")
  varMapInt8Int.Set("-1:-1")
  if varMapInt8Int.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt8Int,
      varMapInt8Int.String(), "-1:-1")
  }
	// fmt.Printf("int mapInt8Int\n%v\n-1:-1\n", *varMapInt8Int)
}


func TestParseMapInt8Int8(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Int8Value{}).Elem()
	var varMapInt8Int8 = new(mapInt8Int8Value)
	Var(varMapInt8Int8, "varMapInt8Int8", "Use mapInt8Int8")
  varMapInt8Int8.Set("-1:-1")
  if varMapInt8Int8.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt8Int8,
      varMapInt8Int8.String(), "-1:-1")
  }
	// fmt.Printf("int8 mapInt8Int8\n%v\n-1:-1\n", *varMapInt8Int8)
}


func TestParseMapInt8Int16(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Int16Value{}).Elem()
	var varMapInt8Int16 = new(mapInt8Int16Value)
	Var(varMapInt8Int16, "varMapInt8Int16", "Use mapInt8Int16")
  varMapInt8Int16.Set("-1:-1")
  if varMapInt8Int16.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt8Int16,
      varMapInt8Int16.String(), "-1:-1")
  }
	// fmt.Printf("int16 mapInt8Int16\n%v\n-1:-1\n", *varMapInt8Int16)
}


func TestParseMapInt8Int32(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Int32Value{}).Elem()
	var varMapInt8Int32 = new(mapInt8Int32Value)
	Var(varMapInt8Int32, "varMapInt8Int32", "Use mapInt8Int32")
  varMapInt8Int32.Set("-1:-1")
  if varMapInt8Int32.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt8Int32,
      varMapInt8Int32.String(), "-1:-1")
  }
	// fmt.Printf("int32 mapInt8Int32\n%v\n-1:-1\n", *varMapInt8Int32)
}


func TestParseMapInt8Int64(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Int64Value{}).Elem()
	var varMapInt8Int64 = new(mapInt8Int64Value)
	Var(varMapInt8Int64, "varMapInt8Int64", "Use mapInt8Int64")
  varMapInt8Int64.Set("-1:-1")
  if varMapInt8Int64.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt8Int64,
      varMapInt8Int64.String(), "-1:-1")
  }
	// fmt.Printf("int64 mapInt8Int64\n%v\n-1:-1\n", *varMapInt8Int64)
}


func TestParseMapInt8Uint(t *testing.T) {
	// T := reflect.TypeOf(mapInt8UintValue{}).Elem()
	var varMapInt8Uint = new(mapInt8UintValue)
	Var(varMapInt8Uint, "varMapInt8Uint", "Use mapInt8Uint")
  varMapInt8Uint.Set("-1:2")
  if varMapInt8Uint.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt8Uint,
      varMapInt8Uint.String(), "-1:2")
  }
	// fmt.Printf("uint mapInt8Uint\n%v\n-1:2\n", *varMapInt8Uint)
}


func TestParseMapInt8Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Uint8Value{}).Elem()
	var varMapInt8Uint8 = new(mapInt8Uint8Value)
	Var(varMapInt8Uint8, "varMapInt8Uint8", "Use mapInt8Uint8")
  varMapInt8Uint8.Set("-1:2")
  if varMapInt8Uint8.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt8Uint8,
      varMapInt8Uint8.String(), "-1:2")
  }
	// fmt.Printf("uint8 mapInt8Uint8\n%v\n-1:2\n", *varMapInt8Uint8)
}


func TestParseMapInt8Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Uint16Value{}).Elem()
	var varMapInt8Uint16 = new(mapInt8Uint16Value)
	Var(varMapInt8Uint16, "varMapInt8Uint16", "Use mapInt8Uint16")
  varMapInt8Uint16.Set("-1:2")
  if varMapInt8Uint16.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt8Uint16,
      varMapInt8Uint16.String(), "-1:2")
  }
	// fmt.Printf("uint16 mapInt8Uint16\n%v\n-1:2\n", *varMapInt8Uint16)
}


func TestParseMapInt8Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Uint32Value{}).Elem()
	var varMapInt8Uint32 = new(mapInt8Uint32Value)
	Var(varMapInt8Uint32, "varMapInt8Uint32", "Use mapInt8Uint32")
  varMapInt8Uint32.Set("-1:2")
  if varMapInt8Uint32.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt8Uint32,
      varMapInt8Uint32.String(), "-1:2")
  }
	// fmt.Printf("uint32 mapInt8Uint32\n%v\n-1:2\n", *varMapInt8Uint32)
}


func TestParseMapInt8Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Uint64Value{}).Elem()
	var varMapInt8Uint64 = new(mapInt8Uint64Value)
	Var(varMapInt8Uint64, "varMapInt8Uint64", "Use mapInt8Uint64")
  varMapInt8Uint64.Set("-1:2")
  if varMapInt8Uint64.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt8Uint64,
      varMapInt8Uint64.String(), "-1:2")
  }
	// fmt.Printf("uint64 mapInt8Uint64\n%v\n-1:2\n", *varMapInt8Uint64)
}


func TestParseMapInt8Float64(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Float64Value{}).Elem()
	var varMapInt8Float64 = new(mapInt8Float64Value)
	Var(varMapInt8Float64, "varMapInt8Float64", "Use mapInt8Float64")
  varMapInt8Float64.Set("-1:2.71828")
  if varMapInt8Float64.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt8Float64,
      varMapInt8Float64.String(), "-1:2.71828")
  }
	// fmt.Printf("float64 mapInt8Float64\n%v\n-1:2.71828\n", *varMapInt8Float64)
}


func TestParseMapInt8Float32(t *testing.T) {
	// T := reflect.TypeOf(mapInt8Float32Value{}).Elem()
	var varMapInt8Float32 = new(mapInt8Float32Value)
	Var(varMapInt8Float32, "varMapInt8Float32", "Use mapInt8Float32")
  varMapInt8Float32.Set("-1:2.71828")
  if varMapInt8Float32.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt8Float32,
      varMapInt8Float32.String(), "-1:2.71828")
  }
	// fmt.Printf("float32 mapInt8Float32\n%v\n-1:2.71828\n", *varMapInt8Float32)
}


func TestParseMapInt8Bool(t *testing.T) {
	// T := reflect.TypeOf(mapInt8BoolValue{}).Elem()
	var varMapInt8Bool = new(mapInt8BoolValue)
	Var(varMapInt8Bool, "varMapInt8Bool", "Use mapInt8Bool")
  varMapInt8Bool.Set("-1:true")
  if varMapInt8Bool.String() != "-1:true" {
    t.Fatalf("%v %v %v", varMapInt8Bool,
      varMapInt8Bool.String(), "-1:true")
  }
	// fmt.Printf("bool mapInt8Bool\n%v\n-1:true\n", *varMapInt8Bool)
}


func TestParseMapInt8String(t *testing.T) {
	// T := reflect.TypeOf(mapInt8StringValue{}).Elem()
	var varMapInt8String = new(mapInt8StringValue)
	Var(varMapInt8String, "varMapInt8String", "Use mapInt8String")
  varMapInt8String.Set("-1:one")
  if varMapInt8String.String() != "-1:one" {
    t.Fatalf("%v %v %v", varMapInt8String,
      varMapInt8String.String(), "-1:one")
  }
	// fmt.Printf("string mapInt8String\n%v\n-1:one\n", *varMapInt8String)
}


func TestParseMapInt16Duration(t *testing.T) {
	// T := reflect.TypeOf(mapInt16DurationValue{}).Elem()
	var varMapInt16Duration = new(mapInt16DurationValue)
	Var(varMapInt16Duration, "varMapInt16Duration", "Use mapInt16Duration")
  varMapInt16Duration.Set("-1:1h2m3s")
  if varMapInt16Duration.String() != "-1:1h2m3s" {
    t.Fatalf("%v %v %v", varMapInt16Duration,
      varMapInt16Duration.String(), "-1:1h2m3s")
  }
	// fmt.Printf("duration mapInt16Duration\n%v\n-1:1h2m3s\n", *varMapInt16Duration)
}


func TestParseMapInt16Int(t *testing.T) {
	// T := reflect.TypeOf(mapInt16IntValue{}).Elem()
	var varMapInt16Int = new(mapInt16IntValue)
	Var(varMapInt16Int, "varMapInt16Int", "Use mapInt16Int")
  varMapInt16Int.Set("-1:-1")
  if varMapInt16Int.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt16Int,
      varMapInt16Int.String(), "-1:-1")
  }
	// fmt.Printf("int mapInt16Int\n%v\n-1:-1\n", *varMapInt16Int)
}


func TestParseMapInt16Int8(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Int8Value{}).Elem()
	var varMapInt16Int8 = new(mapInt16Int8Value)
	Var(varMapInt16Int8, "varMapInt16Int8", "Use mapInt16Int8")
  varMapInt16Int8.Set("-1:-1")
  if varMapInt16Int8.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt16Int8,
      varMapInt16Int8.String(), "-1:-1")
  }
	// fmt.Printf("int8 mapInt16Int8\n%v\n-1:-1\n", *varMapInt16Int8)
}


func TestParseMapInt16Int16(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Int16Value{}).Elem()
	var varMapInt16Int16 = new(mapInt16Int16Value)
	Var(varMapInt16Int16, "varMapInt16Int16", "Use mapInt16Int16")
  varMapInt16Int16.Set("-1:-1")
  if varMapInt16Int16.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt16Int16,
      varMapInt16Int16.String(), "-1:-1")
  }
	// fmt.Printf("int16 mapInt16Int16\n%v\n-1:-1\n", *varMapInt16Int16)
}


func TestParseMapInt16Int32(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Int32Value{}).Elem()
	var varMapInt16Int32 = new(mapInt16Int32Value)
	Var(varMapInt16Int32, "varMapInt16Int32", "Use mapInt16Int32")
  varMapInt16Int32.Set("-1:-1")
  if varMapInt16Int32.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt16Int32,
      varMapInt16Int32.String(), "-1:-1")
  }
	// fmt.Printf("int32 mapInt16Int32\n%v\n-1:-1\n", *varMapInt16Int32)
}


func TestParseMapInt16Int64(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Int64Value{}).Elem()
	var varMapInt16Int64 = new(mapInt16Int64Value)
	Var(varMapInt16Int64, "varMapInt16Int64", "Use mapInt16Int64")
  varMapInt16Int64.Set("-1:-1")
  if varMapInt16Int64.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt16Int64,
      varMapInt16Int64.String(), "-1:-1")
  }
	// fmt.Printf("int64 mapInt16Int64\n%v\n-1:-1\n", *varMapInt16Int64)
}


func TestParseMapInt16Uint(t *testing.T) {
	// T := reflect.TypeOf(mapInt16UintValue{}).Elem()
	var varMapInt16Uint = new(mapInt16UintValue)
	Var(varMapInt16Uint, "varMapInt16Uint", "Use mapInt16Uint")
  varMapInt16Uint.Set("-1:2")
  if varMapInt16Uint.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt16Uint,
      varMapInt16Uint.String(), "-1:2")
  }
	// fmt.Printf("uint mapInt16Uint\n%v\n-1:2\n", *varMapInt16Uint)
}


func TestParseMapInt16Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Uint8Value{}).Elem()
	var varMapInt16Uint8 = new(mapInt16Uint8Value)
	Var(varMapInt16Uint8, "varMapInt16Uint8", "Use mapInt16Uint8")
  varMapInt16Uint8.Set("-1:2")
  if varMapInt16Uint8.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt16Uint8,
      varMapInt16Uint8.String(), "-1:2")
  }
	// fmt.Printf("uint8 mapInt16Uint8\n%v\n-1:2\n", *varMapInt16Uint8)
}


func TestParseMapInt16Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Uint16Value{}).Elem()
	var varMapInt16Uint16 = new(mapInt16Uint16Value)
	Var(varMapInt16Uint16, "varMapInt16Uint16", "Use mapInt16Uint16")
  varMapInt16Uint16.Set("-1:2")
  if varMapInt16Uint16.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt16Uint16,
      varMapInt16Uint16.String(), "-1:2")
  }
	// fmt.Printf("uint16 mapInt16Uint16\n%v\n-1:2\n", *varMapInt16Uint16)
}


func TestParseMapInt16Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Uint32Value{}).Elem()
	var varMapInt16Uint32 = new(mapInt16Uint32Value)
	Var(varMapInt16Uint32, "varMapInt16Uint32", "Use mapInt16Uint32")
  varMapInt16Uint32.Set("-1:2")
  if varMapInt16Uint32.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt16Uint32,
      varMapInt16Uint32.String(), "-1:2")
  }
	// fmt.Printf("uint32 mapInt16Uint32\n%v\n-1:2\n", *varMapInt16Uint32)
}


func TestParseMapInt16Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Uint64Value{}).Elem()
	var varMapInt16Uint64 = new(mapInt16Uint64Value)
	Var(varMapInt16Uint64, "varMapInt16Uint64", "Use mapInt16Uint64")
  varMapInt16Uint64.Set("-1:2")
  if varMapInt16Uint64.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt16Uint64,
      varMapInt16Uint64.String(), "-1:2")
  }
	// fmt.Printf("uint64 mapInt16Uint64\n%v\n-1:2\n", *varMapInt16Uint64)
}


func TestParseMapInt16Float64(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Float64Value{}).Elem()
	var varMapInt16Float64 = new(mapInt16Float64Value)
	Var(varMapInt16Float64, "varMapInt16Float64", "Use mapInt16Float64")
  varMapInt16Float64.Set("-1:2.71828")
  if varMapInt16Float64.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt16Float64,
      varMapInt16Float64.String(), "-1:2.71828")
  }
	// fmt.Printf("float64 mapInt16Float64\n%v\n-1:2.71828\n", *varMapInt16Float64)
}


func TestParseMapInt16Float32(t *testing.T) {
	// T := reflect.TypeOf(mapInt16Float32Value{}).Elem()
	var varMapInt16Float32 = new(mapInt16Float32Value)
	Var(varMapInt16Float32, "varMapInt16Float32", "Use mapInt16Float32")
  varMapInt16Float32.Set("-1:2.71828")
  if varMapInt16Float32.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt16Float32,
      varMapInt16Float32.String(), "-1:2.71828")
  }
	// fmt.Printf("float32 mapInt16Float32\n%v\n-1:2.71828\n", *varMapInt16Float32)
}


func TestParseMapInt16Bool(t *testing.T) {
	// T := reflect.TypeOf(mapInt16BoolValue{}).Elem()
	var varMapInt16Bool = new(mapInt16BoolValue)
	Var(varMapInt16Bool, "varMapInt16Bool", "Use mapInt16Bool")
  varMapInt16Bool.Set("-1:true")
  if varMapInt16Bool.String() != "-1:true" {
    t.Fatalf("%v %v %v", varMapInt16Bool,
      varMapInt16Bool.String(), "-1:true")
  }
	// fmt.Printf("bool mapInt16Bool\n%v\n-1:true\n", *varMapInt16Bool)
}


func TestParseMapInt16String(t *testing.T) {
	// T := reflect.TypeOf(mapInt16StringValue{}).Elem()
	var varMapInt16String = new(mapInt16StringValue)
	Var(varMapInt16String, "varMapInt16String", "Use mapInt16String")
  varMapInt16String.Set("-1:one")
  if varMapInt16String.String() != "-1:one" {
    t.Fatalf("%v %v %v", varMapInt16String,
      varMapInt16String.String(), "-1:one")
  }
	// fmt.Printf("string mapInt16String\n%v\n-1:one\n", *varMapInt16String)
}


func TestParseMapInt32Duration(t *testing.T) {
	// T := reflect.TypeOf(mapInt32DurationValue{}).Elem()
	var varMapInt32Duration = new(mapInt32DurationValue)
	Var(varMapInt32Duration, "varMapInt32Duration", "Use mapInt32Duration")
  varMapInt32Duration.Set("-1:1h2m3s")
  if varMapInt32Duration.String() != "-1:1h2m3s" {
    t.Fatalf("%v %v %v", varMapInt32Duration,
      varMapInt32Duration.String(), "-1:1h2m3s")
  }
	// fmt.Printf("duration mapInt32Duration\n%v\n-1:1h2m3s\n", *varMapInt32Duration)
}


func TestParseMapInt32Int(t *testing.T) {
	// T := reflect.TypeOf(mapInt32IntValue{}).Elem()
	var varMapInt32Int = new(mapInt32IntValue)
	Var(varMapInt32Int, "varMapInt32Int", "Use mapInt32Int")
  varMapInt32Int.Set("-1:-1")
  if varMapInt32Int.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt32Int,
      varMapInt32Int.String(), "-1:-1")
  }
	// fmt.Printf("int mapInt32Int\n%v\n-1:-1\n", *varMapInt32Int)
}


func TestParseMapInt32Int8(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Int8Value{}).Elem()
	var varMapInt32Int8 = new(mapInt32Int8Value)
	Var(varMapInt32Int8, "varMapInt32Int8", "Use mapInt32Int8")
  varMapInt32Int8.Set("-1:-1")
  if varMapInt32Int8.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt32Int8,
      varMapInt32Int8.String(), "-1:-1")
  }
	// fmt.Printf("int8 mapInt32Int8\n%v\n-1:-1\n", *varMapInt32Int8)
}


func TestParseMapInt32Int16(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Int16Value{}).Elem()
	var varMapInt32Int16 = new(mapInt32Int16Value)
	Var(varMapInt32Int16, "varMapInt32Int16", "Use mapInt32Int16")
  varMapInt32Int16.Set("-1:-1")
  if varMapInt32Int16.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt32Int16,
      varMapInt32Int16.String(), "-1:-1")
  }
	// fmt.Printf("int16 mapInt32Int16\n%v\n-1:-1\n", *varMapInt32Int16)
}


func TestParseMapInt32Int32(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Int32Value{}).Elem()
	var varMapInt32Int32 = new(mapInt32Int32Value)
	Var(varMapInt32Int32, "varMapInt32Int32", "Use mapInt32Int32")
  varMapInt32Int32.Set("-1:-1")
  if varMapInt32Int32.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt32Int32,
      varMapInt32Int32.String(), "-1:-1")
  }
	// fmt.Printf("int32 mapInt32Int32\n%v\n-1:-1\n", *varMapInt32Int32)
}


func TestParseMapInt32Int64(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Int64Value{}).Elem()
	var varMapInt32Int64 = new(mapInt32Int64Value)
	Var(varMapInt32Int64, "varMapInt32Int64", "Use mapInt32Int64")
  varMapInt32Int64.Set("-1:-1")
  if varMapInt32Int64.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt32Int64,
      varMapInt32Int64.String(), "-1:-1")
  }
	// fmt.Printf("int64 mapInt32Int64\n%v\n-1:-1\n", *varMapInt32Int64)
}


func TestParseMapInt32Uint(t *testing.T) {
	// T := reflect.TypeOf(mapInt32UintValue{}).Elem()
	var varMapInt32Uint = new(mapInt32UintValue)
	Var(varMapInt32Uint, "varMapInt32Uint", "Use mapInt32Uint")
  varMapInt32Uint.Set("-1:2")
  if varMapInt32Uint.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt32Uint,
      varMapInt32Uint.String(), "-1:2")
  }
	// fmt.Printf("uint mapInt32Uint\n%v\n-1:2\n", *varMapInt32Uint)
}


func TestParseMapInt32Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Uint8Value{}).Elem()
	var varMapInt32Uint8 = new(mapInt32Uint8Value)
	Var(varMapInt32Uint8, "varMapInt32Uint8", "Use mapInt32Uint8")
  varMapInt32Uint8.Set("-1:2")
  if varMapInt32Uint8.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt32Uint8,
      varMapInt32Uint8.String(), "-1:2")
  }
	// fmt.Printf("uint8 mapInt32Uint8\n%v\n-1:2\n", *varMapInt32Uint8)
}


func TestParseMapInt32Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Uint16Value{}).Elem()
	var varMapInt32Uint16 = new(mapInt32Uint16Value)
	Var(varMapInt32Uint16, "varMapInt32Uint16", "Use mapInt32Uint16")
  varMapInt32Uint16.Set("-1:2")
  if varMapInt32Uint16.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt32Uint16,
      varMapInt32Uint16.String(), "-1:2")
  }
	// fmt.Printf("uint16 mapInt32Uint16\n%v\n-1:2\n", *varMapInt32Uint16)
}


func TestParseMapInt32Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Uint32Value{}).Elem()
	var varMapInt32Uint32 = new(mapInt32Uint32Value)
	Var(varMapInt32Uint32, "varMapInt32Uint32", "Use mapInt32Uint32")
  varMapInt32Uint32.Set("-1:2")
  if varMapInt32Uint32.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt32Uint32,
      varMapInt32Uint32.String(), "-1:2")
  }
	// fmt.Printf("uint32 mapInt32Uint32\n%v\n-1:2\n", *varMapInt32Uint32)
}


func TestParseMapInt32Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Uint64Value{}).Elem()
	var varMapInt32Uint64 = new(mapInt32Uint64Value)
	Var(varMapInt32Uint64, "varMapInt32Uint64", "Use mapInt32Uint64")
  varMapInt32Uint64.Set("-1:2")
  if varMapInt32Uint64.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt32Uint64,
      varMapInt32Uint64.String(), "-1:2")
  }
	// fmt.Printf("uint64 mapInt32Uint64\n%v\n-1:2\n", *varMapInt32Uint64)
}


func TestParseMapInt32Float64(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Float64Value{}).Elem()
	var varMapInt32Float64 = new(mapInt32Float64Value)
	Var(varMapInt32Float64, "varMapInt32Float64", "Use mapInt32Float64")
  varMapInt32Float64.Set("-1:2.71828")
  if varMapInt32Float64.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt32Float64,
      varMapInt32Float64.String(), "-1:2.71828")
  }
	// fmt.Printf("float64 mapInt32Float64\n%v\n-1:2.71828\n", *varMapInt32Float64)
}


func TestParseMapInt32Float32(t *testing.T) {
	// T := reflect.TypeOf(mapInt32Float32Value{}).Elem()
	var varMapInt32Float32 = new(mapInt32Float32Value)
	Var(varMapInt32Float32, "varMapInt32Float32", "Use mapInt32Float32")
  varMapInt32Float32.Set("-1:2.71828")
  if varMapInt32Float32.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt32Float32,
      varMapInt32Float32.String(), "-1:2.71828")
  }
	// fmt.Printf("float32 mapInt32Float32\n%v\n-1:2.71828\n", *varMapInt32Float32)
}


func TestParseMapInt32Bool(t *testing.T) {
	// T := reflect.TypeOf(mapInt32BoolValue{}).Elem()
	var varMapInt32Bool = new(mapInt32BoolValue)
	Var(varMapInt32Bool, "varMapInt32Bool", "Use mapInt32Bool")
  varMapInt32Bool.Set("-1:true")
  if varMapInt32Bool.String() != "-1:true" {
    t.Fatalf("%v %v %v", varMapInt32Bool,
      varMapInt32Bool.String(), "-1:true")
  }
	// fmt.Printf("bool mapInt32Bool\n%v\n-1:true\n", *varMapInt32Bool)
}


func TestParseMapInt32String(t *testing.T) {
	// T := reflect.TypeOf(mapInt32StringValue{}).Elem()
	var varMapInt32String = new(mapInt32StringValue)
	Var(varMapInt32String, "varMapInt32String", "Use mapInt32String")
  varMapInt32String.Set("-1:one")
  if varMapInt32String.String() != "-1:one" {
    t.Fatalf("%v %v %v", varMapInt32String,
      varMapInt32String.String(), "-1:one")
  }
	// fmt.Printf("string mapInt32String\n%v\n-1:one\n", *varMapInt32String)
}


func TestParseMapInt64Duration(t *testing.T) {
	// T := reflect.TypeOf(mapInt64DurationValue{}).Elem()
	var varMapInt64Duration = new(mapInt64DurationValue)
	Var(varMapInt64Duration, "varMapInt64Duration", "Use mapInt64Duration")
  varMapInt64Duration.Set("-1:1h2m3s")
  if varMapInt64Duration.String() != "-1:1h2m3s" {
    t.Fatalf("%v %v %v", varMapInt64Duration,
      varMapInt64Duration.String(), "-1:1h2m3s")
  }
	// fmt.Printf("duration mapInt64Duration\n%v\n-1:1h2m3s\n", *varMapInt64Duration)
}


func TestParseMapInt64Int(t *testing.T) {
	// T := reflect.TypeOf(mapInt64IntValue{}).Elem()
	var varMapInt64Int = new(mapInt64IntValue)
	Var(varMapInt64Int, "varMapInt64Int", "Use mapInt64Int")
  varMapInt64Int.Set("-1:-1")
  if varMapInt64Int.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt64Int,
      varMapInt64Int.String(), "-1:-1")
  }
	// fmt.Printf("int mapInt64Int\n%v\n-1:-1\n", *varMapInt64Int)
}


func TestParseMapInt64Int8(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Int8Value{}).Elem()
	var varMapInt64Int8 = new(mapInt64Int8Value)
	Var(varMapInt64Int8, "varMapInt64Int8", "Use mapInt64Int8")
  varMapInt64Int8.Set("-1:-1")
  if varMapInt64Int8.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt64Int8,
      varMapInt64Int8.String(), "-1:-1")
  }
	// fmt.Printf("int8 mapInt64Int8\n%v\n-1:-1\n", *varMapInt64Int8)
}


func TestParseMapInt64Int16(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Int16Value{}).Elem()
	var varMapInt64Int16 = new(mapInt64Int16Value)
	Var(varMapInt64Int16, "varMapInt64Int16", "Use mapInt64Int16")
  varMapInt64Int16.Set("-1:-1")
  if varMapInt64Int16.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt64Int16,
      varMapInt64Int16.String(), "-1:-1")
  }
	// fmt.Printf("int16 mapInt64Int16\n%v\n-1:-1\n", *varMapInt64Int16)
}


func TestParseMapInt64Int32(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Int32Value{}).Elem()
	var varMapInt64Int32 = new(mapInt64Int32Value)
	Var(varMapInt64Int32, "varMapInt64Int32", "Use mapInt64Int32")
  varMapInt64Int32.Set("-1:-1")
  if varMapInt64Int32.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt64Int32,
      varMapInt64Int32.String(), "-1:-1")
  }
	// fmt.Printf("int32 mapInt64Int32\n%v\n-1:-1\n", *varMapInt64Int32)
}


func TestParseMapInt64Int64(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Int64Value{}).Elem()
	var varMapInt64Int64 = new(mapInt64Int64Value)
	Var(varMapInt64Int64, "varMapInt64Int64", "Use mapInt64Int64")
  varMapInt64Int64.Set("-1:-1")
  if varMapInt64Int64.String() != "-1:-1" {
    t.Fatalf("%v %v %v", varMapInt64Int64,
      varMapInt64Int64.String(), "-1:-1")
  }
	// fmt.Printf("int64 mapInt64Int64\n%v\n-1:-1\n", *varMapInt64Int64)
}


func TestParseMapInt64Uint(t *testing.T) {
	// T := reflect.TypeOf(mapInt64UintValue{}).Elem()
	var varMapInt64Uint = new(mapInt64UintValue)
	Var(varMapInt64Uint, "varMapInt64Uint", "Use mapInt64Uint")
  varMapInt64Uint.Set("-1:2")
  if varMapInt64Uint.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt64Uint,
      varMapInt64Uint.String(), "-1:2")
  }
	// fmt.Printf("uint mapInt64Uint\n%v\n-1:2\n", *varMapInt64Uint)
}


func TestParseMapInt64Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Uint8Value{}).Elem()
	var varMapInt64Uint8 = new(mapInt64Uint8Value)
	Var(varMapInt64Uint8, "varMapInt64Uint8", "Use mapInt64Uint8")
  varMapInt64Uint8.Set("-1:2")
  if varMapInt64Uint8.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt64Uint8,
      varMapInt64Uint8.String(), "-1:2")
  }
	// fmt.Printf("uint8 mapInt64Uint8\n%v\n-1:2\n", *varMapInt64Uint8)
}


func TestParseMapInt64Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Uint16Value{}).Elem()
	var varMapInt64Uint16 = new(mapInt64Uint16Value)
	Var(varMapInt64Uint16, "varMapInt64Uint16", "Use mapInt64Uint16")
  varMapInt64Uint16.Set("-1:2")
  if varMapInt64Uint16.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt64Uint16,
      varMapInt64Uint16.String(), "-1:2")
  }
	// fmt.Printf("uint16 mapInt64Uint16\n%v\n-1:2\n", *varMapInt64Uint16)
}


func TestParseMapInt64Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Uint32Value{}).Elem()
	var varMapInt64Uint32 = new(mapInt64Uint32Value)
	Var(varMapInt64Uint32, "varMapInt64Uint32", "Use mapInt64Uint32")
  varMapInt64Uint32.Set("-1:2")
  if varMapInt64Uint32.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt64Uint32,
      varMapInt64Uint32.String(), "-1:2")
  }
	// fmt.Printf("uint32 mapInt64Uint32\n%v\n-1:2\n", *varMapInt64Uint32)
}


func TestParseMapInt64Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Uint64Value{}).Elem()
	var varMapInt64Uint64 = new(mapInt64Uint64Value)
	Var(varMapInt64Uint64, "varMapInt64Uint64", "Use mapInt64Uint64")
  varMapInt64Uint64.Set("-1:2")
  if varMapInt64Uint64.String() != "-1:2" {
    t.Fatalf("%v %v %v", varMapInt64Uint64,
      varMapInt64Uint64.String(), "-1:2")
  }
	// fmt.Printf("uint64 mapInt64Uint64\n%v\n-1:2\n", *varMapInt64Uint64)
}


func TestParseMapInt64Float64(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Float64Value{}).Elem()
	var varMapInt64Float64 = new(mapInt64Float64Value)
	Var(varMapInt64Float64, "varMapInt64Float64", "Use mapInt64Float64")
  varMapInt64Float64.Set("-1:2.71828")
  if varMapInt64Float64.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt64Float64,
      varMapInt64Float64.String(), "-1:2.71828")
  }
	// fmt.Printf("float64 mapInt64Float64\n%v\n-1:2.71828\n", *varMapInt64Float64)
}


func TestParseMapInt64Float32(t *testing.T) {
	// T := reflect.TypeOf(mapInt64Float32Value{}).Elem()
	var varMapInt64Float32 = new(mapInt64Float32Value)
	Var(varMapInt64Float32, "varMapInt64Float32", "Use mapInt64Float32")
  varMapInt64Float32.Set("-1:2.71828")
  if varMapInt64Float32.String() != "-1:2.71828" {
    t.Fatalf("%v %v %v", varMapInt64Float32,
      varMapInt64Float32.String(), "-1:2.71828")
  }
	// fmt.Printf("float32 mapInt64Float32\n%v\n-1:2.71828\n", *varMapInt64Float32)
}


func TestParseMapInt64Bool(t *testing.T) {
	// T := reflect.TypeOf(mapInt64BoolValue{}).Elem()
	var varMapInt64Bool = new(mapInt64BoolValue)
	Var(varMapInt64Bool, "varMapInt64Bool", "Use mapInt64Bool")
  varMapInt64Bool.Set("-1:true")
  if varMapInt64Bool.String() != "-1:true" {
    t.Fatalf("%v %v %v", varMapInt64Bool,
      varMapInt64Bool.String(), "-1:true")
  }
	// fmt.Printf("bool mapInt64Bool\n%v\n-1:true\n", *varMapInt64Bool)
}


func TestParseMapInt64String(t *testing.T) {
	// T := reflect.TypeOf(mapInt64StringValue{}).Elem()
	var varMapInt64String = new(mapInt64StringValue)
	Var(varMapInt64String, "varMapInt64String", "Use mapInt64String")
  varMapInt64String.Set("-1:one")
  if varMapInt64String.String() != "-1:one" {
    t.Fatalf("%v %v %v", varMapInt64String,
      varMapInt64String.String(), "-1:one")
  }
	// fmt.Printf("string mapInt64String\n%v\n-1:one\n", *varMapInt64String)
}


func TestParseMapUintDuration(t *testing.T) {
	// T := reflect.TypeOf(mapUintDurationValue{}).Elem()
	var varMapUintDuration = new(mapUintDurationValue)
	Var(varMapUintDuration, "varMapUintDuration", "Use mapUintDuration")
  varMapUintDuration.Set("2:1h2m3s")
  if varMapUintDuration.String() != "2:1h2m3s" {
    t.Fatalf("%v %v %v", varMapUintDuration,
      varMapUintDuration.String(), "2:1h2m3s")
  }
	// fmt.Printf("duration mapUintDuration\n%v\n2:1h2m3s\n", *varMapUintDuration)
}


func TestParseMapUintInt(t *testing.T) {
	// T := reflect.TypeOf(mapUintIntValue{}).Elem()
	var varMapUintInt = new(mapUintIntValue)
	Var(varMapUintInt, "varMapUintInt", "Use mapUintInt")
  varMapUintInt.Set("2:-1")
  if varMapUintInt.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUintInt,
      varMapUintInt.String(), "2:-1")
  }
	// fmt.Printf("int mapUintInt\n%v\n2:-1\n", *varMapUintInt)
}


func TestParseMapUintInt8(t *testing.T) {
	// T := reflect.TypeOf(mapUintInt8Value{}).Elem()
	var varMapUintInt8 = new(mapUintInt8Value)
	Var(varMapUintInt8, "varMapUintInt8", "Use mapUintInt8")
  varMapUintInt8.Set("2:-1")
  if varMapUintInt8.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUintInt8,
      varMapUintInt8.String(), "2:-1")
  }
	// fmt.Printf("int8 mapUintInt8\n%v\n2:-1\n", *varMapUintInt8)
}


func TestParseMapUintInt16(t *testing.T) {
	// T := reflect.TypeOf(mapUintInt16Value{}).Elem()
	var varMapUintInt16 = new(mapUintInt16Value)
	Var(varMapUintInt16, "varMapUintInt16", "Use mapUintInt16")
  varMapUintInt16.Set("2:-1")
  if varMapUintInt16.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUintInt16,
      varMapUintInt16.String(), "2:-1")
  }
	// fmt.Printf("int16 mapUintInt16\n%v\n2:-1\n", *varMapUintInt16)
}


func TestParseMapUintInt32(t *testing.T) {
	// T := reflect.TypeOf(mapUintInt32Value{}).Elem()
	var varMapUintInt32 = new(mapUintInt32Value)
	Var(varMapUintInt32, "varMapUintInt32", "Use mapUintInt32")
  varMapUintInt32.Set("2:-1")
  if varMapUintInt32.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUintInt32,
      varMapUintInt32.String(), "2:-1")
  }
	// fmt.Printf("int32 mapUintInt32\n%v\n2:-1\n", *varMapUintInt32)
}


func TestParseMapUintInt64(t *testing.T) {
	// T := reflect.TypeOf(mapUintInt64Value{}).Elem()
	var varMapUintInt64 = new(mapUintInt64Value)
	Var(varMapUintInt64, "varMapUintInt64", "Use mapUintInt64")
  varMapUintInt64.Set("2:-1")
  if varMapUintInt64.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUintInt64,
      varMapUintInt64.String(), "2:-1")
  }
	// fmt.Printf("int64 mapUintInt64\n%v\n2:-1\n", *varMapUintInt64)
}


func TestParseMapUintUint(t *testing.T) {
	// T := reflect.TypeOf(mapUintUintValue{}).Elem()
	var varMapUintUint = new(mapUintUintValue)
	Var(varMapUintUint, "varMapUintUint", "Use mapUintUint")
  varMapUintUint.Set("2:2")
  if varMapUintUint.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUintUint,
      varMapUintUint.String(), "2:2")
  }
	// fmt.Printf("uint mapUintUint\n%v\n2:2\n", *varMapUintUint)
}


func TestParseMapUintUint8(t *testing.T) {
	// T := reflect.TypeOf(mapUintUint8Value{}).Elem()
	var varMapUintUint8 = new(mapUintUint8Value)
	Var(varMapUintUint8, "varMapUintUint8", "Use mapUintUint8")
  varMapUintUint8.Set("2:2")
  if varMapUintUint8.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUintUint8,
      varMapUintUint8.String(), "2:2")
  }
	// fmt.Printf("uint8 mapUintUint8\n%v\n2:2\n", *varMapUintUint8)
}


func TestParseMapUintUint16(t *testing.T) {
	// T := reflect.TypeOf(mapUintUint16Value{}).Elem()
	var varMapUintUint16 = new(mapUintUint16Value)
	Var(varMapUintUint16, "varMapUintUint16", "Use mapUintUint16")
  varMapUintUint16.Set("2:2")
  if varMapUintUint16.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUintUint16,
      varMapUintUint16.String(), "2:2")
  }
	// fmt.Printf("uint16 mapUintUint16\n%v\n2:2\n", *varMapUintUint16)
}


func TestParseMapUintUint32(t *testing.T) {
	// T := reflect.TypeOf(mapUintUint32Value{}).Elem()
	var varMapUintUint32 = new(mapUintUint32Value)
	Var(varMapUintUint32, "varMapUintUint32", "Use mapUintUint32")
  varMapUintUint32.Set("2:2")
  if varMapUintUint32.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUintUint32,
      varMapUintUint32.String(), "2:2")
  }
	// fmt.Printf("uint32 mapUintUint32\n%v\n2:2\n", *varMapUintUint32)
}


func TestParseMapUintUint64(t *testing.T) {
	// T := reflect.TypeOf(mapUintUint64Value{}).Elem()
	var varMapUintUint64 = new(mapUintUint64Value)
	Var(varMapUintUint64, "varMapUintUint64", "Use mapUintUint64")
  varMapUintUint64.Set("2:2")
  if varMapUintUint64.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUintUint64,
      varMapUintUint64.String(), "2:2")
  }
	// fmt.Printf("uint64 mapUintUint64\n%v\n2:2\n", *varMapUintUint64)
}


func TestParseMapUintFloat64(t *testing.T) {
	// T := reflect.TypeOf(mapUintFloat64Value{}).Elem()
	var varMapUintFloat64 = new(mapUintFloat64Value)
	Var(varMapUintFloat64, "varMapUintFloat64", "Use mapUintFloat64")
  varMapUintFloat64.Set("2:2.71828")
  if varMapUintFloat64.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUintFloat64,
      varMapUintFloat64.String(), "2:2.71828")
  }
	// fmt.Printf("float64 mapUintFloat64\n%v\n2:2.71828\n", *varMapUintFloat64)
}


func TestParseMapUintFloat32(t *testing.T) {
	// T := reflect.TypeOf(mapUintFloat32Value{}).Elem()
	var varMapUintFloat32 = new(mapUintFloat32Value)
	Var(varMapUintFloat32, "varMapUintFloat32", "Use mapUintFloat32")
  varMapUintFloat32.Set("2:2.71828")
  if varMapUintFloat32.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUintFloat32,
      varMapUintFloat32.String(), "2:2.71828")
  }
	// fmt.Printf("float32 mapUintFloat32\n%v\n2:2.71828\n", *varMapUintFloat32)
}


func TestParseMapUintBool(t *testing.T) {
	// T := reflect.TypeOf(mapUintBoolValue{}).Elem()
	var varMapUintBool = new(mapUintBoolValue)
	Var(varMapUintBool, "varMapUintBool", "Use mapUintBool")
  varMapUintBool.Set("2:true")
  if varMapUintBool.String() != "2:true" {
    t.Fatalf("%v %v %v", varMapUintBool,
      varMapUintBool.String(), "2:true")
  }
	// fmt.Printf("bool mapUintBool\n%v\n2:true\n", *varMapUintBool)
}


func TestParseMapUintString(t *testing.T) {
	// T := reflect.TypeOf(mapUintStringValue{}).Elem()
	var varMapUintString = new(mapUintStringValue)
	Var(varMapUintString, "varMapUintString", "Use mapUintString")
  varMapUintString.Set("2:one")
  if varMapUintString.String() != "2:one" {
    t.Fatalf("%v %v %v", varMapUintString,
      varMapUintString.String(), "2:one")
  }
	// fmt.Printf("string mapUintString\n%v\n2:one\n", *varMapUintString)
}


func TestParseMapUint8Duration(t *testing.T) {
	// T := reflect.TypeOf(mapUint8DurationValue{}).Elem()
	var varMapUint8Duration = new(mapUint8DurationValue)
	Var(varMapUint8Duration, "varMapUint8Duration", "Use mapUint8Duration")
  varMapUint8Duration.Set("2:1h2m3s")
  if varMapUint8Duration.String() != "2:1h2m3s" {
    t.Fatalf("%v %v %v", varMapUint8Duration,
      varMapUint8Duration.String(), "2:1h2m3s")
  }
	// fmt.Printf("duration mapUint8Duration\n%v\n2:1h2m3s\n", *varMapUint8Duration)
}


func TestParseMapUint8Int(t *testing.T) {
	// T := reflect.TypeOf(mapUint8IntValue{}).Elem()
	var varMapUint8Int = new(mapUint8IntValue)
	Var(varMapUint8Int, "varMapUint8Int", "Use mapUint8Int")
  varMapUint8Int.Set("2:-1")
  if varMapUint8Int.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint8Int,
      varMapUint8Int.String(), "2:-1")
  }
	// fmt.Printf("int mapUint8Int\n%v\n2:-1\n", *varMapUint8Int)
}


func TestParseMapUint8Int8(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Int8Value{}).Elem()
	var varMapUint8Int8 = new(mapUint8Int8Value)
	Var(varMapUint8Int8, "varMapUint8Int8", "Use mapUint8Int8")
  varMapUint8Int8.Set("2:-1")
  if varMapUint8Int8.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint8Int8,
      varMapUint8Int8.String(), "2:-1")
  }
	// fmt.Printf("int8 mapUint8Int8\n%v\n2:-1\n", *varMapUint8Int8)
}


func TestParseMapUint8Int16(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Int16Value{}).Elem()
	var varMapUint8Int16 = new(mapUint8Int16Value)
	Var(varMapUint8Int16, "varMapUint8Int16", "Use mapUint8Int16")
  varMapUint8Int16.Set("2:-1")
  if varMapUint8Int16.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint8Int16,
      varMapUint8Int16.String(), "2:-1")
  }
	// fmt.Printf("int16 mapUint8Int16\n%v\n2:-1\n", *varMapUint8Int16)
}


func TestParseMapUint8Int32(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Int32Value{}).Elem()
	var varMapUint8Int32 = new(mapUint8Int32Value)
	Var(varMapUint8Int32, "varMapUint8Int32", "Use mapUint8Int32")
  varMapUint8Int32.Set("2:-1")
  if varMapUint8Int32.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint8Int32,
      varMapUint8Int32.String(), "2:-1")
  }
	// fmt.Printf("int32 mapUint8Int32\n%v\n2:-1\n", *varMapUint8Int32)
}


func TestParseMapUint8Int64(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Int64Value{}).Elem()
	var varMapUint8Int64 = new(mapUint8Int64Value)
	Var(varMapUint8Int64, "varMapUint8Int64", "Use mapUint8Int64")
  varMapUint8Int64.Set("2:-1")
  if varMapUint8Int64.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint8Int64,
      varMapUint8Int64.String(), "2:-1")
  }
	// fmt.Printf("int64 mapUint8Int64\n%v\n2:-1\n", *varMapUint8Int64)
}


func TestParseMapUint8Uint(t *testing.T) {
	// T := reflect.TypeOf(mapUint8UintValue{}).Elem()
	var varMapUint8Uint = new(mapUint8UintValue)
	Var(varMapUint8Uint, "varMapUint8Uint", "Use mapUint8Uint")
  varMapUint8Uint.Set("2:2")
  if varMapUint8Uint.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint8Uint,
      varMapUint8Uint.String(), "2:2")
  }
	// fmt.Printf("uint mapUint8Uint\n%v\n2:2\n", *varMapUint8Uint)
}


func TestParseMapUint8Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Uint8Value{}).Elem()
	var varMapUint8Uint8 = new(mapUint8Uint8Value)
	Var(varMapUint8Uint8, "varMapUint8Uint8", "Use mapUint8Uint8")
  varMapUint8Uint8.Set("2:2")
  if varMapUint8Uint8.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint8Uint8,
      varMapUint8Uint8.String(), "2:2")
  }
	// fmt.Printf("uint8 mapUint8Uint8\n%v\n2:2\n", *varMapUint8Uint8)
}


func TestParseMapUint8Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Uint16Value{}).Elem()
	var varMapUint8Uint16 = new(mapUint8Uint16Value)
	Var(varMapUint8Uint16, "varMapUint8Uint16", "Use mapUint8Uint16")
  varMapUint8Uint16.Set("2:2")
  if varMapUint8Uint16.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint8Uint16,
      varMapUint8Uint16.String(), "2:2")
  }
	// fmt.Printf("uint16 mapUint8Uint16\n%v\n2:2\n", *varMapUint8Uint16)
}


func TestParseMapUint8Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Uint32Value{}).Elem()
	var varMapUint8Uint32 = new(mapUint8Uint32Value)
	Var(varMapUint8Uint32, "varMapUint8Uint32", "Use mapUint8Uint32")
  varMapUint8Uint32.Set("2:2")
  if varMapUint8Uint32.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint8Uint32,
      varMapUint8Uint32.String(), "2:2")
  }
	// fmt.Printf("uint32 mapUint8Uint32\n%v\n2:2\n", *varMapUint8Uint32)
}


func TestParseMapUint8Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Uint64Value{}).Elem()
	var varMapUint8Uint64 = new(mapUint8Uint64Value)
	Var(varMapUint8Uint64, "varMapUint8Uint64", "Use mapUint8Uint64")
  varMapUint8Uint64.Set("2:2")
  if varMapUint8Uint64.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint8Uint64,
      varMapUint8Uint64.String(), "2:2")
  }
	// fmt.Printf("uint64 mapUint8Uint64\n%v\n2:2\n", *varMapUint8Uint64)
}


func TestParseMapUint8Float64(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Float64Value{}).Elem()
	var varMapUint8Float64 = new(mapUint8Float64Value)
	Var(varMapUint8Float64, "varMapUint8Float64", "Use mapUint8Float64")
  varMapUint8Float64.Set("2:2.71828")
  if varMapUint8Float64.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint8Float64,
      varMapUint8Float64.String(), "2:2.71828")
  }
	// fmt.Printf("float64 mapUint8Float64\n%v\n2:2.71828\n", *varMapUint8Float64)
}


func TestParseMapUint8Float32(t *testing.T) {
	// T := reflect.TypeOf(mapUint8Float32Value{}).Elem()
	var varMapUint8Float32 = new(mapUint8Float32Value)
	Var(varMapUint8Float32, "varMapUint8Float32", "Use mapUint8Float32")
  varMapUint8Float32.Set("2:2.71828")
  if varMapUint8Float32.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint8Float32,
      varMapUint8Float32.String(), "2:2.71828")
  }
	// fmt.Printf("float32 mapUint8Float32\n%v\n2:2.71828\n", *varMapUint8Float32)
}


func TestParseMapUint8Bool(t *testing.T) {
	// T := reflect.TypeOf(mapUint8BoolValue{}).Elem()
	var varMapUint8Bool = new(mapUint8BoolValue)
	Var(varMapUint8Bool, "varMapUint8Bool", "Use mapUint8Bool")
  varMapUint8Bool.Set("2:true")
  if varMapUint8Bool.String() != "2:true" {
    t.Fatalf("%v %v %v", varMapUint8Bool,
      varMapUint8Bool.String(), "2:true")
  }
	// fmt.Printf("bool mapUint8Bool\n%v\n2:true\n", *varMapUint8Bool)
}


func TestParseMapUint8String(t *testing.T) {
	// T := reflect.TypeOf(mapUint8StringValue{}).Elem()
	var varMapUint8String = new(mapUint8StringValue)
	Var(varMapUint8String, "varMapUint8String", "Use mapUint8String")
  varMapUint8String.Set("2:one")
  if varMapUint8String.String() != "2:one" {
    t.Fatalf("%v %v %v", varMapUint8String,
      varMapUint8String.String(), "2:one")
  }
	// fmt.Printf("string mapUint8String\n%v\n2:one\n", *varMapUint8String)
}


func TestParseMapUint16Duration(t *testing.T) {
	// T := reflect.TypeOf(mapUint16DurationValue{}).Elem()
	var varMapUint16Duration = new(mapUint16DurationValue)
	Var(varMapUint16Duration, "varMapUint16Duration", "Use mapUint16Duration")
  varMapUint16Duration.Set("2:1h2m3s")
  if varMapUint16Duration.String() != "2:1h2m3s" {
    t.Fatalf("%v %v %v", varMapUint16Duration,
      varMapUint16Duration.String(), "2:1h2m3s")
  }
	// fmt.Printf("duration mapUint16Duration\n%v\n2:1h2m3s\n", *varMapUint16Duration)
}


func TestParseMapUint16Int(t *testing.T) {
	// T := reflect.TypeOf(mapUint16IntValue{}).Elem()
	var varMapUint16Int = new(mapUint16IntValue)
	Var(varMapUint16Int, "varMapUint16Int", "Use mapUint16Int")
  varMapUint16Int.Set("2:-1")
  if varMapUint16Int.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint16Int,
      varMapUint16Int.String(), "2:-1")
  }
	// fmt.Printf("int mapUint16Int\n%v\n2:-1\n", *varMapUint16Int)
}


func TestParseMapUint16Int8(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Int8Value{}).Elem()
	var varMapUint16Int8 = new(mapUint16Int8Value)
	Var(varMapUint16Int8, "varMapUint16Int8", "Use mapUint16Int8")
  varMapUint16Int8.Set("2:-1")
  if varMapUint16Int8.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint16Int8,
      varMapUint16Int8.String(), "2:-1")
  }
	// fmt.Printf("int8 mapUint16Int8\n%v\n2:-1\n", *varMapUint16Int8)
}


func TestParseMapUint16Int16(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Int16Value{}).Elem()
	var varMapUint16Int16 = new(mapUint16Int16Value)
	Var(varMapUint16Int16, "varMapUint16Int16", "Use mapUint16Int16")
  varMapUint16Int16.Set("2:-1")
  if varMapUint16Int16.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint16Int16,
      varMapUint16Int16.String(), "2:-1")
  }
	// fmt.Printf("int16 mapUint16Int16\n%v\n2:-1\n", *varMapUint16Int16)
}


func TestParseMapUint16Int32(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Int32Value{}).Elem()
	var varMapUint16Int32 = new(mapUint16Int32Value)
	Var(varMapUint16Int32, "varMapUint16Int32", "Use mapUint16Int32")
  varMapUint16Int32.Set("2:-1")
  if varMapUint16Int32.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint16Int32,
      varMapUint16Int32.String(), "2:-1")
  }
	// fmt.Printf("int32 mapUint16Int32\n%v\n2:-1\n", *varMapUint16Int32)
}


func TestParseMapUint16Int64(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Int64Value{}).Elem()
	var varMapUint16Int64 = new(mapUint16Int64Value)
	Var(varMapUint16Int64, "varMapUint16Int64", "Use mapUint16Int64")
  varMapUint16Int64.Set("2:-1")
  if varMapUint16Int64.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint16Int64,
      varMapUint16Int64.String(), "2:-1")
  }
	// fmt.Printf("int64 mapUint16Int64\n%v\n2:-1\n", *varMapUint16Int64)
}


func TestParseMapUint16Uint(t *testing.T) {
	// T := reflect.TypeOf(mapUint16UintValue{}).Elem()
	var varMapUint16Uint = new(mapUint16UintValue)
	Var(varMapUint16Uint, "varMapUint16Uint", "Use mapUint16Uint")
  varMapUint16Uint.Set("2:2")
  if varMapUint16Uint.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint16Uint,
      varMapUint16Uint.String(), "2:2")
  }
	// fmt.Printf("uint mapUint16Uint\n%v\n2:2\n", *varMapUint16Uint)
}


func TestParseMapUint16Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Uint8Value{}).Elem()
	var varMapUint16Uint8 = new(mapUint16Uint8Value)
	Var(varMapUint16Uint8, "varMapUint16Uint8", "Use mapUint16Uint8")
  varMapUint16Uint8.Set("2:2")
  if varMapUint16Uint8.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint16Uint8,
      varMapUint16Uint8.String(), "2:2")
  }
	// fmt.Printf("uint8 mapUint16Uint8\n%v\n2:2\n", *varMapUint16Uint8)
}


func TestParseMapUint16Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Uint16Value{}).Elem()
	var varMapUint16Uint16 = new(mapUint16Uint16Value)
	Var(varMapUint16Uint16, "varMapUint16Uint16", "Use mapUint16Uint16")
  varMapUint16Uint16.Set("2:2")
  if varMapUint16Uint16.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint16Uint16,
      varMapUint16Uint16.String(), "2:2")
  }
	// fmt.Printf("uint16 mapUint16Uint16\n%v\n2:2\n", *varMapUint16Uint16)
}


func TestParseMapUint16Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Uint32Value{}).Elem()
	var varMapUint16Uint32 = new(mapUint16Uint32Value)
	Var(varMapUint16Uint32, "varMapUint16Uint32", "Use mapUint16Uint32")
  varMapUint16Uint32.Set("2:2")
  if varMapUint16Uint32.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint16Uint32,
      varMapUint16Uint32.String(), "2:2")
  }
	// fmt.Printf("uint32 mapUint16Uint32\n%v\n2:2\n", *varMapUint16Uint32)
}


func TestParseMapUint16Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Uint64Value{}).Elem()
	var varMapUint16Uint64 = new(mapUint16Uint64Value)
	Var(varMapUint16Uint64, "varMapUint16Uint64", "Use mapUint16Uint64")
  varMapUint16Uint64.Set("2:2")
  if varMapUint16Uint64.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint16Uint64,
      varMapUint16Uint64.String(), "2:2")
  }
	// fmt.Printf("uint64 mapUint16Uint64\n%v\n2:2\n", *varMapUint16Uint64)
}


func TestParseMapUint16Float64(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Float64Value{}).Elem()
	var varMapUint16Float64 = new(mapUint16Float64Value)
	Var(varMapUint16Float64, "varMapUint16Float64", "Use mapUint16Float64")
  varMapUint16Float64.Set("2:2.71828")
  if varMapUint16Float64.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint16Float64,
      varMapUint16Float64.String(), "2:2.71828")
  }
	// fmt.Printf("float64 mapUint16Float64\n%v\n2:2.71828\n", *varMapUint16Float64)
}


func TestParseMapUint16Float32(t *testing.T) {
	// T := reflect.TypeOf(mapUint16Float32Value{}).Elem()
	var varMapUint16Float32 = new(mapUint16Float32Value)
	Var(varMapUint16Float32, "varMapUint16Float32", "Use mapUint16Float32")
  varMapUint16Float32.Set("2:2.71828")
  if varMapUint16Float32.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint16Float32,
      varMapUint16Float32.String(), "2:2.71828")
  }
	// fmt.Printf("float32 mapUint16Float32\n%v\n2:2.71828\n", *varMapUint16Float32)
}


func TestParseMapUint16Bool(t *testing.T) {
	// T := reflect.TypeOf(mapUint16BoolValue{}).Elem()
	var varMapUint16Bool = new(mapUint16BoolValue)
	Var(varMapUint16Bool, "varMapUint16Bool", "Use mapUint16Bool")
  varMapUint16Bool.Set("2:true")
  if varMapUint16Bool.String() != "2:true" {
    t.Fatalf("%v %v %v", varMapUint16Bool,
      varMapUint16Bool.String(), "2:true")
  }
	// fmt.Printf("bool mapUint16Bool\n%v\n2:true\n", *varMapUint16Bool)
}


func TestParseMapUint16String(t *testing.T) {
	// T := reflect.TypeOf(mapUint16StringValue{}).Elem()
	var varMapUint16String = new(mapUint16StringValue)
	Var(varMapUint16String, "varMapUint16String", "Use mapUint16String")
  varMapUint16String.Set("2:one")
  if varMapUint16String.String() != "2:one" {
    t.Fatalf("%v %v %v", varMapUint16String,
      varMapUint16String.String(), "2:one")
  }
	// fmt.Printf("string mapUint16String\n%v\n2:one\n", *varMapUint16String)
}


func TestParseMapUint32Duration(t *testing.T) {
	// T := reflect.TypeOf(mapUint32DurationValue{}).Elem()
	var varMapUint32Duration = new(mapUint32DurationValue)
	Var(varMapUint32Duration, "varMapUint32Duration", "Use mapUint32Duration")
  varMapUint32Duration.Set("2:1h2m3s")
  if varMapUint32Duration.String() != "2:1h2m3s" {
    t.Fatalf("%v %v %v", varMapUint32Duration,
      varMapUint32Duration.String(), "2:1h2m3s")
  }
	// fmt.Printf("duration mapUint32Duration\n%v\n2:1h2m3s\n", *varMapUint32Duration)
}


func TestParseMapUint32Int(t *testing.T) {
	// T := reflect.TypeOf(mapUint32IntValue{}).Elem()
	var varMapUint32Int = new(mapUint32IntValue)
	Var(varMapUint32Int, "varMapUint32Int", "Use mapUint32Int")
  varMapUint32Int.Set("2:-1")
  if varMapUint32Int.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint32Int,
      varMapUint32Int.String(), "2:-1")
  }
	// fmt.Printf("int mapUint32Int\n%v\n2:-1\n", *varMapUint32Int)
}


func TestParseMapUint32Int8(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Int8Value{}).Elem()
	var varMapUint32Int8 = new(mapUint32Int8Value)
	Var(varMapUint32Int8, "varMapUint32Int8", "Use mapUint32Int8")
  varMapUint32Int8.Set("2:-1")
  if varMapUint32Int8.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint32Int8,
      varMapUint32Int8.String(), "2:-1")
  }
	// fmt.Printf("int8 mapUint32Int8\n%v\n2:-1\n", *varMapUint32Int8)
}


func TestParseMapUint32Int16(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Int16Value{}).Elem()
	var varMapUint32Int16 = new(mapUint32Int16Value)
	Var(varMapUint32Int16, "varMapUint32Int16", "Use mapUint32Int16")
  varMapUint32Int16.Set("2:-1")
  if varMapUint32Int16.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint32Int16,
      varMapUint32Int16.String(), "2:-1")
  }
	// fmt.Printf("int16 mapUint32Int16\n%v\n2:-1\n", *varMapUint32Int16)
}


func TestParseMapUint32Int32(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Int32Value{}).Elem()
	var varMapUint32Int32 = new(mapUint32Int32Value)
	Var(varMapUint32Int32, "varMapUint32Int32", "Use mapUint32Int32")
  varMapUint32Int32.Set("2:-1")
  if varMapUint32Int32.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint32Int32,
      varMapUint32Int32.String(), "2:-1")
  }
	// fmt.Printf("int32 mapUint32Int32\n%v\n2:-1\n", *varMapUint32Int32)
}


func TestParseMapUint32Int64(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Int64Value{}).Elem()
	var varMapUint32Int64 = new(mapUint32Int64Value)
	Var(varMapUint32Int64, "varMapUint32Int64", "Use mapUint32Int64")
  varMapUint32Int64.Set("2:-1")
  if varMapUint32Int64.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint32Int64,
      varMapUint32Int64.String(), "2:-1")
  }
	// fmt.Printf("int64 mapUint32Int64\n%v\n2:-1\n", *varMapUint32Int64)
}


func TestParseMapUint32Uint(t *testing.T) {
	// T := reflect.TypeOf(mapUint32UintValue{}).Elem()
	var varMapUint32Uint = new(mapUint32UintValue)
	Var(varMapUint32Uint, "varMapUint32Uint", "Use mapUint32Uint")
  varMapUint32Uint.Set("2:2")
  if varMapUint32Uint.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint32Uint,
      varMapUint32Uint.String(), "2:2")
  }
	// fmt.Printf("uint mapUint32Uint\n%v\n2:2\n", *varMapUint32Uint)
}


func TestParseMapUint32Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Uint8Value{}).Elem()
	var varMapUint32Uint8 = new(mapUint32Uint8Value)
	Var(varMapUint32Uint8, "varMapUint32Uint8", "Use mapUint32Uint8")
  varMapUint32Uint8.Set("2:2")
  if varMapUint32Uint8.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint32Uint8,
      varMapUint32Uint8.String(), "2:2")
  }
	// fmt.Printf("uint8 mapUint32Uint8\n%v\n2:2\n", *varMapUint32Uint8)
}


func TestParseMapUint32Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Uint16Value{}).Elem()
	var varMapUint32Uint16 = new(mapUint32Uint16Value)
	Var(varMapUint32Uint16, "varMapUint32Uint16", "Use mapUint32Uint16")
  varMapUint32Uint16.Set("2:2")
  if varMapUint32Uint16.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint32Uint16,
      varMapUint32Uint16.String(), "2:2")
  }
	// fmt.Printf("uint16 mapUint32Uint16\n%v\n2:2\n", *varMapUint32Uint16)
}


func TestParseMapUint32Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Uint32Value{}).Elem()
	var varMapUint32Uint32 = new(mapUint32Uint32Value)
	Var(varMapUint32Uint32, "varMapUint32Uint32", "Use mapUint32Uint32")
  varMapUint32Uint32.Set("2:2")
  if varMapUint32Uint32.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint32Uint32,
      varMapUint32Uint32.String(), "2:2")
  }
	// fmt.Printf("uint32 mapUint32Uint32\n%v\n2:2\n", *varMapUint32Uint32)
}


func TestParseMapUint32Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Uint64Value{}).Elem()
	var varMapUint32Uint64 = new(mapUint32Uint64Value)
	Var(varMapUint32Uint64, "varMapUint32Uint64", "Use mapUint32Uint64")
  varMapUint32Uint64.Set("2:2")
  if varMapUint32Uint64.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint32Uint64,
      varMapUint32Uint64.String(), "2:2")
  }
	// fmt.Printf("uint64 mapUint32Uint64\n%v\n2:2\n", *varMapUint32Uint64)
}


func TestParseMapUint32Float64(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Float64Value{}).Elem()
	var varMapUint32Float64 = new(mapUint32Float64Value)
	Var(varMapUint32Float64, "varMapUint32Float64", "Use mapUint32Float64")
  varMapUint32Float64.Set("2:2.71828")
  if varMapUint32Float64.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint32Float64,
      varMapUint32Float64.String(), "2:2.71828")
  }
	// fmt.Printf("float64 mapUint32Float64\n%v\n2:2.71828\n", *varMapUint32Float64)
}


func TestParseMapUint32Float32(t *testing.T) {
	// T := reflect.TypeOf(mapUint32Float32Value{}).Elem()
	var varMapUint32Float32 = new(mapUint32Float32Value)
	Var(varMapUint32Float32, "varMapUint32Float32", "Use mapUint32Float32")
  varMapUint32Float32.Set("2:2.71828")
  if varMapUint32Float32.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint32Float32,
      varMapUint32Float32.String(), "2:2.71828")
  }
	// fmt.Printf("float32 mapUint32Float32\n%v\n2:2.71828\n", *varMapUint32Float32)
}


func TestParseMapUint32Bool(t *testing.T) {
	// T := reflect.TypeOf(mapUint32BoolValue{}).Elem()
	var varMapUint32Bool = new(mapUint32BoolValue)
	Var(varMapUint32Bool, "varMapUint32Bool", "Use mapUint32Bool")
  varMapUint32Bool.Set("2:true")
  if varMapUint32Bool.String() != "2:true" {
    t.Fatalf("%v %v %v", varMapUint32Bool,
      varMapUint32Bool.String(), "2:true")
  }
	// fmt.Printf("bool mapUint32Bool\n%v\n2:true\n", *varMapUint32Bool)
}


func TestParseMapUint32String(t *testing.T) {
	// T := reflect.TypeOf(mapUint32StringValue{}).Elem()
	var varMapUint32String = new(mapUint32StringValue)
	Var(varMapUint32String, "varMapUint32String", "Use mapUint32String")
  varMapUint32String.Set("2:one")
  if varMapUint32String.String() != "2:one" {
    t.Fatalf("%v %v %v", varMapUint32String,
      varMapUint32String.String(), "2:one")
  }
	// fmt.Printf("string mapUint32String\n%v\n2:one\n", *varMapUint32String)
}


func TestParseMapUint64Duration(t *testing.T) {
	// T := reflect.TypeOf(mapUint64DurationValue{}).Elem()
	var varMapUint64Duration = new(mapUint64DurationValue)
	Var(varMapUint64Duration, "varMapUint64Duration", "Use mapUint64Duration")
  varMapUint64Duration.Set("2:1h2m3s")
  if varMapUint64Duration.String() != "2:1h2m3s" {
    t.Fatalf("%v %v %v", varMapUint64Duration,
      varMapUint64Duration.String(), "2:1h2m3s")
  }
	// fmt.Printf("duration mapUint64Duration\n%v\n2:1h2m3s\n", *varMapUint64Duration)
}


func TestParseMapUint64Int(t *testing.T) {
	// T := reflect.TypeOf(mapUint64IntValue{}).Elem()
	var varMapUint64Int = new(mapUint64IntValue)
	Var(varMapUint64Int, "varMapUint64Int", "Use mapUint64Int")
  varMapUint64Int.Set("2:-1")
  if varMapUint64Int.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint64Int,
      varMapUint64Int.String(), "2:-1")
  }
	// fmt.Printf("int mapUint64Int\n%v\n2:-1\n", *varMapUint64Int)
}


func TestParseMapUint64Int8(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Int8Value{}).Elem()
	var varMapUint64Int8 = new(mapUint64Int8Value)
	Var(varMapUint64Int8, "varMapUint64Int8", "Use mapUint64Int8")
  varMapUint64Int8.Set("2:-1")
  if varMapUint64Int8.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint64Int8,
      varMapUint64Int8.String(), "2:-1")
  }
	// fmt.Printf("int8 mapUint64Int8\n%v\n2:-1\n", *varMapUint64Int8)
}


func TestParseMapUint64Int16(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Int16Value{}).Elem()
	var varMapUint64Int16 = new(mapUint64Int16Value)
	Var(varMapUint64Int16, "varMapUint64Int16", "Use mapUint64Int16")
  varMapUint64Int16.Set("2:-1")
  if varMapUint64Int16.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint64Int16,
      varMapUint64Int16.String(), "2:-1")
  }
	// fmt.Printf("int16 mapUint64Int16\n%v\n2:-1\n", *varMapUint64Int16)
}


func TestParseMapUint64Int32(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Int32Value{}).Elem()
	var varMapUint64Int32 = new(mapUint64Int32Value)
	Var(varMapUint64Int32, "varMapUint64Int32", "Use mapUint64Int32")
  varMapUint64Int32.Set("2:-1")
  if varMapUint64Int32.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint64Int32,
      varMapUint64Int32.String(), "2:-1")
  }
	// fmt.Printf("int32 mapUint64Int32\n%v\n2:-1\n", *varMapUint64Int32)
}


func TestParseMapUint64Int64(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Int64Value{}).Elem()
	var varMapUint64Int64 = new(mapUint64Int64Value)
	Var(varMapUint64Int64, "varMapUint64Int64", "Use mapUint64Int64")
  varMapUint64Int64.Set("2:-1")
  if varMapUint64Int64.String() != "2:-1" {
    t.Fatalf("%v %v %v", varMapUint64Int64,
      varMapUint64Int64.String(), "2:-1")
  }
	// fmt.Printf("int64 mapUint64Int64\n%v\n2:-1\n", *varMapUint64Int64)
}


func TestParseMapUint64Uint(t *testing.T) {
	// T := reflect.TypeOf(mapUint64UintValue{}).Elem()
	var varMapUint64Uint = new(mapUint64UintValue)
	Var(varMapUint64Uint, "varMapUint64Uint", "Use mapUint64Uint")
  varMapUint64Uint.Set("2:2")
  if varMapUint64Uint.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint64Uint,
      varMapUint64Uint.String(), "2:2")
  }
	// fmt.Printf("uint mapUint64Uint\n%v\n2:2\n", *varMapUint64Uint)
}


func TestParseMapUint64Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Uint8Value{}).Elem()
	var varMapUint64Uint8 = new(mapUint64Uint8Value)
	Var(varMapUint64Uint8, "varMapUint64Uint8", "Use mapUint64Uint8")
  varMapUint64Uint8.Set("2:2")
  if varMapUint64Uint8.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint64Uint8,
      varMapUint64Uint8.String(), "2:2")
  }
	// fmt.Printf("uint8 mapUint64Uint8\n%v\n2:2\n", *varMapUint64Uint8)
}


func TestParseMapUint64Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Uint16Value{}).Elem()
	var varMapUint64Uint16 = new(mapUint64Uint16Value)
	Var(varMapUint64Uint16, "varMapUint64Uint16", "Use mapUint64Uint16")
  varMapUint64Uint16.Set("2:2")
  if varMapUint64Uint16.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint64Uint16,
      varMapUint64Uint16.String(), "2:2")
  }
	// fmt.Printf("uint16 mapUint64Uint16\n%v\n2:2\n", *varMapUint64Uint16)
}


func TestParseMapUint64Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Uint32Value{}).Elem()
	var varMapUint64Uint32 = new(mapUint64Uint32Value)
	Var(varMapUint64Uint32, "varMapUint64Uint32", "Use mapUint64Uint32")
  varMapUint64Uint32.Set("2:2")
  if varMapUint64Uint32.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint64Uint32,
      varMapUint64Uint32.String(), "2:2")
  }
	// fmt.Printf("uint32 mapUint64Uint32\n%v\n2:2\n", *varMapUint64Uint32)
}


func TestParseMapUint64Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Uint64Value{}).Elem()
	var varMapUint64Uint64 = new(mapUint64Uint64Value)
	Var(varMapUint64Uint64, "varMapUint64Uint64", "Use mapUint64Uint64")
  varMapUint64Uint64.Set("2:2")
  if varMapUint64Uint64.String() != "2:2" {
    t.Fatalf("%v %v %v", varMapUint64Uint64,
      varMapUint64Uint64.String(), "2:2")
  }
	// fmt.Printf("uint64 mapUint64Uint64\n%v\n2:2\n", *varMapUint64Uint64)
}


func TestParseMapUint64Float64(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Float64Value{}).Elem()
	var varMapUint64Float64 = new(mapUint64Float64Value)
	Var(varMapUint64Float64, "varMapUint64Float64", "Use mapUint64Float64")
  varMapUint64Float64.Set("2:2.71828")
  if varMapUint64Float64.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint64Float64,
      varMapUint64Float64.String(), "2:2.71828")
  }
	// fmt.Printf("float64 mapUint64Float64\n%v\n2:2.71828\n", *varMapUint64Float64)
}


func TestParseMapUint64Float32(t *testing.T) {
	// T := reflect.TypeOf(mapUint64Float32Value{}).Elem()
	var varMapUint64Float32 = new(mapUint64Float32Value)
	Var(varMapUint64Float32, "varMapUint64Float32", "Use mapUint64Float32")
  varMapUint64Float32.Set("2:2.71828")
  if varMapUint64Float32.String() != "2:2.71828" {
    t.Fatalf("%v %v %v", varMapUint64Float32,
      varMapUint64Float32.String(), "2:2.71828")
  }
	// fmt.Printf("float32 mapUint64Float32\n%v\n2:2.71828\n", *varMapUint64Float32)
}


func TestParseMapUint64Bool(t *testing.T) {
	// T := reflect.TypeOf(mapUint64BoolValue{}).Elem()
	var varMapUint64Bool = new(mapUint64BoolValue)
	Var(varMapUint64Bool, "varMapUint64Bool", "Use mapUint64Bool")
  varMapUint64Bool.Set("2:true")
  if varMapUint64Bool.String() != "2:true" {
    t.Fatalf("%v %v %v", varMapUint64Bool,
      varMapUint64Bool.String(), "2:true")
  }
	// fmt.Printf("bool mapUint64Bool\n%v\n2:true\n", *varMapUint64Bool)
}


func TestParseMapUint64String(t *testing.T) {
	// T := reflect.TypeOf(mapUint64StringValue{}).Elem()
	var varMapUint64String = new(mapUint64StringValue)
	Var(varMapUint64String, "varMapUint64String", "Use mapUint64String")
  varMapUint64String.Set("2:one")
  if varMapUint64String.String() != "2:one" {
    t.Fatalf("%v %v %v", varMapUint64String,
      varMapUint64String.String(), "2:one")
  }
	// fmt.Printf("string mapUint64String\n%v\n2:one\n", *varMapUint64String)
}


func TestParseMapFloat64Duration(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64DurationValue{}).Elem()
	var varMapFloat64Duration = new(mapFloat64DurationValue)
	Var(varMapFloat64Duration, "varMapFloat64Duration", "Use mapFloat64Duration")
  varMapFloat64Duration.Set("2.71828:1h2m3s")
  if varMapFloat64Duration.String() != "2.71828:1h2m3s" {
    t.Fatalf("%v %v %v", varMapFloat64Duration,
      varMapFloat64Duration.String(), "2.71828:1h2m3s")
  }
	// fmt.Printf("duration mapFloat64Duration\n%v\n2.71828:1h2m3s\n", *varMapFloat64Duration)
}


func TestParseMapFloat64Int(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64IntValue{}).Elem()
	var varMapFloat64Int = new(mapFloat64IntValue)
	Var(varMapFloat64Int, "varMapFloat64Int", "Use mapFloat64Int")
  varMapFloat64Int.Set("2.71828:-1")
  if varMapFloat64Int.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat64Int,
      varMapFloat64Int.String(), "2.71828:-1")
  }
	// fmt.Printf("int mapFloat64Int\n%v\n2.71828:-1\n", *varMapFloat64Int)
}


func TestParseMapFloat64Int8(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Int8Value{}).Elem()
	var varMapFloat64Int8 = new(mapFloat64Int8Value)
	Var(varMapFloat64Int8, "varMapFloat64Int8", "Use mapFloat64Int8")
  varMapFloat64Int8.Set("2.71828:-1")
  if varMapFloat64Int8.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat64Int8,
      varMapFloat64Int8.String(), "2.71828:-1")
  }
	// fmt.Printf("int8 mapFloat64Int8\n%v\n2.71828:-1\n", *varMapFloat64Int8)
}


func TestParseMapFloat64Int16(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Int16Value{}).Elem()
	var varMapFloat64Int16 = new(mapFloat64Int16Value)
	Var(varMapFloat64Int16, "varMapFloat64Int16", "Use mapFloat64Int16")
  varMapFloat64Int16.Set("2.71828:-1")
  if varMapFloat64Int16.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat64Int16,
      varMapFloat64Int16.String(), "2.71828:-1")
  }
	// fmt.Printf("int16 mapFloat64Int16\n%v\n2.71828:-1\n", *varMapFloat64Int16)
}


func TestParseMapFloat64Int32(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Int32Value{}).Elem()
	var varMapFloat64Int32 = new(mapFloat64Int32Value)
	Var(varMapFloat64Int32, "varMapFloat64Int32", "Use mapFloat64Int32")
  varMapFloat64Int32.Set("2.71828:-1")
  if varMapFloat64Int32.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat64Int32,
      varMapFloat64Int32.String(), "2.71828:-1")
  }
	// fmt.Printf("int32 mapFloat64Int32\n%v\n2.71828:-1\n", *varMapFloat64Int32)
}


func TestParseMapFloat64Int64(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Int64Value{}).Elem()
	var varMapFloat64Int64 = new(mapFloat64Int64Value)
	Var(varMapFloat64Int64, "varMapFloat64Int64", "Use mapFloat64Int64")
  varMapFloat64Int64.Set("2.71828:-1")
  if varMapFloat64Int64.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat64Int64,
      varMapFloat64Int64.String(), "2.71828:-1")
  }
	// fmt.Printf("int64 mapFloat64Int64\n%v\n2.71828:-1\n", *varMapFloat64Int64)
}


func TestParseMapFloat64Uint(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64UintValue{}).Elem()
	var varMapFloat64Uint = new(mapFloat64UintValue)
	Var(varMapFloat64Uint, "varMapFloat64Uint", "Use mapFloat64Uint")
  varMapFloat64Uint.Set("2.71828:2")
  if varMapFloat64Uint.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat64Uint,
      varMapFloat64Uint.String(), "2.71828:2")
  }
	// fmt.Printf("uint mapFloat64Uint\n%v\n2.71828:2\n", *varMapFloat64Uint)
}


func TestParseMapFloat64Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Uint8Value{}).Elem()
	var varMapFloat64Uint8 = new(mapFloat64Uint8Value)
	Var(varMapFloat64Uint8, "varMapFloat64Uint8", "Use mapFloat64Uint8")
  varMapFloat64Uint8.Set("2.71828:2")
  if varMapFloat64Uint8.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat64Uint8,
      varMapFloat64Uint8.String(), "2.71828:2")
  }
	// fmt.Printf("uint8 mapFloat64Uint8\n%v\n2.71828:2\n", *varMapFloat64Uint8)
}


func TestParseMapFloat64Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Uint16Value{}).Elem()
	var varMapFloat64Uint16 = new(mapFloat64Uint16Value)
	Var(varMapFloat64Uint16, "varMapFloat64Uint16", "Use mapFloat64Uint16")
  varMapFloat64Uint16.Set("2.71828:2")
  if varMapFloat64Uint16.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat64Uint16,
      varMapFloat64Uint16.String(), "2.71828:2")
  }
	// fmt.Printf("uint16 mapFloat64Uint16\n%v\n2.71828:2\n", *varMapFloat64Uint16)
}


func TestParseMapFloat64Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Uint32Value{}).Elem()
	var varMapFloat64Uint32 = new(mapFloat64Uint32Value)
	Var(varMapFloat64Uint32, "varMapFloat64Uint32", "Use mapFloat64Uint32")
  varMapFloat64Uint32.Set("2.71828:2")
  if varMapFloat64Uint32.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat64Uint32,
      varMapFloat64Uint32.String(), "2.71828:2")
  }
	// fmt.Printf("uint32 mapFloat64Uint32\n%v\n2.71828:2\n", *varMapFloat64Uint32)
}


func TestParseMapFloat64Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Uint64Value{}).Elem()
	var varMapFloat64Uint64 = new(mapFloat64Uint64Value)
	Var(varMapFloat64Uint64, "varMapFloat64Uint64", "Use mapFloat64Uint64")
  varMapFloat64Uint64.Set("2.71828:2")
  if varMapFloat64Uint64.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat64Uint64,
      varMapFloat64Uint64.String(), "2.71828:2")
  }
	// fmt.Printf("uint64 mapFloat64Uint64\n%v\n2.71828:2\n", *varMapFloat64Uint64)
}


func TestParseMapFloat64Float64(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Float64Value{}).Elem()
	var varMapFloat64Float64 = new(mapFloat64Float64Value)
	Var(varMapFloat64Float64, "varMapFloat64Float64", "Use mapFloat64Float64")
  varMapFloat64Float64.Set("2.71828:2.71828")
  if varMapFloat64Float64.String() != "2.71828:2.71828" {
    t.Fatalf("%v %v %v", varMapFloat64Float64,
      varMapFloat64Float64.String(), "2.71828:2.71828")
  }
	// fmt.Printf("float64 mapFloat64Float64\n%v\n2.71828:2.71828\n", *varMapFloat64Float64)
}


func TestParseMapFloat64Float32(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64Float32Value{}).Elem()
	var varMapFloat64Float32 = new(mapFloat64Float32Value)
	Var(varMapFloat64Float32, "varMapFloat64Float32", "Use mapFloat64Float32")
  varMapFloat64Float32.Set("2.71828:2.71828")
  if varMapFloat64Float32.String() != "2.71828:2.71828" {
    t.Fatalf("%v %v %v", varMapFloat64Float32,
      varMapFloat64Float32.String(), "2.71828:2.71828")
  }
	// fmt.Printf("float32 mapFloat64Float32\n%v\n2.71828:2.71828\n", *varMapFloat64Float32)
}


func TestParseMapFloat64Bool(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64BoolValue{}).Elem()
	var varMapFloat64Bool = new(mapFloat64BoolValue)
	Var(varMapFloat64Bool, "varMapFloat64Bool", "Use mapFloat64Bool")
  varMapFloat64Bool.Set("2.71828:true")
  if varMapFloat64Bool.String() != "2.71828:true" {
    t.Fatalf("%v %v %v", varMapFloat64Bool,
      varMapFloat64Bool.String(), "2.71828:true")
  }
	// fmt.Printf("bool mapFloat64Bool\n%v\n2.71828:true\n", *varMapFloat64Bool)
}


func TestParseMapFloat64String(t *testing.T) {
	// T := reflect.TypeOf(mapFloat64StringValue{}).Elem()
	var varMapFloat64String = new(mapFloat64StringValue)
	Var(varMapFloat64String, "varMapFloat64String", "Use mapFloat64String")
  varMapFloat64String.Set("2.71828:one")
  if varMapFloat64String.String() != "2.71828:one" {
    t.Fatalf("%v %v %v", varMapFloat64String,
      varMapFloat64String.String(), "2.71828:one")
  }
	// fmt.Printf("string mapFloat64String\n%v\n2.71828:one\n", *varMapFloat64String)
}


func TestParseMapFloat32Duration(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32DurationValue{}).Elem()
	var varMapFloat32Duration = new(mapFloat32DurationValue)
	Var(varMapFloat32Duration, "varMapFloat32Duration", "Use mapFloat32Duration")
  varMapFloat32Duration.Set("2.71828:1h2m3s")
  if varMapFloat32Duration.String() != "2.71828:1h2m3s" {
    t.Fatalf("%v %v %v", varMapFloat32Duration,
      varMapFloat32Duration.String(), "2.71828:1h2m3s")
  }
	// fmt.Printf("duration mapFloat32Duration\n%v\n2.71828:1h2m3s\n", *varMapFloat32Duration)
}


func TestParseMapFloat32Int(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32IntValue{}).Elem()
	var varMapFloat32Int = new(mapFloat32IntValue)
	Var(varMapFloat32Int, "varMapFloat32Int", "Use mapFloat32Int")
  varMapFloat32Int.Set("2.71828:-1")
  if varMapFloat32Int.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat32Int,
      varMapFloat32Int.String(), "2.71828:-1")
  }
	// fmt.Printf("int mapFloat32Int\n%v\n2.71828:-1\n", *varMapFloat32Int)
}


func TestParseMapFloat32Int8(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Int8Value{}).Elem()
	var varMapFloat32Int8 = new(mapFloat32Int8Value)
	Var(varMapFloat32Int8, "varMapFloat32Int8", "Use mapFloat32Int8")
  varMapFloat32Int8.Set("2.71828:-1")
  if varMapFloat32Int8.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat32Int8,
      varMapFloat32Int8.String(), "2.71828:-1")
  }
	// fmt.Printf("int8 mapFloat32Int8\n%v\n2.71828:-1\n", *varMapFloat32Int8)
}


func TestParseMapFloat32Int16(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Int16Value{}).Elem()
	var varMapFloat32Int16 = new(mapFloat32Int16Value)
	Var(varMapFloat32Int16, "varMapFloat32Int16", "Use mapFloat32Int16")
  varMapFloat32Int16.Set("2.71828:-1")
  if varMapFloat32Int16.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat32Int16,
      varMapFloat32Int16.String(), "2.71828:-1")
  }
	// fmt.Printf("int16 mapFloat32Int16\n%v\n2.71828:-1\n", *varMapFloat32Int16)
}


func TestParseMapFloat32Int32(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Int32Value{}).Elem()
	var varMapFloat32Int32 = new(mapFloat32Int32Value)
	Var(varMapFloat32Int32, "varMapFloat32Int32", "Use mapFloat32Int32")
  varMapFloat32Int32.Set("2.71828:-1")
  if varMapFloat32Int32.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat32Int32,
      varMapFloat32Int32.String(), "2.71828:-1")
  }
	// fmt.Printf("int32 mapFloat32Int32\n%v\n2.71828:-1\n", *varMapFloat32Int32)
}


func TestParseMapFloat32Int64(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Int64Value{}).Elem()
	var varMapFloat32Int64 = new(mapFloat32Int64Value)
	Var(varMapFloat32Int64, "varMapFloat32Int64", "Use mapFloat32Int64")
  varMapFloat32Int64.Set("2.71828:-1")
  if varMapFloat32Int64.String() != "2.71828:-1" {
    t.Fatalf("%v %v %v", varMapFloat32Int64,
      varMapFloat32Int64.String(), "2.71828:-1")
  }
	// fmt.Printf("int64 mapFloat32Int64\n%v\n2.71828:-1\n", *varMapFloat32Int64)
}


func TestParseMapFloat32Uint(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32UintValue{}).Elem()
	var varMapFloat32Uint = new(mapFloat32UintValue)
	Var(varMapFloat32Uint, "varMapFloat32Uint", "Use mapFloat32Uint")
  varMapFloat32Uint.Set("2.71828:2")
  if varMapFloat32Uint.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat32Uint,
      varMapFloat32Uint.String(), "2.71828:2")
  }
	// fmt.Printf("uint mapFloat32Uint\n%v\n2.71828:2\n", *varMapFloat32Uint)
}


func TestParseMapFloat32Uint8(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Uint8Value{}).Elem()
	var varMapFloat32Uint8 = new(mapFloat32Uint8Value)
	Var(varMapFloat32Uint8, "varMapFloat32Uint8", "Use mapFloat32Uint8")
  varMapFloat32Uint8.Set("2.71828:2")
  if varMapFloat32Uint8.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat32Uint8,
      varMapFloat32Uint8.String(), "2.71828:2")
  }
	// fmt.Printf("uint8 mapFloat32Uint8\n%v\n2.71828:2\n", *varMapFloat32Uint8)
}


func TestParseMapFloat32Uint16(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Uint16Value{}).Elem()
	var varMapFloat32Uint16 = new(mapFloat32Uint16Value)
	Var(varMapFloat32Uint16, "varMapFloat32Uint16", "Use mapFloat32Uint16")
  varMapFloat32Uint16.Set("2.71828:2")
  if varMapFloat32Uint16.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat32Uint16,
      varMapFloat32Uint16.String(), "2.71828:2")
  }
	// fmt.Printf("uint16 mapFloat32Uint16\n%v\n2.71828:2\n", *varMapFloat32Uint16)
}


func TestParseMapFloat32Uint32(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Uint32Value{}).Elem()
	var varMapFloat32Uint32 = new(mapFloat32Uint32Value)
	Var(varMapFloat32Uint32, "varMapFloat32Uint32", "Use mapFloat32Uint32")
  varMapFloat32Uint32.Set("2.71828:2")
  if varMapFloat32Uint32.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat32Uint32,
      varMapFloat32Uint32.String(), "2.71828:2")
  }
	// fmt.Printf("uint32 mapFloat32Uint32\n%v\n2.71828:2\n", *varMapFloat32Uint32)
}


func TestParseMapFloat32Uint64(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Uint64Value{}).Elem()
	var varMapFloat32Uint64 = new(mapFloat32Uint64Value)
	Var(varMapFloat32Uint64, "varMapFloat32Uint64", "Use mapFloat32Uint64")
  varMapFloat32Uint64.Set("2.71828:2")
  if varMapFloat32Uint64.String() != "2.71828:2" {
    t.Fatalf("%v %v %v", varMapFloat32Uint64,
      varMapFloat32Uint64.String(), "2.71828:2")
  }
	// fmt.Printf("uint64 mapFloat32Uint64\n%v\n2.71828:2\n", *varMapFloat32Uint64)
}


func TestParseMapFloat32Float64(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Float64Value{}).Elem()
	var varMapFloat32Float64 = new(mapFloat32Float64Value)
	Var(varMapFloat32Float64, "varMapFloat32Float64", "Use mapFloat32Float64")
  varMapFloat32Float64.Set("2.71828:2.71828")
  if varMapFloat32Float64.String() != "2.71828:2.71828" {
    t.Fatalf("%v %v %v", varMapFloat32Float64,
      varMapFloat32Float64.String(), "2.71828:2.71828")
  }
	// fmt.Printf("float64 mapFloat32Float64\n%v\n2.71828:2.71828\n", *varMapFloat32Float64)
}


func TestParseMapFloat32Float32(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32Float32Value{}).Elem()
	var varMapFloat32Float32 = new(mapFloat32Float32Value)
	Var(varMapFloat32Float32, "varMapFloat32Float32", "Use mapFloat32Float32")
  varMapFloat32Float32.Set("2.71828:2.71828")
  if varMapFloat32Float32.String() != "2.71828:2.71828" {
    t.Fatalf("%v %v %v", varMapFloat32Float32,
      varMapFloat32Float32.String(), "2.71828:2.71828")
  }
	// fmt.Printf("float32 mapFloat32Float32\n%v\n2.71828:2.71828\n", *varMapFloat32Float32)
}


func TestParseMapFloat32Bool(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32BoolValue{}).Elem()
	var varMapFloat32Bool = new(mapFloat32BoolValue)
	Var(varMapFloat32Bool, "varMapFloat32Bool", "Use mapFloat32Bool")
  varMapFloat32Bool.Set("2.71828:true")
  if varMapFloat32Bool.String() != "2.71828:true" {
    t.Fatalf("%v %v %v", varMapFloat32Bool,
      varMapFloat32Bool.String(), "2.71828:true")
  }
	// fmt.Printf("bool mapFloat32Bool\n%v\n2.71828:true\n", *varMapFloat32Bool)
}


func TestParseMapFloat32String(t *testing.T) {
	// T := reflect.TypeOf(mapFloat32StringValue{}).Elem()
	var varMapFloat32String = new(mapFloat32StringValue)
	Var(varMapFloat32String, "varMapFloat32String", "Use mapFloat32String")
  varMapFloat32String.Set("2.71828:one")
  if varMapFloat32String.String() != "2.71828:one" {
    t.Fatalf("%v %v %v", varMapFloat32String,
      varMapFloat32String.String(), "2.71828:one")
  }
	// fmt.Printf("string mapFloat32String\n%v\n2.71828:one\n", *varMapFloat32String)
}


func TestParseMapBoolDuration(t *testing.T) {
	// T := reflect.TypeOf(mapBoolDurationValue{}).Elem()
	var varMapBoolDuration = new(mapBoolDurationValue)
	Var(varMapBoolDuration, "varMapBoolDuration", "Use mapBoolDuration")
  varMapBoolDuration.Set("true:1h2m3s")
  if varMapBoolDuration.String() != "true:1h2m3s" {
    t.Fatalf("%v %v %v", varMapBoolDuration,
      varMapBoolDuration.String(), "true:1h2m3s")
  }
	// fmt.Printf("duration mapBoolDuration\n%v\ntrue:1h2m3s\n", *varMapBoolDuration)
}


func TestParseMapBoolInt(t *testing.T) {
	// T := reflect.TypeOf(mapBoolIntValue{}).Elem()
	var varMapBoolInt = new(mapBoolIntValue)
	Var(varMapBoolInt, "varMapBoolInt", "Use mapBoolInt")
  varMapBoolInt.Set("true:-1")
  if varMapBoolInt.String() != "true:-1" {
    t.Fatalf("%v %v %v", varMapBoolInt,
      varMapBoolInt.String(), "true:-1")
  }
	// fmt.Printf("int mapBoolInt\n%v\ntrue:-1\n", *varMapBoolInt)
}


func TestParseMapBoolInt8(t *testing.T) {
	// T := reflect.TypeOf(mapBoolInt8Value{}).Elem()
	var varMapBoolInt8 = new(mapBoolInt8Value)
	Var(varMapBoolInt8, "varMapBoolInt8", "Use mapBoolInt8")
  varMapBoolInt8.Set("true:-1")
  if varMapBoolInt8.String() != "true:-1" {
    t.Fatalf("%v %v %v", varMapBoolInt8,
      varMapBoolInt8.String(), "true:-1")
  }
	// fmt.Printf("int8 mapBoolInt8\n%v\ntrue:-1\n", *varMapBoolInt8)
}


func TestParseMapBoolInt16(t *testing.T) {
	// T := reflect.TypeOf(mapBoolInt16Value{}).Elem()
	var varMapBoolInt16 = new(mapBoolInt16Value)
	Var(varMapBoolInt16, "varMapBoolInt16", "Use mapBoolInt16")
  varMapBoolInt16.Set("true:-1")
  if varMapBoolInt16.String() != "true:-1" {
    t.Fatalf("%v %v %v", varMapBoolInt16,
      varMapBoolInt16.String(), "true:-1")
  }
	// fmt.Printf("int16 mapBoolInt16\n%v\ntrue:-1\n", *varMapBoolInt16)
}


func TestParseMapBoolInt32(t *testing.T) {
	// T := reflect.TypeOf(mapBoolInt32Value{}).Elem()
	var varMapBoolInt32 = new(mapBoolInt32Value)
	Var(varMapBoolInt32, "varMapBoolInt32", "Use mapBoolInt32")
  varMapBoolInt32.Set("true:-1")
  if varMapBoolInt32.String() != "true:-1" {
    t.Fatalf("%v %v %v", varMapBoolInt32,
      varMapBoolInt32.String(), "true:-1")
  }
	// fmt.Printf("int32 mapBoolInt32\n%v\ntrue:-1\n", *varMapBoolInt32)
}


func TestParseMapBoolInt64(t *testing.T) {
	// T := reflect.TypeOf(mapBoolInt64Value{}).Elem()
	var varMapBoolInt64 = new(mapBoolInt64Value)
	Var(varMapBoolInt64, "varMapBoolInt64", "Use mapBoolInt64")
  varMapBoolInt64.Set("true:-1")
  if varMapBoolInt64.String() != "true:-1" {
    t.Fatalf("%v %v %v", varMapBoolInt64,
      varMapBoolInt64.String(), "true:-1")
  }
	// fmt.Printf("int64 mapBoolInt64\n%v\ntrue:-1\n", *varMapBoolInt64)
}


func TestParseMapBoolUint(t *testing.T) {
	// T := reflect.TypeOf(mapBoolUintValue{}).Elem()
	var varMapBoolUint = new(mapBoolUintValue)
	Var(varMapBoolUint, "varMapBoolUint", "Use mapBoolUint")
  varMapBoolUint.Set("true:2")
  if varMapBoolUint.String() != "true:2" {
    t.Fatalf("%v %v %v", varMapBoolUint,
      varMapBoolUint.String(), "true:2")
  }
	// fmt.Printf("uint mapBoolUint\n%v\ntrue:2\n", *varMapBoolUint)
}


func TestParseMapBoolUint8(t *testing.T) {
	// T := reflect.TypeOf(mapBoolUint8Value{}).Elem()
	var varMapBoolUint8 = new(mapBoolUint8Value)
	Var(varMapBoolUint8, "varMapBoolUint8", "Use mapBoolUint8")
  varMapBoolUint8.Set("true:2")
  if varMapBoolUint8.String() != "true:2" {
    t.Fatalf("%v %v %v", varMapBoolUint8,
      varMapBoolUint8.String(), "true:2")
  }
	// fmt.Printf("uint8 mapBoolUint8\n%v\ntrue:2\n", *varMapBoolUint8)
}


func TestParseMapBoolUint16(t *testing.T) {
	// T := reflect.TypeOf(mapBoolUint16Value{}).Elem()
	var varMapBoolUint16 = new(mapBoolUint16Value)
	Var(varMapBoolUint16, "varMapBoolUint16", "Use mapBoolUint16")
  varMapBoolUint16.Set("true:2")
  if varMapBoolUint16.String() != "true:2" {
    t.Fatalf("%v %v %v", varMapBoolUint16,
      varMapBoolUint16.String(), "true:2")
  }
	// fmt.Printf("uint16 mapBoolUint16\n%v\ntrue:2\n", *varMapBoolUint16)
}


func TestParseMapBoolUint32(t *testing.T) {
	// T := reflect.TypeOf(mapBoolUint32Value{}).Elem()
	var varMapBoolUint32 = new(mapBoolUint32Value)
	Var(varMapBoolUint32, "varMapBoolUint32", "Use mapBoolUint32")
  varMapBoolUint32.Set("true:2")
  if varMapBoolUint32.String() != "true:2" {
    t.Fatalf("%v %v %v", varMapBoolUint32,
      varMapBoolUint32.String(), "true:2")
  }
	// fmt.Printf("uint32 mapBoolUint32\n%v\ntrue:2\n", *varMapBoolUint32)
}


func TestParseMapBoolUint64(t *testing.T) {
	// T := reflect.TypeOf(mapBoolUint64Value{}).Elem()
	var varMapBoolUint64 = new(mapBoolUint64Value)
	Var(varMapBoolUint64, "varMapBoolUint64", "Use mapBoolUint64")
  varMapBoolUint64.Set("true:2")
  if varMapBoolUint64.String() != "true:2" {
    t.Fatalf("%v %v %v", varMapBoolUint64,
      varMapBoolUint64.String(), "true:2")
  }
	// fmt.Printf("uint64 mapBoolUint64\n%v\ntrue:2\n", *varMapBoolUint64)
}


func TestParseMapBoolFloat64(t *testing.T) {
	// T := reflect.TypeOf(mapBoolFloat64Value{}).Elem()
	var varMapBoolFloat64 = new(mapBoolFloat64Value)
	Var(varMapBoolFloat64, "varMapBoolFloat64", "Use mapBoolFloat64")
  varMapBoolFloat64.Set("true:2.71828")
  if varMapBoolFloat64.String() != "true:2.71828" {
    t.Fatalf("%v %v %v", varMapBoolFloat64,
      varMapBoolFloat64.String(), "true:2.71828")
  }
	// fmt.Printf("float64 mapBoolFloat64\n%v\ntrue:2.71828\n", *varMapBoolFloat64)
}


func TestParseMapBoolFloat32(t *testing.T) {
	// T := reflect.TypeOf(mapBoolFloat32Value{}).Elem()
	var varMapBoolFloat32 = new(mapBoolFloat32Value)
	Var(varMapBoolFloat32, "varMapBoolFloat32", "Use mapBoolFloat32")
  varMapBoolFloat32.Set("true:2.71828")
  if varMapBoolFloat32.String() != "true:2.71828" {
    t.Fatalf("%v %v %v", varMapBoolFloat32,
      varMapBoolFloat32.String(), "true:2.71828")
  }
	// fmt.Printf("float32 mapBoolFloat32\n%v\ntrue:2.71828\n", *varMapBoolFloat32)
}


func TestParseMapBoolBool(t *testing.T) {
	// T := reflect.TypeOf(mapBoolBoolValue{}).Elem()
	var varMapBoolBool = new(mapBoolBoolValue)
	Var(varMapBoolBool, "varMapBoolBool", "Use mapBoolBool")
  varMapBoolBool.Set("true:true")
  if varMapBoolBool.String() != "true:true" {
    t.Fatalf("%v %v %v", varMapBoolBool,
      varMapBoolBool.String(), "true:true")
  }
	// fmt.Printf("bool mapBoolBool\n%v\ntrue:true\n", *varMapBoolBool)
}


func TestParseMapBoolString(t *testing.T) {
	// T := reflect.TypeOf(mapBoolStringValue{}).Elem()
	var varMapBoolString = new(mapBoolStringValue)
	Var(varMapBoolString, "varMapBoolString", "Use mapBoolString")
  varMapBoolString.Set("true:one")
  if varMapBoolString.String() != "true:one" {
    t.Fatalf("%v %v %v", varMapBoolString,
      varMapBoolString.String(), "true:one")
  }
	// fmt.Printf("string mapBoolString\n%v\ntrue:one\n", *varMapBoolString)
}


func TestParseMapStringDuration(t *testing.T) {
	// T := reflect.TypeOf(mapStringDurationValue{}).Elem()
	var varMapStringDuration = new(mapStringDurationValue)
	Var(varMapStringDuration, "varMapStringDuration", "Use mapStringDuration")
  varMapStringDuration.Set("one:1h2m3s")
  if varMapStringDuration.String() != "one:1h2m3s" {
    t.Fatalf("%v %v %v", varMapStringDuration,
      varMapStringDuration.String(), "one:1h2m3s")
  }
	// fmt.Printf("duration mapStringDuration\n%v\none:1h2m3s\n", *varMapStringDuration)
}


func TestParseMapStringInt(t *testing.T) {
	// T := reflect.TypeOf(mapStringIntValue{}).Elem()
	var varMapStringInt = new(mapStringIntValue)
	Var(varMapStringInt, "varMapStringInt", "Use mapStringInt")
  varMapStringInt.Set("one:-1")
  if varMapStringInt.String() != "one:-1" {
    t.Fatalf("%v %v %v", varMapStringInt,
      varMapStringInt.String(), "one:-1")
  }
	// fmt.Printf("int mapStringInt\n%v\none:-1\n", *varMapStringInt)
}


func TestParseMapStringInt8(t *testing.T) {
	// T := reflect.TypeOf(mapStringInt8Value{}).Elem()
	var varMapStringInt8 = new(mapStringInt8Value)
	Var(varMapStringInt8, "varMapStringInt8", "Use mapStringInt8")
  varMapStringInt8.Set("one:-1")
  if varMapStringInt8.String() != "one:-1" {
    t.Fatalf("%v %v %v", varMapStringInt8,
      varMapStringInt8.String(), "one:-1")
  }
	// fmt.Printf("int8 mapStringInt8\n%v\none:-1\n", *varMapStringInt8)
}


func TestParseMapStringInt16(t *testing.T) {
	// T := reflect.TypeOf(mapStringInt16Value{}).Elem()
	var varMapStringInt16 = new(mapStringInt16Value)
	Var(varMapStringInt16, "varMapStringInt16", "Use mapStringInt16")
  varMapStringInt16.Set("one:-1")
  if varMapStringInt16.String() != "one:-1" {
    t.Fatalf("%v %v %v", varMapStringInt16,
      varMapStringInt16.String(), "one:-1")
  }
	// fmt.Printf("int16 mapStringInt16\n%v\none:-1\n", *varMapStringInt16)
}


func TestParseMapStringInt32(t *testing.T) {
	// T := reflect.TypeOf(mapStringInt32Value{}).Elem()
	var varMapStringInt32 = new(mapStringInt32Value)
	Var(varMapStringInt32, "varMapStringInt32", "Use mapStringInt32")
  varMapStringInt32.Set("one:-1")
  if varMapStringInt32.String() != "one:-1" {
    t.Fatalf("%v %v %v", varMapStringInt32,
      varMapStringInt32.String(), "one:-1")
  }
	// fmt.Printf("int32 mapStringInt32\n%v\none:-1\n", *varMapStringInt32)
}


func TestParseMapStringInt64(t *testing.T) {
	// T := reflect.TypeOf(mapStringInt64Value{}).Elem()
	var varMapStringInt64 = new(mapStringInt64Value)
	Var(varMapStringInt64, "varMapStringInt64", "Use mapStringInt64")
  varMapStringInt64.Set("one:-1")
  if varMapStringInt64.String() != "one:-1" {
    t.Fatalf("%v %v %v", varMapStringInt64,
      varMapStringInt64.String(), "one:-1")
  }
	// fmt.Printf("int64 mapStringInt64\n%v\none:-1\n", *varMapStringInt64)
}


func TestParseMapStringUint(t *testing.T) {
	// T := reflect.TypeOf(mapStringUintValue{}).Elem()
	var varMapStringUint = new(mapStringUintValue)
	Var(varMapStringUint, "varMapStringUint", "Use mapStringUint")
  varMapStringUint.Set("one:2")
  if varMapStringUint.String() != "one:2" {
    t.Fatalf("%v %v %v", varMapStringUint,
      varMapStringUint.String(), "one:2")
  }
	// fmt.Printf("uint mapStringUint\n%v\none:2\n", *varMapStringUint)
}


func TestParseMapStringUint8(t *testing.T) {
	// T := reflect.TypeOf(mapStringUint8Value{}).Elem()
	var varMapStringUint8 = new(mapStringUint8Value)
	Var(varMapStringUint8, "varMapStringUint8", "Use mapStringUint8")
  varMapStringUint8.Set("one:2")
  if varMapStringUint8.String() != "one:2" {
    t.Fatalf("%v %v %v", varMapStringUint8,
      varMapStringUint8.String(), "one:2")
  }
	// fmt.Printf("uint8 mapStringUint8\n%v\none:2\n", *varMapStringUint8)
}


func TestParseMapStringUint16(t *testing.T) {
	// T := reflect.TypeOf(mapStringUint16Value{}).Elem()
	var varMapStringUint16 = new(mapStringUint16Value)
	Var(varMapStringUint16, "varMapStringUint16", "Use mapStringUint16")
  varMapStringUint16.Set("one:2")
  if varMapStringUint16.String() != "one:2" {
    t.Fatalf("%v %v %v", varMapStringUint16,
      varMapStringUint16.String(), "one:2")
  }
	// fmt.Printf("uint16 mapStringUint16\n%v\none:2\n", *varMapStringUint16)
}


func TestParseMapStringUint32(t *testing.T) {
	// T := reflect.TypeOf(mapStringUint32Value{}).Elem()
	var varMapStringUint32 = new(mapStringUint32Value)
	Var(varMapStringUint32, "varMapStringUint32", "Use mapStringUint32")
  varMapStringUint32.Set("one:2")
  if varMapStringUint32.String() != "one:2" {
    t.Fatalf("%v %v %v", varMapStringUint32,
      varMapStringUint32.String(), "one:2")
  }
	// fmt.Printf("uint32 mapStringUint32\n%v\none:2\n", *varMapStringUint32)
}


func TestParseMapStringUint64(t *testing.T) {
	// T := reflect.TypeOf(mapStringUint64Value{}).Elem()
	var varMapStringUint64 = new(mapStringUint64Value)
	Var(varMapStringUint64, "varMapStringUint64", "Use mapStringUint64")
  varMapStringUint64.Set("one:2")
  if varMapStringUint64.String() != "one:2" {
    t.Fatalf("%v %v %v", varMapStringUint64,
      varMapStringUint64.String(), "one:2")
  }
	// fmt.Printf("uint64 mapStringUint64\n%v\none:2\n", *varMapStringUint64)
}


func TestParseMapStringFloat64(t *testing.T) {
	// T := reflect.TypeOf(mapStringFloat64Value{}).Elem()
	var varMapStringFloat64 = new(mapStringFloat64Value)
	Var(varMapStringFloat64, "varMapStringFloat64", "Use mapStringFloat64")
  varMapStringFloat64.Set("one:2.71828")
  if varMapStringFloat64.String() != "one:2.71828" {
    t.Fatalf("%v %v %v", varMapStringFloat64,
      varMapStringFloat64.String(), "one:2.71828")
  }
	// fmt.Printf("float64 mapStringFloat64\n%v\none:2.71828\n", *varMapStringFloat64)
}


func TestParseMapStringFloat32(t *testing.T) {
	// T := reflect.TypeOf(mapStringFloat32Value{}).Elem()
	var varMapStringFloat32 = new(mapStringFloat32Value)
	Var(varMapStringFloat32, "varMapStringFloat32", "Use mapStringFloat32")
  varMapStringFloat32.Set("one:2.71828")
  if varMapStringFloat32.String() != "one:2.71828" {
    t.Fatalf("%v %v %v", varMapStringFloat32,
      varMapStringFloat32.String(), "one:2.71828")
  }
	// fmt.Printf("float32 mapStringFloat32\n%v\none:2.71828\n", *varMapStringFloat32)
}


func TestParseMapStringBool(t *testing.T) {
	// T := reflect.TypeOf(mapStringBoolValue{}).Elem()
	var varMapStringBool = new(mapStringBoolValue)
	Var(varMapStringBool, "varMapStringBool", "Use mapStringBool")
  varMapStringBool.Set("one:true")
  if varMapStringBool.String() != "one:true" {
    t.Fatalf("%v %v %v", varMapStringBool,
      varMapStringBool.String(), "one:true")
  }
	// fmt.Printf("bool mapStringBool\n%v\none:true\n", *varMapStringBool)
}


func TestParseMapStringString(t *testing.T) {
	// T := reflect.TypeOf(mapStringStringValue{}).Elem()
	var varMapStringString = new(mapStringStringValue)
	Var(varMapStringString, "varMapStringString", "Use mapStringString")
  varMapStringString.Set("one:one")
  if varMapStringString.String() != "one:one" {
    t.Fatalf("%v %v %v", varMapStringString,
      varMapStringString.String(), "one:one")
  }
	// fmt.Printf("string mapStringString\n%v\none:one\n", *varMapStringString)
}

