package usecase

import "fidibo/book/domain"

func (uc *usecase) Insert(book *domain.Book) (*domain.Book, error) {
	return uc.repo.Insert(book)
}
