package Behaviors

import "L2/patterns/Strategy/Animals"

type Action struct {
	Behavior Behavior
}

func (a *Action) DoSomething(animal Animals.Animal) {
	a.Behavior.DoSomething(animal)
}
