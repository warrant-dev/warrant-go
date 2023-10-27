package warrant

import (
	"encoding/json"
)

type Warrant struct {
	ObjectType string  `json:"objectType"`
	ObjectId   string  `json:"objectId"`
	Relation   string  `json:"relation"`
	Subject    Subject `json:"subject"`
	Policy     string  `json:"policy,omitempty"`
	IsImplicit bool    `json:"isImplicit,omitempty"`
	Wookie     string  `json:"wookie,omitempty"`
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

func (subject Subject) GetRelation() string {
	return subject.Relation
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
	ObjectType string                 `json:"objectType"`
	ObjectId   string                 `json:"objectId"`
	Meta       map[string]interface{} `json:"meta"`
}

func (object Object) GetObjectType() string {
	return object.ObjectType
}

func (object Object) GetObjectId() string {
	return object.ObjectId
}

type ObjectParams struct {
	RequestOptions
	ObjectType string                 `json:"objectType"`
	ObjectId   string                 `json:"objectId,omitempty"`
	Meta       map[string]interface{} `json:"meta,omitempty"`
}

type ListObjectParams struct {
	ListParams
	ObjectType string `json:"objectType,omitempty" url:"objectType,omitempty"`
	Query      string `json:"q,omitempty" url:"q,omitempty"`
}

type WarrantObject interface {
	GetObjectType() string
	GetObjectId() string
}

type WarrantCheck struct {
	Object   WarrantObject `json:"object"`
	Relation string        `json:"relation"`
	Subject  WarrantObject `json:"subject"`
	Context  PolicyContext `json:"context,omitempty"`
}

func (warrantCheck WarrantCheck) MarshalJSON() ([]byte, error) {
	var m map[string]interface{}
	subject, ok := warrantCheck.Subject.(*Subject)
	if ok {
		m = map[string]interface{}{
			"objectType": warrantCheck.Object.GetObjectType(),
			"objectId":   warrantCheck.Object.GetObjectId(),
			"relation":   warrantCheck.Relation,
			"subject": map[string]interface{}{
				"objectType": subject.GetObjectType(),
				"objectId":   subject.GetObjectId(),
				"relation":   subject.GetRelation(),
			},
			"context": warrantCheck.Context,
		}
	} else {
		m = map[string]interface{}{
			"objectType": warrantCheck.Object.GetObjectType(),
			"objectId":   warrantCheck.Object.GetObjectId(),
			"relation":   warrantCheck.Relation,
			"subject": map[string]interface{}{
				"objectType": warrantCheck.Subject.GetObjectType(),
				"objectId":   warrantCheck.Subject.GetObjectId(),
			},
			"context": warrantCheck.Context,
		}
	}

	return json.Marshal(m)
}

type WarrantCheckParams struct {
	RequestOptions
	WarrantCheck WarrantCheck `json:"warrantCheck"`
	Debug        bool         `json:"debug,omitempty"`
}

type WarrantCheckManyParams struct {
	RequestOptions
	Op       string         `json:"op"`
	Warrants []WarrantCheck `json:"warrants"`
	Debug    bool           `json:"debug,omitempty"`
}

type WarrantCheckResult struct {
	Code   int64  `json:"code"`
	Result string `json:"result"`
}

type PermissionCheckParams struct {
	RequestOptions
	PermissionId string        `json:"permissionId"`
	UserId       string        `json:"userId"`
	Context      PolicyContext `json:"context,omitempty"`
	Debug        bool          `json:"debug,omitempty"`
}

type RoleCheckParams struct {
	RequestOptions
	RoleId  string        `json:"roleId"`
	UserId  string        `json:"userId"`
	Context PolicyContext `json:"context,omitempty"`
	Debug   bool          `json:"debug,omitempty"`
}

type FeatureCheckParams struct {
	RequestOptions
	FeatureId string        `json:"featureId"`
	Subject   Subject       `json:"subject"`
	Context   PolicyContext `json:"context,omitempty"`
	Debug     bool          `json:"debug,omitempty"`
}

type AccessCheckRequest struct {
	RequestOptions
	Op       string         `json:"op"`
	Warrants []WarrantCheck `json:"warrants"`
	Debug    bool           `json:"debug,omitempty"`
}
