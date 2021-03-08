package mayuresh

import (
	"fmt"
	"time"
)

func BasicSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("It's Saturday")
	case time.Friday:
		fmt.Println("It's Friday")

	default:
		fmt.Println("Between Sunday and Thrusday")
	}
}

func NoCondSwitch() {
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}

/*  A fallthrough statement transfers control to the next case.
    It may be used only as the final statement in a clause. */

func FallThroughSwitch() {
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}
}

func SwitchAsTypeAssertions() {
	var i interface{} = 5
	switch v := i.(type) {
	case int:
		fmt.Println("Its an int: ", v)
	case string:
		fmt.Println("Its a string: ", v)
	default:
		fmt.Println("Don't know about type")
	}
}
