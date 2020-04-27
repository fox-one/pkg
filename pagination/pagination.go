package pagination

type Pagination struct {
	NextCursor string `json:"next_cursor"`
	HasNext    bool   `json:"has_next"`
}

func (p Pagination) Next() string {
	if p.HasNext {
		return p.NextCursor
	}

	return ""
}
