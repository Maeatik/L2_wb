package main

import (
	"L2/pattern/Strategy/Animals"
	"L2/pattern/Strategy/Behaviors"
	"L2/pattern/Strategy/Robots"
)

func main() {
	//создаем зверей
	lion := &Animals.Lion{Name: "Alex"}
	duck := &Animals.Duck{Name: "Donald"}
	rhino := &Animals.Rhino{Name: "Muchacho"}
	//создаем функции заботы над животными
	feed := Behaviors.Action{Behavior: &Behaviors.Feeder{}}
	wash := Behaviors.Action{Behavior: &Behaviors.Washer{}}
	heal := Behaviors.Action{Behavior: &Behaviors.Healer{}}

	//создаем робота-няньку
	robot := Robots.GoodRobot{}

	//задаем роботу функции заботы над животными и заботимся по очереди с каждым зверьком
	robot.SetBehavior(&wash)
	robot.DoSomething(lion)

	robot.SetBehavior(&feed)
	robot.DoSomething(duck)

	robot.SetBehavior(&heal)
	robot.DoSomething(rhino)

	//
	//
	//robot.DoSomething(rhino)
}
