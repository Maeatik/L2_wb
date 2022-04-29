package Factory

import "fmt"

type electricityRedCargoPlane struct {
}

func (e *electricityRedCargoPlane) move() {
	fmt.Println("Летит красный электрический грузовой самолет")
}
