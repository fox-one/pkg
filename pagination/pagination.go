package pagination

type Pagination struct {
	NextCursor string `json:"next_cursor,omitempty"`
	HasNext    bool   `json:"has_next,omitempty"`
}

func (p Pagination) Next() string {
	if p.HasNext {
		return p.NextCursor
	}

	return ""
}
