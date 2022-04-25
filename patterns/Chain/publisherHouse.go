package main

import "fmt"

type publisherHouse struct {
	next handler
}

func (p *publisherHouse) execute(b *book) {
	if b.publishDone {
		fmt.Println("Издательство опубликовало книгу")
		p.next.execute(b)
		return
	}
	fmt.Println("Издательство начало печатать книгу")
	b.publishDone = true
	p.next.execute(b)
}

func (p *publisherHouse) setNext(next handler) {
	p.next = next
}
