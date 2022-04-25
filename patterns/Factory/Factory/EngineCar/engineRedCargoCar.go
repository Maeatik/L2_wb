package Factory

import "fmt"

type EngineRedCargoCar struct {
}

func (e *EngineRedCargoCar) Move() {
	fmt.Println("Едет дизельная красная грузовая машина")
}
