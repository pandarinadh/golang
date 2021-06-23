package main

import (
	"fmt"
) 	



func main() {
	anInt := 42
		var p  = &anInt
	fmt.Println("value of P:", *p)

	*p = 44;
	fmt.Println("value of P:", *p)
	fmt.Println("value of anint:", anInt)

}