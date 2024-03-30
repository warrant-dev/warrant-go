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
	resp, err := c.apiClient.MakeRequest("POST", "/v2/warrants", params, &RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var createdWarrant Warrant
	err = json.Unmarshal([]byte(body), &createdWarrant)
	if err != nil {
		return nil, WrapError("Invalid response from server", err)
	}
	warrantToken := resp.Header.Get("Warrant-Token")
	createdWarrant.WarrantToken = warrantToken
	return &createdWarrant, nil
}

func Create(params *WarrantParams) (*Warrant, error) {
	return getClient().Create(params)
}

func (c WarrantClient) BatchCreate(params []WarrantParams) ([]Warrant, error) {
	resp, err := c.apiClient.MakeRequest("POST", "/v2/warrants", params, &RequestOptions{})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	var createdWarrants []Warrant
	err = json.Unmarshal([]byte(body), &createdWarrants)
	if err != nil {
		return nil, WrapError("Invalid response from server", err)
	}
	warrantToken := resp.Header.Get("Warrant-Token")
	for i := range createdWarrants {
		createdWarrants[i].WarrantToken = warrantToken
	}
	return createdWarrants, nil
}

func BatchCreate(params []WarrantParams) ([]Warrant, error) {
	return getClient().BatchCreate(params)
}

func (c WarrantClient) Delete(params *WarrantParams) (string, error) {
	resp, err := c.apiClient.MakeRequest("DELETE", "/v2/warrants", params, &RequestOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	warrantToken := resp.Header.Get("Warrant-Token")
	return warrantToken, nil
}

func Delete(params *WarrantParams) (string, error) {
	return getClient().Delete(params)
}

func (c WarrantClient) BatchDelete(params []WarrantParams) (string, error) {
	resp, err := c.apiClient.MakeRequest("DELETE", "/v2/warrants", params, &RequestOptions{})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	warrantToken := resp.Header.Get("Warrant-Token")
	return warrantToken, nil
}

func BatchDelete(params []WarrantParams) (string, error) {
	return getClient().BatchDelete(params)
}

func (c WarrantClient) ListWarrants(listParams *ListWarrantParams) (ListResponse[Warrant], error) {
	if listParams == nil {
		listParams = &ListWarrantParams{}
	}
	var warrantsListResponse ListResponse[Warrant]
	queryParams, err := query.Values(listParams)
	if err != nil {
		return warrantsListResponse, WrapError("Could not parse listParams", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v2/warrants?%s", queryParams.Encode()), warrantsListResponse, &listParams.RequestOptions)
	if err != nil {
		return warrantsListResponse, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return warrantsListResponse, WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	err = json.Unmarshal([]byte(body), &warrantsListResponse)
	if err != nil {
		return warrantsListResponse, WrapError("Invalid response from server", err)
	}
	return warrantsListResponse, nil
}

func ListWarrants(listParams *ListWarrantParams) (ListResponse[Warrant], error) {
	return getClient().ListWarrants(listParams)
}

func (c WarrantClient) Query(queryString string, params *QueryParams) (ListResponse[QueryResult], error) {
	if params == nil {
		params = &QueryParams{}
	}
	var queryResponse ListResponse[QueryResult]
	queryParams, err := query.Values(params)
	if err != nil {
		return queryResponse, WrapError("Could not parse params", err)
	}

	resp, err := c.apiClient.MakeRequest("GET", fmt.Sprintf("/v2/query?q=%s&%s", url.QueryEscape(queryString), queryParams.Encode()), queryResponse, &params.RequestOptions)
	if err != nil {
		return queryResponse, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return queryResponse, WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
	err = json.Unmarshal([]byte(body), &queryResponse)
	if err != nil {
		return queryResponse, WrapError("Invalid response from server", err)
	}
	return queryResponse, nil
}

func Query(queryString string, params *QueryParams) (ListResponse[QueryResult], error) {
	return getClient().Query(queryString, params)
}

func (c WarrantClient) Check(params *WarrantCheckParams) (bool, error) {
	if params == nil {
		params = &WarrantCheckParams{}
	}
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
	if params == nil {
		params = &WarrantCheckManyParams{}
	}
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
	if params == nil {
		params = &PermissionCheckParams{}
	}
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
	if params == nil {
		params = &RoleCheckParams{}
	}
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
	if params == nil {
		params = &FeatureCheckParams{}
	}
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
	resp, err := c.apiClient.MakeRequest("POST", "/v2/check", params, &params.RequestOptions)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, WrapError("Error reading response", err)
	}
	defer resp.Body.Close()
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
