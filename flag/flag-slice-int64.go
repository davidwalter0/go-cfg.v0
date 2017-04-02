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

// SliceInt64Value []Int64
type SliceInt64Value []int64

func newSliceInt64Value(val SliceInt64Value, p *SliceInt64Value) *SliceInt64Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceInt64Value) Set(s string) error {
	T := reflect.TypeOf(SliceInt64Value{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start int64
    n, _ = strconv.ParseInt(text, 0, T.Bits())
    *slc = append(*slc, (int64)(n.(int64)))
  // end int64
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceInt64Value) Get() interface{} { return ([]int64)(*slc) }

// String join a string from slice
func (slc *SliceInt64Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceInt64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceInt64Var(p *SliceInt64Value, name string, value SliceInt64Value, usage string) {
	f.Var(newSliceInt64Value(value, p), name, usage)
}

// SliceInt64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceInt64Var(p *SliceInt64Value, name string, value SliceInt64Value, usage string) {
	CommandLine.Var(newSliceInt64Value(value, p), name, usage)
}

// SliceInt64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceInt64(name string, value SliceInt64Value, usage string) *SliceInt64Value {
	p := new(SliceInt64Value)
	f.SliceInt64Var(p, name, value, usage)
	return p
}

// SliceInt64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceInt64(name string, value SliceInt64Value, usage string) *SliceInt64Value {
	return CommandLine.SliceInt64(name, value, usage)
}

