package Factory

import "fmt"

type engineRedPassengerPlane struct {
}

func (e *engineRedPassengerPlane) move() {
	fmt.Println("Летит дизельный красный пассажирский самолет")
}
