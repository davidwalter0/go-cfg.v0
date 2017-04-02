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

// SliceIntValue []Int
type SliceIntValue []int

func newSliceIntValue(val SliceIntValue, p *SliceIntValue) *SliceIntValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}

// Set a slice after parsing a string
func (slc *SliceIntValue) Set(s string) error {
	T := reflect.TypeOf(SliceIntValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)
	var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
  // start int
    n, _ = strconv.ParseInt(text, 0, T.Bits())
    *slc = append(*slc, (int)(n.(int64)))
  // end int
	}
	return nil
}

// Get get a slice interface from the value
func (slc *SliceIntValue) Get() interface{} { return ([]int)(*slc) }

// String join a string from slice
func (slc *SliceIntValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// SliceIntVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) SliceIntVar(p *SliceIntValue, name string, value SliceIntValue, usage string) {
	f.Var(newSliceIntValue(value, p), name, usage)
}

// SliceIntVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func SliceIntVar(p *SliceIntValue, name string, value SliceIntValue, usage string) {
	CommandLine.Var(newSliceIntValue(value, p), name, usage)
}

// SliceInt defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) SliceInt(name string, value SliceIntValue, usage string) *SliceIntValue {
	p := new(SliceIntValue)
	f.SliceIntVar(p, name, value, usage)
	return p
}

// SliceInt defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func SliceInt(name string, value SliceIntValue, usage string) *SliceIntValue {
	return CommandLine.SliceInt(name, value, usage)
}

