package Factory

import "L2/patterns/Factory/Factory/Interfaces"

type EngineStyleFactory struct {
}

func (e *EngineStyleFactory) CreateTransportFactory(style string) Interfaces.TransportEngineTypeFactory {
	if style == "electricity" {
		return &ElectricityFactory{}
	} else if style == "engine" {
		return &EngineFactory{}
	} else {
		return nil
	}
}
