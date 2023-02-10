package warrant

import "time"

type Tenant struct {
	TenantId  string    `json:"tenantId"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListTenantParams struct {
	ListParams
}

type TenantParams struct {
	TenantId string `json:"tenantId,omitempty"`
	Name     string `json:"name,omitempty"`
}
