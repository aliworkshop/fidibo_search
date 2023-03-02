package repository

import "fidibo/book/domain"

func (repo *repo) Insert(book *domain.Book) (*domain.Book, error) {
	err := repo.db.Model(book).Create(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}
