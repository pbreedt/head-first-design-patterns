package factory

import (
	"fmt"
	"strings"
)

// ChicagoPizzaStore implements the PizzaStore interface since it inherits this behavior from BasePizzaStore
type ChicagoPizzaStore struct {
	BasePizzaStore
}

func NewChicagoPizzaStore() ChicagoPizzaStore {
	return ChicagoPizzaStore{}
}

// ChicagoPizzaStore is also a PizzaCreator which implements the Factory Method to create Chicago-style Pizzas
func (cps ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	var p Pizza

	switch strings.ToLower(pizzaType) {
	case "cheese":
		p = NewChicagoStyleCheesePizza()
		return p
	case "pepperoni":
		p = NewChicagoStylePepperoniPizza()
		return p
	case "clam":
		p = NewChicagoStyleClamPizza()
		return p
	case "veggie":
		p = NewChicagoStyleVeggiePizza()
		return p
	}
	p = BasePizza{}
	return p
}

// ChicagoStyleCheesePizza inherits from BasePizza & thus implements the Pizza interface
type ChicagoStyleCheesePizza struct {
	BasePizza
}

func NewChicagoStyleCheesePizza() ChicagoStyleCheesePizza {
	cp := ChicagoStyleCheesePizza{}
	cp.name = "Chicago Style Deep Dish Cheese Pizza"
	cp.dough = "Extra Thick Crust Dough"
	cp.sauce = "Plum Tomato Sauce"
	cp.toppings = append(cp.toppings, "Shredded Mozzarella Cheese")
	return cp
}

// For some reason, only ChicagoStyleCheesePizzas are cut into squares 🤔
func (cp ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}

type ChicagoStylePepperoniPizza struct {
	BasePizza
}

func NewChicagoStylePepperoniPizza() ChicagoStylePepperoniPizza {
	//...not implemented
	return ChicagoStylePepperoniPizza{}
}

type ChicagoStyleClamPizza struct {
	BasePizza
}

func NewChicagoStyleClamPizza() ChicagoStyleClamPizza {
	//...not implemented
	return ChicagoStyleClamPizza{}
}

type ChicagoStyleVeggiePizza struct {
	BasePizza
}

func NewChicagoStyleVeggiePizza() ChicagoStyleVeggiePizza {
	//...not implemented
	return ChicagoStyleVeggiePizza{}
}
