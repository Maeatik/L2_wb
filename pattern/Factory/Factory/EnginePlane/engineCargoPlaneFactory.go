package Factory

import (
	"L2/pattern/Factory/Factory/Interfaces"
)

type engineCargoPlaneFactory struct {
}

func (e *engineCargoPlaneFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &engineBlueCargoPlane{}
	} else if color == "red" {
		return &engineRedCargoPlane{}
	} else {
		return nil
	}
}
