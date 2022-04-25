package main

import "fmt"

type computedShape struct {
	area int
}

//посещаем фигуры при их подтверждении из функции мейн
func (c *computedShape) visitDot(d *dot) {

	fmt.Println("Dot #", d.getId(), ": x:", d.getX(), "y:", d.getY())
}

func (c *computedShape) visitCircle(cc *circle) {
	fmt.Println("Circle with radius", cc.getRadius(), ": x:", cc.getX(), "y:", cc.getY())
}

func (c *computedShape) visitRectangle(r *rectangle) {
	fmt.Println("Rectangle #", r.getId(), ": x:", r.getX(), "y:", r.getY())
}
