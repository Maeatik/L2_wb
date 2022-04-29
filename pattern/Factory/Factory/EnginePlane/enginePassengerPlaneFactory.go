package Factory

import (
	"L2/pattern/Factory/Factory/Interfaces"
)

type enginePassengerPlaneFactory struct {
}

func (e *enginePassengerPlaneFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &engineBluePassengerPlane{}
	} else if color == "red" {
		return &engineRedPassengerPlane{}
	} else {
		return nil
	}
}
