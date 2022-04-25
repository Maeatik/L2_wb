package Peripherial

import "fmt"

type DvdPlayer struct {
	Name  string
	movie string
}

func (d *DvdPlayer) setName(name string) {
	d.Name = name
}

func (d *DvdPlayer) on() {
	fmt.Println("Включен фильм")
}
func (d *DvdPlayer) off() {
	fmt.Println("Фильм выключен")
}
func (d *DvdPlayer) setMovie(name string) {
	fmt.Println("Загружен фильм:", name)
}
