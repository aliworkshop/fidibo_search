package usecase

import (
	"encoding/json"
	"fidibo/book/domain"
	"fmt"
	"log"
	"time"
)

func (uc *usecase) Search(keyword string) ([]domain.Book, error) {
	//return uc.repo.Search(keyword)

	books := make([]domain.Book, 0)
	redisKey := fmt.Sprintf("search_data_for_%s", keyword)
	if uc.redis.Exists(redisKey) {
		data, err := uc.redis.Get(redisKey)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &books)
		if err != nil {
			return nil, err
		}
		return books, nil
	}
	res, err := uc.fidibo.Search(keyword)
	if err != nil {
		return nil, err
	}
	for _, item := range res.Books.Hits.Hits {
		books = append(books, domain.Book{
			ImageName:  item.Source.ImageName,
			Publishers: item.Source.Publishers,
			Id:         item.Source.Id,
			Title:      item.Source.Title,
			Content:    item.Source.Content,
			Slug:       item.Source.Slug,
			Authors:    item.Source.Authors,
		})
	}

	if err = uc.redis.Store(redisKey, books, 10*time.Minute); err != nil {
		log.Println("error on storing book results")
	}

	return books, nil
}
