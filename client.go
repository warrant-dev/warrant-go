package warrant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type WarrantClient struct {
	apiClient *ApiClient
}

func NewClient(config ClientConfig) WarrantClient {
	return WarrantClient{
		apiClient: &ApiClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}

func (c WarrantClient) Create(params *WarrantParams) (*Warrant, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v1/warrants", params, &RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapError("Error reading response", err)
	}
	var createdWarrant Warrant
	err = json.Unmarshal([]byte(body), &createdWarrant)
	if err != nil {
		return nil, WrapError("Invalid response from server", err)
	}
	return &createdWarrant, nil
}

func Create(params *WarrantParams) (*Warrant, error) {
	return getClient().Create(params)
}

func (c WarrantClient) Delete(params *WarrantParams) error {
	_, err := c.apiClient.MakeRequest("DELETE", "/v1/warrants", params, &RequestOptions{})
	if err != nil {
		return err
	}
	return nil
}

func Delete(params *WarrantParams) error {
	return getClient().Delete(params)
}

func (c WarrantClient) Query(queryString string, listParams *ListWarrantParams) (*QueryWarrantResult, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v1/query?q=%s&%s", url.QueryEscape(queryString), queryParams.Encode()), nil, &listParams.RequestOptions)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapError("Error reading response", err)
	}
	var queryResult QueryWarrantResult
	err = json.Unmarshal([]byte(body), &queryResult)
	if err != nil {
		return nil, WrapError("Invalid response from server", err)
	}
	return &queryResult, nil
}

func Query(queryString string, params *ListWarrantParams) (*QueryWarrantResult, error) {
	return getClient().Query(queryString, params)
}

func (c WarrantClient) Check(params *WarrantCheckParams) (bool, error) {
	accessCheckRequest := AccessCheckRequest{
		RequestOptions: params.RequestOptions,
		Warrants:       []WarrantCheck{params.WarrantCheck},
		Debug:          params.Debug,
	}

	checkResult, err := c.makeAuthorizeRequest(&accessCheckRequest)
	if err != nil {
		return false, err
	}

	if checkResult.Result == "Authorized" {
		return true, nil
	} else {
		return false, nil
	}
}

func Check(params *WarrantCheckParams) (bool, error) {
	return getClient().Check(params)
}

func (c WarrantClient) CheckMany(params *WarrantCheckManyParams) (bool, error) {
	warrants := make([]WarrantCheck, 0)
	for _, warrantCheck := range params.Warrants {
		warrants = append(warrants, warrantCheck)
	}

	accessCheckRequest := AccessCheckRequest{
		RequestOptions: params.RequestOptions,
		Op:             params.Op,
		Warrants:       warrants,
		Debug:          params.Debug,
	}

	checkResult, err := c.makeAuthorizeRequest(&accessCheckRequest)
	if err != nil {
		return false, err
	}

	if checkResult.Result == "Authorized" {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckMany(params *WarrantCheckManyParams) (bool, error) {
	return getClient().CheckMany(params)
}

func (c WarrantClient) CheckUserHasPermission(params *PermissionCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
		RequestOptions: params.RequestOptions,
		WarrantCheck: WarrantCheck{
			Object: Object{
				ObjectType: ObjectTypePermission,
				ObjectId:   params.PermissionId,
			},
			Relation: "member",
			Subject: Subject{
				ObjectType: ObjectTypeUser,
				ObjectId:   params.UserId,
			},
			Context: params.Context,
		},
		Debug: params.Debug,
	})
}

func CheckUserHasPermission(params *PermissionCheckParams) (bool, error) {
	return getClient().CheckUserHasPermission(params)
}

func (c WarrantClient) CheckUserHasRole(params *RoleCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
		RequestOptions: params.RequestOptions,
		WarrantCheck: WarrantCheck{
			Object: Object{
				ObjectType: ObjectTypeRole,
				ObjectId:   params.RoleId,
			},
			Relation: "member",
			Subject: Subject{
				ObjectType: ObjectTypeUser,
				ObjectId:   params.UserId,
			},
			Context: params.Context,
		},
		Debug: params.Debug,
	})
}

func CheckUserHasRole(params *RoleCheckParams) (bool, error) {
	return getClient().CheckUserHasRole(params)
}

func (c WarrantClient) CheckHasFeature(params *FeatureCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
		RequestOptions: params.RequestOptions,
		WarrantCheck: WarrantCheck{
			Object: Object{
				ObjectType: ObjectTypeFeature,
				ObjectId:   params.FeatureId,
			},
			Relation: "member",
			Subject:  params.Subject,
			Context:  params.Context,
		},
		Debug: params.Debug,
	})
}

func CheckHasFeature(params *FeatureCheckParams) (bool, error) {
	return getClient().CheckHasFeature(params)
}

func (c WarrantClient) makeAuthorizeRequest(params *AccessCheckRequest) (*WarrantCheckResult, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v2/authorize", params, &params.RequestOptions)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapError("Error reading response", err)
	}
	var result WarrantCheckResult
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, WrapError("Invalid response from server", err)
	}
	return &result, nil
}

func getClient() WarrantClient {
	config := ClientConfig{
		ApiKey:                  ApiKey,
		ApiEndpoint:             ApiEndpoint,
		AuthorizeEndpoint:       AuthorizeEndpoint,
		SelfServiceDashEndpoint: SelfServiceDashEndpoint,
		HttpClient:              HttpClient,
	}

	return WarrantClient{
		&ApiClient{
			HttpClient: HttpClient,
			Config:     config,
		},
	}
}
