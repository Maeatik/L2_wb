package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type EnginePlaneFactory struct {
}

func (e *EnginePlaneFactory) CreateStyleFactory(style string) Interfaces.ConcreteTransportStyleFactory {
	if style == "passanger" {
		return &enginePassengerPlaneFactory{}
	} else if style == "cargo" {
		return &engineCargoPlaneFactory{}
	} else {
		return nil
	}
}
