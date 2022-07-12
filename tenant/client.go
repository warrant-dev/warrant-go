package tenant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/go-querystring/query"
	warrantdev "github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/user"
	"github.com/warrant-dev/warrant-go/warrant"
)

type Tenant struct {
	TenantId  string    `json:"tenantId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type TenantParams struct {
	TenantId string `json:"tenantId"`
	Name     string `json:"name"`
}

type TenantListParams struct {
	warrantdev.ListParams
}

func New(params *TenantParams) (*Tenant, error) {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("POST", "/v1/tenants", params)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrantdev.WrapError("Error reading response", err)
	}
	var newTenant Tenant
	err = json.Unmarshal([]byte(body), &newTenant)
	if err != nil {
		return nil, warrantdev.WrapError("Invalid response from server", err)
	}
	return &newTenant, nil
}

func Get(id string) (*Tenant, error) {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s", id), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrantdev.WrapError("Error reading response", err)
	}
	var tenant Tenant
	err = json.Unmarshal([]byte(body), &tenant)
	if err != nil {
		return nil, warrantdev.WrapError("Invalid response from server", err)
	}
	return &tenant, nil
}

func Delete(id string) error {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s", id), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}

	return nil
}

func List(listParams *TenantListParams) ([]*Tenant, error) {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	filterQuery, err := query.Values(listParams)
	if err != nil {
		return nil, warrantdev.WrapError("Could not parse filters", err)
	}

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/tenants?%s", filterQuery.Encode()), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrantdev.WrapError("Error reading response", err)
	}
	var tenants []*Tenant
	err = json.Unmarshal([]byte(body), &tenants)
	if err != nil {
		return nil, warrantdev.WrapError("Invalid response from server", err)
	}
	return tenants, nil
}

func Update(id string, params *TenantParams) (*Tenant, error) {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("PUT", fmt.Sprintf("/v1/tenants/%s", id), params)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrantdev.WrapError("Error reading response", err)
	}
	var updatedTenant Tenant
	err = json.Unmarshal([]byte(body), &updatedTenant)
	if err != nil {
		return nil, warrantdev.WrapError("Invalid response from server", err)
	}
	return &updatedTenant, nil
}

func (tenant *Tenant) Update(params *TenantParams) (*Tenant, error) {
	return Update(tenant.TenantId, params)
}

func (tenant *Tenant) ListUsers() ([]*user.User, error) {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s/users", tenant.TenantId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrantdev.WrapError("Error reading response", err)
	}
	var tenantUsers []*user.User
	err = json.Unmarshal([]byte(body), &tenantUsers)
	if err != nil {
		return nil, warrantdev.WrapError("Invalid response from server", err)
	}
	return tenantUsers, nil
}

func (tenant *Tenant) AddUser(userId string) (*warrant.Warrant, error) {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("POST", fmt.Sprintf("/v1/tenants/%s/users/%s", tenant.TenantId, userId), nil)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrantdev.WrapError("Error reading response", err)
	}
	var assignedWarrant warrant.Warrant
	err = json.Unmarshal([]byte(body), &assignedWarrant)
	if err != nil {
		return nil, warrantdev.WrapError("Invalid response from server", err)
	}
	return &assignedWarrant, nil
}

func (tenant *Tenant) RemoveUser(userId string) error {
	client := warrantdev.NewClient(warrantdev.ClientConfig{
		ApiKey: warrantdev.ApiKey,
	})

	resp, err := client.MakeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s/users/%s", tenant.TenantId, userId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return warrantdev.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}

	return nil
}
