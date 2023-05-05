package main

import (
	"fmt"
	"unicode/utf8"
)

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
	
	// condition()
	// helooArr()
	sliceList()
}

func condition() {
	point := 10

	if point >= 50 && point <= 100  {
		fmt.Printf("Pass")
	} else if point < 20 {
		fmt.Printf("fail")

	}
}

func helooArr() {
	//declaration fix array
	var x[3] int= [3]int{1,2,3}

	//short declaration
	y := [3]int{10,20,30}
	y[0] = 111

	//2D
	z := [3][2]int{
		{1,2},
		{10,20},
		{111,222},
	}

	//dynamic array
	aa := [...]int{1,2,3,4,5}

	//empty array
	bb := []int{}

	fmt.Printf("%#v\n %v\n %v\n %v\n %v",x, y, z, aa, bb)

}

func sliceList () {
	//emply array is slice
	//					0	 1	2	 3	4	 5	6	 7  8   9
	x := []int{10,20,30,40,50,60,70,80,90,100}

	//select slice 
	aa := x[8:10]

	//selct to le
	bb := x[1:]

	//list can append
	x = append(x, 4)
	y := append(x,10)
	z := len(y)
	
	ch := "กขค"
	cs := utf8.RuneCountInString(ch)

	fmt.Printf("%v \n %v \n %v \n", x, y, z)

	fmt.Printf("\n\n %v \n", cs)

	fmt.Printf("\n %v \n %v \n", aa, bb)

}