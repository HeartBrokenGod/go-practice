package neel

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Constants() {

	const n = (1 << 63)

	var num uint64 = n

	fmt.Println(num)

	fmt.Println("num type is ", reflect.TypeOf(num))

	fmt.Println("num size is ", unsafe.Sizeof(num), " bytes")

}
