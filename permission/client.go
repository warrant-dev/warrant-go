package permission

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

func (c Client) Create(params *warrant.PermissionParams) (*warrant.Permission, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/permissions", params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var newPermission warrant.Permission
	err = json.Unmarshal([]byte(body), &newPermission)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &newPermission, nil
}

func Create(params *warrant.PermissionParams) (*warrant.Permission, error) {
	return getClient().Create(params)
}

func (c Client) Get(permissionId string) (*warrant.Permission, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/permissions/%s", permissionId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var foundPermission warrant.Permission
	err = json.Unmarshal([]byte(body), &foundPermission)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &foundPermission, nil
}

func Get(permissionId string) (*warrant.Permission, error) {
	return getClient().Get(permissionId)
}

func (c Client) Update(permissionId string, params *warrant.PermissionParams) (*warrant.Permission, error) {
	resp, err := c.warrantClient.MakeRequest("PUT", fmt.Sprintf("/v1/permissions/%s", permissionId), params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var updatedPermission warrant.Permission
	err = json.Unmarshal([]byte(body), &updatedPermission)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &updatedPermission, nil
}

func Update(permissionId string, params *warrant.PermissionParams) (*warrant.Permission, error) {
	return getClient().Update(permissionId, params)
}

func (c Client) Delete(permissionId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/permissions/%s", permissionId), nil)
	if err != nil {
		return err
	}
	return nil
}

func Delete(permissionId string) error {
	return getClient().Delete(permissionId)
}

func (c Client) ListPermissions(listParams *warrant.ListPermissionParams) ([]warrant.Permission, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/permissions?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var permissions []warrant.Permission
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return permissions, nil
}

func ListPermissions(listParams *warrant.ListPermissionParams) ([]warrant.Permission, error) {
	return getClient().ListPermissions(listParams)
}

func (c Client) ListPermissionsForRole(roleId string, listParams *warrant.ListPermissionParams) ([]warrant.Permission, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/roles/%s/permissions?%s", roleId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var permissions []warrant.Permission
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return permissions, nil
}

func ListPermissionsForRole(roleId string, listParams *warrant.ListPermissionParams) ([]warrant.Permission, error) {
	return getClient().ListPermissionsForRole(roleId, listParams)
}

func (c Client) AssignPermissionToRole(permissionId string, roleId string) (*warrant.Permission, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/roles/%s/permissions/%s", roleId, permissionId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var assignedPermission warrant.Permission
	err = json.Unmarshal([]byte(body), &assignedPermission)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &assignedPermission, nil
}

func AssignPermissionToRole(permissionId string, roleId string) (*warrant.Permission, error) {
	return getClient().AssignPermissionToRole(permissionId, roleId)
}

func (c Client) RemovePermissionFromRole(permissionId string, roleId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/roles/%s/permissions/%s", roleId, permissionId), nil)
	if err != nil {
		return err
	}
	return nil
}

func RemovePermissionFromRole(permissionId string, roleId string) error {
	return getClient().RemovePermissionFromRole(permissionId, roleId)
}

func (c Client) ListPermissionsForUser(userId string, listParams *warrant.ListPermissionParams) ([]warrant.Permission, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/permissions?%s", userId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var permissions []warrant.Permission
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return permissions, nil
}

func ListPermissionsForUser(userId string, listParams *warrant.ListPermissionParams) ([]warrant.Permission, error) {
	return getClient().ListPermissionsForUser(userId, listParams)
}

func (c Client) AssignPermissionToUser(permissionId string, userId string) (*warrant.Permission, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/users/%s/permissions/%s", userId, permissionId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var assignedPermission warrant.Permission
	err = json.Unmarshal([]byte(body), &assignedPermission)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &assignedPermission, nil
}

func AssignPermissionToUser(permissionId string, userId string) (*warrant.Permission, error) {
	return getClient().AssignPermissionToUser(permissionId, userId)
}

func (c Client) RemovePermissionFromUser(permissionId string, userId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s/permissions/%s", userId, permissionId), nil)
	if err != nil {
		return err
	}
	return nil
}

func RemovePermissionFromUser(permissionId string, userId string) error {
	return getClient().RemovePermissionFromUser(permissionId, userId)
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
