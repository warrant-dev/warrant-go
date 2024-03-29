package warrant

const ObjectTypeRole = "role"

type Role struct {
	RoleId string                 `json:"roleId"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}

func (role Role) GetObjectType() string {
	return "role"
}

func (role Role) GetObjectId() string {
	return role.RoleId
}

type ListRoleParams struct {
	ListParams
}

type RoleParams struct {
	RequestOptions
	RoleId string                 `json:"roleId"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}
