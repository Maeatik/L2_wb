package Factory

import "fmt"

type engineRedCargoCar struct {
}

func (e *engineRedCargoCar) move() {
	fmt.Println("Едет дизельная красная грузовая машина")
}
