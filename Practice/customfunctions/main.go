package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hellow from go")
	doSomething()

	sum := addValues(5, 10)
	fmt.Println(sum)

	multiSum, multiCount := addAllValues(4, 7, 9)
	fmt.Println(multiSum)
	fmt.Println(multiCount)
}

func doSomething() {
	fmt.Println("doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addValues1(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, len(values)
}
