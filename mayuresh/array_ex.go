package mayuresh

import "fmt"

// The elements of an array or struct will have its fields zeroed if no value is specified

// CreateArray to ways in which we can declare array
func CreateArray() {
	var a [3]int
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	usingEllipsis := [...]int{1, 2, 3}

	fmt.Println(a, q, r, usingEllipsis)
}
