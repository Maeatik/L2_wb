package Factory

type concreteTransportStyleFactory interface {
	createTransport(color string) transport
}
