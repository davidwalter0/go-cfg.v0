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

// SliceUint64Value []Uint64
type SliceUint64Value []uint64

func newSliceUint64Value(val SliceUint64Value, p *SliceUint64Value) *SliceUint64Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceUint64Value) Set(s string) error {
	T := reflect.TypeOf(SliceUint64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start uint64
    n, _ = strconv.ParseUint(text, 0, T.Bits())
    *slc = append(*slc, (uint64)(n.(uint64)))
  // end uint64
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceUint64Value) Get() interface{} { return ([]uint64)(*slc) }

// String join a string from slice
func (slc *SliceUint64Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceUint64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceUint64Var(p *SliceUint64Value, name string, value SliceUint64Value, usage string) {
	f.Var(newSliceUint64Value(value, p), name, usage)
}

// SliceUint64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceUint64Var(p *SliceUint64Value, name string, value SliceUint64Value, usage string) {
	CommandLine.Var(newSliceUint64Value(value, p), name, usage)
}

// SliceUint64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceUint64(name string, value SliceUint64Value, usage string) *SliceUint64Value {
	p := new(SliceUint64Value)
	f.SliceUint64Var(p, name, value, usage)
	return p
}

// SliceUint64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceUint64(name string, value SliceUint64Value, usage string) *SliceUint64Value {
	return CommandLine.SliceUint64(name, value, usage)
}

