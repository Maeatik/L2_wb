package Factory

import "fmt"

type engineBlueCargoPlane struct {
}

func (e *engineBlueCargoPlane) Move() {
	fmt.Println("Летит синий дизельный грузовой самолет")
}
