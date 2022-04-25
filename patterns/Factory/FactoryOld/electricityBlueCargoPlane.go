package Factory

import "fmt"

type electricityBlueCargoPlane struct {
}

func (e *electricityBlueCargoPlane) move() {
	fmt.Println("Летит синий электрический грузовой самолет")
}
