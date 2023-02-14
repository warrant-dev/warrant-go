package role

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

func (c Client) Create(params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/roles", params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var newRole warrant.Role
	err = json.Unmarshal([]byte(body), &newRole)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &newRole, nil
}

func Create(params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Create(params)
}

func (c Client) Get(roleId string) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/roles/%s", roleId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var foundRole warrant.Role
	err = json.Unmarshal([]byte(body), &foundRole)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &foundRole, nil
}

func Get(roleId string) (*warrant.Role, error) {
	return getClient().Get(roleId)
}

func (c Client) Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("PUT", fmt.Sprintf("/v1/roles/%s", roleId), params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var updatedRole warrant.Role
	err = json.Unmarshal([]byte(body), &updatedRole)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &updatedRole, nil
}

func Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Update(roleId, params)
}

func (c Client) Delete(roleId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/roles/%s", roleId), nil)
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
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/roles?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var roles []warrant.Role
	err = json.Unmarshal([]byte(body), &roles)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return roles, nil
}

func ListRoles(listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	return getClient().ListRoles(listParams)
}

func (c Client) ListRolesForUser(userId string, listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/roles?%s", userId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var roles []warrant.Role
	err = json.Unmarshal([]byte(body), &roles)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return roles, nil
}

func ListRolesForUser(userId string, listParams *warrant.ListRoleParams) ([]warrant.Role, error) {
	return getClient().ListRolesForUser(userId, listParams)
}

func (c Client) AssignRoleToUser(roleId string, userId string) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var assignedRole warrant.Role
	err = json.Unmarshal([]byte(body), &assignedRole)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &assignedRole, nil
}

func AssignRoleToUser(roleId string, userId string) (*warrant.Role, error) {
	return getClient().AssignRoleToUser(roleId, userId)
}

func (c Client) RemoveRoleFromUser(roleId string, userId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId), nil)
	if err != nil {
		return err
	}
	return nil
}

func RemoveRoleFromUser(roleId string, userId string) error {
	return getClient().RemoveRoleFromUser(roleId, userId)
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
