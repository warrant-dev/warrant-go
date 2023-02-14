package warrant

type User struct {
	UserId string `json:"userId"`
	Email  string `json:"email,omitempty"`
}

func (user User) GetObjectType() string {
	return "user"
}

func (user User) GetObjectId() string {
	return user.UserId
}

type ListUserParams struct {
	ListParams
}

type UserParams struct {
	UserId string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
}
