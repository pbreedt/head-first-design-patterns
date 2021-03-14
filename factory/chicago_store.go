package factory

import "fmt"

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

	switch pizzaType {
	case "cheese":
		p = NewChicagoStyleCheesePizza()
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
	//...not implemented
}
type ChicagoStyleClamPizza struct {
	//...not implemented
}
type ChicagoStyleVeggiePizza struct {
	//...not implemented
}
