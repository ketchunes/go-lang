package main

import "fmt"

type Auto interface {
	StepOnGas()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("вы нажали на газ bmv")
}

type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("я жига из танков лол")
}

func ride(auto Auto) {
	fmt.Println("я водитель")
	fmt.Println("я сажусь в машину")
	fmt.Println("и нажимаю на газ")
	auto.StepOnGas()
}
func main() {
	// bmw := BMW{}
	zhiga := Zhiga{}
	ride(zhiga)
}
