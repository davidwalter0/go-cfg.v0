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

// SliceUintValue []Uint
type SliceUintValue []uint

func newSliceUintValue(val SliceUintValue, p *SliceUintValue) *SliceUintValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceUintValue) Set(s string) error {
	T := reflect.TypeOf(SliceUintValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start uint
    n, _ = strconv.ParseUint(text, 0, T.Bits())
    *slc = append(*slc, (uint)(n.(uint64)))
  // end uint
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceUintValue) Get() interface{} { return ([]uint)(*slc) }

// String join a string from slice
func (slc *SliceUintValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceUintVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceUintVar(p *SliceUintValue, name string, value SliceUintValue, usage string) {
	f.Var(newSliceUintValue(value, p), name, usage)
}

// SliceUintVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceUintVar(p *SliceUintValue, name string, value SliceUintValue, usage string) {
	CommandLine.Var(newSliceUintValue(value, p), name, usage)
}

// SliceUint defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceUint(name string, value SliceUintValue, usage string) *SliceUintValue {
	p := new(SliceUintValue)
	f.SliceUintVar(p, name, value, usage)
	return p
}

// SliceUint defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceUint(name string, value SliceUintValue, usage string) *SliceUintValue {
	return CommandLine.SliceUint(name, value, usage)
}

