package models

import (
	"context"

	"gorm.io/gorm"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
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
