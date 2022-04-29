package Factory

import "fmt"

type engineBluePassengerCar struct {
}

func (e *engineBluePassengerCar) move() {
	fmt.Println("Едет дизельная синяя пассажирная машина")
}
