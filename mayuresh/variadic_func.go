package mayuresh

import "fmt"

// BasicVariadic MyBasicVariadic("MyName", "abc", "cde", "hyu")
func BasicVariadic(name string, i ...string) {
	fmt.Println(name, i)
}

/*
syntax error: cannot use ... with non-final parameter i
func FinalParameterSpread(i ...string, y ...int) {
	fmt.Println(y, i)
}
*/


