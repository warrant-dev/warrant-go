package warrant

type RequestOptions struct {
	WarrantToken string `json:"warrantToken,omitempty" url:"warrantToken,omitempty"`
}

func (requestOptions *RequestOptions) SetWarrantToken(token string) {
	requestOptions.WarrantToken = token
}
