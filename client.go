package warrant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const API_URL_BASE = "https://api.warrant.dev"
const API_VERSION = "/v1"
const SELF_SERVICE_DASH_URL_BASE = "https://self-serve.warrant.dev"

type ClientConfig struct {
	ApiKey string
}

type WarrantClient struct {
	httpClient *http.Client
	config     ClientConfig
}

func NewClient(config ClientConfig) WarrantClient {
	if config.ApiKey == "" {
		panic("You must provide an ApiKey to initialize the Warrant Client")
	}

	return WarrantClient{
		httpClient: http.DefaultClient,
		config:     config,
	}
}

func (client WarrantClient) CreateTenant(tenant Tenant) (*Tenant, error) {
	resp, err := client.makeRequest("POST", "/v1/tenants", tenant)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newTenant Tenant
	err = json.Unmarshal([]byte(body), &newTenant)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newTenant, nil
}

func (client WarrantClient) UpdateTenant(tenantId string, tenant Tenant) (*Tenant, error) {
	resp, err := client.makeRequest("PUT", fmt.Sprintf("/v1/tenants/%s", tenantId), tenant)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var updatedTenant Tenant
	err = json.Unmarshal([]byte(body), &updatedTenant)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &updatedTenant, nil
}

func (client WarrantClient) ListTenants() ([]Tenant, error) {
	resp, err := client.makeRequest("GET", "/v1/tenants", nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var tenants []Tenant
	err = json.Unmarshal([]byte(body), &tenants)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return tenants, nil
}

func (client WarrantClient) GetTenant(tenantId string) (*Tenant, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/tenants/%s", tenantId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var foundTenant Tenant
	err = json.Unmarshal([]byte(body), &foundTenant)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &foundTenant, nil
}

func (client WarrantClient) GetUsersForTenant(tenantId string) ([]User, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/tenants/%s/users", tenantId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var tenantUsers []User
	err = json.Unmarshal([]byte(body), &tenantUsers)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return tenantUsers, nil
}

func (client WarrantClient) AssignUserToTenant(tenantId string, userId string) (*Warrant, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/v1/tenants/%s/users/%s", tenantId, userId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newWarrant Warrant
	err = json.Unmarshal([]byte(body), &newWarrant)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newWarrant, nil
}

func (client WarrantClient) RemoveUserFromTenant(tenantId string, userId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s/users/%s", tenantId, userId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) DeleteTenant(tenantId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s", tenantId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) CreateUser(user User) (*User, error) {
	resp, err := client.makeRequest("POST", "/v1/users", user)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newUser User
	err = json.Unmarshal([]byte(body), &newUser)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newUser, nil
}

func (client WarrantClient) UpdateUser(userId string, user User) (*User, error) {
	resp, err := client.makeRequest("PUT", fmt.Sprintf("/v1/users/%s", userId), user)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var updatedUser User
	err = json.Unmarshal([]byte(body), &updatedUser)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &updatedUser, nil
}

func (client WarrantClient) ListUsers() ([]User, error) {
	resp, err := client.makeRequest("GET", "/v1/users", nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var users []User
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return users, nil
}

func (client WarrantClient) GetUser(userId string) (*User, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/users/%s", userId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var foundUser User
	err = json.Unmarshal([]byte(body), &foundUser)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &foundUser, nil
}

func (client WarrantClient) GetTenantsForUser(userId string) ([]Tenant, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/users/%s/tenants", userId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var userTenants []Tenant
	err = json.Unmarshal([]byte(body), &userTenants)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return userTenants, nil
}

func (client WarrantClient) GetRolesForUser(userId string) ([]Role, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/users/%s/roles", userId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var userRoles []Role
	err = json.Unmarshal([]byte(body), &userRoles)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return userRoles, nil
}

func (client WarrantClient) AssignRoleToUser(userId string, roleId string) (*Role, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newRole Role
	err = json.Unmarshal([]byte(body), &newRole)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newRole, nil
}

func (client WarrantClient) RemoveRoleFromUser(userId string, roleId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) AssignPermissionToUser(userId string, permissionId string) (*Permission, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/v1/users/%s/permissions/%s", userId, permissionId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newPermission Permission
	err = json.Unmarshal([]byte(body), &newPermission)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newPermission, nil
}

func (client WarrantClient) RemovePermissionFromUser(userId string, permissionId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/users/%s/permissions/%s", userId, permissionId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) GetPermissionsForUser(userId string) ([]Permission, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/users/%s/permissions", userId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var userPermissions []Permission
	err = json.Unmarshal([]byte(body), &userPermissions)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return userPermissions, nil
}

func (client WarrantClient) DeleteUser(userId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/users/%s", userId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) CreateRole(role Role) (*Role, error) {
	resp, err := client.makeRequest("POST", "/v1/roles", role)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newRole Role
	err = json.Unmarshal([]byte(body), &newRole)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newRole, nil
}

func (client WarrantClient) ListRoles() ([]Role, error) {
	resp, err := client.makeRequest("GET", "/v1/roles", nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var roles []Role
	err = json.Unmarshal([]byte(body), &roles)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return roles, nil
}

func (client WarrantClient) GetRole(roleId string) (*Role, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/roles/%s", roleId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var foundRole Role
	err = json.Unmarshal([]byte(body), &foundRole)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &foundRole, nil
}

func (client WarrantClient) AssignPermissionToRole(roleId string, permissionId string) (*Permission, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/v1/roles/%s/permissions/%s", roleId, permissionId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newPermission Permission
	err = json.Unmarshal([]byte(body), &newPermission)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newPermission, nil
}

func (client WarrantClient) RemovePermissionFromRole(roleId string, permissionId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/roles/%s/permissions/%s", roleId, permissionId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) GetPermissionsForRole(roleId string) ([]Permission, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/roles/%s/permissions", roleId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var rolePermissions []Permission
	err = json.Unmarshal([]byte(body), &rolePermissions)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return rolePermissions, nil
}

func (client WarrantClient) DeleteRole(roleId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/roles/%s", roleId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) CreatePermission(permission Permission) (*Permission, error) {
	resp, err := client.makeRequest("POST", "/v1/permissions", permission)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newPermission Permission
	err = json.Unmarshal([]byte(body), &newPermission)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newPermission, nil
}

func (client WarrantClient) ListPermissions() ([]Permission, error) {
	resp, err := client.makeRequest("GET", "/v1/permissions", nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var permissions []Permission
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return permissions, nil
}

func (client WarrantClient) GetPermission(permissionId string) (*Permission, error) {
	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/permissions/%s", permissionId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var foundPermission Permission
	err = json.Unmarshal([]byte(body), &foundPermission)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &foundPermission, nil
}

func (client WarrantClient) DeletePermission(permissionId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/v1/permissions/%s", permissionId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) CreateWarrant(warrantToCreate Warrant) (*Warrant, error) {
	resp, err := client.makeRequest("POST", "/v1/warrants", warrantToCreate)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newWarrant Warrant
	err = json.Unmarshal([]byte(body), &newWarrant)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newWarrant, nil
}

func (client WarrantClient) ListWarrants(warrantFilters ListWarrantFilters) ([]Warrant, error) {
	filterQuery, err := query.Values(warrantFilters)
	if err != nil {
		return nil, wrapError("Could not parse filters", err)
	}

	resp, err := client.makeRequest("GET", fmt.Sprintf("/v1/warrants?%s", filterQuery.Encode()), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}

	var warrants []Warrant
	err = json.Unmarshal([]byte(body), &warrants)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}

	return warrants, nil
}

func (client WarrantClient) DeleteWarrant(warrantToDelete Warrant) error {
	resp, err := client.makeRequest("DELETE", "/v1/warrants", warrantToDelete)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func (client WarrantClient) CreateAuthorizationSession(session Session) (string, error) {
	requestBody := map[string]string{
		"type":   "sess",
		"userId": session.UserId,
	}

	resp, err := client.makeRequest("POST", "/v1/sessions", requestBody)
	if err != nil {
		return "", err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return "", Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", wrapError("Error reading response", err)
	}
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", wrapError("Invalid response from server", err)
	}
	return response["token"], nil
}

func (client WarrantClient) CreateSelfServiceSession(session Session, redirectUrl string) (string, error) {
	requestBody := map[string]string{
		"type":   "ssdash",
		"userId": session.UserId,
	}
	if session.TenantId != "" {
		requestBody["tenantId"] = session.TenantId
	}

	resp, err := client.makeRequest("POST", "/v1/sessions", requestBody)
	if err != nil {
		return "", err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return "", Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", wrapError("Error reading response", err)
	}
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", wrapError("Invalid response from server", err)
	}
	return fmt.Sprintf("%s/%s?redirectUrl=%s", SELF_SERVICE_DASH_URL_BASE, response["token"], redirectUrl), nil
}

func (client WarrantClient) IsAuthorized(toCheck WarrantCheckParams) (bool, error) {
	resp, err := client.makeRequest("POST", "/v2/authorize", toCheck)
	if err != nil {
		return false, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return false, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, wrapError("Error reading response", err)
	}

	var result WarrantCheckResult
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return false, wrapError("Invalid response from server", err)
	}

	if result.Result == "Authorized" {
		return true, nil
	} else {
		return false, nil
	}
}

func (client WarrantClient) HasPermission(permissionId string, userId string) (bool, error) {
	return client.IsAuthorized(WarrantCheckParams{
		Warrants: []Warrant{
			{
				ObjectType: "permission",
				ObjectId:   permissionId,
				Relation:   "member",
				Subject: Subject{
					ObjectType: "user",
					ObjectId:   userId,
				},
			},
		},
	})
}

func (client WarrantClient) makeRequest(method string, requestUri string, payload interface{}) (*http.Response, error) {
	postBody, err := json.Marshal(payload)
	if err != nil {
		return nil, wrapError("Invalid request payload", err)
	}
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(method, API_URL_BASE+requestUri, requestBody)
	if err != nil {
		return nil, wrapError("Unable to create request", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", client.config.ApiKey))
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, wrapError("Error making request", err)
	}
	return resp, nil
}
