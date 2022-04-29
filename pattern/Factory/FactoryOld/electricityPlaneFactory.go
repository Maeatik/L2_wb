package Factory

type electricityPlaneFactory struct {
}

func (e *electricityPlaneFactory) createStyleFactory(style string) concreteTransportStyleFactory {
	if style == "passanger" {
		return &electricityPassengerPlaneFactory{}
	} else if style == "cargo" {
		return &electricityCargoPlaneFactory{}
	} else {
		return nil
	}
}
