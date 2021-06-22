package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)




func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered: ", input)

	fmt.Println("Enter a number ")
	numinput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numinput), 64)
	if err!= nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of number", aFloat)
	}

}