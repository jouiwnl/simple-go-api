package commons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPage(t *testing.T) {
	t.Run("should instantiate a new instace of an page", func(t *testing.T) {
		offset := 0
		limit := 20
		total := 20

		pageable := Pageable{
			Offset: offset,
			Limit:  limit,
		}

		page := Page[any]{
			Total:    total,
			Content:  []any{},
			Pageable: &pageable,
		}

		assert.Equal(t, offset, page.Pageable.Offset)
		assert.Equal(t, limit, page.Pageable.Limit)
		assert.Equal(t, total, page.Total)
	})
}
