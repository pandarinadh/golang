package main

import 	"fmt"

const aConst int = 64

func main() {
	var aString string = "this is GO!"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "this is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable type is %T\n", anotherString)

	//:= no need for var

	myString := "this is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable type is %T\n", myString)


	fmt.Println(aConst)
	fmt.Printf("The variable type is %T\n", aConst)

	
}