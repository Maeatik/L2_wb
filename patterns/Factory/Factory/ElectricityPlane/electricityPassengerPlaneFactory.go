package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type electricityPassengerPlaneFactory struct {
}

func (e *electricityPassengerPlaneFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &electricityBluePassengerPlane{}
	} else if color == "red" {
		return &electricityRedPassengerPlane{}
	} else {
		return nil
	}
}
