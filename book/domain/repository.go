package domain

type BookRepo interface {
	Insert(book *Book) (*Book, error)
	Search(keyword string) ([]Book, error)
}
