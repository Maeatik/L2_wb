package main

import "L2/pattern/Factory/Factory"

func main() {
	//создаем обычную фабрику - она решает какой тип двигателя будет в транспорте
	styleFactory := Factory.EngineStyleFactory{}
	transportEngineStyleFactory := styleFactory.CreateTransportFactory("electricity")
	//фабрика которая решает строить машину или самолт
	transportFactory := transportEngineStyleFactory.CreateTransportFactory("car")
	//машина пассажирская или грузовая
	concreteTransportStyleFactory := transportFactory.CreateStyleFactory("cargo")
	//цвет машины
	transport := concreteTransportStyleFactory.CreateTransport("red")
	// запускаем транспорт
	transport.Move()
}
