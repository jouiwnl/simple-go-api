package commons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPageable(t *testing.T) {
	t.Run("should instantiate a new instace of an pageable", func(t *testing.T) {
		offset := 0
		limit := 20

		pageable := Pageable{
			Offset: offset,
			Limit:  limit,
		}

		assert.Equal(t, offset, pageable.Offset)
		assert.Equal(t, limit, pageable.Limit)
	})
}
