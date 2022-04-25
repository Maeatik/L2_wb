package Behaviors

import (
	"L2/patterns/Strategy/Animals"
	"fmt"
)

type Washer struct {
}

func (w *Washer) DoSomething(animal Animals.Animal) {
	name := animal.GetName()
	fmt.Println("Я хороший робот. Я мою", name)
}
