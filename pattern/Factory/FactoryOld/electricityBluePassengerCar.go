package Factory

import "fmt"

type electricityBluePassengerCar struct {
}

func (e *electricityBluePassengerCar) move() {
	fmt.Println("Едет электрическая синяя пассажирная машина")
}
