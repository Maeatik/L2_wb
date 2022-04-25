package main

type builder interface {
	setMicrowave()
	setRefrigerator()
	setPlate()
	getKitchen() kitchen
}

func getBuilder(build string) builder {
	if build == "ok" {
		return &okKitchen{}
	}
	if build == "lux" {
		return &luxKitchen{}
	}
	return nil
}
