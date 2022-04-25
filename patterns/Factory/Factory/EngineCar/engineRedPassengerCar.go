package Factory

import "fmt"

type EngineRedPassengerCar struct {
}

func (e *EngineRedPassengerCar) Move() {
	fmt.Println("Едет дизельная красная пассажирная машина")
}
