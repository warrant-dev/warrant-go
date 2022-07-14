package warrant

type Tenant struct {
	TenantId string `json:"tenantId"`
	Name     string `json:"name"`
}

type ListTenantParams struct {
	ListParams
}
