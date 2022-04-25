package main

import "fmt"

func main() {
	okKitBuilder := getBuilder("ok")
	luxKitBuilder := getBuilder("lux")

	dude := newDudeBuilder(okKitBuilder)
	okKit := dude.buildKitchen()

	fmt.Println("НОРМАЛЬНАЯ кухня состоит из:", okKit.microwave, ",", okKit.refrigerator, "and", okKit.plate)

	dude.setDudeBuilder(luxKitBuilder)
	luxKit := dude.buildKitchen()
	fmt.Println("СУПЕР кухня состоит из:", luxKit.microwave, ",", luxKit.refrigerator, "and", luxKit.plate)

}
