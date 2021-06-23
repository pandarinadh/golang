package main

import (
	"fmt"
) 	



func main() {
	fmt.Println("structs")

	camry := Vehcle {"Camry", 4}

	fmt.Println(camry)

	fmt.Printf("%v\n", camry)
	fmt.Printf("Name:  %v\n tires: %v\n", camry.name,camry.numberOfTires)
	camry.numberOfTires = 5
	fmt.Printf("Name:  %v\n tires: %v\n", camry.name,camry.numberOfTires)

}
// upper case means public or export, lower case not export or private
type Vehcle struct {
	name string 
	numberOfTires int
}