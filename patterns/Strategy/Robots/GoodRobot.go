package Robots

import (
	"L2/patterns/Strategy/Animals"
	"L2/patterns/Strategy/Behaviors"
)

type GoodRobot struct {
	behavior Behaviors.Behavior
}

func (g *GoodRobot) DoSomething(animal Animals.Animal) {
	g.behavior.DoSomething(animal)
}

func (g *GoodRobot) SetBehavior(behavior Behaviors.Behavior) {
	g.behavior = behavior
}
