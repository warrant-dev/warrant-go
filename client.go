package warrant

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/warrant-dev/warrant-go/client"
)

type Client struct {
	warrantClient *client.WarrantClient
}

func (c Client) Create(params *WarrantParams) (*Warrant, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/warrants", params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
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

func (c Client) Check(params *WarrantCheckParams) (bool, error) {
	accessCheckRequest := AccessCheckRequest{
		Warrants: []Warrant{
			{
				ObjectType: params.Object.ObjectType,
				ObjectId:   params.Object.ObjectId,
				Relation:   params.Relation,
				Subject:    *params.Subject,
				Context:    params.Context,
			},
		},
		ConsistentRead: params.ConsistentRead,
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

func (c Client) CheckMany(params *WarrantCheckManyParams) (bool, error) {
	warrants := make([]Warrant, 0)
	for _, warrantCheck := range params.Warrants {
		warrants = append(warrants, warrantCheck.ToWarrant())
	}

	accessCheckRequest := AccessCheckRequest{
		Op:             params.Op,
		Warrants:       warrants,
		ConsistentRead: params.ConsistentRead,
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

func (c Client) CheckUserHasPermission(params *PermissionCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
		Object: &WarrantObject{
			ObjectType: "permission",
			ObjectId:   params.PermissionId,
		},
		Relation: "member",
		Subject: &Subject{
			ObjectType: "user",
			ObjectId:   params.UserId,
		},
		Context:        params.Context,
		ConsistentRead: params.ConsistentRead,
		Debug:          params.Debug,
	})
}

func CheckUserHasPermission(params *PermissionCheckParams) (bool, error) {
	return getClient().CheckUserHasPermission(params)
}

func (c Client) CheckUserHasRole(params *RoleCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
		Object: &WarrantObject{
			ObjectType: "role",
			ObjectId:   params.RoleId,
		},
		Relation: "member",
		Subject: &Subject{
			ObjectType: "user",
			ObjectId:   params.UserId,
		},
		Context:        params.Context,
		ConsistentRead: params.ConsistentRead,
		Debug:          params.Debug,
	})
}

func CheckUserHasRole(params *RoleCheckParams) (bool, error) {
	return getClient().CheckUserHasRole(params)
}

func (c Client) CheckHasFeature(params *FeatureCheckParams) (bool, error) {
	return c.Check(&WarrantCheckParams{
		Object: &WarrantObject{
			ObjectType: "feature",
			ObjectId:   params.FeatureId,
		},
		Relation:       "member",
		Subject:        params.Subject,
		Context:        params.Context,
		ConsistentRead: params.ConsistentRead,
		Debug:          params.Debug,
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

	body, err := ioutil.ReadAll(resp.Body)
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
	if ApiKey == "" {
		panic("You must provide an ApiKey to initialize the Warrant Client")
	}

	config := client.ClientConfig{
		ApiKey:            ApiKey,
		AuthorizeEndpoint: AuthorizeEndpoint,
	}

	return Client{
		&client.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}
