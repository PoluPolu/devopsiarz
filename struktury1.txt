package main

import "fmt"

type Car struct {
	brand  string
	model  string
	year   int
	newest bool
	prices map[string]float64
	Engine
}

type Engine struct {
	name     string
	power    int
	capacity float64
}

func struktury() {
	PandasEnggine := Engine{
		name:     "Suszarka",
		power:    50,
		capacity: 0.599,
	}

	PiotrsCar := Car{
		brand:  "Fiat",
		model:  "Pandzia",
		year:   2010,
		newest: false,
		prices: map[string]float64{
			"DE":  5000.99,
			"PL":  6000.99,
			"USA": 7000.99,
		},
		Engine: PandasEnggine,
	}

	fmt.Printf("%+v\n", PiotrsCar)

	fmt.Printf("=> Sam. Piotra to %s %s\n", PiotrsCar.brand, PiotrsCar.model)
	fmt.Printf("=> Silnik sam. Piotra to %f l\n", PiotrsCar.Engine.capacity)

}
