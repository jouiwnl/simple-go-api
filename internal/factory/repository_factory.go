package factory

import (
	"context"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"github.com/jouiwnl/simple-go-api/internal/user/repository"
)

type RepositoryFactory struct {
	db      *sqlc.DBTX
	queries *sqlc.Queries
	ctx     *context.Context
}

func NewRepositoryFactory(db sqlc.DBTX, queries sqlc.Queries, ctx context.Context) *RepositoryFactory {
	return &RepositoryFactory{
		db:      &db,
		queries: &queries,
		ctx:     &ctx,
	}
}

func (f *RepositoryFactory) NewUserRepository() *repository.UserRepository {
	return repository.NewUserRepository(*f.db, *f.queries, *f.ctx)
}
