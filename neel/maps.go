package neel
import "fmt"
func Maps() {
	// creating map
	var mapVar map[int]string

	// initiate the map
	mapVar =  make(map[int]string)

	mapVar[1]="one"
	mapVar[2]="two"

	fmt.Println(mapVar[1])

	type Person struct{
		Name string
		Age int
	}
	type Rank int

	mapVar2 := make(map[Rank]Person)

	p1 := Person{"neel",25}
	p2 := Person{"sushil",24}

	mapVar2[0] = p1
	mapVar2[1] = p2

	fmt.Println(mapVar2[1])

	mapVar3 := make(map[Person]Rank)
	mapVar3[p1] = 0
	mapVar3[p2] = 1

	fmt.Println(mapVar3[p1])

	p3 := Person{"sushil",24}

	fmt.Println(mapVar3[p3])


	fmt.Println(mapVar3[struct{Name string; Age int}{"neel",25}])


}
