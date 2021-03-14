package factory

import (
	"fmt"
	"testing"
)

func TestPizzaStore(t *testing.T) {
	nyps := NewNYPizzaStore()
	p := nyps.OrderPizza(nyps, "cheese")
	fmt.Println("Ethan ordered a", p.GetName())

	fmt.Println("")

	cps := NewChicagoPizzaStore()
	p = cps.OrderPizza(cps, "cheese")
	fmt.Println("Joel ordered a", p.GetName())
}
