package warrant

type RequestOptions struct {
	WarrantToken string
}

func (requestOptions *RequestOptions) SetWarrantToken(token string) {
	requestOptions.WarrantToken = token
}
