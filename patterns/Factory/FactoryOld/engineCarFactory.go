package Factory

type engineCarFactory struct {
}

func (e *engineCarFactory) createStyleFactory(style string) concreteTransportStyleFactory {
	if style == "passanger" {
		return &enginePassengerCarFactory{}
	} else if style == "cargo" {
		return &engineCargoCarFactory{}
	} else {
		return nil
	}
}
