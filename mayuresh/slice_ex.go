package mayuresh

import "fmt"

func SliceLenAndCap() {
	x := make([]int, 4, 5)

	fmt.Println("Initial len and cap: ", len(x), cap(x))

	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 90)

	fmt.Println("Append {1,2,3,4,5,6} ==> len and cap: ", x, len(x), cap(x))

	x = append(x, 8,9,7)

	fmt.Println("Aga in append {7,8,9} len and cap: ", x, len(x), cap(x))

	/* Initial len and cap:  4 5
	Append {1,2,3,4,5,6}  ==> len and cap:  10 10  
	Again append {8,9,7} len and cap:  13 20 */
}

func SliceTwoD() {
	a := make([][]uint8, 5)

	for i, v := range a {
		fmt.Println(i, v)
	}
}

// SliceUsingNew - It will also create slice but returns pointers to initialized memory.
func SliceUsingNew() {
	slice := new([20]int)[0:10]

	fmt.Println(slice)
}

