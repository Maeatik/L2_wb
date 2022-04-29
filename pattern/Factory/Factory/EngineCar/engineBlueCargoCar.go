package Factory

import "fmt"

type EngineBlueCargoCar struct {
}

func (e *EngineBlueCargoCar) Move() {
	fmt.Println("Едет дизельная синяя грузовая машина")
}
