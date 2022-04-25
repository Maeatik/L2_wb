package main

import "fmt"

type shop struct {
	next handler
}

func (s *shop) execute(b *book) {
	if b.salesDone {
		fmt.Println("Магазин начал продавать книгу", b.title)
		s.next.execute(b)
		return
	}
	fmt.Println("Магазин закупил книги")
	b.salesDone = true
	s.next.execute(b)
}

func (s *shop) setNext(next handler) {
	s.next = next
}
