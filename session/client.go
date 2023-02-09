package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/client"
)

type Client struct {
	warrantClient *client.WarrantClient
}

func (c Client) CreateAuthorizationSession(params *warrant.AuthorizationSessionParams) (string, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/sessions", params)
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
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/sessions", params)
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

	config := client.ClientConfig{
		ApiKey:            warrant.ApiKey,
		AuthorizeEndpoint: warrant.AuthorizeEndpoint,
	}

	return Client{
		&client.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}