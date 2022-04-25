package Factory

import "fmt"

type engineBlueCargoPlane struct {
}

func (e *engineBlueCargoPlane) move() {
	fmt.Println("Летит синий дизельный грузовой самолет")
}
