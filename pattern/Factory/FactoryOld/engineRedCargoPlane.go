package Factory

import "fmt"

type engineRedCargoPlane struct {
}

func (e *engineRedCargoPlane) move() {
	fmt.Println("Летит красный дизельный грузовой самолет")
}
