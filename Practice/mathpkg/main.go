package main

import (
	"fmt"
	"math"
)



func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Sum of int" , intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Sum of float" , floatSum)

	// no :=  because it is reassigned
	floatSum = math.Round(floatSum * 100) /100
	fmt.Println("Sum of now" , floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumference: %.2f\n" , circumference)

}