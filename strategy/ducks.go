package strategy

import (
	"fmt"
)

// Ducker is "idiomatic" Go naming, but silly here
//- did not matter since we're just proving concepts here
type Ducker interface {
	Display()
	Swim()
	Fly()   //Book sample use PerformFly()
	Quack() //Book sample use PerformQuack()
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) Display() {
	fmt.Println("I'm a duck")
}

func (duck *Duck) SetFlyBehavior(flyBehavior FlyBehavior) {
	duck.flyBehavior = flyBehavior
}

func (duck *Duck) SetQuackBehavior(quackBehavior QuackBehavior) {
	duck.quackBehavior = quackBehavior
}

func (duck Duck) Fly() {
	// Delegate Fly to flyBehavior
	duck.flyBehavior.Fly()
}

func (duck Duck) Quack() {
	// Delegate Quack to quackBehavior
	duck.quackBehavior.Quack()
}

func (*Duck) Swim() {
	// Default swim behaviour for all ducks
	fmt.Println("All ducks float, even decoys")
}

type MallardDuck struct {
	Duck
}

func (MallardDuck) Display() {
	fmt.Println("I'm a Mallarduck duck")
}

func NewMallardDuck() *MallardDuck {
	mallardDuck := MallardDuck{}
	mallardDuck.SetFlyBehavior(FlyWithWings{})
	mallardDuck.SetQuackBehavior(RealQuack{})
	return &mallardDuck
}

type RedheadDuck struct {
	Duck
}

func (RedheadDuck) Display() {
	fmt.Println("I'm a Redhead duck")
}

func NewRedheadDuck() *RedheadDuck {
	redheadDuck := RedheadDuck{}
	redheadDuck.SetFlyBehavior(FlyWithWings{})
	redheadDuck.SetQuackBehavior(RealQuack{})
	return &redheadDuck
}

type RubberDuck struct {
	Duck
}

func (RubberDuck) Display() {
	fmt.Println("I'm a rubber duck")
}

func NewRubberDuck() *RubberDuck {
	rubberDuck := RubberDuck{}
	rubberDuck.SetFlyBehavior(FlyNoWay{})
	rubberDuck.SetQuackBehavior(Squeak{})
	return &rubberDuck
}

type DecoyDuck struct {
	Duck
}

func (DecoyDuck) Display() {
	fmt.Println("I'm a duck decoy")
}

func NewDecoyDuck() *DecoyDuck {
	decoyDuck := DecoyDuck{}
	decoyDuck.SetFlyBehavior(FlyNoWay{})
	decoyDuck.SetQuackBehavior(MuteQuack{})
	return &decoyDuck
}

type ModelDuck struct {
	Duck
}

func (ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

func NewModelDuck() *ModelDuck {
	modelDuck := ModelDuck{}
	modelDuck.SetFlyBehavior(FlyNoWay{})
	modelDuck.SetQuackBehavior(MuteQuack{})
	return &modelDuck
}
