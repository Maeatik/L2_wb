package main

// Чувак Карим-строитель
type dudeBuilder struct {
	dude builder
}

//создаем нового Карима-строителя?
func newDudeBuilder(b builder) *dudeBuilder {
	return &dudeBuilder{
		dude: b,
	}
}

//функция для смены профиля
func (d *dudeBuilder) setDudeBuilder(b builder) {
	d.dude = b
}

//Карим строитель строит кухню
func (d *dudeBuilder) buildKitchen() kitchen {
	d.dude.setMicrowave()
	d.dude.setRefrigerator()
	d.dude.setPlate()
	return d.dude.getKitchen()
}
