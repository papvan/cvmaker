package author

import (
	"context"
	"papvan/cvmaker/pkg/logging"
)

type Service struct {
	repository Repository
	logger     *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateAuthorDTO) (u Author, err error) {
	return
}
