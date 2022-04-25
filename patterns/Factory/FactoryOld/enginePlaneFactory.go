package Factory

type enginePlaneFactory struct {
}

func (e *enginePlaneFactory) createStyleFactory(style string) concreteTransportStyleFactory {
	if style == "passanger" {
		return &enginePassengerPlaneFactory{}
	} else if style == "cargo" {
		return &engineCargoPlaneFactory{}
	} else {
		return nil
	}
}
