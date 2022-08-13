package book

import (
	"context"
	"github.com/jackc/pgx/v4"
	"papvan/cvmaker/internal/author/model"
	"papvan/cvmaker/internal/book"
	"papvan/cvmaker/pkg/client/postgresql"
	"papvan/cvmaker/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, book *book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context) ([]book.Book, error) {
	q := `
		SELECT id, name, age
		FROM public.book
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)
	for rows.Next() {
		var bookItem Book
		err := rows.Scan(&bookItem.ID, &bookItem.Name, &bookItem.Age)
		if err != nil {
			return nil, err
		}

		sq := `
			SELECT a.id, a.name
			FROM book_authors ba
			JOIN public.author a ON a.id=ba.author_id
			WHERE book_id=$1
		`
		var authorRows pgx.Rows
		authorRows, err = r.client.Query(ctx, sq, bookItem.ID)
		if err != nil {
			return nil, err
		}

		authors := make([]model.Author, 0)
		for authorRows.Next() {
			var ath model.Author
			err := authorRows.Scan(&ath.ID, &ath.Name)
			if err != nil {
				return nil, err
			}

			authors = append(authors, ath)
		}
		bookItem.Authors = authors

		books = append(books, bookItem.ToDomain())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (book.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, book book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
