package Factory

type electricityCargoCarFactory struct {
}

func (e *electricityCargoCarFactory) createTransport(color string) transport {
	if color == "blue" {
		return &electricityBlueCargoCar{}
	} else if color == "red" {
		return &electricityRedCargoCar{}
	} else {
		return nil
	}
}
