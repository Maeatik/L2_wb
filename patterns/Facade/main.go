package main

import "L2/patterns/Facade/Peripherial"

func main() {
	//создаем перифирию домашнего кинотеатра
	amplifier := Peripherial.Amplifier{Name: "Усилитель 1"}
	projector := Peripherial.Projector{Name: "Проектор 2"}
	dvd := Peripherial.DvdPlayer{Name: "Проектор 3"}
	sound := Peripherial.SoundPlayer{Name: "Колонки 4"}
	screen := Peripherial.Screen{Name: "Экран 5"}

	//оборудуем кинотеатр
	homeTheatre := Peripherial.HomeTheatreFacade{amplifier, projector, screen, sound, dvd}
	//включаем фильм
	homeTheatre.WatchFilm("GO-GO-GOPHER")
	//выключаем кинотеатр
	homeTheatre.Stop()
}
