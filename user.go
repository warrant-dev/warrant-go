package warrant

type User struct {
	UserId string `json:"userId"`
	Email  string `json:"email,omitempty"`
}

type ListUserParams struct {
	ListParams
}

type UserParams struct {
	UserId string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
}
