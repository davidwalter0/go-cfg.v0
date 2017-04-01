package main

import (
	"log"
	"os"

	cfg "github.com/davidwalter0/envflagstructconfig"
)

var tracer = cfg.Tracer()

func init() {
	tracer.Disable()
}

func main() {
	defer tracer.ScopedTrace()()
	var prefix = os.Getenv("APP_OVERRIDE_PREFIX")
	if len(prefix) == 0 {
		prefix = "myapp"
	}
	var specification Specification
	cfg.Initialize(prefix, &specification)
	log.Println(specification)
}
