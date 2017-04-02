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

// SliceFloat32Value []Float32
type SliceFloat32Value []float32

func newSliceFloat32Value(val SliceFloat32Value, p *SliceFloat32Value) *SliceFloat32Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceFloat32Value) Set(s string) error {
	T := reflect.TypeOf(SliceFloat32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start float32
    n, _ = strconv.ParseFloat(text, T.Bits())
    *slc = append(*slc, (float32)(n.(float64)))
  // end float32
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceFloat32Value) Get() interface{} { return ([]float32)(*slc) }

// String join a string from slice
func (slc *SliceFloat32Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceFloat32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceFloat32Var(p *SliceFloat32Value, name string, value SliceFloat32Value, usage string) {
	f.Var(newSliceFloat32Value(value, p), name, usage)
}

// SliceFloat32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceFloat32Var(p *SliceFloat32Value, name string, value SliceFloat32Value, usage string) {
	CommandLine.Var(newSliceFloat32Value(value, p), name, usage)
}

// SliceFloat32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceFloat32(name string, value SliceFloat32Value, usage string) *SliceFloat32Value {
	p := new(SliceFloat32Value)
	f.SliceFloat32Var(p, name, value, usage)
	return p
}

// SliceFloat32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceFloat32(name string, value SliceFloat32Value, usage string) *SliceFloat32Value {
	return CommandLine.SliceFloat32(name, value, usage)
}

