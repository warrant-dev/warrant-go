package warrant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v3/client"
	"github.com/warrant-dev/warrant-go/v3/config"
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

func (c Client) Create(params *WarrantParams) (*Warrant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/warrants", params)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var createdWarrant Warrant
	err = json.Unmarshal([]byte(body), &createdWarrant)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &createdWarrant, nil
}

func Create(params *WarrantParams) (*Warrant, error) {
	return getClient().Create(params)
}

func (c Client) Delete(params *WarrantParams) error {
	_, err := c.warrantClient.MakeRequest("DELETE", "/v1/warrants", params)
	if err != nil {
		return err
	}
	return nil
}

func Delete(params *WarrantParams) error {
	return getClient().Delete(params)
}

func (c Client) Query(queryString string, listParams *ListWarrantParams) (*QueryWarrantResult, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/query?q=%s&%s", url.QueryEscape(queryString), queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var queryResult QueryWarrantResult
	err = json.Unmarshal([]byte(body), &queryResult)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &queryResult, nil
}

func Query(queryString string, params *ListWarrantParams) (*QueryWarrantResult, error) {
	return getClient().Query(queryString, params)
}

func (c Client) Check(params *WarrantCheckParams) (bool, error) {
	accessCheckRequest := AccessCheckRequest{
		Warrants: []WarrantCheck{params.WarrantCheck},
		Debug:    params.Debug,
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

func (c Client) CheckMany(params *WarrantCheckManyParams) (bool, error) {
	warrants := make([]WarrantCheck, 0)
	for _, warrantCheck := range params.Warrants {
		warrants = append(warrants, warrantCheck)
	}

	accessCheckRequest := AccessCheckRequest{
		Op:       params.Op,
		Warrants: warrants,
		Debug:    params.Debug,
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

func (c Client) CheckUserHasPermission(params *PermissionCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
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

func (c Client) CheckUserHasRole(params *RoleCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
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

func (c Client) CheckHasFeature(params *FeatureCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
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

func (c Client) makeAuthorizeRequest(params *AccessCheckRequest) (*WarrantCheckResult, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v2/authorize", params)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var result WarrantCheckResult
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &result, nil
}

func getClient() Client {
	config := config.ClientConfig{
		ApiKey:                  ApiKey,
		ApiEndpoint:             ApiEndpoint,
		AuthorizeEndpoint:       AuthorizeEndpoint,
		SelfServiceDashEndpoint: SelfServiceDashEndpoint,
	}

	return Client{
		&client.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}
