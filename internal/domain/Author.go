package domain

type Author struct {
	ID    uint
	Name  string
	Books []Book
}

func NewAuthor(name string) *Author {
	return &Author{
		Name:  name,
		Books: []Book{},
	}
}

func (a *Author) AddBook(book Book) {
	a.Books = append(a.Books, book)
}
