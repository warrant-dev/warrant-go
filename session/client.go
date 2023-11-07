package session

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/warrant-dev/warrant-go/v5"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: warrant.NewApiClient(config),
	}
}

func (c Client) CreateAuthorizationSession(params *warrant.AuthorizationSessionParams) (string, error) {
	sessionParams := map[string]interface{}{
		"type":   "sess",
		"userId": params.UserId,
		"ttl":    params.TTL,
	}
	resp, err := c.apiClient.MakeRequest("POST", "/v1/sessions", sessionParams, &warrant.RequestOptions{})
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", warrant.WrapError("Invalid response from server", err)
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

	resp, err := c.apiClient.MakeRequest("POST", "/v1/sessions", sessionParams, &warrant.RequestOptions{})
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", warrant.WrapError("Invalid response from server", err)
	}
	return fmt.Sprintf("%s/%s?redirectUrl=%s", warrant.SelfServiceDashEndpoint, response["token"], params.RedirectUrl), nil
}

func CreateSelfServiceSession(params *warrant.SelfServiceSessionParams) (string, error) {
	return getClient().CreateSelfServiceSession(params)
}

func getClient() Client {
	config := warrant.ClientConfig{
		ApiKey:                  warrant.ApiKey,
		ApiEndpoint:             warrant.ApiEndpoint,
		AuthorizeEndpoint:       warrant.AuthorizeEndpoint,
		SelfServiceDashEndpoint: warrant.SelfServiceDashEndpoint,
		HttpClient:              warrant.HttpClient,
	}

	return Client{
		&warrant.ApiClient{
			HttpClient: warrant.HttpClient,
			Config:     config,
		},
	}
}
