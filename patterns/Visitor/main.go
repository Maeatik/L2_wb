package main

func main() {
	circle := &circle{
		radius: 5,
		x:      5,
		y:      6,
	}

	dot := &dot{
		id: 1,
		x:  3,
		y:  2,
	}

	rectangle := &rectangle{
		id:     2,
		x:      6,
		y:      7,
		wight:  10,
		height: 15,
	}
	computingShape := &computedShape{}

	circle.accept(computingShape)
	dot.accept(computingShape)
	rectangle.accept(computingShape)

}
