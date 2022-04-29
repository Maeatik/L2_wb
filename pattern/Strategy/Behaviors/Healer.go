package Behaviors

import (
	"L2/pattern/Strategy/Animals"
	"fmt"
)

type Healer struct {
}

func (h *Healer) DoSomething(animal Animals.Animal) {
	name := animal.GetName()
	fmt.Println("Я хороший робот. Я лечу", name)
}
