package Factory

type engineCargoCarFactory struct {
}

func (e *engineCargoCarFactory) createTransport(color string) transport {
	if color == "blue" {
		return &engineBlueCargoCar{}
	} else if color == "red" {
		return &engineRedCargoCar{}
	} else {
		return nil
	}
}
