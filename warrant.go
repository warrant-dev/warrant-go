package warrant

type Warrant struct {
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	Relation   string  `json:"relation"`
	Subject    Subject `json:"subject"`
	Context    Context `json:"context,omitempty"`
	IsImplicit bool    `json:"isImplicit,omitempty"`
}

type Context map[string]string

type Subject struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Relation   string `json:"relation,omitempty"`
}

func (subject Subject) GetObjectType() string {
	return subject.ObjectType
}

func (subject Subject) GetObjectId() string {
	return subject.ObjectId
}

type WarrantParams struct {
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	Relation   string  `json:"relation"`
	Subject    Subject `json:"subject"`
	Context    Context `json:"context,omitempty"`
}

type ListWarrantParams struct {
	ListParams
}

type Object struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
}

func (object Object) GetObjectType() string {
	return object.ObjectType
}

func (object Object) GetObjectId() string {
	return object.ObjectId
}

type WarrantObject interface {
	GetObjectType() string
	GetObjectId() string
}

type QueryWarrantResult struct {
	Result interface{}            `json:"result"`
	Meta   map[string]interface{} `json:"meta"`
}

type WarrantCheck struct {
	Object   WarrantObject `json:"object"`
	Relation string        `json:"relation"`
	Subject  WarrantObject `json:"subject"`
	Context  Context       `json:"context,omitempty"`
}

func (warrantCheck WarrantCheck) ToWarrant() Warrant {
	subject, ok := warrantCheck.Subject.(Subject)
	if ok {
		return Warrant{
			ObjectType: warrantCheck.Object.GetObjectType(),
			ObjectId:   warrantCheck.Object.GetObjectId(),
			Relation:   warrantCheck.Relation,
			Subject:    subject,
			Context:    warrantCheck.Context,
		}
	}

	return Warrant{
		ObjectType: warrantCheck.Object.GetObjectType(),
		ObjectId:   warrantCheck.Object.GetObjectId(),
		Relation:   warrantCheck.Relation,
		Subject: Subject{
			ObjectType: warrantCheck.Subject.GetObjectType(),
			ObjectId:   warrantCheck.Subject.GetObjectId(),
		},
		Context: warrantCheck.Context,
	}
}

type WarrantCheckParams struct {
	WarrantCheck   WarrantCheck `json:"warrantCheck"`
	ConsistentRead bool         `json:"consistentRead,omitempty"`
	Debug          bool         `json:"debug,omitempty"`
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
	FeatureId      string  `json:"featureId"`
	Subject        Subject `json:"subject"`
	Context        Context `json:"context,omitempty"`
	ConsistentRead bool    `json:"consistentRead,omitempty"`
	Debug          bool    `json:"debug,omitempty"`
}

type AccessCheckRequest struct {
	Op             string    `json:"op"`
	Warrants       []Warrant `json:"warrants"`
	ConsistentRead bool      `json:"consistentRead,omitempty"`
	Debug          bool      `json:"debug,omitempty"`
}
