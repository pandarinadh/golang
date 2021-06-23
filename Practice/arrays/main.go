package main

import (
	"fmt"
) 	



func main() {
	var colors [3] string 
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Yellow"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5] int {4,5,6,7}
	fmt.Println(numbers)
	fmt.Println("lenth :", len(numbers))
}