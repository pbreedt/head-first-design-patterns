package factory

// NYPizzaStore implements the PizzaStore interface since it inherits this behavior from BasePizzaStore
type NYPizzaStore struct {
	BasePizzaStore
}

func NewNYPizzaStore() NYPizzaStore {
	return NYPizzaStore{}
}

// NYPizzaStore is also a PizzaCreator which implements the Factory Method to create NY-style Pizzas
func (cps NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	var p Pizza

	switch pizzaType {
	case "cheese":
		p = NewNYStyleCheesePizza()
		return p
	}
	p = BasePizza{}
	return p
}

// NYStyleCheesePizza inherits from BasePizza & thus implements the Pizza interface
type NYStyleCheesePizza struct {
	BasePizza
}

func NewNYStyleCheesePizza() NYStyleCheesePizza {
	cp := NYStyleCheesePizza{}
	cp.name = "NY Style Sauce and Cheese Pizza"
	cp.dough = "Thin Crust Dough"
	cp.sauce = "Marinara Sauce"
	cp.toppings = append(cp.toppings, "Grated Reggiano Cheese")
	return cp
}

type NYStylePepperoniPizza struct {
	//...not implemented
}
type NYStyleClamPizza struct {
	//...not implemented
}
type NYStyleVeggiePizza struct {
	//...not implemented
}
