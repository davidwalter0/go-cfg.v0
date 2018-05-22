// Run with
// go run simple.go

package main

import (
	"encoding/json"
	"fmt"
	"github.com/davidwalter0/go-cfg"
)

type myApp struct {
	I      int `default:"-1"`
	Nested struct {
		Y float64
	}
}

func main() {
	var myapp myApp

	var sti *cfg.StructInfo = &cfg.StructInfo{
		StructPtr: &myapp,
	}

	if err := sti.Parse(); err != nil { // parse tags, environment, flags
		fmt.Errorf("%v\n", err)
	}
	fmt.Printf("%v %T\n", myapp, myapp)
	jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
	fmt.Printf("\n%v\n", string(jsonText))
}
