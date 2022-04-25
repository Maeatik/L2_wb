package main

type luxKitchen struct {
	microwave    string
	refrigerator string
	plate        string
}

func newLuxKitchen() *luxKitchen {
	return &luxKitchen{}
}

func (l *luxKitchen) setMicrowave() {
	l.microwave = "wow microwave"
}

func (l *luxKitchen) setRefrigerator() {
	l.refrigerator = "wow refrigerator"
}

func (l *luxKitchen) setPlate() {
	l.plate = "wow plate"
}

func (l *luxKitchen) getKitchen() kitchen {
	return kitchen{
		microwave:    l.microwave,
		refrigerator: l.refrigerator,
		plate:        l.plate,
	}
}
