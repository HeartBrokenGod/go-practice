package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var b [3]int
	b[0] = 10
	b[1] = 20

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println(" ")
	for i, v := range a {
		fmt.Println(i, "--", v)
	}
	fmt.Println(" ")
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println(" ")
	j := 0
	for range a {
		fmt.Println(a[j])
		j++
	}
	fmt.Println(" ")
	fmt.Println(b)
	fmt.Println(" ")
	c := [2]string{"hello", "World"}
	fmt.Print(c)
	fmt.Println(" ")
}
