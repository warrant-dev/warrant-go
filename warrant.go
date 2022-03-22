package warrant

type Warrant struct {
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	Relation   string      `json:"relation"`
	User       WarrantUser `json:"user"`
}

type WarrantUser struct {
	UserId string `json:"userId,omitempty"`
	*Userset
}

type Userset struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Relation   string `json:"relation"`
}
