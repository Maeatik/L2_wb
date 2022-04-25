package Factory

import "fmt"

type EngineBluePassengerCar struct {
}

func (e *EngineBluePassengerCar) Move() {
	fmt.Println("Едет дизельная синяя пассажирная машина")
}
