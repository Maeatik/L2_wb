package Factory

import (
	"L2/patterns/Factory/Factory/Interfaces"
)

type enginePassengerCarFactory struct {
}

func (e *enginePassengerCarFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &EngineBluePassengerCar{}
	} else if color == "red" {
		return &EngineRedPassengerCar{}
	} else {
		return nil
	}
}
