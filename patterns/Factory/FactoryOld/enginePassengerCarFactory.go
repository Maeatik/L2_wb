package Factory

type enginePassengerCarFactory struct {
}

func (e *enginePassengerCarFactory) createTransport(color string) transport {
	if color == "blue" {
		return &engineBluePassengerCar{}
	} else if color == "red" {
		return &engineRedPassengerCar{}
	} else {
		return nil
	}
}
