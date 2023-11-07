package tenant

import (
	"fmt"

	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/object"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: warrant.NewApiClient(config),
	}
}

func (c Client) Create(params *warrant.TenantParams) (*warrant.Tenant, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeTenant,
		RequestOptions: params.RequestOptions,
	}
	if params.TenantId != "" {
		objectParams.ObjectId = params.TenantId
	}
	if params.Meta != nil {
		objectParams.Meta = params.Meta
	}
	object, err := object.Create(&objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Tenant{
		TenantId: object.ObjectId,
		Meta:     object.Meta,
	}, nil
}

func Create(params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.TenantParams) ([]warrant.Tenant, error) {
	objectsToCreate := make([]warrant.ObjectParams, 0)
	for _, tenantParam := range params {
		objectsToCreate = append(objectsToCreate, warrant.ObjectParams{
			RequestOptions: tenantParam.RequestOptions,
			ObjectType:     warrant.ObjectTypeTenant,
			ObjectId:       tenantParam.TenantId,
			Meta:           tenantParam.Meta,
		})
	}

	createdObjects, err := object.BatchCreate(objectsToCreate)
	if err != nil {
		return nil, err
	}

	tenants := make([]warrant.Tenant, 0)
	for _, createdObject := range createdObjects {
		tenants = append(tenants, warrant.Tenant{
			TenantId: createdObject.ObjectId,
			Meta:     createdObject.Meta,
		})
	}

	return tenants, nil
}

func BatchCreate(params []warrant.TenantParams) ([]warrant.Tenant, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeTenant,
		ObjectId:       tenantId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Get(warrant.ObjectTypeTenant, tenantId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Tenant{
		TenantId: object.ObjectId,
		Meta:     object.Meta,
	}, nil
}

func Get(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Get(tenantId, params)
}

func (c Client) Update(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeTenant,
		ObjectId:       tenantId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Update(warrant.ObjectTypeTenant, tenantId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Tenant{
		TenantId: object.ObjectId,
		Meta:     object.Meta,
	}, nil
}

func Update(tenantId string, params *warrant.TenantParams) (*warrant.Tenant, error) {
	return getClient().Update(tenantId, params)
}

func (c Client) Delete(tenantId string) (string, error) {
	return object.Delete(warrant.ObjectTypeTenant, tenantId)
}

func Delete(tenantId string) (string, error) {
	return getClient().Delete(tenantId)
}

func (c Client) BatchDelete(params []warrant.TenantParams) (string, error) {
	objectsToDelete := make([]warrant.ObjectParams, 0)
	for _, tenantParam := range params {
		objectsToDelete = append(objectsToDelete, warrant.ObjectParams{
			RequestOptions: tenantParam.RequestOptions,
			ObjectType:     warrant.ObjectTypeTenant,
			ObjectId:       tenantParam.TenantId,
			Meta:           tenantParam.Meta,
		})
	}

	warrantToken, err := object.BatchDelete(objectsToDelete)
	if err != nil {
		return "", err
	}

	return warrantToken, nil
}

func BatchDelete(params []warrant.TenantParams) (string, error) {
	return getClient().BatchDelete(params)
}

func (c Client) ListTenants(listParams *warrant.ListTenantParams) (warrant.ListResponse[warrant.Tenant], error) {
	var tenantsListResponse warrant.ListResponse[warrant.Tenant]

	objectsListResponse, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: listParams.ListParams,
		ObjectType: warrant.ObjectTypeTenant,
	})
	if err != nil {
		return tenantsListResponse, err
	}

	tenants := make([]warrant.Tenant, 0)
	for _, object := range objectsListResponse.Results {
		tenants = append(tenants, warrant.Tenant{
			TenantId: object.ObjectId,
			Meta:     object.Meta,
		})
	}

	tenantsListResponse = warrant.ListResponse[warrant.Tenant]{
		Results:    tenants,
		PrevCursor: objectsListResponse.PrevCursor,
		NextCursor: objectsListResponse.NextCursor,
	}

	return tenantsListResponse, nil
}

func ListTenants(listParams *warrant.ListTenantParams) (warrant.ListResponse[warrant.Tenant], error) {
	return getClient().ListTenants(listParams)
}

func (c Client) ListTenantsForUser(userId string, listParams *warrant.ListTenantParams) (warrant.ListResponse[warrant.Tenant], error) {
	var tenantsListResponse warrant.ListResponse[warrant.Tenant]

	queryResponse, err := warrant.Query(fmt.Sprintf("select tenant where user:%s is *", userId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return tenantsListResponse, err
	}

	tenants := make([]warrant.Tenant, 0)
	for _, queryResult := range queryResponse.Results {
		tenants = append(tenants, warrant.Tenant{
			TenantId: queryResult.ObjectId,
			Meta:     queryResult.Meta,
		})
	}

	tenantsListResponse = warrant.ListResponse[warrant.Tenant]{
		Results:    tenants,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return tenantsListResponse, nil
}

func ListTenantsForUser(userId string, listParams *warrant.ListTenantParams) (warrant.ListResponse[warrant.Tenant], error) {
	return getClient().ListTenantsForUser(userId, listParams)
}

func getClient() Client {
	config := warrant.ClientConfig{
		ApiKey:                  warrant.ApiKey,
		ApiEndpoint:             warrant.ApiEndpoint,
		AuthorizeEndpoint:       warrant.AuthorizeEndpoint,
		SelfServiceDashEndpoint: warrant.SelfServiceDashEndpoint,
		HttpClient:              warrant.HttpClient,
	}

	return Client{
		&warrant.ApiClient{
			HttpClient: warrant.HttpClient,
			Config:     config,
		},
	}
}
