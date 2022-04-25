package main

import (
	"L2/patterns/Strategy/Animals"
	"L2/patterns/Strategy/Behaviors"
	"L2/patterns/Strategy/Robots"
)

func main() {
	lion := &Animals.Lion{Name: "Alex"}
	duck := &Animals.Duck{Name: "Donald"}
	rhino := &Animals.Rhino{Name: "Muchacho"}
	feed := Behaviors.Action{Behavior: &Behaviors.Feeder{}}
	wash := Behaviors.Action{Behavior: &Behaviors.Washer{}}
	heal := Behaviors.Action{Behavior: &Behaviors.Healer{}}

	robot := Robots.GoodRobot{}

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
