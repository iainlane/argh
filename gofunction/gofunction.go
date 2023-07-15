package main

import "C"

//export GoFunction
func GoFunction() int {
	return 1337
}

func main() {}
