package simplefactory

type PizzaStore interface {
	OrderPizza(pizzaType string) Pizza
}

type SimplePizzaStore struct {
	PizzaFactory PizzaFactory
}

func (sps SimplePizzaStore) OrderPizza(pizzaType string) Pizza {
	p := sps.PizzaFactory.CreatePizza(pizzaType)
	p.Prepare()
	p.Bake()
	p.Cut()
	p.Box()
	return p
}
