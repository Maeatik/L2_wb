package Factory

import "fmt"

type engineBluePassengerPlane struct {
}

func (e *engineBluePassengerPlane) Move() {
	fmt.Println("Летит дизельный синий пассажирский самолет")
}
