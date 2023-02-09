package role

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go"
)

type Client struct {
	warrantClient *warrant.WarrantClient
}

func (c Client) Create(params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/roles", params)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
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

func (c Client) Get(roleId string) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/roles/%s", roleId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
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

func Get(roleId string) (*warrant.Role, error) {
	return getClient().Get(roleId)
}

func (c Client) Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("PUT", fmt.Sprintf("/v1/roles/%s", roleId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
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
	resp, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/roles/%s", roleId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
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

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/roles?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
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

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/roles?%s", userId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
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

func (c Client) AssignRoleForUser(roleId string, userId string) (*warrant.Role, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var assignedRole warrant.Role
	err = json.Unmarshal([]byte(body), &assignedRole)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &assignedRole, nil
}

func AssignRoleForUser(roleId string, userId string) (*warrant.Role, error) {
	return getClient().AssignRoleForUser(userId, roleId)
}

func (c Client) RemoveRoleForUser(roleId string, userId string) error {
	resp, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func RemoveRoleForUser(roleId string, userId string) error {
	return getClient().RemoveRoleForUser(userId, roleId)
}

func getClient() Client {
	if warrant.ApiKey == "" {
		panic("You must provide an ApiKey to initialize the Warrant Client")
	}

	config := warrant.ClientConfig{
		ApiKey:            warrant.ApiKey,
		AuthorizeEndpoint: warrant.AuthorizeEndpoint,
	}

	return Client{
		&warrant.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}
