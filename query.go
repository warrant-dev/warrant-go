package warrant

type QueryParams struct {
	ListParams
}

type QueryResponse struct {
	Results []QueryResult `json:"results"`
	LastId  *string       `json:"lastId,omitempty"`
}

type QueryResult struct {
	ObjectType string                 `json:"objectType"`
	ObjectId   string                 `json:"objectId"`
	Relation   string                 `json:"relation"`
	Warrant    Warrant                `json:"warrant"`
	IsImplicit bool                   `json:"isImplicit"`
	Meta       map[string]interface{} `json:"meta"`
}
