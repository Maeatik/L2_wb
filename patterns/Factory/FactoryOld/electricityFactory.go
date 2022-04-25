package Factory

type electricityFactory struct {
}

func (e *electricityFactory) createTransportFactory(style string) transportFactory {
	if style == "car" {
		return &electricityCarFactory{}
	} else if style == "plane" {
		return &electricityPlaneFactory{}
	}
	return nil
}
