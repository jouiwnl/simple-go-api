package page

type Pageable struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func PageableOf(offset int, limit int) Pageable {
	return Pageable{Offset: offset, Limit: limit}
}
