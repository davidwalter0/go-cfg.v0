package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"

	"github.com/davidwalter0/envflagstructconfig"
	// "github.com/davidwalter0/envflagstructconfig/flag"
)

type Key string
type Value float64

func main() {

	{
		var myapp myAPP

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
		/*
			flag.Usage()
			// Error can't call parse again
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("\n\n\n\n")
					fmt.Printf("***Error*** %v\n", err)
					fmt.Printf("\n\n\n\n")
					// sti.Parse()
				}
			}()
			sti.Parse()
		*/
	}

	os.Exit(0)
}
