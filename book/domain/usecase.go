package domain

type BookUc interface {
	Insert(book *Book) (*Book, error)
	Search(keyword string) ([]Book, error)
}
