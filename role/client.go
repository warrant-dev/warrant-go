package role

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/go-querystring/query"
	warrant "github.com/warrant-dev/warrant-go"
)

type Role struct {
	RoleId string `json:"roleId"`
}

type RoleParams struct {
	RoleId string `json:"roleId"`
}

type RoleListParams struct {
	warrant.ListParams
}

func New(params *RoleParams) (*Role, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("POST", "/v1/roles", params)
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
	var newRole Role
	err = json.Unmarshal([]byte(body), &newRole)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newRole, nil
}

func Get(id string) (*Role, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/roles/%s", id), nil)
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
	var role Role
	err = json.Unmarshal([]byte(body), &role)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &role, nil
}

func Delete(id string) error {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/roles/%s", id), nil)
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

func List(listParams *RoleListParams) ([]*Role, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	filterQuery, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse filters", err)
	}

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/roles?%s", filterQuery.Encode()), nil)
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
	var roles []*Role
	err = json.Unmarshal([]byte(body), &roles)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return roles, nil
}
