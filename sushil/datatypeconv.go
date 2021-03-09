package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main(){

	fmt.Println("conversion of String to any datatype :-")

	a,err := strconv.Atoi("35") // return two value (int, err)
	if err != nil{
		fmt.Println("error")
	}
	fmt.Printf("%d is of %T\n",a,a)

	b,_ := strconv.ParseBool("1") // we can pass string like '1','t','T','true','True','TRUE' for true value and same for false
	fmt.Printf("%t is of %T\n",b,b)

	c,_ := strconv.ParseFloat("3.146782",64) // takes two paramerter (str, bitsize)
	fmt.Printf("%f is of %T\n",c,c)

	d,_ := strconv.ParseInt("24",10,64) // takes three parameters (str,basesize,bitsize)
	fmt.Printf("%d is of %T\n",d,d)

	e,_ := strconv.ParseUint("12",10,32) // takes three parameters (str,basesize,bitsize)
	fmt.Printf("%d is of %T\n",e,e)

	fmt.Println("conversion of any datatype to string :- ")

	f := strconv.FormatBool(true)
	fmt.Printf("%q is of %T\n",f,f)

	g := strconv.FormatInt(-34,8) // takes two parameter (int,basesize)
	fmt.Printf("%q is of %T\n",g,g)

	h := strconv.FormatFloat(23.2423,'f',-1,32) // four paramerters (float,fmt size,prec,bitsize)
	fmt.Printf("%q is of %T\n",h,h)

	i := strconv.FormatUint(224,10) // takes two parameter (uint,basesize)
	fmt.Printf("%q is of %T\n",i,i)
	fmt.Println(reflect.ValueOf(i).Kind())

}