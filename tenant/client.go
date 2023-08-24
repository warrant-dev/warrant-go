package tenant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v4"
)

type Client struct {
	warrantClient *warrant.WarrantClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		warrantClient: &warrant.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}

func (c Client) Create(params *warrant.TenantParams) (*warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/tenants", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var newTenant warrant.Tenant
	err = json.Unmarshal([]byte(body), &newTenant)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newTenant, nil
}

func Create(params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.TenantParams) ([]warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/tenants", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var createdTenants []warrant.Tenant
	err = json.Unmarshal([]byte(body), &createdTenants)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return createdTenants, nil
}

func BatchCreate(params []warrant.TenantParams) ([]warrant.Tenant, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s", tenantId), nil, &params.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var foundTenant warrant.Tenant
	err = json.Unmarshal([]byte(body), &foundTenant)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &foundTenant, nil
}

func Get(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Get(tenantId, params)
}

func (c Client) Update(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	resp, err := c.warrantClient.MakeRequest("PUT", fmt.Sprintf("/v1/tenants/%s", tenantId), params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var updatedTenant warrant.Tenant
	err = json.Unmarshal([]byte(body), &updatedTenant)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &updatedTenant, nil
}

func Update(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Update(tenantId, params)
}

func (c Client) Delete(tenantId string) error {
	resp, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s", tenantId), nil, &warrant.RequestOptions{})
	if err != nil {
		return err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := io.ReadAll(resp.Body)
		return warrant.Error{
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
		return nil, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants?%s", queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var tenants []warrant.Tenant
	err = json.Unmarshal([]byte(body), &tenants)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return tenants, nil
}

func ListTenants(listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	return getClient().ListTenants(listParams)
}

func (c Client) ListTenantsForUser(userId string, listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/tenants?%s", userId, queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var tenants []warrant.Tenant
	err = json.Unmarshal([]byte(body), &tenants)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return tenants, nil
}

func ListTenantsForUser(userId string, listParams *warrant.ListTenantParams) ([]warrant.Tenant, error) {
	return getClient().ListTenantsForUser(userId, listParams)
}

func getClient() Client {
	config := warrant.ClientConfig{
		ApiKey:                  warrant.ApiKey,
		ApiEndpoint:             warrant.ApiEndpoint,
		AuthorizeEndpoint:       warrant.AuthorizeEndpoint,
		SelfServiceDashEndpoint: warrant.SelfServiceDashEndpoint,
	}

	return Client{
		&warrant.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}
