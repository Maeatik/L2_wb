package Interfaces

type TransportEngineTypeFactory interface {
	CreateTransportFactory(style string) TransportFactory
}
