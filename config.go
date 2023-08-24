package warrant

var ApiKey string
var ApiEndpoint string = "https://api.warrant.dev"
var AuthorizeEndpoint string = "https://api.warrant.dev"
var SelfServiceDashEndpoint string = "https://self-serve.warrant.dev"

type ClientConfig struct {
	ApiKey                  string
	ApiEndpoint             string
	AuthorizeEndpoint       string
	SelfServiceDashEndpoint string
}
