package Factory

type electricityCarFactory struct {
}

func (e *electricityCarFactory) createStyleFactory(style string) concreteTransportStyleFactory {
	if style == "passanger" {
		return &electricityPassengerCarFactory{}
	} else if style == "cargo" {
		return &electricityCargoCarFactory{}
	} else {
		return nil
	}
}
