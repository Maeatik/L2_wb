package Factory

import "fmt"

type electricityRedCargoPlane struct {
}

func (e *electricityRedCargoPlane) Move() {
	fmt.Println("Летит красный электрический грузовой самолет")
}
