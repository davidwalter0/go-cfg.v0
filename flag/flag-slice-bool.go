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

// SliceBoolValue []Bool
type SliceBoolValue []bool

func newSliceBoolValue(val SliceBoolValue, p *SliceBoolValue) *SliceBoolValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceBoolValue) Set(s string) error {
	T := reflect.TypeOf(SliceBoolValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start bool
    n, _ = strconv.ParseBool(text)
    *slc = append(*slc, (bool)(n.(bool)))
  // end bool
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceBoolValue) Get() interface{} { return ([]bool)(*slc) }

// String join a string from slice
func (slc *SliceBoolValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceBoolVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceBoolVar(p *SliceBoolValue, name string, value SliceBoolValue, usage string) {
	f.Var(newSliceBoolValue(value, p), name, usage)
}

// SliceBoolVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceBoolVar(p *SliceBoolValue, name string, value SliceBoolValue, usage string) {
	CommandLine.Var(newSliceBoolValue(value, p), name, usage)
}

// SliceBool defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceBool(name string, value SliceBoolValue, usage string) *SliceBoolValue {
	p := new(SliceBoolValue)
	f.SliceBoolVar(p, name, value, usage)
	return p
}

// SliceBool defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceBool(name string, value SliceBoolValue, usage string) *SliceBoolValue {
	return CommandLine.SliceBool(name, value, usage)
}

