package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type engineCargoCarFactory struct {
}

func (e *engineCargoCarFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &EngineBlueCargoCar{}
	} else if color == "red" {
		return &EngineRedCargoCar{}
	} else {
		return nil
	}
}
