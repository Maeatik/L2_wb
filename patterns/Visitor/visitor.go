package main

//посетитель фигур
type visitor interface {
	visitDot(*dot)
	visitCircle(*circle)
	visitRectangle(*rectangle)
}
