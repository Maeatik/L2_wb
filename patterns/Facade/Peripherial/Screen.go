package Peripherial

import "fmt"

type Screen struct {
	Name string
}

func (s *Screen) setName(name string) {
	s.Name = name
}

func (s *Screen) on() {
	fmt.Println("Экран опущен")
}
func (s *Screen) off() {
	fmt.Println("Экран поднят")
}
