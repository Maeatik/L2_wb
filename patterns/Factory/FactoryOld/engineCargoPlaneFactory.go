package Factory

type engineCargoPlaneFactory struct {
}

func (e *engineCargoPlaneFactory) createTransport(color string) transport {
	if color == "blue" {
		return &engineBlueCargoPlane{}
	} else if color == "red" {
		return &engineRedCargoPlane{}
	} else {
		return nil
	}
}
