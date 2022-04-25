package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type EngineCarFactory struct {
}

func (e *EngineCarFactory) CreateStyleFactory(style string) Interfaces.ConcreteTransportStyleFactory {
	if style == "passanger" {
		return &enginePassengerCarFactory{}
	} else if style == "cargo" {
		return &engineCargoCarFactory{}
	} else {
		return nil
	}
}
