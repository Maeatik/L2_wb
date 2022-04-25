package Factory

import "fmt"

type engineRedPassengerPlane struct {
}

func (e *engineRedPassengerPlane) Move() {
	fmt.Println("Летит дизельный красный пассажирский самолет")
}
