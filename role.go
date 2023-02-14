package warrant

type Role struct {
	RoleId      string `json:"roleId"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
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
	RoleId      string `json:"roleId"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
