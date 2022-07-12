package warrantdev

type ListParams struct {
	Limit int64 `json:"limit" url:"limit,omitempty"`
	Page  int64 `json:"page" url:"page,omitempty"`
}
