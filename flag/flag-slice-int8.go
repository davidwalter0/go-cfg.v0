package flag

////////////////////////////////////////////////////////////////////////
//
////////////////////////////////////////////////////////////////////////
import (
	"fmt"
	"strings"
  "strconv"
  "reflect"
)

// SliceInt8Value []Int8
type SliceInt8Value []int8

func newSliceInt8Value(val SliceInt8Value, p *SliceInt8Value) *SliceInt8Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceInt8Value) Set(s string) error {
	T := reflect.TypeOf(SliceInt8Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start int8
    n, _ = strconv.ParseInt(text, 0, T.Bits())
    *slc = append(*slc, (int8)(n.(int64)))
  // end int8
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceInt8Value) Get() interface{} { return ([]int8)(*slc) }

// String join a string from slice
func (slc *SliceInt8Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceInt8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceInt8Var(p *SliceInt8Value, name string, value SliceInt8Value, usage string) {
	f.Var(newSliceInt8Value(value, p), name, usage)
}

// SliceInt8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceInt8Var(p *SliceInt8Value, name string, value SliceInt8Value, usage string) {
	CommandLine.Var(newSliceInt8Value(value, p), name, usage)
}

// SliceInt8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceInt8(name string, value SliceInt8Value, usage string) *SliceInt8Value {
	p := new(SliceInt8Value)
	f.SliceInt8Var(p, name, value, usage)
	return p
}

// SliceInt8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceInt8(name string, value SliceInt8Value, usage string) *SliceInt8Value {
	return CommandLine.SliceInt8(name, value, usage)
}

