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

// SliceUint16Value []Uint16
type SliceUint16Value []uint16

func newSliceUint16Value(val SliceUint16Value, p *SliceUint16Value) *SliceUint16Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceUint16Value) Set(s string) error {
	T := reflect.TypeOf(SliceUint16Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start uint16
    n, _ = strconv.ParseUint(text, 0, T.Bits())
    *slc = append(*slc, (uint16)(n.(uint64)))
  // end uint16
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceUint16Value) Get() interface{} { return ([]uint16)(*slc) }

// String join a string from slice
func (slc *SliceUint16Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceUint16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceUint16Var(p *SliceUint16Value, name string, value SliceUint16Value, usage string) {
	f.Var(newSliceUint16Value(value, p), name, usage)
}

// SliceUint16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceUint16Var(p *SliceUint16Value, name string, value SliceUint16Value, usage string) {
	CommandLine.Var(newSliceUint16Value(value, p), name, usage)
}

// SliceUint16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceUint16(name string, value SliceUint16Value, usage string) *SliceUint16Value {
	p := new(SliceUint16Value)
	f.SliceUint16Var(p, name, value, usage)
	return p
}

// SliceUint16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceUint16(name string, value SliceUint16Value, usage string) *SliceUint16Value {
	return CommandLine.SliceUint16(name, value, usage)
}

