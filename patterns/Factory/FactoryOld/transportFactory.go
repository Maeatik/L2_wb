package Factory

type transportFactory interface {
	createStyleFactory(style string) concreteTransportStyleFactory
}
