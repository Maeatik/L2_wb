package Factory

import "fmt"

type electricityBluePassengerCar struct {
}

func (e *electricityBluePassengerCar) Move() {
	fmt.Println("Едет электрическая синяя пассажирная машина")
}
