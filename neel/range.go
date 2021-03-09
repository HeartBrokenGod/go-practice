package neel

import "fmt"

func Range() {

	//	range over map
	mapVar := make(map[int]int)
	mapVar[1] = 1
	mapVar[2] = 4
	mapVar[3] = 9
	for k, v := range mapVar {
		fmt.Println("mapVar[", k, "] : ", v)
	}

	//	range over slice
	sliceVar := make([]int, 5)
	for i, v := range sliceVar {
		fmt.Println("sliceVar[", i, "] : ", v)
	}

	//	range over array
	arrayVar := [10]int{}
	for i, v := range arrayVar {
		fmt.Println("arrayVar[", i, "] : ", v)
	}

	//	compile time error :	cannot range over struct

	//	type Person struct{
	//		Name string
	//		Age int
	//	}
	//
	//	p := Person{"Neel",25}
	//
	//	for s := range p {
	//		fmt.Println("struct Person",s)
	//	}

	//	range over string
	stringVar := "Hello"

	//	i is index of type int and v is character of type rune i,e int32
	for i, v := range stringVar {
		fmt.Println("stingVar[", i, "] : ", v)
	}

}
