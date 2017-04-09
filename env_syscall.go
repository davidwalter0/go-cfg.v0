// +build !appengine

package envflagstructconfig

import "syscall"

var lookupEnv = syscall.Getenv
