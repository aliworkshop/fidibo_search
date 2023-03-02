package domain

import "time"

type Book struct {
	ImageName  string `json:"image_name"`
	Publishers struct {
		Title string `json:"title"`
	} `json:"publishers"`
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
}

type DBBook struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`

	Publishers []Publisher `json:"publishers,omitempty" gorm:"many2many:book_publisher;"`
	Authors    []Author    `json:"authors,omitempty" gorm:"many2many:book_author;"`
}

func (DBBook) TableName() string {
	return "books"
}
