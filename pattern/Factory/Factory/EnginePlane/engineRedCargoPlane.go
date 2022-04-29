package Factory

import "fmt"

type engineRedCargoPlane struct {
}

func (e *engineRedCargoPlane) Move() {
	fmt.Println("Летит красный дизельный грузовой самолет")
}
