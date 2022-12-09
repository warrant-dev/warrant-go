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
	ApiKey            string
	AuthorizeEndpoint string
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
	requestUrl := client.buildRequestUrl("/v1/tenants")
	resp, err := client.makeRequest("POST", requestUrl, tenant)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants/%s", tenantId))
	resp, err := client.makeRequest("PUT", requestUrl, tenant)
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

func (client WarrantClient) ListTenants(listParams ListTenantParams) ([]Tenant, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants?%s", queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants/%s", tenantId))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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

func (client WarrantClient) ListUsersForTenant(tenantId string) ([]User, error) {
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants/%s/users", tenantId))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants/%s/users/%s", tenantId, userId))
	resp, err := client.makeRequest("POST", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants/%s/users/%s", tenantId, userId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/tenants/%s", tenantId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl("/v1/users")
	resp, err := client.makeRequest("POST", requestUrl, user)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s", userId))
	resp, err := client.makeRequest("PUT", requestUrl, user)
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

func (client WarrantClient) ListUsers(listParams ListUserParams) ([]User, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users?%s", queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s", userId))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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

func (client WarrantClient) ListTenantsForUser(userId string, listParams ListTenantParams) ([]Tenant, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/tenants?%s", userId, queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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

func (client WarrantClient) ListRolesForUser(userId string, listParams ListRoleParams) ([]Role, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/roles?%s", userId, queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId))
	resp, err := client.makeRequest("POST", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/roles/%s", userId, roleId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/permissions/%s", userId, permissionId))
	resp, err := client.makeRequest("POST", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/permissions/%s", userId, permissionId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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

func (client WarrantClient) ListPermissionsForUser(userId string, listParams ListPermissionParams) ([]Permission, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s/permissions?%s", userId, queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/users/%s", userId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl("/v1/roles")
	resp, err := client.makeRequest("POST", requestUrl, role)
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

func (client WarrantClient) ListRoles(listParams ListRoleParams) ([]Role, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/roles?%s", queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/roles/%s", roleId))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/roles/%s/permissions/%s", roleId, permissionId))
	resp, err := client.makeRequest("POST", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/roles/%s/permissions/%s", roleId, permissionId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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

func (client WarrantClient) ListPermissionsForRole(roleId string, listParams ListPermissionParams) ([]Permission, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/roles/%s/permissions?%s", roleId, queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/roles/%s", roleId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl("/v1/permissions")
	resp, err := client.makeRequest("POST", requestUrl, permission)
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

func (client WarrantClient) ListPermissions(listParams ListPermissionParams) ([]Permission, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/permissions?%s", queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/permissions/%s", permissionId))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/permissions/%s", permissionId))
	resp, err := client.makeRequest("DELETE", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl("/v1/warrants")
	resp, err := client.makeRequest("POST", requestUrl, warrantToCreate)
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

func (client WarrantClient) ListWarrants(listParams ListWarrantParams) ([]Warrant, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, wrapError("Could not parse listParams", err)
	}

	requestUrl := client.buildRequestUrl(fmt.Sprintf("/v1/warrants?%s", queryParams.Encode()))
	resp, err := client.makeRequest("GET", requestUrl, nil)
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
	requestUrl := client.buildRequestUrl("/v1/warrants")
	resp, err := client.makeRequest("DELETE", requestUrl, warrantToDelete)
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

	requestUrl := client.buildRequestUrl("/v1/sessions")
	resp, err := client.makeRequest("POST", requestUrl, requestBody)
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

	requestUrl := client.buildRequestUrl("/v1/sessions")
	resp, err := client.makeRequest("POST", requestUrl, requestBody)
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
	if client.config.AuthorizeEndpoint != "" {
		return client.edgeAuthorize(toCheck)
	}

	return client.authorize(toCheck)
}

func (client WarrantClient) HasPermission(toCheck PermissionCheckParams) (bool, error) {
	return client.IsAuthorized(WarrantCheckParams{
		Warrants: []Warrant{
			{
				ObjectType: "permission",
				ObjectId:   toCheck.PermissionId,
				Relation:   "member",
				Subject: Subject{
					ObjectType: "user",
					ObjectId:   toCheck.UserId,
				},
			},
		},
		ConsistentRead: toCheck.ConsistentRead,
		Debug:          toCheck.Debug,
	})
}

func (client WarrantClient) authorize(toCheck WarrantCheckParams) (bool, error) {
	requestUrl := client.buildRequestUrl("/v2/authorize")
	resp, err := client.makeRequest("POST", requestUrl, toCheck)
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

func (client WarrantClient) edgeAuthorize(toCheck WarrantCheckParams) (bool, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("%s%s", client.config.AuthorizeEndpoint, "/v2/authorize"), toCheck)
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

func (client WarrantClient) makeRequest(method string, url string, payload interface{}) (*http.Response, error) {
	if payload == nil {
		req, err := http.NewRequest(method, url, nil)
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

	postBody, err := json.Marshal(payload)
	if err != nil {
		return nil, wrapError("Invalid request payload", err)
	}
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(method, url, requestBody)
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

func (client WarrantClient) buildRequestUrl(requestUri string) string {
	return fmt.Sprintf("%s%s", API_URL_BASE, requestUri)
}
