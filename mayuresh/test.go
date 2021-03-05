package mayuresh

import "fmt"

type person struct {
	name string
	age  int
}

type child struct {
	name string
	age  int
}

type pet struct {
	name string
}

func TestMyCode() {
	bob := person{
		name: "bob",
		age:  15,
	}
	babyBob := child(bob)
	// "babyBob := pet(bob)" would result in a compilation error
	fmt.Println(bob, babyBob)
}
