package specificationHelper

import (
	"github.com/jouiwnl/go-crud/commons"
)

func RunSpecs[T any](candidate T, specs []commons.Specification[T]) []string {
	var messages []string

	for _, spec := range specs {
		valid, message := spec.IsSatisfiedBy(candidate)

		if valid {
			continue
		}

		messages = append(messages, message)
	}

	return messages
}
