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

// SliceUint32Value []Uint32
type SliceUint32Value []uint32

func newSliceUint32Value(val SliceUint32Value, p *SliceUint32Value) *SliceUint32Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceUint32Value) Set(s string) error {
	T := reflect.TypeOf(SliceUint32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start uint32
    n, _ = strconv.ParseUint(text, 0, T.Bits())
    *slc = append(*slc, (uint32)(n.(uint64)))
  // end uint32
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceUint32Value) Get() interface{} { return ([]uint32)(*slc) }

// String join a string from slice
func (slc *SliceUint32Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceUint32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceUint32Var(p *SliceUint32Value, name string, value SliceUint32Value, usage string) {
	f.Var(newSliceUint32Value(value, p), name, usage)
}

// SliceUint32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceUint32Var(p *SliceUint32Value, name string, value SliceUint32Value, usage string) {
	CommandLine.Var(newSliceUint32Value(value, p), name, usage)
}

// SliceUint32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceUint32(name string, value SliceUint32Value, usage string) *SliceUint32Value {
	p := new(SliceUint32Value)
	f.SliceUint32Var(p, name, value, usage)
	return p
}

// SliceUint32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceUint32(name string, value SliceUint32Value, usage string) *SliceUint32Value {
	return CommandLine.SliceUint32(name, value, usage)
}

