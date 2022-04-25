package Behaviors

import (
	"L2/patterns/Strategy/Animals"
	"fmt"
)

type Healer struct {
}

func (h *Healer) DoSomething(animal Animals.Animal) {
	name := animal.GetName()
	fmt.Println("Я хороший робот. Я лечу", name)
}
