package Peripherial

import "fmt"

type Amplifier struct {
	dvd   DvdPlayer
	sound SoundPlayer
	Name  string
}

func (a *Amplifier) setName(name string) {
	a.Name = name
}

func (a *Amplifier) on() {
	fmt.Println("Включен усилитель")
}
func (a *Amplifier) off() {
	fmt.Println("Усилитель выключен")
}
