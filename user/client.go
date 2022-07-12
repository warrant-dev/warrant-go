package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/go-querystring/query"
	warrant "github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/permission"
	"github.com/warrant-dev/warrant-go/role"
	"github.com/warrant-dev/warrant-go/tenant"
	warrantClient "github.com/warrant-dev/warrant-go/warrant"
)

type User struct {
	UserId    string    `json:"userId"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserParams struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

type UserListParams struct {
	warrant.ListParams
}

func New(params *UserParams) (*User, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("POST", "/v1/users", params)
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
	var newUser User
	err = json.Unmarshal([]byte(body), &newUser)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newUser, nil
}

func Get(id string) (*User, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/users/%s", id), nil)
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
	var user User
	err = json.Unmarshal([]byte(body), &user)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &user, nil
}

func Delete(id string) error {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s", id), nil)
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

func List(listParams *UserListParams) ([]*User, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	filterQuery, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse filters", err)
	}

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/users?%s", filterQuery.Encode()), nil)
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
	var users []*User
	err = json.Unmarshal([]byte(body), &users)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return users, nil
}

func Update(id string, params *UserParams) (*User, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("PUT", fmt.Sprintf("/v1/users/%s", id), params)
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
	var updatedTenant User
	err = json.Unmarshal([]byte(body), &updatedTenant)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &updatedTenant, nil
}

func (user *User) Update(params *UserParams) (*User, error) {
	return Update(user.UserId, params)
}

func (user *User) ListRoles() ([]*role.Role, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/roles", user.UserId), nil)
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
	var userRoles []*role.Role
	err = json.Unmarshal([]byte(body), &userRoles)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return userRoles, nil
}

func (user *User) AssignRole(roleId string) (*role.Role, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("POST", fmt.Sprintf("/v1/users/%s/roles/%s", user.UserId, roleId), nil)
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
	var assignedRole role.Role
	err = json.Unmarshal([]byte(body), &assignedRole)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &assignedRole, nil
}

func (user *User) RemoveRole(roleId string) error {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s/roles/%s", user.UserId, roleId), nil)
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

func (user *User) ListPermissions() ([]*permission.Permission, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/permissions", user.UserId), nil)
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
	var userPermissions []*permission.Permission
	err = json.Unmarshal([]byte(body), &userPermissions)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return userPermissions, nil
}

func (user *User) AssignPermission(permissionId string) (*permission.Permission, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("POST", fmt.Sprintf("/v1/users/%s/permissions/%s", user.UserId, permissionId), nil)
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
	var assignedPermission permission.Permission
	err = json.Unmarshal([]byte(body), &assignedPermission)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &assignedPermission, nil
}

func (user *User) RemovePermission(permissionId string) error {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s/permissions/%s", user.UserId, permissionId), nil)
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

func (user *User) ListTenants() ([]*tenant.Tenant, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/tenants", user.UserId), nil)
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
	var userTenants []*tenant.Tenant
	err = json.Unmarshal([]byte(body), &userTenants)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return userTenants, nil
}

func (user *User) HasPermission(permissionId string) (bool, error) {
	return warrantClient.IsAuthorized(&warrantClient.WarrantCheckParams{
		Warrants: []*warrantClient.WarrantParams{
			{
				ObjectType: "permission",
				ObjectId:   permissionId,
				Relation:   "member",
				Subject: &warrantClient.SubjectParams{
					ObjectType: "user",
					ObjectId:   user.UserId,
				},
			},
		},
	})
}
