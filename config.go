package warrant

import "net/http"

var ApiKey string
var ApiEndpoint string = "https://api.warrant.dev"
var AuthorizeEndpoint string = "https://api.warrant.dev"
var SelfServiceDashEndpoint string = "https://self-serve.warrant.dev"
var HttpClient *http.Client = http.DefaultClient

type ClientConfig struct {
	ApiKey                  string
	ApiEndpoint             string
	AuthorizeEndpoint       string
	SelfServiceDashEndpoint string
	HttpClient              *http.Client
}
