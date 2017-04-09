package main

import (
	"log"
	"os"

	_ "github.com/davidwalter0/envflagstructconfig"
	"github.com/davidwalter0/envflagstructconfig/flag"
)

type Key string
type Value float64

func main() {
	text := "a:2.71828,c:3.1415926,e:6.022140857e23"
	var s []Key
	// Text2TypeSlice(&s, text)
	if err := envflagstructconfig.Text2TypeSlice(&s, text); err != nil {
		log.Fatalf("Text2TypeSlice: %q\n", err)
	} else {
		log.Printf("Text2TypeSlice: %v, %T\n", s, s)
	}
	if err := Text2TypeSlice(s, text); err != nil {
		log.Printf("Error Text2TypeSlice: %q\n", err)
	} else {
		log.Printf("Text2TypeSlice: %v, %T\n", s, s)
	}

	if err := Text2TypedVal(&s, text); err != nil {
		log.Fatalf("Text2TypedVal: %q\n", err)
	} else {
		log.Printf("Text2TypedVal: %v, %T\n", s, s)
	}
	if err := Text2TypedVal(s, text); err != nil {
		log.Printf("Error Text2TypedVal: %q\n", err)
	} else {
		log.Printf("Text2TypedVal: %v, %T\n", s, s)
	}

	log.Printf("Text2TypeSlice: %v, %T\n", s, s)

	var m map[Key]Value
	Text2TypedValMap(&m, text)
	log.Printf("Text2TypedValMap: %v, %T\n", m, m)
	{
		var m map[string]string
		text := "a:0xb,c:0xd,e:0xf"
		Text2TypedValMap(&m, text)
		log.Printf("m:  %v %T\n", m, m)
	}
	{
		var m map[string]int
		text := "a:0xb,c:0xd,e:0xf"
		Text2TypedValMap(&m, text)
		log.Printf("m:  %v %T\n", m, m)
	}
	fmt.Printf("\n\n")

	if err := Text2TypedVal(&m, text); err != nil {
		log.Printf("Text2TypedVal: %v, %T\n", m, m)
	} else {
		log.Printf("Text2TypedVal: %v, %T\n", m, m)
	}
	{
		var m map[string]string
		text := "a:0xb,c:0xd,e:0xf"
		if err := Text2TypedVal(&m, text); err != nil {
			log.Printf("Text2TypedVal: %v, %T\n", m, m)
		} else {
			log.Printf("Text2TypedVal: %v, %T\n", m, m)
		}
		if err := Text2TypedVal(m, text); err != nil {
			log.Printf("Not a pointer ? Text2TypedVal: %v, %T\n", m, m)
		} else {
			log.Printf("Text2TypedVal: %v, %T\n", m, m)
		}
		log.Printf("m:  %v %T\n", m, m)
	}
	{
		var m map[string]int
		text := "a:0xb,c:0xd,e:0xf"
		if err := Text2TypedVal(&m, text); err != nil {
			log.Printf("Text2TypedVal: %v, %T\n", m, m)
		} else {
			log.Printf("Text2TypedVal: %v, %T\n", m, m)
		}
		if err := Text2TypedVal(m, text); err != nil {
			log.Printf("Not a pointer ? Text2TypedVal: %v, %T\n", m, m)
		} else {
			log.Printf("Text2TypedVal: %v, %T\n", m, m)
		}
		log.Printf("m:  %v %T\n", m, m)
	}
	fmt.Printf("\n\n")

	{
		var myapp myAPP
		if err := ParseStruct2TypedVal(&myapp); err != nil {
			log.Printf("Error %v %q\n", myapp, err)
		} else {
			jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
			log.Printf("\n%v\n", string(jsonText))
		}

		log.Println(myapp)
		log.Printf("%v %T\n", myapp, myapp)

		var sti *StructInfo = &StructInfo{
			StructPtr: &myapp,
		}
		sti.Process()
		log.Println(myapp)
		log.Printf("%v %T\n", myapp, myapp)
		jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
		log.Printf("\n%v\n", string(jsonText))

		flag.Parse()
		log.Printf("%v %T\n", myapp, myapp)
		jsonText, _ = json.MarshalIndent(&myapp, "", "  ")
		log.Printf("\n%v\n", string(jsonText))

	}
	flag.Usage()
	os.Exit(0)
}
