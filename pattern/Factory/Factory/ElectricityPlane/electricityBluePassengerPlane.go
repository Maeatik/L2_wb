package Factory

import "fmt"

type electricityBluePassengerPlane struct {
}

func (e *electricityBluePassengerPlane) Move() {
	fmt.Println("Летит электрический синий пассажирский самолет")
}
