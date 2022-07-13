package warrant

type Session struct {
	UserId   string `json:"userId"`
	TenantId string `json:"tenantId"`
	TTL      int64  `json:"ttl"`
}
