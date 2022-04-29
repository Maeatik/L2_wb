package main

//интерфейс строителя
type builder interface {
	setMicrowave()
	setRefrigerator()
	setPlate()
	getKitchen() kitchen
}

//устанавливаем тип кухни, которую надо строить
func getBuilder(build string) builder {
	if build == "ok" {
		return &okKitchen{}
	}
	if build == "lux" {
		return &luxKitchen{}
	}
	return nil
}
