package warrant

type ObjectType struct {
	Type         string                 `json:"type"`
	Relations    map[string]interface{} `json:"relations"`
	WarrantToken string                 `json:"warrantToken,omitempty"`
}

type ListObjectTypeParams struct {
	ListParams
}

type ObjectTypeParams struct {
	RequestOptions
	Type      string                 `json:"type"`
	Relations map[string]interface{} `json:"relations"`
}
