package Factory

import (
	FactoryElectricityCar "L2/pattern/Factory/Factory/ElectricityCar"
	FactoryElectricityPlane "L2/pattern/Factory/Factory/ElectricityPlane"
	"L2/pattern/Factory/Factory/Interfaces"
)

type ElectricityFactory struct {
}

func (e *ElectricityFactory) CreateTransportFactory(style string) Interfaces.TransportFactory {
	if style == "car" {
		return &FactoryElectricityCar.ElectricityCarFactory{}
	} else if style == "plane" {
		return &FactoryElectricityPlane.ElectricityPlaneFactory{}
	}
	return nil
}
