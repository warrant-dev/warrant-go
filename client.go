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
	resp, err := client.makeRequest("POST", "/tenants", tenant)
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

func (client WarrantClient) DeleteTenant(tenantId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/tenants/%s", tenantId), nil)
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
	resp, err := client.makeRequest("POST", "/users", user)
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

func (client WarrantClient) DeleteUser(userId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/users/%s", userId), nil)
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

func (client WarrantClient) AssignUserToTenant(tenantId string, userId string) (*Warrant, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/tenants/%s/users/%s", tenantId, userId), nil)
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
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/tenants/%s/users/%s", tenantId, userId), nil)
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

func (client WarrantClient) CreateRole(roleId string) (*Role, error) {
	resp, err := client.makeRequest("POST", "/roles", Role{
		RoleId: roleId,
	})
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

func (client WarrantClient) DeleteRole(roleId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/roles/%s", roleId), nil)
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

func (client WarrantClient) CreatePermission(permissionId string) (*Permission, error) {
	resp, err := client.makeRequest("POST", "/permissions", Permission{
		PermissionId: permissionId,
	})
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

func (client WarrantClient) DeletePermission(permissionId string) error {
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/permissions/%s", permissionId), nil)
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

func (client WarrantClient) AssignRoleToUser(userId string, roleId string) (*Role, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/users/%s/roles/%s", userId, roleId), nil)
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
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/users/%s/roles/%s", userId, roleId), nil)
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
	resp, err := client.makeRequest("POST", fmt.Sprintf("/users/%s/permissions/%s", userId, permissionId), nil)
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
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/users/%s/permissions/%s", userId, permissionId), nil)
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

func (client WarrantClient) AssignPermissionToRole(roleId string, permissionId string) (*Permission, error) {
	resp, err := client.makeRequest("POST", fmt.Sprintf("/roles/%s/permissions/%s", roleId, permissionId), nil)
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
	resp, err := client.makeRequest("DELETE", fmt.Sprintf("/roles/%s/permissions/%s", roleId, permissionId), nil)
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
	resp, err := client.makeRequest("POST", "/warrants", warrantToCreate)
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

func (client WarrantClient) ListWarrants(warrantFilters ListWarrantFilters) ([]*Warrant, error) {
	filterQuery, err := query.Values(warrantFilters)
	if err != nil {
		return nil, wrapError("Could not parse filters", err)
	}

	resp, err := client.makeRequest("GET", fmt.Sprintf("/warrants?%s", filterQuery.Encode()), nil)
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

	var warrants []*Warrant
	err = json.Unmarshal([]byte(body), &warrants)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}

	return warrants, nil
}

func (client WarrantClient) CreateAuthorizationSession(session Session) (string, error) {
	requestBody := map[string]string{
		"type":   "sess",
		"userId": session.UserId,
	}
	if session.TenantId != "" {
		requestBody["tenantId"] = session.TenantId
	}

	resp, err := client.makeRequest("POST", "/sessions", requestBody)
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

	resp, err := client.makeRequest("POST", "/sessions", requestBody)
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

func (client WarrantClient) IsAuthorized(toCheck Warrant) (bool, error) {
	resp, err := client.makeRequest("POST", "/authorize", toCheck)
	if err != nil {
		return false, err
	}
	respStatus := resp.StatusCode
	if respStatus == 200 {
		return true, nil
	} else if respStatus == 401 {
		return false, nil
	}
	msg, _ := ioutil.ReadAll(resp.Body)
	return false, Error{
		Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
	}
}

func (client WarrantClient) HasPermission(permissionId string, userId string) (bool, error) {
	return client.IsAuthorized(Warrant{
		ObjectType: "permission",
		ObjectId:   permissionId,
		Relation:   "member",
		Subject: WarrantSubject{
			ObjectType: "user",
			ObjectId:   userId,
		},
	})
}

func (client WarrantClient) makeRequest(method string, requestUri string, payload interface{}) (*http.Response, error) {
	postBody, err := json.Marshal(payload)
	if err != nil {
		return nil, wrapError("Invalid request payload", err)
	}
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(method, API_URL_BASE+API_VERSION+requestUri, requestBody)
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
