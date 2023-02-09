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

func (c Client) Check(object WarrantObject, relation string, subject Subject) (bool, error) {
	warrantCheckParams := WarrantCheckParams{
		Warrants: []Warrant{
			{
				ObjectType: object.ObjectType,
				ObjectId:   object.ObjectId,
				Relation:   relation,
				Subject:    subject,
			},
		},
	}

	checkResult, err := c.makeAuthorizeRequest(&warrantCheckParams)
	if err != nil {
		return false, err
	}

	if checkResult.Result == "Authorized" {
		return true, nil
	} else {
		return false, nil
	}
}

func Check(object WarrantObject, relation string, subject Subject) (bool, error) {
	return getClient().Check(object, relation, subject)
}

func (c Client) CheckMany(op string, warrants []Warrant) (bool, error) {
	warrantCheckParams := WarrantCheckParams{
		Op:       op,
		Warrants: warrants,
	}

	checkResult, err := c.makeAuthorizeRequest(&warrantCheckParams)
	if err != nil {
		return false, err
	}

	if checkResult.Result == "Authorized" {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckMany(op string, warrants []Warrant) (bool, error) {
	return getClient().CheckMany(op, warrants)
}

func (c Client) CheckUserHasPermission(user *User, permissionId string) (bool, error) {
	return c.Check(
		WarrantObject{
			ObjectType: "permission",
			ObjectId:   permissionId,
		},
		"member",
		Subject{
			ObjectType: "user",
			ObjectId:   user.UserId,
		},
	)
}

func CheckUserHasPermission(user *User, permissionId string) (bool, error) {
	return getClient().CheckUserHasPermission(user, permissionId)
}

func (c Client) CheckUserHasRole(user *User, roleId string) (bool, error) {
	return c.Check(
		WarrantObject{
			ObjectType: "role",
			ObjectId:   roleId,
		},
		"member",
		Subject{
			ObjectType: "user",
			ObjectId:   user.UserId,
		},
	)
}

func CheckUserHasRole(user *User, roleId string) (bool, error) {
	return getClient().CheckUserHasRole(user, roleId)
}

func (c Client) CheckHasFeature(subject *Subject, featureId string) (bool, error) {
	return c.Check(
		WarrantObject{
			ObjectType: "feature",
			ObjectId:   featureId,
		},
		"member",
		*subject,
	)
}

func CheckHasFeature(subject *Subject, featureId string) (bool, error) {
	return getClient().CheckHasFeature(subject, featureId)
}

func (c Client) makeAuthorizeRequest(params *WarrantCheckParams) (*WarrantCheckResult, error) {
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
