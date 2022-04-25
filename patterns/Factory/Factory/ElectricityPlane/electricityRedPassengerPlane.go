package Factory

import "fmt"

type electricityRedPassengerPlane struct {
}

func (e *electricityRedPassengerPlane) Move() {
	fmt.Println("Летит электрический красный пассажирский самолет")
}
