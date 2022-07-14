package warrant

type User struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

type ListUserParams struct {
	ListParams
}
