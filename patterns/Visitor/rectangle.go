package main

import "fmt"

type rectangle struct {
	id     int
	x      int
	y      int
	wight  int
	height int
}

func (r *rectangle) accept(v visitor) {
	v.visitRectangle(r)
}

func (r *rectangle) move(x int, y int) {
	fmt.Println("x:", x, "y:", y)
}

func (r *rectangle) getX() int {
	return r.x
}

func (r *rectangle) getY() int {
	return r.y
}

func (r *rectangle) getId() int {
	return r.id
}

func (r *rectangle) getWight() int {
	return r.wight
}

func (r *rectangle) getHeight() int {
	return r.height
}
