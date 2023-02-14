package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/client"
	"github.com/warrant-dev/warrant-go/config"
)

type Client struct {
	warrantClient *client.WarrantClient
}

func NewClient(config config.ClientConfig) Client {
	return Client{
		warrantClient: &client.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}

func (c Client) CreateAuthorizationSession(params *warrant.AuthorizationSessionParams) (string, error) {
	sessionParams := map[string]interface{}{
		"type":   "sess",
		"userId": params.UserId,
		"ttl":    params.TTL,
	}
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/sessions", sessionParams)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", client.WrapError("Error reading response", err)
	}
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", client.WrapError("Invalid response from server", err)
	}
	return response["token"], nil
}

func CreateAuthorizationSession(params *warrant.AuthorizationSessionParams) (string, error) {
	return getClient().CreateAuthorizationSession(params)
}

func (c Client) CreateSelfServiceSession(params *warrant.SelfServiceSessionParams) (string, error) {
	sessionParams := map[string]interface{}{
		"type":                "ssdash",
		"userId":              params.UserId,
		"tenantId":            params.TenantId,
		"selfServiceStrategy": params.SelfServiceStrategy,
		"ttl":                 params.TTL,
	}
	if params.ObjectType != "" {
		sessionParams["objectType"] = params.ObjectType
	}
	if params.ObjectId != "" {
		sessionParams["objectId"] = params.ObjectId
	}

	resp, err := c.warrantClient.MakeRequest("POST", "/v1/sessions", sessionParams)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", client.WrapError("Error reading response", err)
	}
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", client.WrapError("Invalid response from server", err)
	}
	return fmt.Sprintf("%s/%s?redirectUrl=%s", warrant.SelfServiceDashEndpoint, response["token"], params.RedirectUrl), nil
}

func CreateSelfServiceSession(params *warrant.SelfServiceSessionParams) (string, error) {
	return getClient().CreateSelfServiceSession(params)
}

func getClient() Client {
	if warrant.ApiKey == "" {
		panic("You must provide an ApiKey to initialize the Warrant Client")
	}

	config := config.ClientConfig{
		ApiKey:                  warrant.ApiKey,
		ApiEndpoint:             warrant.ApiEndpoint,
		AuthorizeEndpoint:       warrant.AuthorizeEndpoint,
		SelfServiceDashEndpoint: warrant.SelfServiceDashEndpoint,
	}

	return Client{
		&client.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}
