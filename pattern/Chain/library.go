package main

import "fmt"

type library struct {
	next handler
}

func (l *library) execute(b *book) {
	if b.readingDone {
		fmt.Println(b.title, "- ее кто-то прочитал")
	}
	fmt.Println("Книгу купили в библиотеку", b.title)
}

func (l *library) setNext(next handler) {
	l.next = next
}
