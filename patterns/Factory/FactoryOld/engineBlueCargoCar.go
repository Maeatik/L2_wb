package Factory

import "fmt"

type engineBlueCargoCar struct {
}

func (e *engineBlueCargoCar) move() {
	fmt.Println("Едет дизельная синяя грузовая машина")
}
