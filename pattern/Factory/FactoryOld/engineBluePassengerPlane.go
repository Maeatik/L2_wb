package Factory

import "fmt"

type engineBluePassengerPlane struct {
}

func (e *engineBluePassengerPlane) move() {
	fmt.Println("Летит дизельный синий пассажирский самолет")
}
