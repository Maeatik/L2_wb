package Factory

import "fmt"

type electricityRedCargoCar struct {
}

func (e *electricityRedCargoCar) Move() {
	fmt.Println("Едет электрическая красная грузовая машина")
}
