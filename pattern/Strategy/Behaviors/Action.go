package Behaviors

import "L2/pattern/Strategy/Animals"

type Action struct {
	Behavior Behavior
}

func (a *Action) DoSomething(animal Animals.Animal) {
	a.Behavior.DoSomething(animal)
}
