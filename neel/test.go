package neel

import "fmt"

func changearr(i *([5][5]int)) {

	(*i)[0][0] = 67

}

func Test() {

	var mat [5][5]int

	changearr(&mat)

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Printf("\n")

	}

}
