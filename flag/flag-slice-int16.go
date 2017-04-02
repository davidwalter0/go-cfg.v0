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

// SliceInt16Value []Int16
type SliceInt16Value []int16

func newSliceInt16Value(val SliceInt16Value, p *SliceInt16Value) *SliceInt16Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceInt16Value) Set(s string) error {
	T := reflect.TypeOf(SliceInt16Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start int16
    n, _ = strconv.ParseInt(text, 0, T.Bits())
    *slc = append(*slc, (int16)(n.(int64)))
  // end int16
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceInt16Value) Get() interface{} { return ([]int16)(*slc) }

// String join a string from slice
func (slc *SliceInt16Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceInt16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceInt16Var(p *SliceInt16Value, name string, value SliceInt16Value, usage string) {
	f.Var(newSliceInt16Value(value, p), name, usage)
}

// SliceInt16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceInt16Var(p *SliceInt16Value, name string, value SliceInt16Value, usage string) {
	CommandLine.Var(newSliceInt16Value(value, p), name, usage)
}

// SliceInt16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceInt16(name string, value SliceInt16Value, usage string) *SliceInt16Value {
	p := new(SliceInt16Value)
	f.SliceInt16Var(p, name, value, usage)
	return p
}

// SliceInt16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceInt16(name string, value SliceInt16Value, usage string) *SliceInt16Value {
	return CommandLine.SliceInt16(name, value, usage)
}

