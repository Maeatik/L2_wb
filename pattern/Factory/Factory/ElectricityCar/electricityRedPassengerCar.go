package Factory

import "fmt"

type electricityRedPassengerCar struct {
}

func (e *electricityRedPassengerCar) Move() {
	fmt.Println("Едет электрическая красная пассажирная машина")
}
