package Factory

import "fmt"

type electricityRedCargoCar struct {
}

func (e *electricityRedCargoCar) move() {
	fmt.Println("Едет электрическая красная грузовая машина")
}
