package main

import (
	"log"

	cfg "github.com/davidwalter0/envflagstructconfig"
)

var tracer = cfg.Tracer()

func init() {
	tracer.Disable()
}

func main() {
	defer tracer.ScopedTrace()()
	prefix := "myapp"
	var specification Specification
	cfg.Initialize(prefix, &specification)
	log.Println(specification)
}
