package factory

import "fmt"

/*
PizzaCreator defines the "Factory method"
Why an additional interface for CreatePizza?
	If CreatePizza is added to PizzaStore, we also have provide an implementation with BasePizzaStore as receiver to satify the interface
	This causes the OrderPizza to call the BasePizzaStore.CreatePizza method which not create the different pizza types.
	To fix that we have to cater for the specific type (ChicagoPizzaStore / NYPizzaStore) in order to create the correct type of pizza - which is definitely not what we want.
*/
type PizzaCreator interface {
	CreatePizza(pizzaType string) Pizza
}

// PizzaStore no longer have the CreatePizza, but OrderPizza now takes a PizzaCreator as input
type PizzaStore interface {
	OrderPizza(PizzaCreator, pizzaType string) Pizza
}

type BasePizzaStore struct{}

//Base implementation for a PizzaStore
func (bps BasePizzaStore) OrderPizza(pc PizzaCreator, pizzaType string) Pizza {
	p := pc.CreatePizza(pizzaType)
	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return p
}

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()

	GetName() string
	GetDough() string
	GetSauce() string
	GetToppings() []string
}

type BasePizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p BasePizza) GetName() string {
	return p.name
}
func (p BasePizza) GetDough() string {
	return p.dough
}
func (p BasePizza) GetSauce() string {
	return p.sauce
}
func (p BasePizza) GetToppings() []string {
	return p.toppings
}

// Base behavior for preparing a Pizza
func (p BasePizza) Prepare() {
	fmt.Println("Preparing", p.GetName())
	fmt.Println("Tossing dough...", p.GetDough())
	fmt.Println("Adding sauce...", p.GetSauce())
	fmt.Println("Adding toppings...")
	for _, top := range p.GetToppings() {
		fmt.Println("\t", top)
	}
}

func (p BasePizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p BasePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p BasePizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}
