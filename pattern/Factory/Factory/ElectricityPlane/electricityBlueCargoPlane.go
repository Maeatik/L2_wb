package Factory

import "fmt"

type electricityBlueCargoPlane struct {
}

func (e *electricityBlueCargoPlane) Move() {
	fmt.Println("Летит синий электрический грузовой самолет")
}
