package warrant

type Warrant struct {
	ObjectType string         `json:"objectType"`
	ObjectId   string         `json:"objectId"`
	Relation   string         `json:"relation"`
	Subject    WarrantSubject `json:"subject"`
}

type WarrantSubject struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
}

type ListWarrantFilters struct {
	ObjectType string `json:"objectType" url:"objectType,omitempty"`
	ObjectId   string `json:"objectId" url:"objectId,omitempty"`
	Relation   string `json:"relation" url:"relation,omitempty"`
	UserId     string `json:"userId" url:"userId,omitempty"`
}
