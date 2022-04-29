package Behaviors

import (
	"L2/pattern/Strategy/Animals"
	"fmt"
)

type Washer struct {
}

func (w *Washer) DoSomething(animal Animals.Animal) {
	name := animal.GetName()
	fmt.Println("Я хороший робот. Я мою", name)
}
