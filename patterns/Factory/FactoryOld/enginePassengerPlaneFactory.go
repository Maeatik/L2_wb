package Factory

type enginePassengerPlaneFactory struct {
}

func (e *enginePassengerPlaneFactory) createTransport(color string) transport {
	if color == "blue" {
		return &engineBluePassengerPlane{}
	} else if color == "red" {
		return &engineRedPassengerPlane{}
	} else {
		return nil
	}
}
