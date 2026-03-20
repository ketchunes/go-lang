package main

import "fmt"

type Auto interface {
	StepOnGas()
}

type BMW struct{}

type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Жига")
}

func (b BMW) StepOnGas() {
	fmt.Println("Я бмв")
}

func rideBmw(bmw BMW) {
	fmt.Println("Я водитель")
	fmt.Println("Я сажусь в машину")
	fmt.Println("Я нажмаю на газ")
	bmw.StepOnGas()
}

func rideZhiga(zhiga Zhiga) {
	fmt.Println("Я водитель")
	fmt.Println("Я сажусь в машину")
	fmt.Println("Я нажмаю на газ")
	zhiga.StepOnGas()
}

func ride(auto Auto) {
	fmt.Println("Я водитель")
	fmt.Println("Я сажусь в машину")
	fmt.Println("Я нажмаю на газ")
	auto.StepOnGas()
}
func main() {
	bmw := BMW{}
	ride(bmw)
}
