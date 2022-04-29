package Factory

import "fmt"

type engineRedPassengerCar struct {
}

func (e *engineRedPassengerCar) move() {
	fmt.Println("Едет дизельная красная пассажирная машина")
}
