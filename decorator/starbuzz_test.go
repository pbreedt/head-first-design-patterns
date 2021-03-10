package decorator

import (
	"fmt"
	"testing"
)

func TestBeverage(t *testing.T) {
	// var bev Beverage
	es := Espresso{}
	fmt.Println(es.Description(), "@", es.Cost())

	var bev2 Beverage
	bev2 = DarkRoast{}
	bev2 = NewMocha(bev2)
	bev2 = NewMocha(bev2)
	bev2 = NewWhip(bev2)
	fmt.Println(bev2.Description(), "@", bev2.Cost())

	var bev3 Beverage
	bev3 = HouseBlend{}
	bev3 = NewSoy(bev3)
	bev3 = NewMocha(bev3)
	bev3 = NewWhip(bev3)
	fmt.Println(bev3.Description(), "@", bev3.Cost())

	var bev4 Beverage
	bev4 = Decaf{}
	bev4 = NewSoy(bev4)
	fmt.Println(bev4.Description(), "@", bev4.Cost())
}
