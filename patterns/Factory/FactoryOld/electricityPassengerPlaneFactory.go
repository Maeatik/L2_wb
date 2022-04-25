package Factory

type electricityPassengerPlaneFactory struct {
}

func (e *electricityPassengerPlaneFactory) createTransport(color string) transport {
	if color == "blue" {
		return &electricityBluePassengerPlane{}
	} else if color == "red" {
		return &electricityRedPassengerPlane{}
	} else {
		return nil
	}
}
