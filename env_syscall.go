// +build !appengine

package cfg // import "github.com/davidwalter0/go-cfg"

import "syscall"

var lookupEnv = syscall.Getenv
