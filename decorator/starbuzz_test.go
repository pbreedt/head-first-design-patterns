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
	bev2 = NewMocha(Beverage(bev2))
	bev2 = NewMocha(Beverage(bev2))
	bev2 = NewWhip(Beverage(bev2))
	fmt.Println(bev2.Description(), "@", bev2.Cost())

	var bev3 Beverage
	bev3 = HouseBlend{}
	bev3 = NewSoy(Beverage(bev3))
	bev3 = NewMocha(Beverage(bev3))
	bev3 = NewWhip(Beverage(bev3))
	fmt.Println(bev3.Description(), "@", bev3.Cost())

}
