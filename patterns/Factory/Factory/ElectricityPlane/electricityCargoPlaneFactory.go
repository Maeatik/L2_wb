package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type electricityCargoPlaneFactory struct {
}

func (e *electricityCargoPlaneFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &electricityBlueCargoPlane{}
	} else if color == "red" {
		return &electricityRedCargoPlane{}
	} else {
		return nil
	}
}
