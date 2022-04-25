package Factory

import (
	FactoryEngineCar "L2/patterns/Factory/Factory/EngineCar"
	FactoryEnginePlane "L2/patterns/Factory/Factory/EnginePlane"
	"L2/patterns/Factory/Factory/Interfaces"
)

type EngineFactory struct {
}

func (e EngineFactory) CreateTransportFactory(style string) Interfaces.TransportFactory {
	if style == "car" {
		return &FactoryEngineCar.EngineCarFactory{}
	} else if style == "plane" {
		return &FactoryEnginePlane.EnginePlaneFactory{}
	}
	return nil
}
