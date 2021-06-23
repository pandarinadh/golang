package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)	



func main() {
	fmt.Println("Simple Calculator")
	fmt.Println("Enter value 1 :")
	reader := bufio.NewReader(os.Stdin)

	value1, _ := reader.ReadString('\n')

	value1F,err := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	if err!= nil {
		fmt.Println(err)
		panic("value 1 must be number")
	}

	fmt.Println("Enter value 2 :")
	reader = bufio.NewReader(os.Stdin)

	value2, _ := reader.ReadString('\n')

	value2F,err := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	if err!= nil {
		fmt.Println(err)
		panic("value 2 must be number")
	}

	valueSum := value1F + value2F

	fmt.Println("Sum of both: ", valueSum)


}