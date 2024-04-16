package main

import "fmt"

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Car is driving (fw nad bkw)!")
}

func (c Car) Break() {
	fmt.Println("Car is breaking!")
}

func (c Car) Turn() {
	fmt.Println("Car is turning!")
}

type Truck struct{}

func (t Truck) Drive() {
	fmt.Println("Car is driving (fw nad bkw)!")
}

func (t Truck) Break() {
	fmt.Println("Car is breaking!")
}

func (t Truck) Turn() {
	fmt.Println("Car is turning!")
}

type Bicycle struct{}

func (b Bicycle) Drive() {
	fmt.Println("Car is driving (fw nad bkw)!")
}

func (b Bicycle) Break() {
	fmt.Println("Car is breaking!")
}

func (b Bicycle) Turn() {
	fmt.Println("Car is turning!")
}

func DoSthWVehicke(c Car, t Truck, b Bicycle){}

func interfejsy() {

}
