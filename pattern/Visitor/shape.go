package main

type shape interface {
	accept(visitor)
	move(x int, y int)
}
