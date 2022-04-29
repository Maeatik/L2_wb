package Peripherial

import "fmt"

type SoundPlayer struct {
	Name string
}

func (s *SoundPlayer) setName(name string) {
	s.Name = name
}

func (s *SoundPlayer) on() {
	fmt.Println("Звук включен")
}
func (s *SoundPlayer) off() {
	fmt.Println("Звук выключен")
}
