package warrant

import "time"

const ObjectTypeTenant = "tenant"

type Tenant struct {
	TenantId  string    `json:"tenantId"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func (tenant Tenant) GetObjectType() string {
	return "tenant"
}

func (tenant Tenant) GetObjectId() string {
	return tenant.TenantId
}

type ListTenantParams struct {
	ListParams
}

type TenantParams struct {
	RequestOptions
	TenantId string `json:"tenantId,omitempty"`
	Name     string `json:"name,omitempty"`
}
