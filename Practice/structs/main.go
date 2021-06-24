package main

import (
	"fmt"
)

func main() {
	fmt.Println("structs")

	camry := Vehcle{"Camry", 4, "silver"}

	fmt.Println(camry)

	fmt.Printf("%v\n", camry)
	fmt.Printf("Name:  %v\n tires: %v\n", camry.name, camry.numberOfTires)
	camry.numberOfTires = 5
	fmt.Printf("Name:  %v\n tires: %v\n", camry.name, camry.numberOfTires)

	camry.DisplayTires()
	camry.numberOfTires = 6
	camry.DisplayTires()
	camry.DisplayThreeTimes()
	camry.DisplayThreeTimes()

}

// upper case means public or export, lower case not export or private
type Vehcle struct {
	name          string
	numberOfTires int
	color         string
}

//DisplayTires number of tires
func (d Vehcle) DisplayTires() {
	fmt.Println(d.numberOfTires)
}

func (d Vehcle) DisplayThreeTimes() {
	d.name = fmt.Sprintf("%v %v %v", d.name, d.name, d.name)
	fmt.Println(d.name)

}
