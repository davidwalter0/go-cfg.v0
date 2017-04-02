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

// SliceInt32Value []Int32
type SliceInt32Value []int32

func newSliceInt32Value(val SliceInt32Value, p *SliceInt32Value) *SliceInt32Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceInt32Value) Set(s string) error {
	T := reflect.TypeOf(SliceInt32Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start int32
    n, _ = strconv.ParseInt(text, 0, T.Bits())
    *slc = append(*slc, (int32)(n.(int64)))
  // end int32
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceInt32Value) Get() interface{} { return ([]int32)(*slc) }

// String join a string from slice
func (slc *SliceInt32Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceInt32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceInt32Var(p *SliceInt32Value, name string, value SliceInt32Value, usage string) {
	f.Var(newSliceInt32Value(value, p), name, usage)
}

// SliceInt32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceInt32Var(p *SliceInt32Value, name string, value SliceInt32Value, usage string) {
	CommandLine.Var(newSliceInt32Value(value, p), name, usage)
}

// SliceInt32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceInt32(name string, value SliceInt32Value, usage string) *SliceInt32Value {
	p := new(SliceInt32Value)
	f.SliceInt32Var(p, name, value, usage)
	return p
}

// SliceInt32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceInt32(name string, value SliceInt32Value, usage string) *SliceInt32Value {
	return CommandLine.SliceInt32(name, value, usage)
}

