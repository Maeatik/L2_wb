package main

import "fmt"

type projector struct {
	isOn bool
}

func (p *projector) on() {
	if p.isOn == true {
		fmt.Println("Проектор уже включен")
	} else {
		p.isOn = true
		fmt.Println("Вы включили проектор")
	}
}

func (p *projector) off() {
	if p.isOn == false {
		fmt.Println("Проектор итак выключен")
	} else {
		p.isOn = false
		fmt.Println("Вы выключили проектор")
	}

}
