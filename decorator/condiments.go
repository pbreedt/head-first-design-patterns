package decorator

// There is no template / base "Condiment"
// Start defining condiments, each applied to a beverage

// SteamedMilk applied to a beverage
type SteamedMilk struct {
	beverage Beverage
}

// NewSteamedMilk: SteamedMilk can only be created with a beverage
func NewSteamedMilk(b Beverage) Beverage {
	sm := SteamedMilk{}
	sm.beverage = b
	return Beverage(sm)
}

// Description of Beverage is altered by the SteamedMilk
func (sm SteamedMilk) Description() string {
	return sm.beverage.Description() + ", Steamed Milk"
}

// Cost of Beverage is altered by the SteamedMilk
func (sm SteamedMilk) Cost() float64 {
	return sm.beverage.Cost() + 0.10
}

// Mocha applied to a beverage
type Mocha struct {
	beverage Beverage
}

func NewMocha(b Beverage) Beverage {
	m := Mocha{}
	m.beverage = b
	return Beverage(m)
}

func (m Mocha) Description() string {
	return m.beverage.Description() + ", Mocha"
}

func (m Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.20
}

// Soy applied to a beverage
type Soy struct {
	beverage Beverage
}

func NewSoy(b Beverage) Beverage {
	s := Soy{}
	s.beverage = b
	return Beverage(s)
}

func (s Soy) Description() string {
	return s.beverage.Description() + ", Soy"
}

func (s Soy) Cost() float64 {
	return s.beverage.Cost() + 0.15
}

// Whip (whipped cream, I suppose) applied to a beverage
type Whip struct {
	beverage Beverage
}

func NewWhip(b Beverage) Beverage {
	w := Whip{}
	w.beverage = b
	return Beverage(w)
}

func (w Whip) Description() string {
	return w.beverage.Description() + ", Whip"
}

func (w Whip) Cost() float64 {
	return w.beverage.Cost() + 0.10
}
