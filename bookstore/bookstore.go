package bookstore

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b1 Book) Book {
	return Book{
		Title:  b1.Title,
		Author: b1.Author,
		Copies: b1.Copies - 1,
	}
}
