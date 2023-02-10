package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
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

func (c Client) Create(params *warrant.UserParams) (*warrant.User, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/users", params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var newUser warrant.User
	err = json.Unmarshal([]byte(body), &newUser)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &newUser, nil
}

func Create(params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.UserParams) ([]warrant.User, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/users", params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var createdUsers []warrant.User
	err = json.Unmarshal([]byte(body), &createdUsers)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return createdUsers, nil
}

func BatchCreate(params []warrant.UserParams) ([]warrant.User, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(userId string) (*warrant.User, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s", userId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var foundUser warrant.User
	err = json.Unmarshal([]byte(body), &foundUser)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &foundUser, nil
}

func Get(userId string) (*warrant.User, error) {
	return getClient().Get(userId)
}

func (c Client) Update(userId string, params *warrant.UserParams) (*warrant.User, error) {
	resp, err := c.warrantClient.MakeRequest("PUT", fmt.Sprintf("/v1/users/%s", userId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var updatedUser warrant.User
	err = json.Unmarshal([]byte(body), &updatedUser)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &updatedUser, nil
}

func Update(userId string, params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Update(userId, params)
}

func (c Client) Delete(userId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s", userId), nil)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	return getClient().Delete(userId)
}

func (c Client) ListUsers(listParams *warrant.ListUserParams) ([]warrant.User, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var users []warrant.User
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return users, nil
}

func ListUsers(listParams *warrant.ListUserParams) ([]warrant.User, error) {
	return getClient().ListUsers(listParams)
}

func (c Client) ListUsersForTenant(tenantId string, listParams *warrant.ListUserParams) ([]warrant.User, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s/users?%s", tenantId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var users []warrant.User
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return users, nil
}

func ListUsersForTenant(tenantId string, listParams *warrant.ListUserParams) ([]warrant.User, error) {
	return getClient().ListUsersForTenant(tenantId, listParams)
}

func (c Client) AssignUserToTenant(userId string, tenantId string) (*warrant.Warrant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/tenants/%s/users/%s", tenantId, userId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var assignedWarrant warrant.Warrant
	err = json.Unmarshal([]byte(body), &assignedWarrant)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &assignedWarrant, nil
}

func AssignUserToTenant(userId string, tenantId string) (*warrant.Warrant, error) {
	return getClient().AssignUserToTenant(tenantId, userId)
}

func (c Client) RemoveUserFromTenant(userId string, tenantId string) error {
	resp, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s/users/%s", tenantId, userId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return client.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func RemoveUserFromTenant(userId string, tenantId string) error {
	return getClient().RemoveUserFromTenant(tenantId, userId)
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
