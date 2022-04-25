package main

type visitor interface {
	visitDot(*dot)
	visitCircle(*circle)
	visitRectangle(*rectangle)
}
