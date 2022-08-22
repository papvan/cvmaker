package service

import (
	"context"
	"fmt"
	"papvan/cvmaker/internal/author/model"
	"papvan/cvmaker/internal/author/storage"
	model2 "papvan/cvmaker/internal/author/storage/model"
	"papvan/cvmaker/pkg/api/filter"
	"papvan/cvmaker/pkg/api/sorting"
	"papvan/cvmaker/pkg/logging"
)

type Service struct {
	repository storage.Repository
	logger     *logging.Logger
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{repository: repository, logger: logger}
}

func (s *Service) GetAll(ctx context.Context, filterOptions filter.Options, sortOptions sorting.Options) ([]model.Author, error) {
	options := model2.NewSortOptions(sortOptions.Field, sortOptions.Order)
	all, err := s.repository.FindAll(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("fail to get all authors due to error: %v", err)
	}

	return all, nil
}
