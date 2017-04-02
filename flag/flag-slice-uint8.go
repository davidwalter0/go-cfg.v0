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

// SliceUint8Value []Uint8
type SliceUint8Value []uint8

func newSliceUint8Value(val SliceUint8Value, p *SliceUint8Value) *SliceUint8Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceUint8Value) Set(s string) error {
	T := reflect.TypeOf(SliceUint8Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start uint8
    n, _ = strconv.ParseUint(text, 0, T.Bits())
    *slc = append(*slc, (uint8)(n.(uint64)))
  // end uint8
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceUint8Value) Get() interface{} { return ([]uint8)(*slc) }

// String join a string from slice
func (slc *SliceUint8Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceUint8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceUint8Var(p *SliceUint8Value, name string, value SliceUint8Value, usage string) {
	f.Var(newSliceUint8Value(value, p), name, usage)
}

// SliceUint8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceUint8Var(p *SliceUint8Value, name string, value SliceUint8Value, usage string) {
	CommandLine.Var(newSliceUint8Value(value, p), name, usage)
}

// SliceUint8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceUint8(name string, value SliceUint8Value, usage string) *SliceUint8Value {
	p := new(SliceUint8Value)
	f.SliceUint8Var(p, name, value, usage)
	return p
}

// SliceUint8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceUint8(name string, value SliceUint8Value, usage string) *SliceUint8Value {
	return CommandLine.SliceUint8(name, value, usage)
}

