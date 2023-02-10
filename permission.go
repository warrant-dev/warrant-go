package warrant

type Permission struct {
	PermissionId string `json:"permissionId"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}

type ListPermissionParams struct {
	ListParams
}

type PermissionParams struct {
	PermissionId string `json:"permissionId"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}
