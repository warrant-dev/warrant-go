package object

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v5"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: warrant.NewApiClient(config),
	}
}

func (c Client) Create(params *warrant.ObjectParams) (*warrant.Object, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v2/objects", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var newObject warrant.Object
	err = json.Unmarshal([]byte(body), &newObject)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newObject, nil
}

func Create(params *warrant.ObjectParams) (*warrant.Object, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.ObjectParams) ([]warrant.Object, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v2/objects", params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var newObjects []warrant.Object
	err = json.Unmarshal([]byte(body), &newObjects)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return newObjects, nil
}

func BatchCreate(params []warrant.ObjectParams) ([]warrant.Object, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(objectType string, objectId string, params *warrant.ObjectParams) (*warrant.Object, error) {
	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v2/objects/%s/%s", objectType, objectId), nil, &params.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var foundObject warrant.Object
	err = json.Unmarshal([]byte(body), &foundObject)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &foundObject, nil
}

func Get(objectType string, objectId string, params *warrant.ObjectParams) (*warrant.Object, error) {
	return getClient().Get(objectType, objectId, params)
}

func (c Client) Update(objectType string, objectId string, params *warrant.ObjectParams) (*warrant.Object, error) {
	resp, err := c.apiClient.MakeRequest("PUT", fmt.Sprintf("/v2/objects/%s/%s", objectType, objectId), params, &warrant.RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var updatedObject warrant.Object
	err = json.Unmarshal([]byte(body), &updatedObject)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &updatedObject, nil
}

func Update(objectType string, objectId string, params *warrant.ObjectParams) (*warrant.Object, error) {
	return getClient().Update(objectType, objectId, params)
}

func (c Client) Delete(objectType string, objectId string) (string, error) {
	resp, err := c.apiClient.MakeRequest("DELETE", fmt.Sprintf("/v2/objects/%s/%s", objectType, objectId), nil, &warrant.RequestOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	warrantToken := resp.Header.Get("Warrant-Token")
	return warrantToken, nil
}

func Delete(objectType string, objectId string) (string, error) {
	return getClient().Delete(objectType, objectId)
}

func (c Client) BatchDelete(params []warrant.ObjectParams) (string, error) {
	resp, err := c.apiClient.MakeRequest("DELETE", "/v2/objects", params, &warrant.RequestOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	warrantToken := resp.Header.Get("Warrant-Token")
	return warrantToken, nil
}

func BatchDelete(params []warrant.ObjectParams) (string, error) {
	return getClient().BatchDelete(params)
}

func (c Client) ListObjects(listParams *warrant.ListObjectParams) (warrant.ListResponse[warrant.Object], error) {
	var objectsListResponse warrant.ListResponse[warrant.Object]
	queryParams, err := query.Values(listParams)
	if err != nil {
		return objectsListResponse, warrant.WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v2/objects?%s", queryParams.Encode()), objectsListResponse, &listParams.RequestOptions)
	if err != nil {
		return objectsListResponse, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return objectsListResponse, warrant.WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	err = json.Unmarshal([]byte(body), &objectsListResponse)
	if err != nil {
		return objectsListResponse, warrant.WrapError("Invalid response from server", err)
	}
	return objectsListResponse, nil
}

func ListObjects(listParams *warrant.ListObjectParams) (warrant.ListResponse[warrant.Object], error) {
	return getClient().ListObjects(listParams)
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
