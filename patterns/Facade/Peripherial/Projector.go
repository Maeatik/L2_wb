package Peripherial

import "fmt"

type Projector struct {
	Name string
}

func (p *Projector) setName(name string) {
	p.Name = name
}

func (p *Projector) on() {
	fmt.Println("Проектор включен")
}
func (p *Projector) off() {
	fmt.Println("Проектор выключен")
}
