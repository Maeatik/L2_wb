package main

type handler interface {
	execute(*book)
	setNext(handler)
}
