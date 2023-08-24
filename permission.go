package warrant

const ObjectTypePermission = "permission"

type Permission struct {
	PermissionId string `json:"permissionId"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}

func (permission Permission) GetObjectType() string {
	return "permission"
}

func (permission Permission) GetObjectId() string {
	return permission.PermissionId
}

type ListPermissionParams struct {
	ListParams
}

type PermissionParams struct {
	RequestOptions
	PermissionId string `json:"permissionId"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}
