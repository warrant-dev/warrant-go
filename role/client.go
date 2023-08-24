package role

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v5"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: &warrant.ApiClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}

func (c Client) Create(params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v1/roles", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var newRole warrant.Role
	err = json.Unmarshal([]byte(body), &newRole)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newRole, nil
}

func Create(params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Create(params)
}

func (c Client) Get(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/roles/%s", roleId), nil, &params.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var foundRole warrant.Role
	err = json.Unmarshal([]byte(body), &foundRole)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &foundRole, nil
}

func Get(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Get(roleId, params)
}

func (c Client) Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.apiClient.MakeRequest("PUT", fmt.Sprintf("/v1/roles/%s", roleId), params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var updatedRole warrant.Role
	err = json.Unmarshal([]byte(body), &updatedRole)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &updatedRole, nil
}

func Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Update(roleId, params)
}

func (c Client) Delete(roleId string) error {
	_, err := c.apiClient.MakeRequest("DELETE", fmt.Sprintf("/v1/roles/%s", roleId), nil, &warrant.RequestOptions{})
	if err != nil {
		return err
	}
	return nil
}

func Delete(roleId string) error {
	return getClient().Delete(roleId)
}

func (c Client) ListRoles(listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/roles?%s", queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var roles []warrant.Role
	err = json.Unmarshal([]byte(body), &roles)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return roles, nil
}

func ListRoles(listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	return getClient().ListRoles(listParams)
}

func (c Client) ListRolesForUser(userId string, listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/roles?%s", userId, queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var roles []warrant.Role
	err = json.Unmarshal([]byte(body), &roles)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return roles, nil
}

func ListRolesForUser(userId string, listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	return getClient().ListRolesForUser(userId, listParams)
}

func (c Client) AssignRoleToUser(roleId string, userId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeRole,
		ObjectId:   roleId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignRoleToUser(roleId string, userId string) (*warrant.Warrant, error) {
	return getClient().AssignRoleToUser(roleId, userId)
}

func (c Client) RemoveRoleFromUser(roleId string, userId string) error {
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeRole,
		ObjectId:   roleId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemoveRoleFromUser(roleId string, userId string) error {
	return getClient().RemoveRoleFromUser(roleId, userId)
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
