package warrant

type Permission struct {
	PermissionId string `json:"permissionId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

type ListPermissionParams struct {
	ListParams
}

type PermissionParams struct {
	PermissionId string `json:"permissionId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
