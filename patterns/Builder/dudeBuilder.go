package main

type dudeBuilder struct {
	dude builder
}

func newDudeBuilder(b builder) *dudeBuilder {
	return &dudeBuilder{
		dude: b,
	}
}

func (d *dudeBuilder) setDudeBuilder(b builder) {
	d.dude = b
}

func (d *dudeBuilder) buildKitchen() kitchen {
	d.dude.setMicrowave()
	d.dude.setRefrigerator()
	d.dude.setPlate()
	return d.dude.getKitchen()
}
