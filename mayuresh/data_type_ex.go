package mayuresh

import "fmt"

func ZeroValuesDefault() {

	var intZero int
	var strZero string
	var floatZero float32
	var boolZero bool
	newZeroValue := new(int)

	fmt.Println("new Zero Value: ", *newZeroValue)
	fmt.Printf("Int Zero Value: %d, %T\n", intZero, intZero)
	fmt.Printf("Str Zero Value: %s, %T\n", strZero, strZero)
	fmt.Printf("Float Zero Value: %f, %T\n", floatZero, floatZero)
	fmt.Printf("Bool Zero Value: %t, %T\n", boolZero, boolZero)

	/*Int Zero Value: 0, int
	Str Zero Value: , string
	Float Zero Value: 0.000000, float32
	Bool Zero Value: false, bool*/
}

func ZeroValuesComposite() {

	var pointerZero *int
	var interfaceZero interface{}
	var s []int
	if s == nil {
		fmt.Println("Slice is nil: ", s)
	}
	var arrayInt [3]int = [3]int{}
	type person struct {
		name        string
		div         string
		year        int
		isGraduated bool
	}

	var mapZero map[int]int
	if mapZero == nil {
		fmt.Println("Map is nil: ", mapZero)
	}

	var personZeroVal person

	fmt.Println("Ponter Zero Value: ", pointerZero)
	fmt.Println("Interface Zero Value: ", interfaceZero)
	fmt.Printf("Array Int Zero Value: %d, %T\n", arrayInt, arrayInt)
	fmt.Println("Type Struct Zero Value: ", personZeroVal)

	/*Slice is nil:  []
	Map is nil:  map[]
	Ponter Zero Value:  <nil>
	Interface Zero Value:  <nil>
	Array Int Zero Value: [0 0 0], [3]int
	Type Struct Zero Value:  {  0 false}*/
}
