package warrant

type ClientConfig struct {
	ApiKey            string
	AuthorizeEndpoint string
}

var ApiKey string
var ApiEndpoint = "https://api.warrant.dev"

var AuthorizeEndpoint string
var SelfServiceDashEndpoint = "https://self-serve.warrant.dev"
