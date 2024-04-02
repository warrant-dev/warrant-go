package objecttype

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v6"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: warrant.NewApiClient(config),
	}
}

func (c Client) Create(params *warrant.ObjectTypeParams) (*warrant.ObjectType, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v2/object-types", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var newObjectType warrant.ObjectType
	err = json.Unmarshal([]byte(body), &newObjectType)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	warrantToken := resp.Header.Get("Warrant-Token")
	newObjectType.WarrantToken = warrantToken
	return &newObjectType, nil
}

func Create(params *warrant.ObjectTypeParams) (*warrant.ObjectType, error) {
	return getClient().Create(params)
}

func (c Client) Get(objectTypeId string, params *warrant.ObjectTypeParams) (*warrant.ObjectType, error) {
	if params == nil {
		params = &warrant.ObjectTypeParams{}
	}
	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v2/object-types/%s", objectTypeId), nil, &params.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var foundObjectType warrant.ObjectType
	err = json.Unmarshal([]byte(body), &foundObjectType)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &foundObjectType, nil
}

func Get(objectTypeId string, params *warrant.ObjectTypeParams) (*warrant.ObjectType, error) {
	return getClient().Get(objectTypeId, params)
}

func (c Client) Update(objectTypeId string, params *warrant.ObjectTypeParams) (*warrant.ObjectType, error) {
	resp, err := c.apiClient.MakeRequest("PUT", fmt.Sprintf("/v2/object-types/%s", objectTypeId), params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var updatedObjectType warrant.ObjectType
	err = json.Unmarshal([]byte(body), &updatedObjectType)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	warrantToken := resp.Header.Get("Warrant-Token")
	updatedObjectType.WarrantToken = warrantToken
	return &updatedObjectType, nil
}

func Update(objectTypeId string, params *warrant.ObjectTypeParams) (*warrant.ObjectType, error) {
	return getClient().Update(objectTypeId, params)
}

func (c Client) BatchUpdate(params []warrant.ObjectTypeParams) ([]warrant.ObjectType, error) {
	resp, err := c.apiClient.MakeRequest("PUT", "/v2/object-types", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var updatedObjectTypes []warrant.ObjectType
	err = json.Unmarshal([]byte(body), &updatedObjectTypes)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	warrantToken := resp.Header.Get("Warrant-Token")
	for i := range updatedObjectTypes {
		updatedObjectTypes[i].WarrantToken = warrantToken
	}
	return updatedObjectTypes, nil
}

func BatchUpdate(params []warrant.ObjectTypeParams) ([]warrant.ObjectType, error) {
	return getClient().BatchUpdate(params)
}

func (c Client) Delete(objectTypeId string) (string, error) {
	resp, err := c.apiClient.MakeRequest("DELETE", fmt.Sprintf("/v2/object-types/%s", objectTypeId), nil, &warrant.RequestOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	warrantToken := resp.Header.Get("Warrant-Token")
	return warrantToken, nil
}

func Delete(objectTypeId string) (string, error) {
	return getClient().Delete(objectTypeId)
}

func (c Client) ListObjectTypes(listParams *warrant.ListObjectTypeParams) (warrant.ListResponse[warrant.ObjectType], error) {
	if listParams == nil {
		listParams = &warrant.ListObjectTypeParams{}
	}
	var objectTypesListResponse warrant.ListResponse[warrant.ObjectType]
	queryParams, err := query.Values(listParams)
	if err != nil {
		return objectTypesListResponse, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v2/object-types?%s", queryParams.Encode()), objectTypesListResponse, &listParams.RequestOptions)
	if err != nil {
		return objectTypesListResponse, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return objectTypesListResponse, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	err = json.Unmarshal([]byte(body), &objectTypesListResponse)
	if err != nil {
		return objectTypesListResponse, warrant.WrapError("Invalid response from server", err)
	}
	return objectTypesListResponse, nil
}

func ListObjectTypes(listParams *warrant.ListObjectTypeParams) (warrant.ListResponse[warrant.ObjectType], error) {
	return getClient().ListObjectTypes(listParams)
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
