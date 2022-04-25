package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type ElectricityCarFactory struct {
}

func (e *ElectricityCarFactory) CreateStyleFactory(style string) Interfaces.ConcreteTransportStyleFactory {
	if style == "passanger" {
		return &electricityPassengerCarFactory{}
	} else if style == "cargo" {
		return &electricityCargoCarFactory{}
	} else {
		return nil
	}
}
