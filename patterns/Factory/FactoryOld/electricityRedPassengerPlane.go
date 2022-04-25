package Factory

import "fmt"

type electricityRedPassengerPlane struct {
}

func (e *electricityRedPassengerPlane) move() {
	fmt.Println("Летит электрический красный пассажирский самолет")
}
