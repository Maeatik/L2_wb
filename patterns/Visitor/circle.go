package main

import "fmt"

type circle struct {
	radius int
	x      int
	y      int
}

//подтверждаем фигуру и посещаем
func (c *circle) accept(v visitor) {
	v.visitCircle(c)
}

func (c *circle) move(x int, y int) {
	fmt.Println("x:", x, "y:", y)
}

func (c *circle) getX() int {
	return c.x
}

func (c *circle) getY() int {
	return c.y
}

func (c *circle) getRadius() int {
	return c.radius
}
