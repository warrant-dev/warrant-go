package warrant

const ObjectTypeUser = "user"

type User struct {
	UserId string                 `json:"userId"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
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
	UserId string                 `json:"userId,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}
