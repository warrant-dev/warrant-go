package warrant

type Warrant struct {
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	Relation   string  `json:"relation"`
	Subject    Subject `json:"subject"`
	Context    Context `json:"context,omitempty"`
	IsImplicit bool    `json:"isImplicit,omitempty"`
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

type QueryWarrantResult struct {
	Result interface{}            `json:"result"`
	Meta   map[string]interface{} `json:"meta"`
}

type WarrantCheck struct {
	Object   *WarrantObject `json:"object"`
	Relation string         `json:"relation"`
	Subject  *Subject       `json:"subject"`
	Context  Context        `json:"context,omitempty"`
}

func (warrantCheck WarrantCheck) ToWarrant() Warrant {
	return Warrant{
		ObjectType: warrantCheck.Object.ObjectType,
		ObjectId:   warrantCheck.Object.ObjectId,
		Relation:   warrantCheck.Relation,
		Subject:    *warrantCheck.Subject,
		Context:    warrantCheck.Context,
	}
}

type WarrantCheckParams struct {
	Object         *WarrantObject `json:"object"`
	Relation       string         `json:"relation"`
	Subject        *Subject       `json:"subject"`
	Context        Context        `json:"context,omitempty"`
	ConsistentRead bool           `json:"consistentRead,omitempty"`
	Debug          bool           `json:"debug,omitempty"`
}

type WarrantCheckManyParams struct {
	Op             string         `json:"op"`
	Warrants       []WarrantCheck `json:"warrants"`
	ConsistentRead bool           `json:"consistentRead,omitempty"`
	Debug          bool           `json:"debug,omitempty"`
}

type WarrantCheckResult struct {
	Code   int64  `json:"code"`
	Result string `json:"result"`
}

type PermissionCheckParams struct {
	PermissionId   string  `json:"permissionId"`
	UserId         string  `json:"userId"`
	Context        Context `json:"context,omitempty"`
	ConsistentRead bool    `json:"consistentRead,omitempty"`
	Debug          bool    `json:"debug,omitempty"`
}

type RoleCheckParams struct {
	RoleId         string  `json:"roleId"`
	UserId         string  `json:"userId"`
	Context        Context `json:"context,omitempty"`
	ConsistentRead bool    `json:"consistentRead,omitempty"`
	Debug          bool    `json:"debug,omitempty"`
}

type FeatureCheckParams struct {
	FeatureId      string   `json:"featureId"`
	Subject        *Subject `json:"subject"`
	Context        Context  `json:"context,omitempty"`
	ConsistentRead bool     `json:"consistentRead,omitempty"`
	Debug          bool     `json:"debug,omitempty"`
}

type AccessCheckRequest struct {
	Op             string    `json:"op"`
	Warrants       []Warrant `json:"warrants"`
	ConsistentRead bool      `json:"consistentRead,omitempty"`
	Debug          bool      `json:"debug,omitempty"`
}

// func (subject Subject) EncodeValues(key string, v *url.Values) error {
// 	v.Set(key, fmt.Sprintf("%s:%s", subject.ObjectType, subject.ObjectId))

// 	return nil
// }
