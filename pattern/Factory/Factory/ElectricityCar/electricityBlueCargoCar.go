package Factory

import "fmt"

type electricityBlueCargoCar struct {
}

func (e *electricityBlueCargoCar) Move() {
	fmt.Println("Едет электрическая синяя грузовая машина")
}
