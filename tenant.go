package warrant

const ObjectTypeTenant = "tenant"

type Tenant struct {
	TenantId string                 `json:"tenantId"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
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
	TenantId string                 `json:"tenantId,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}
