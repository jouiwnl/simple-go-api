package repository

import (
	"context"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
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

func (f *RepositoryFactory) NewUserRepository() *UserRepository {
	return NewUserRepository(*f.db, *f.queries, *f.ctx)
}
