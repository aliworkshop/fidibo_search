package domain

import "time"

type Publisher struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`

	Books []Book `json:"books,omitempty" gorm:"many2many:book_publisher;"`
}

func (Publisher) TableName() string {
	return "publishers"
}
