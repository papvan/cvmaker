package postgresql

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"papvan/cvmaker/internal/author/model"
	"papvan/cvmaker/internal/author/storage"
	"papvan/cvmaker/pkg/client/postgresql"
	"papvan/cvmaker/pkg/logging"
	repeatable "papvan/cvmaker/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(ctx context.Context, author *model.Author) error {
	q := `
		INSERT INTO author (name) 
		VALUES ($1) 
		RETURNING id
	`

	r.logger.Tracef("SQL: %s", repeatable.FormatSQLToLine(q))
	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
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

func (r *repository) FindAll(ctx context.Context, sortOptions storage.SortOptions) ([]model.Author, error) {
	qb := sq.Select("id", "name", "age", "is_alive", "created_at").From("public.author")
	if sortOptions != nil {
		qb = qb.OrderBy(sortOptions.GetOrderBy())
	}
	sql, args, err := qb.ToSql()
	r.logger.Infof("SQL: %s", sql)
	if err != nil {
		return nil, err
	}

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)
	for rows.Next() {
		var ath model.Author
		err := rows.Scan(&ath.ID, &ath.Name, &ath.Age, &ath.IsAlive, &ath.CreatedAt)
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

func (r *repository) FindOne(ctx context.Context, id string) (model.Author, error) {
	q := `
		SELECT id, name
		FROM public.author
		WHERE id = $1
	`

	var ath model.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return model.Author{}, err
	}

	return ath, nil
}

func (r *repository) Update(ctx context.Context, author model.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
