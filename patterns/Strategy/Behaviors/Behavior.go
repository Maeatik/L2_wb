package Behaviors

import (
	"L2/patterns/Strategy/Animals"
)

type Behavior interface {
	DoSomething(animal Animals.Animal)
}
