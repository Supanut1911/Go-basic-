package main

import "fmt"

//package scope / global var
var z int

func main() {
	//declaration
	var x int = 10


	//short declaration
	y := 10

	//default value
	aa := z

	// fmt.Printf("%d %d %f %d", x,y,z, aa)
	fmt.Printf("%d %d %d %d", x,y,z, aa)

}
