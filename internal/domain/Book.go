package domain

type Book struct {
	ID       uint
	Title    string
	AuthorID uint
}

func NewBook(title string, authorID uint) *Book {
	return &Book{
		Title:    title,
		AuthorID: authorID,
	}
}
