package main

import (
	"fmt"
	"sort"
)

func main() {
	//this is array
	var colors = [3]string{"Red", "Green", "Yellow"}
	fmt.Println(colors)
	fmt.Println(colors[1])

	// remove index then slice

	var colorsSlice = []string{"Red", "Green", "Yellow"}
	fmt.Println(colorsSlice)
	colorsSlice = append(colorsSlice, "Purple")
	fmt.Println(colorsSlice)

	colorsSlice = append(colorsSlice[1:len(colorsSlice)])

	fmt.Println(colorsSlice)

	colorsSlice = append(colorsSlice[:len(colorsSlice)-1])

	fmt.Println(colorsSlice)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 135
	numbers[2] = 136
	numbers[3] = 137
	numbers[4] = 138
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)


}
