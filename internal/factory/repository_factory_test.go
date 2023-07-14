package factory

import (
	"context"
	"database/sql"
	"github.com/jouiwnl/simple-go-api/internal/config/database/sqlc"
	"testing"
)

func TestRepositoryFactory(t *testing.T) {
	t.Run("should instantiate a new repository factory without errors", func(t *testing.T) {
		fakeDbConnection := &sql.DB{}
		fakeSqlC := *sqlc.New(fakeDbConnection)
		fakeCtx := context.Background()

		expected := NewRepositoryFactory(fakeDbConnection, fakeSqlC, fakeCtx)

		if expected == nil {
			t.Errorf("Expected '%s', but got '%s'", "not null repository", "null repository")
		}
	})
}
