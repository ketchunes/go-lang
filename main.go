package main

import "fmt"

type Auto interface {
	StepOnGas()
	StepOnBrake()
}

type BMW struct{}

type Zhiga struct{}

type Lamba struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Жига")
}

func (b BMW) StepOnGas() {
	fmt.Println("Я бмв")
}

func (l Lamba) StepOnGas() {
	fmt.Println("ламба")
}

func (l Lamba) StepOnBrake() {
	fmt.Println("Я ламба жму на тормоз")
}
func (z Zhiga) StepOnBrake() {
	fmt.Println("жига жмет на тормоз")
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
	auto.StepOnBrake()
}
func main() {
	// lamba := Lamba{}
	// ride(lamba)
	zhiga := Zhiga{}
	ride(zhiga)
}
