package warrant

import "encoding/json"

type Warrant struct {
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	Relation   string  `json:"relation"`
	Subject    Subject `json:"subject"`
	Policy     string  `json:"policy,omitempty"`
	IsImplicit bool    `json:"isImplicit,omitempty"`
}

type PolicyContext map[string]interface{}

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
	Policy     string  `json:"policy,omitempty"`
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
	Context  PolicyContext `json:"context,omitempty"`
}

func (warrantCheck WarrantCheck) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"objectType": warrantCheck.Object.GetObjectType(),
		"objectId":   warrantCheck.Object.GetObjectId(),
		"relation":   warrantCheck.Relation,
		"subject":    warrantCheck.Subject,
		"context":    warrantCheck.Context,
	}

	return json.Marshal(m)
}

type WarrantCheckParams struct {
	WarrantCheck WarrantCheck `json:"warrantCheck"`
	Debug        bool         `json:"debug,omitempty"`
}

type WarrantCheckManyParams struct {
	Op       string         `json:"op"`
	Warrants []WarrantCheck `json:"warrants"`
	Debug    bool           `json:"debug,omitempty"`
}

type WarrantCheckResult struct {
	Code   int64  `json:"code"`
	Result string `json:"result"`
}

type PermissionCheckParams struct {
	PermissionId string        `json:"permissionId"`
	UserId       string        `json:"userId"`
	Context      PolicyContext `json:"context,omitempty"`
	Debug        bool          `json:"debug,omitempty"`
}

type RoleCheckParams struct {
	RoleId  string        `json:"roleId"`
	UserId  string        `json:"userId"`
	Context PolicyContext `json:"context,omitempty"`
	Debug   bool          `json:"debug,omitempty"`
}

type FeatureCheckParams struct {
	FeatureId string        `json:"featureId"`
	Subject   Subject       `json:"subject"`
	Context   PolicyContext `json:"context,omitempty"`
	Debug     bool          `json:"debug,omitempty"`
}

type AccessCheckRequest struct {
	Op       string         `json:"op"`
	Warrants []WarrantCheck `json:"warrants"`
	Debug    bool           `json:"debug,omitempty"`
}
