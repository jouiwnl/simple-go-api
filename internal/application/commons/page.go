package commons

type Page[T any] struct {
	Total    int       `json:"total"`
	Content  []T       `json:"content"`
	Pageable *Pageable `json:"pageable"`
}
