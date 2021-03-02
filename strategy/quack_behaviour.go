package strategy

import "fmt"

type QuackBehavior interface {
	Quack()
}

type RealQuack struct{}

func (RealQuack) Quack() {
	fmt.Println("Quack")
}

type Squeak struct{}

func (Squeak) Quack() {
	fmt.Println("Squeak")
}

type MuteQuack struct{}

func (MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}
