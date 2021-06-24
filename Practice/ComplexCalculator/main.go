package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
) 	

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("ComlexCalculator")

	value1F := getInput("value 1")
	value2F := getInput("value 2")
	var result float64
	operator := getOperation()
	switch  operator {
	case "+" :
		result = add(value1F, value2F)
	case "-" :
		result = substract(value1F, value2F)
	case "*" :
		result = multiply(value1F, value2F)
	case "/" :
		result = divide(value1F, value2F)

	default :
		fmt.Println("Pls choose correct operator")
	}

	if result != 0 {
		result = (math.Round(result * 100)) / 100
		fmt.Printf(" %v  %v  %v = %v\n", value1F,operator,  value2F, result)
	}
}

func getInput(prompt string) float64 {
	fmt.Println(prompt)
	value, _ := reader.ReadString('\n')

	valueF,err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err!= nil {
		fmt.Println(err)
		panic("wrong input, please correct it")
	}
	return valueF
}

func getOperation() string {
	fmt.Println("Operator (+, -, *, /) :")
	value, _ := reader.ReadString('\n')

	return strings.TrimSpace(value)
}

func  add(value1, value2 float64) float64 {
	return value1 + value2
}

func  substract(value1, value2 float64) float64 {
	return value1 - value2
}

func  multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func  divide(value1, value2 float64) float64 {
	return value1 / value2
}

