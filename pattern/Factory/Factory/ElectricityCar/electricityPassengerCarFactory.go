package Factory

import (
	"L2/pattern/Factory/Factory/Interfaces"
)

type electricityPassengerCarFactory struct {
}

func (e *electricityPassengerCarFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &electricityBluePassengerCar{}
	} else if color == "red" {
		return &electricityRedPassengerCar{}
	} else {
		return nil
	}
}
