package Factory

type engineFactory struct {
}

func (e *engineFactory) createTransportFactory(style string) transportFactory {
	if style == "car" {
		return &engineCarFactory{}
	} else if style == "plane" {
		return &enginePlaneFactory{}
	}
	return nil
}
