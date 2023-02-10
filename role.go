package warrant

type Role struct {
	RoleId      string `json:"roleId"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ListRoleParams struct {
	ListParams
}

type RoleParams struct {
	RoleId      string `json:"roleId"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
