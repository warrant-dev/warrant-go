package warrant

type Warrant struct {
	ObjectType    string  `json:"objectType"`
	ObjectId      string  `json:"objectId"`
	Relation      string  `json:"relation"`
	Subject       Subject `json:"subject"`
	IsDirectMatch bool    `json:"isDirectMatch,omitempty"`
}

type Subject struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Relation   string `json:"relation,omitempty"`
}

type ListWarrantParams struct {
	ListParams
	ObjectType string `json:"objectType" url:"objectType,omitempty"`
	ObjectId   string `json:"objectId" url:"objectId,omitempty"`
	Relation   string `json:"relation" url:"relation,omitempty"`
	UserId     string `json:"userId" url:"userId,omitempty"`
}

type QueryWarrantParams struct {
	ObjectType string `json:"objectType" url:"objectType,omitempty"`
	Relation   string `json:"relation" url:"relation,omitempty"`
	Subject    string `json:"subject" url:"subject,omitempty"`
}

type WarrantCheckParams struct {
	Op             string    `json:"op"`
	Warrants       []Warrant `json:"warrants"`
	ConsistentRead bool      `json:"consistentRead"`
	Debug          bool      `json:"debug"`
}

type WarrantCheckResult struct {
	Code   int64  `json:"code"`
	Result string `json:"result"`
}

type PermissionCheckParams struct {
	PermissionId   string `json:"permissionId"`
	UserId         string `json:"userId"`
	ConsistentRead bool   `json:"consistentRead"`
	Debug          bool   `json:"debug"`
}
