package Factory

import "fmt"

type electricityRedPassengerCar struct {
}

func (e *electricityRedPassengerCar) move() {
	fmt.Println("Едет электрическая красная пассажирная машина")
}
