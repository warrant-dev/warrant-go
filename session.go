package warrant

type Session struct {
	UserId   string `json:"userId"`
	TenantId string `json:"tenantId"`
	TTL      int64  `json:"ttl"`
}

type AuthorizationSessionParams struct {
	UserId string `json:"userId"`
	TTL    int64  `json:"ttl"`
}

type SelfServiceSessionParams struct {
	UserId      string `json:"userId"`
	TenantId    string `json:"tenantId"`
	TTL         int64  `json:"ttl,omitempty"`
	RedirectUrl string `json:"redirectUrl"`
}
