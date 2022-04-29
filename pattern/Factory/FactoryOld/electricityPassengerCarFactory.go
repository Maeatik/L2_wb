package Factory

type electricityPassengerCarFactory struct {
}

func (e *electricityPassengerCarFactory) createTransport(color string) transport {
	if color == "blue" {
		return &electricityBluePassengerCar{}
	} else if color == "red" {
		return &electricityRedPassengerCar{}
	} else {
		return nil
	}
}
