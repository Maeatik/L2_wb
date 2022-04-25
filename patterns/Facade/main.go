package main

import "L2/patterns/Facade/Peripherial"

func main() {
	amplifier := Peripherial.Amplifier{Name: "Усилитель 1"}
	projector := Peripherial.Projector{Name: "Проектор 2"}
	dvd := Peripherial.DvdPlayer{Name: "Проектор 3"}
	sound := Peripherial.SoundPlayer{Name: "Колонки 4"}
	screen := Peripherial.Screen{Name: "Экран 5"}

	homeTheatre := Peripherial.HomeTheatreFacade{amplifier, projector, screen, sound, dvd}
	homeTheatre.WatchFilm("GO-GO-GOPHER")
	homeTheatre.Stop()
}
