package simplefactory

import (
	"fmt"
	"strings"
)

type PizzaFactory interface {
	CreatePizza(pizzaType string) Pizza
}

type SimplePizzaFactory struct{}

func (pf SimplePizzaFactory) CreatePizza(pizzaType string) Pizza {
	var p Pizza

	switch strings.ToLower(pizzaType) {
	case "cheese":
		p = CheesePizza{pizzaType}
		return p
	case "pepperoni":
		p = PepperoniPizza{pizzaType}
		return p
		// case "clam":
		// 	p = ClamPizza{pizzaType}
		// 	return p
		// case "veggie":
		// 	p = VeggiePizza{pizzaType}
		// 	return p
	}

	return p
}

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type CheesePizza struct {
	name string
}

func (p CheesePizza) Prepare() {
	fmt.Println("Preparing", p.name)
}

func (p CheesePizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p CheesePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p CheesePizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

type PepperoniPizza struct {
	name string
}

func (p PepperoniPizza) Prepare() {
	fmt.Println("Preparing", p.name)
}

func (p PepperoniPizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p PepperoniPizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p PepperoniPizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

// type ClamPizza struct {
// 	name string
// }
// type VeggiePizza struct {
// 	name string
// }
