package models

import (
	"context"

	"gorm.io/gorm"
)

type Book struct {
	Isbn   string  `json:"isbn,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author string  `json:"author,omitempty"`
	Price  float32 `json:"price,omitempty"`
}

// Create a custom BookModel type which wraps the sql.DB connection pool.
type BookModel struct {
	DB *gorm.DB
}

// Use a method on the custom BookModel type to run the SQL query.
func (m BookModel) All(ctx context.Context) ([]Book, error) {
	var bks []Book
	if err := m.DB.Find(&bks).Error; err != nil {
		return nil, err
	}

	return bks, nil
}

func (m BookModel) Create(ctx context.Context, book *Book) error {
	if err := m.DB.Create(book).Error; err != nil {
		return err
	}

	return nil
}
