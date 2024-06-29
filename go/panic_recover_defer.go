package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println("this is the panic msg: ", str)
	}()
	panic("not implemented yet ..")
}

/*
the defer insures some function executs on any situations ! it executs befor any other functions or returnings or paincs or errors
recover in other hand records the cause of the panic
and panic clearly stops the execution when ever it comes along
A panic generally indicates a programmer error (for
example attempting to access an index of an array
that's out of bounds, forgetting to initialize a map, etc.)
*/
