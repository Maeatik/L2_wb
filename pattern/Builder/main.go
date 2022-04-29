package main

import "fmt"

func main() {
	//создаем двух билдеров, каждому предаем по типу кухни, который каждый их них должен строить
	okKitBuilder := getBuilder("ok")
	luxKitBuilder := getBuilder("lux")

	//создаем чувака, который будет строить обычную кухню
	dude := newDudeBuilder(okKitBuilder)
	//чувак строит обычную кухню
	okKit := dude.buildKitchen()

	fmt.Println("НОРМАЛЬНАЯ кухня состоит из:", okKit.microwave, ",", okKit.refrigerator, "and", okKit.plate)
	//чуваку дали новый заказ на крутую люксовую кухню
	dude.setDudeBuilder(luxKitBuilder)
	//чувак строит люксовую кухню
	luxKit := dude.buildKitchen()
	fmt.Println("СУПЕР кухня состоит из:", luxKit.microwave, ",", luxKit.refrigerator, "and", luxKit.plate)

}
