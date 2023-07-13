package tenant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v4"
	"github.com/warrant-dev/warrant-go/v4/client"
	"github.com/warrant-dev/warrant-go/v4/config"
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

func (c Client) Create(params *warrant.TenantParams) (*warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/tenants", params)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var newTenant warrant.Tenant
	err = json.Unmarshal([]byte(body), &newTenant)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &newTenant, nil
}

func Create(params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.TenantParams) ([]warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/tenants", params)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var createdTenants []warrant.Tenant
	err = json.Unmarshal([]byte(body), &createdTenants)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return createdTenants, nil
}

func BatchCreate(params []warrant.TenantParams) ([]warrant.Tenant, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(tenantId string) (*warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s", tenantId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var foundTenant warrant.Tenant
	err = json.Unmarshal([]byte(body), &foundTenant)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &foundTenant, nil
}

func Get(tenantId string) (*warrant.Tenant, error) {
	return getClient().Get(tenantId)
}

func (c Client) Update(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("PUT", fmt.Sprintf("/v1/tenants/%s", tenantId), params)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var updatedTenant warrant.Tenant
	err = json.Unmarshal([]byte(body), &updatedTenant)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &updatedTenant, nil
}

func Update(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Update(tenantId, params)
}

func (c Client) Delete(tenantId string) error {
	resp, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s", tenantId), nil)
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := io.ReadAll(resp.Body)
		return client.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	return nil
}

func Delete(tenantId string) error {
	return getClient().Delete(tenantId)
}

func (c Client) ListTenants(listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var tenants []warrant.Tenant
	err = json.Unmarshal([]byte(body), &tenants)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return tenants, nil
}

func ListTenants(listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	return getClient().ListTenants(listParams)
}

func (c Client) ListTenantsForUser(userId string, listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/tenants?%s", userId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var tenants []warrant.Tenant
	err = json.Unmarshal([]byte(body), &tenants)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return tenants, nil
}

func ListTenantsForUser(userId string, listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	return getClient().ListTenantsForUser(userId, listParams)
}

func getClient() Client {
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
