package main

type okKitchen struct {
	microwave    string
	refrigerator string
	plate        string
}

func newOkKitchen() *okKitchen {
	return &okKitchen{}
}

func (o *okKitchen) setMicrowave() {
	o.microwave = "ok microwave"
}

func (o *okKitchen) setRefrigerator() {
	o.refrigerator = "ok refrigerator"
}

func (o *okKitchen) setPlate() {
	o.plate = "ok plate"
}

func (o *okKitchen) getKitchen() kitchen {
	return kitchen{
		microwave:    o.microwave,
		refrigerator: o.refrigerator,
		plate:        o.plate,
	}
}
