package warrant

type ObjectType struct {
	Type      string                 `json:"type"`
	Relations map[string]interface{} `json:"relations"`
}

type ListObjectTypeParams struct {
	ListParams
}

type ObjectTypeParams struct {
	RequestOptions
	Type      string                 `json:"type"`
	Relations map[string]interface{} `json:"relations"`
}
