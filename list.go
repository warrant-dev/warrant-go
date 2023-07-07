package warrant

type ListParams struct {
	BeforeId    string `json:"beforeId"`
	BeforeValue string `json:"beforeValue"`
	AfterId     string `json:"afterId"`
	AfterValue  string `json:"afterValue"`
	SortBy      string `json:"sortBy"`
	SortOrder   string `json:"sortOrder"`
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
}
