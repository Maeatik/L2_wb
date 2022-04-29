package Factory

type transportEngineTypeFactory interface {
	createTransportFactory(style string) transportFactory
}
