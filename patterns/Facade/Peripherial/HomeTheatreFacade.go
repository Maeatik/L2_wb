package Peripherial

import "fmt"

type HomeTheatreFacade struct {
	Amplifier Amplifier
	Projector Projector
	Screen    Screen
	Sound     SoundPlayer
	Dvd       DvdPlayer
}

//включаем аппаратуру домашнего кинотеатра
func (h *HomeTheatreFacade) WatchFilm(movie string) {
	h.Dvd.setMovie(movie)
	h.Dvd.on()
	h.Amplifier.on()
	h.Screen.on()
	h.Projector.on()
	fmt.Println("\nПриятного просмотра!\n")
}

//выключаем аппаратуру домашнего кинотеатра
func (h *HomeTheatreFacade) Stop() {
	h.Dvd.off()
	h.Projector.off()
	h.Screen.off()
	h.Amplifier.off()
}
