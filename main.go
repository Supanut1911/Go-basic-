package main

import (
	"basis/person"
	"fmt"
	"unicode/utf8"
)

//package scope / global var
var z int

// type Person struct {
// 	Id int
// 	Name string
// }

// func (p Person) Hello() string {
// 	return "heelo" + p.Name
// }

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
	// sliceList()
	// mapObject()
	// forLoop()
	c := sumFunction(10,20)
	_ = c

	//case use all tuple
	add, sub, text := tupleReturn(33,66)

	//case use some tuple
	xy, _, _ := tupleReturn(22,11)

	// fmt.Print(add, sub, text)
	// fmt.Println(xy)

	_ = add
	_ = sub
	_ = text
	_ = xy


	//anonymus func
	af := func(a, b int) int {
		return a * b
	}

	// fmt.Println(af(3,8))
	_ = af

	resCal := cal(sumFunction)
	_ = resCal
	// fmt.Printf("resCal: %v\n", resCal)

	ano := func(a,b int)int {
		return a / b
	}

	resCalAno := cal(ano)
	// fmt.Printf("resCalAno: %v\n", resCalAno)

	_ = resCalAno

	mySlices := []int{1,2,3,4,5,6,7,8,9,10}
	// fmt.Printf("sumSlice(mySlices): %v\n", sumSlice(mySlices))

	_ = mySlices

	ssv := sumSliceValidict(1,2,3,4,5,6,7,8,9,10)
	// fmt.Printf("ssv: %v\n", ssv)

	_ = ssv

	// fmt.Printf("customer.Name: %v\n", customer.Name)
	
	// pointPointer()

	dealWithStruct()

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

func mapObject(){
	//full declaration map variable
	// var countries map[string]string{}
	
	//short declaration
	countries :=   map[string]string{"greeting": "yoyo"}

	countries["en"] = "Hello"
	countries["th"] = "สวัสดี"



	
	// map แต่ละตัวจะคืนค่ามา 2 อย่างคือ value, boolean_valid_data
	county, statusCountry := countries["jp"]

	//ใช้ map ร่วมกับ if
	if statusCountry {
		fmt.Print("found =>", county)
	} else {
		fmt.Printf("invalid key")
	}
	//การคืนค่าของ go สามารถคืนได้มากกว่า 1 ค่า เรียกว่า tuple 


	fmt.Printf( county, statusCountry)

}

func forLoop() {
	value := []int{1,2,3,4,5}
	
	for i:=0; i < len(value); i++ {
		println(value[i])
	}

	for j,v := range value {
		println(j, v)
	}

	for _,v := range value {
		println("yo =>", v)
	}
}

//single return
func sumFunction(a int,b int) int {
	return a+b
}

//multi return -> tuple  
func tupleReturn(a int, b int) (int, int, string) {
	c := a+b
	d := a-b
	return c,d, "oyo"
}

func cal(fn func(int,int)int) int {
	return fn(70,20)
}

func sumSlice(a []int)int {
	s := 0
	for _, v := range a {
		s+= v
	}
	return s
}

func sumSliceValidict(a ...int) int {
	s := 0
	for _, v := range a {
		s+= v
	}
	return s
} 

func pointPointer() {
	// x := 10
	// var y *int = &x
	// var z *int = y
	// _ = z

	// fmt.Println(x, &x)
	// fmt.Println(*y, y, &y)
	// fmt.Print(*z, z, &z)

	var resSum int 
	sumPointer(&resSum)
	fmt.Printf("resSumPointer: %v\n", resSum)
}

func sumPointer (result *int) {
	a :=10
	b :=20
	*result = a+b
}

func dealWithStruct () {

	//create instance 
	x := person.Person{}
	x.SetName("Jame")

	fmt.Printf("%#v", x.GetName())



	// fmt.Printf("p.Name: %v\n", p.Name)
	// fmt.Println(Hello(p))
	// fmt.Println(p.Hello())
}