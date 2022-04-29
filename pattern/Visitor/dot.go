package main

import "fmt"

type dot struct {
	id int
	x  int
	y  int
}

func (d *dot) move(x int, y int) {
	fmt.Println("x:", x, "y:", y)
}

func (d *dot) draw() {
	fmt.Println("Drawing dot")
}

func (d *dot) accept(v visitor) {
	v.visitDot(d)
}

func (d *dot) getX() int {
	return d.x
}

func (d *dot) getY() int {
	return d.y
}

func (d *dot) getId() int {
	return d.id
}
