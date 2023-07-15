package main

/*
#cgo CFLAGS:-I${SRCDIR}/c/ -I${SRCDIR}/build/
#cgo LDFLAGS:-L${SRCDIR}/build/  -Wl,-rpath,${SRCDIR}/build -lmycso

#include <stdlib.h>
#include "callgofunction.h"
*/
import "C"

// this function is not called
func callCFunctionWhichCallsGoFunction() {
	C.call_go_function()
}

func main() {
	// doesn't make a difference
	// runtime.LockOSThread()
}
