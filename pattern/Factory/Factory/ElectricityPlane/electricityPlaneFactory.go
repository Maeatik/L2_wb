package Factory

import (
	"L2/pattern/Factory/Factory/Interfaces"
)

type ElectricityPlaneFactory struct {
}

func (e *ElectricityPlaneFactory) CreateStyleFactory(style string) Interfaces.ConcreteTransportStyleFactory {
	if style == "passanger" {
		return &electricityPassengerPlaneFactory{}
	} else if style == "cargo" {
		return &electricityCargoPlaneFactory{}
	} else {
		return nil
	}
}
