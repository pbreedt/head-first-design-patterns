package strategy

import "testing"

func TestAllDucks(t *testing.T) {
	ducks := []Ducker{
		NewDecoyDuck(),
		NewMallardDuck(),
		NewModelDuck(),
		NewRubberDuck(),
		NewRedheadDuck(),
	}
	for _, duck := range ducks {
		duck.Display()
		duck.Swim()
		duck.Fly()
		duck.Quack()
	}
}

func TestModelDuck(t *testing.T) {
	modelDuck := NewModelDuck()
	fly := FlyRocketPowered{}
	// quack := MuteQuack{}
	modelDuck.Display()
	modelDuck.Swim()
	modelDuck.Fly()
	modelDuck.SetFlyBehavior(fly)
	modelDuck.Fly()
	modelDuck.Quack()
}
