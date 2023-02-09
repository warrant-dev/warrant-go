package warrant

type Warrant struct {
	ObjectType    string  `json:"objectType"`
	ObjectId      string  `json:"objectId"`
	Relation      string  `json:"relation"`
	Subject       Subject `json:"subject"`
	Context       Context `json:"context"`
	IsDirectMatch bool    `json:"isDirectMatch,omitempty"`
}

type Subject struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Relation   string `json:"relation,omitempty"`
}

type Context map[string]string

type WarrantParams struct {
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	Relation   string  `json:"relation"`
	Subject    Subject `json:"subject"`
}

type WarrantObject struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
}

// type QueryWarrantParams struct {
// 	ObjectType string  `json:"objectType" url:"objectType,omitempty"`
// 	Relation   string  `json:"relation" url:"relation,omitempty"`
// 	Subject    Subject `json:"subject" url:"subject,omitempty"`
// }

type WarrantCheckParams struct {
	Op             string    `json:"op,omitempty"`
	Warrants       []Warrant `json:"warrants"`
	ConsistentRead bool      `json:"consistentRead,omitempty"`
	Debug          bool      `json:"debug,omitempty"`
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

// func (subject Subject) EncodeValues(key string, v *url.Values) error {
// 	v.Set(key, fmt.Sprintf("%s:%s", subject.ObjectType, subject.ObjectId))

// 	return nil
// }
