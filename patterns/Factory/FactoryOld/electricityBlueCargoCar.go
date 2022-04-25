package Factory

import "fmt"

type electricityBlueCargoCar struct {
}

func (e *electricityBlueCargoCar) move() {
	fmt.Println("Едет электрическая синяя грузовая машина")
}
