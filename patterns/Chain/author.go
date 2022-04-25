package main

import "fmt"

type author struct {
	next handler
}

func (a *author) execute(b *book) {
	if b.writingDone {
		fmt.Println(b.author, "дописал книгу")
		a.next.execute(b)
		return
	}
	fmt.Println(b.author, "начал писать книгу")
	b.writingDone = true
	a.next.execute(b)
}

func (a *author) setNext(next handler) {
	a.next = next
}
