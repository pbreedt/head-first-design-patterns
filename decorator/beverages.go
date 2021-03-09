package decorator

// Beverage defined (not using Go -er convention for interfaces)
type Beverage interface {
	Description() string
	Cost() float64
}

// Espresso is a Beverage, with own Description and Cost
type Espresso struct{}

func (b Espresso) Description() string {
	return "Espresso"
}

func (b Espresso) Cost() float64 {
	return 1.99
}

// HouseBlend is a Beverage, with own Description and Cost
type HouseBlend struct{}

func (b HouseBlend) Description() string {
	return "House Blend Coffee"
}

func (b HouseBlend) Cost() float64 {
	return 0.89
}

// DarkRoast is a Beverage, with own Description and Cost
type DarkRoast struct{}

func (b DarkRoast) Description() string {
	return "Dark Roast Coffee"
}

func (b DarkRoast) Cost() float64 {
	return 0.99
}

// Decaf is a Beverage, with own Description and Cost
type Decaf struct{}

func (b Decaf) Description() string {
	return "Decaf Coffee"
}

func (b Decaf) Cost() float64 {
	return 1.05
}
