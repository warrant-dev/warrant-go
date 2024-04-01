package warrant

const (
	SelfServiceStrategyFGAC = "fgac"
	SelfServiceStrategyRBAC = "rbac"
)

type Session struct {
	UserId   string `json:"userId"`
	TenantId string `json:"tenantId"`
	TTL      int64  `json:"ttl"`
}

type AuthorizationSessionParams struct {
	UserId   string        `json:"userId,omitempty"`
	TenantId string        `json:"tenantId,omitempty"`
	TTL      int64         `json:"ttl,omitempty"`
	Context  PolicyContext `json:"context,omitempty"`
}

type SelfServiceSessionParams struct {
	UserId              string        `json:"userId"`
	TenantId            string        `json:"tenantId"`
	TTL                 int64         `json:"ttl,omitempty"`
	Context             PolicyContext `json:"context,omitempty"`
	SelfServiceStrategy string        `json:"selfServiceStrategy"`
	ObjectType          string        `json:"objectType"`
	ObjectId            string        `json:"objectId"`
	RedirectUrl         string        `json:"redirectUrl"`
}
