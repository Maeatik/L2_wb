package Interfaces

type TransportFactory interface {
	CreateStyleFactory(style string) ConcreteTransportStyleFactory
}
