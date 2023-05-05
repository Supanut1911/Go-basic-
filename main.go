package main

import "fmt"

//package scope / global var
var z int

func main() {
	//declaration
	var x int = 10
	_ = x

	//short declaration
	y := 10
	_ = y

	//default value
	aa := z
	_ = aa

	// fmt.Printf("%d %d %f %d", x,y,z, aa)
	// fmt.Printf("%d %d %d %d", x,y,z, aa)

	condition()

}

func condition() {
	point := 10

	if point >= 50 && point <= 100  {
		fmt.Printf("Pass")
	} else if point < 20 {
		fmt.Printf("fail")

	}
}
