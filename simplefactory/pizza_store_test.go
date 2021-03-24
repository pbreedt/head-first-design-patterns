package simplefactory

import (
	"fmt"
	"testing"
)

func TestPizzaStore(t *testing.T) {
	spf := SimplePizzaFactory{}
	var ps PizzaStore = SimplePizzaStore{PizzaFactory: spf}
	p := ps.OrderPizza("cheese")

	fmt.Println("Pizza ready", p)
}
