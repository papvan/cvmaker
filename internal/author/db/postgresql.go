package author

import (
	"context"
	"github.com/jackc/pgconn"
	"papvan/cvmaker/internal/author"
	"papvan/cvmaker/pkg/client/postgresql"
	"papvan/cvmaker/pkg/logging"
	repeatable "papvan/cvmaker/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, author *author.Author) error {
	q := `
		INSERT INTO author (name) 
		VALUES ($1) 
		RETURNING id
	`

	r.logger.Tracef("SQL: %s", repeatable.FormatSQLToLine(q))
	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQL State: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState(),
			)
			return nil
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) ([]author.Author, error) {
	q := `
		SELECT id, name
		FROM public.author
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)
	for rows.Next() {
		var ath author.Author
		err := rows.Scan(&ath.ID, &ath.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `
		SELECT id, name
		FROM public.author
		WHERE id = $1
	`

	var ath author.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return author.Author{}, err
	}

	return ath, nil
}

func (r *repository) Update(ctx context.Context, author author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
