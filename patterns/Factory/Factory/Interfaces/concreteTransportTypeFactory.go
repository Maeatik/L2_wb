package Interfaces

type ConcreteTransportStyleFactory interface {
	CreateTransport(color string) Transport
}
