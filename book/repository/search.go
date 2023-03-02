package repository

import "fidibo/book/domain"

func (repo *repo) Search(keyword string) ([]domain.Book, error) {
	books := new([]domain.Book)
	err := repo.db.Where("title like ?", "%"+keyword+"%").
		Preload("Publishers").Preload("Authors").Find(&books).Error
	if err != nil {
		return nil, err
	}

	return *books, nil
}
