package commons

type Specification[T any] interface {
	IsSatisfiedBy(entity T) (bool, string)
}
