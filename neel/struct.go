package neel

import "fmt"

func Struct(){
	type A struct{
		Num int
		Str string
	}

	type B struct{
		Num int
		Str string
	}

	var aVar A
	var bVar B


	aVar = A{1,"one"}
	// this line will result in compile time error
	// cannot use type A as type B assignment
	// bVar = aVar

	bVar = struct{Num int; Str string}{2,"two"}

	fmt.Println("aVar : ",aVar)
	fmt.Println("bVar : ",bVar)


	aVar2 := A{3,"three"}
	aVar3 := aVar2

	fmt.Println("aVar2 : ",aVar2)
	fmt.Println("aVar3 : ",aVar3)

	aVar2.Num =  9
	aVar2 = A{9,"nine"}

	fmt.Println("aVar2 : ",aVar2)
	fmt.Println("aVar3 : ",aVar3)


}
