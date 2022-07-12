package permission

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/go-querystring/query"
	warrant "github.com/warrant-dev/warrant-go"
)

type Permission struct {
	PermissionId string `json:"permissionId"`
}

type PermissionParams struct {
	PermissionId string `json:"permissionId"`
}

type PermissionListParams struct {
	warrant.ListParams
}

func New(params *PermissionParams) (*Permission, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("POST", "/v1/permissions", params)
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
	var newPermission Permission
	err = json.Unmarshal([]byte(body), &newPermission)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newPermission, nil
}

func Get(id string) (*Permission, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/permissions/%s", id), nil)
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
	var permission Permission
	err = json.Unmarshal([]byte(body), &permission)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &permission, nil
}

func Delete(id string) error {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/permissions/%s", id), nil)
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

func List(listParams *PermissionListParams) ([]*Permission, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	filterQuery, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse filters", err)
	}

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/permissions?%s", filterQuery.Encode()), nil)
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
	var permissions []*Permission
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return permissions, nil
}
