package main

import (
	"fmt"
	"reflect"
)

func variadic(h ...interface{}){
	for _,v := range h{
		fmt.Println(v,"-->",reflect.ValueOf(v).Kind(	))
	}
}

func main(){
	variadic(1,"hello",[3]int{1,2,3},10.23,[]string{"hello","world"},map[string]int{"sushil":24,"neel":24})
}