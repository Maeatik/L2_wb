package Factory

import (
	"L2/pattern/Factory/Factory/Interfaces"
)

type electricityCargoCarFactory struct {
}

func (e *electricityCargoCarFactory) CreateTransport(color string) Interfaces.Transport {
	if color == "blue" {
		return &electricityBlueCargoCar{}
	} else if color == "red" {
		return &electricityRedCargoCar{}
	} else {
		return nil
	}
}
