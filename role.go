package warrant

type Role struct {
	RoleId      string `json:"roleId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListRoleParams struct {
	ListParams
}

type RoleParams struct {
	RoleId      string `json:"roleId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
