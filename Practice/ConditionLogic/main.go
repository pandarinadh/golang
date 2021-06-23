package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Conditions")

	theAnswer := 42
	var result string
	if theAnswer < 0 {
		result = "< 0"
	} else if theAnswer == 0 {
		result = "= 0"
	} else {
		result = "> 0"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "< 0"
	} else if x == 0 {
		result = "= 0"
	} else {
		result = "> 0"
	}
	fmt.Println(result)

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	// dont warp in parentheisis
	//fallthrough to execute other one too
	switch dow {
	case 1:
		result = "it's Sunday"
		//	fallthrough
	case 2:
		result = "it's Monday"
		//   fallthrough
	default:
		result = "its some other day"
	}
	fmt.Println(result)

	colors := []string{"Red", "Green", "Blue"}

	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println("first for ")
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println("second for ")
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println("third for ")
		fmt.Println(color)
	}

	value := 1

	for value < 10 {
		fmt.Println("value: ", value)
		value++
	}

	sum := 1

	for sum < 1000 {
		sum += sum
		fmt.Println("sum = ", sum)

		if sum > 200 {
			goto theEnd
		}
	}
	
	theEnd : 
		fmt.Println("End of program")
	

}
