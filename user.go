package warrant

const ObjectTypeUser = "user"

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
	RequestOptions
	UserId string `json:"userId,omitempty"`
	Email  string `json:"email,omitempty"`
}
