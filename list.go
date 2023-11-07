package warrant

type ListParams struct {
	RequestOptions
	PrevCursor string `json:"prevCursor,omitempty" url:"prevCursor,omitempty"`
	NextCursor string `json:"nextCursor,omitempty" url:"nextCursor,omitempty"`
	SortBy     string `json:"sortBy,omitempty" url:"sortBy,omitempty"`
	SortOrder  string `json:"sortOrder,omitempty" url:"sortOrder,omitempty"`
	Limit      int    `json:"limit,omitempty" url:"limit,omitempty"`
}

type ListResponse[T any] struct {
	Results    []T    `json:"results,omitempty"`
	PrevCursor string `json:"prevCursor,omitempty"`
	NextCursor string `json:"nextCursor,omitempty"`
}
