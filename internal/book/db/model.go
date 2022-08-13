package book

import (
	"database/sql"
	"papvan/cvmaker/internal/author/model"
	"papvan/cvmaker/internal/book"
)

// Book made cause of field Age is nullable in postgreSQL, and this model uses for map from db to main Book model
type Book struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Age     sql.NullInt32  `json:"age"`
	Authors []model.Author `json:"authors"`
}

func (m *Book) ToDomain() book.Book {
	b := book.Book{
		ID:      m.ID,
		Name:    m.Name,
		Authors: m.Authors,
	}

	if m.Age.Valid {
		b.Age = int(m.Age.Int32)
	}

	return b
}
