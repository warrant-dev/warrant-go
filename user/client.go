package user

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
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

func (c Client) Create(params *warrant.UserParams) (*warrant.User, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v1/users", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var newUser warrant.User
	err = json.Unmarshal([]byte(body), &newUser)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newUser, nil
}

func Create(params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.UserParams) ([]warrant.User, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v1/users", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var createdUsers []warrant.User
	err = json.Unmarshal([]byte(body), &createdUsers)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return createdUsers, nil
}

func BatchCreate(params []warrant.UserParams) ([]warrant.User, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(userId string, params *warrant.UserParams) (*warrant.User, error) {
	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s", userId), nil, &params.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var foundUser warrant.User
	err = json.Unmarshal([]byte(body), &foundUser)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &foundUser, nil
}

func Get(userId string, params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Get(userId, params)
}

func (c Client) Update(userId string, params *warrant.UserParams) (*warrant.User, error) {
	resp, err := c.apiClient.MakeRequest("PUT", fmt.Sprintf("/v1/users/%s", userId), params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var updatedUser warrant.User
	err = json.Unmarshal([]byte(body), &updatedUser)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &updatedUser, nil
}

func Update(userId string, params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Update(userId, params)
}

func (c Client) Delete(userId string) error {
	_, err := c.apiClient.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s", userId), nil, &warrant.RequestOptions{})
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
		return nil, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/users?%s", queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var users []warrant.User
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return users, nil
}

func ListUsers(listParams *warrant.ListUserParams) ([]warrant.User, error) {
	return getClient().ListUsers(listParams)
}

func (c Client) ListUsersForTenant(tenantId string, listParams *warrant.ListUserParams) ([]warrant.User, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s/users?%s", tenantId, queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var users []warrant.User
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return users, nil
}

func ListUsersForTenant(tenantId string, listParams *warrant.ListUserParams) ([]warrant.User, error) {
	return getClient().ListUsersForTenant(tenantId, listParams)
}

func (c Client) AssignUserToTenant(userId string, tenantId string, role string) (*warrant.Warrant, error) {
	return warrant.Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeTenant,
		ObjectId:   tenantId,
		Relation:   role,
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignUserToTenant(userId string, tenantId string, role string) (*warrant.Warrant, error) {
	return getClient().AssignUserToTenant(userId, tenantId, role)
}

func (c Client) RemoveUserFromTenant(userId string, tenantId string, role string) error {
	return warrant.Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeTenant,
		ObjectId:   tenantId,
		Relation:   role,
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemoveUserFromTenant(userId string, tenantId string, role string) error {
	return getClient().RemoveUserFromTenant(userId, tenantId, role)
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
