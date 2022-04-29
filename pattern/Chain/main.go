package main

func main() {
	//создаем книгу. Называем ее и пишем автора этой книги
	book := &book{
		title:  "Тестовый автор",
		author: "Вупсень",
	}
	//создаем библиотеку, магазин, издательство и автора
	library := &library{}
	shop := &shop{}
	publisherHouse := &publisherHouse{}
	author := &author{}

	//вызывается в обратном порядке. Сначала автор использует книгу, потом издательство использует книгу,
	//магазин использует издательство, а библиотека использует магазин
	shop.setNext(library)
	publisherHouse.setNext(shop)
	author.setNext(publisherHouse)
	author.execute(book)
}
