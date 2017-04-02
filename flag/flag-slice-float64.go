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

// SliceFloat64Value []Float64
type SliceFloat64Value []float64

func newSliceFloat64Value(val SliceFloat64Value, p *SliceFloat64Value) *SliceFloat64Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceFloat64Value) Set(s string) error {
	T := reflect.TypeOf(SliceFloat64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start float64
    n, _ = strconv.ParseFloat(text, T.Bits())
    *slc = append(*slc, (float64)(n.(float64)))
  // end float64
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceFloat64Value) Get() interface{} { return ([]float64)(*slc) }

// String join a string from slice
func (slc *SliceFloat64Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceFloat64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceFloat64Var(p *SliceFloat64Value, name string, value SliceFloat64Value, usage string) {
	f.Var(newSliceFloat64Value(value, p), name, usage)
}

// SliceFloat64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceFloat64Var(p *SliceFloat64Value, name string, value SliceFloat64Value, usage string) {
	CommandLine.Var(newSliceFloat64Value(value, p), name, usage)
}

// SliceFloat64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceFloat64(name string, value SliceFloat64Value, usage string) *SliceFloat64Value {
	p := new(SliceFloat64Value)
	f.SliceFloat64Var(p, name, value, usage)
	return p
}

// SliceFloat64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceFloat64(name string, value SliceFloat64Value, usage string) *SliceFloat64Value {
	return CommandLine.SliceFloat64(name, value, usage)
}

