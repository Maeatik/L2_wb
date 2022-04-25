package Factory

import "fmt"

type electricityBluePassengerPlane struct {
}

func (e *electricityBluePassengerPlane) move() {
	fmt.Println("Летит электрический синий пассажирский самолет")
}
