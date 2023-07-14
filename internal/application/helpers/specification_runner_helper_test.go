package specificationHelper

import (
	"github.com/jouiwnl/simple-go-api/internal/application/commons"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpecificationRunner(t *testing.T) {
	t.Run("should run all specs from struct without errors", func(t *testing.T) {
		var specs []commons.Specification[any]
		var anyObject any

		errors := RunSpecs(&anyObject, &specs)

		assert.Equal(t, len(errors), 0)
	})

	t.Run("should run all specs from struct with errors", func(t *testing.T) {
		expectedError := "error"
		var anyObject any
		specs := []commons.Specification[any]{
			&MockSpec{
				candidate: &anyObject,
			},
		}

		errors := RunSpecs(&anyObject, &specs)

		assert.Equal(t, len(errors), 1)
		assert.Equal(t, errors[0], expectedError)
	})
}

type MockSpec struct {
	candidate *any
}

func (m *MockSpec) IsSatisfiedBy(candidate any) (bool, string) {
	return false, "error"
}
