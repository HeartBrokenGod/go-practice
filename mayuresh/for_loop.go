package mayuresh

import "fmt"

// ThreeComponent uses three conditions as in java and C
func ThreeComponent() {
	var num int = 10

	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}
}

// Aswhile instead of while loop
func AsWhile() {
	var x int = 1

	for x < 5 {
		fmt.Println(x)
		x++
	}
}

// RangeLoop Looping over elements in slices, arrays, maps, channels or strings is often better done with a range loop.
func RangeLoop() {
	var x string = "ab yxds hjkl tyes"
	y := make(map[string]int)

	fmt.Println("Using ", x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	y = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("Using ", y)

	for k, v := range y {
		fmt.Println(k, v)
	}
}
