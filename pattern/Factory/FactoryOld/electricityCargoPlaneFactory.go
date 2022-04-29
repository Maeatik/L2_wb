package Factory

type electricityCargoPlaneFactory struct {
}

func (e *electricityCargoPlaneFactory) createTransport(color string) transport {
	if color == "blue" {
		return &electricityBlueCargoPlane{}
	} else if color == "red" {
		return &electricityRedCargoPlane{}
	} else {
		return nil
	}
}
