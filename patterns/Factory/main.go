package main

import "L2/patterns/Factory/Factory"

func main() {
	styleFactory := Factory.EngineStyleFactory{}
	transportEngineStyleFactory := styleFactory.CreateTransportFactory("electricity")
	transportFactory := transportEngineStyleFactory.CreateTransportFactory("car")
	concreteTransportStyleFactory := transportFactory.CreateStyleFactory("cargo")
	transport := concreteTransportStyleFactory.CreateTransport("red")

	transport.Move()
}
