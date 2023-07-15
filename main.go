package main

/*
#cgo CFLAGS:-I${SRCDIR}/c/
#cgo LDFLAGS:-L${SRCDIR}/build/ -Wl,-rpath,${SRCDIR}/build/ -lmyc

#include <stdlib.h>
#include "callgofunction.h"
*/
import "C"

// this function is not called
func isDancePerformed() {
	C.call_go_function()
}

func main() {
	// doesn't make a difference
	// runtime.LockOSThread()
}
