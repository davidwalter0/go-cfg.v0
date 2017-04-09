package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"

	"github.com/davidwalter0/envflagstructconfig"
)

type Key string
type Value float64

func main() {
	// text := "a:2.71828,c:3.1415926,e:6.022140857e23"
	// var s []Key
	// // envflagstructconfig.TextTypeSlice(&s, text)
	// if err := envflagstructconfig.TextTypeSlice(&s, text); err != nil {
	// 	log.Fatalf("envflagstructconfig.TextTypeSlice: %q\n", err)
	// } else {
	// 	log.Printf("envflagstructconfig.TextTypeSlice: %v, %T\n", s, s)
	// }
	// if err := envflagstructconfig.TextTypeSlice(s, text); err != nil {
	// 	log.Printf("Error envflagstructconfig.TextTypeSlice: %q\n", err)
	// } else {
	// 	log.Printf("envflagstructconfig.TextTypeSlice: %v, %T\n", s, s)
	// }

	// if err := envflagstructconfig.TextTypedVal(&s, text); err != nil {
	// 	log.Fatalf("envflagstructconfig.TextTypedVal: %q\n", err)
	// } else {
	// 	log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", s, s)
	// }
	// if err := envflagstructconfig.TextTypedVal(s, text); err != nil {
	// 	log.Printf("Error envflagstructconfig.TextTypedVal: %q\n", err)
	// } else {
	// 	log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", s, s)
	// }

	// log.Printf("envflagstructconfig.TextTypeSlice: %v, %T\n", s, s)

	// var m map[Key]Value
	// envflagstructconfig.TextTypedValMap(&m, text)
	// log.Printf("envflagstructconfig.TextTypedValMap: %v, %T\n", m, m)
	// {
	// 	var m map[string]string
	// 	text := "a:0xb,c:0xd,e:0xf"
	// 	envflagstructconfig.TextTypedValMap(&m, text)
	// 	log.Printf("m:  %v %T\n", m, m)
	// }
	// {
	// 	var m map[string]int
	// 	text := "a:0xb,c:0xd,e:0xf"
	// 	envflagstructconfig.TextTypedValMap(&m, text)
	// 	log.Printf("m:  %v %T\n", m, m)
	// }
	// fmt.Printf("\n\n")

	// if err := envflagstructconfig.TextTypedVal(&m, text); err != nil {
	// 	log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// } else {
	// 	log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// }
	// {
	// 	var m map[string]string
	// 	text := "a:0xb,c:0xd,e:0xf"
	// 	if err := envflagstructconfig.TextTypedVal(&m, text); err != nil {
	// 		log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	} else {
	// 		log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	}
	// 	if err := envflagstructconfig.TextTypedVal(m, text); err != nil {
	// 		log.Printf("Not a pointer ? envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	} else {
	// 		log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	}
	// 	log.Printf("m:  %v %T\n", m, m)
	// }
	// {
	// 	var m map[string]int
	// 	text := "a:0xb,c:0xd,e:0xf"
	// 	if err := envflagstructconfig.TextTypedVal(&m, text); err != nil {
	// 		log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	} else {
	// 		log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	}
	// 	if err := envflagstructconfig.TextTypedVal(m, text); err != nil {
	// 		log.Printf("Not a pointer ? envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	} else {
	// 		log.Printf("envflagstructconfig.TextTypedVal: %v, %T\n", m, m)
	// 	}
	// 	log.Printf("m:  %v %T\n", m, m)
	// }
	// fmt.Printf("\n\n")

	{
		var myapp myAPP
		// if err := envflagstructconfig.ParseStructTypedVal(&myapp); err != nil {
		// 	log.Printf("Error %v %q\n", myapp, err)
		// } else {
		// 	jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
		// 	log.Printf("\n%v\n", string(jsonText))
		// }

		// log.Println(myapp)
		// log.Printf("%v %T\n", myapp, myapp)

		var sti *envflagstructconfig.StructInfo = &envflagstructconfig.StructInfo{
			StructPtr: &myapp,
		}
		sti.Parse()
		log.Println(myapp)
		log.Printf("%v %T\n", myapp, myapp)
		jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
		log.Printf("\n%v\n", string(jsonText))

		log.Printf("%v %T\n", myapp, myapp)
		jsonText, _ = json.MarshalIndent(&myapp, "", "  ")
		log.Printf("\n%v\n", string(jsonText))

	}

	os.Exit(0)
}
