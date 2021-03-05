package neel
import "fmt"

func Loops(){
	i := 0

	//	while loop
	for i < 10 {
		fmt.Println("while  i = ", i)
		i++
	}

	//	do while loop
	for {
		fmt.Println("do while i = ",i)
		if(i>10){
			continue
		}
		break
	}


	//	for loop
	for i = 0; i< 10; i++{
		fmt.Println("for i = ",i)
	}



}
