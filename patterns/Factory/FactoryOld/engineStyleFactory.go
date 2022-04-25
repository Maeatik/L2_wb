package Factory

type engineStyleFactory struct {
}

func (e *engineStyleFactory) createTransportFactory(style string) transportEngineTypeFactory {
	if style == "electricity" {
		return &electricityFactory{}
	} else if style == "engine" {
		return &engineFactory{}
	} else {
		return nil
	}
}
