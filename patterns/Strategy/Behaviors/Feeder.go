package Behaviors

import (
	"L2/patterns/Strategy/Animals"
	"fmt"
)

type Feeder struct {
}

func (f *Feeder) DoSomething(animal Animals.Animal) {
	name := animal.GetName()
	fmt.Println("Я хороший робот. Я кормлю", name)
}
