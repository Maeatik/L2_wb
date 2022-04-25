package main

func main() {
	book := &book{
		title:  "Тестовый автор",
		author: "Вупсень",
	}

	library := &library{}
	shop := &shop{}
	publisherHouse := &publisherHouse{}
	author := &author{}

	shop.setNext(library)
	publisherHouse.setNext(shop)
	author.setNext(publisherHouse)
	author.execute(book)
}
