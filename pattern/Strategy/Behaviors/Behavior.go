package Behaviors

import (
	"L2/pattern/Strategy/Animals"
)

type Behavior interface {
	DoSomething(animal Animals.Animal)
}
