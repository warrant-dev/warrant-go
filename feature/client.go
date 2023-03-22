package feature

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go/v3"
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

func (c Client) Create(params *warrant.FeatureParams) (*warrant.Feature, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/features", params)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var newFeature warrant.Feature
	err = json.Unmarshal([]byte(body), &newFeature)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &newFeature, nil
}

func Create(params *warrant.FeatureParams) (*warrant.Feature, error) {
	return getClient().Create(params)
}

func (c Client) Get(featureId string) (*warrant.Feature, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/features/%s", featureId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var foundFeature warrant.Feature
	err = json.Unmarshal([]byte(body), &foundFeature)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &foundFeature, nil
}

func Get(featureId string) (*warrant.Feature, error) {
	return getClient().Get(featureId)
}

func (c Client) Delete(featureId string) error {
	resp, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/features/%s", featureId), nil)
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

func Delete(featureId string) error {
	return getClient().Delete(featureId)
}

func (c Client) ListFeatures(listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/features?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var permissions []warrant.Feature
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return permissions, nil
}

func ListFeatures(listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	return getClient().ListFeatures(listParams)
}

func (c Client) ListFeaturesForPricingTier(pricingTierId string, listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/pricing-tiers/%s/features?%s", pricingTierId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var features []warrant.Feature
	err = json.Unmarshal([]byte(body), &features)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return features, nil
}

func ListFeaturesForPricingTier(pricingTierId string, listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	return getClient().ListFeaturesForPricingTier(pricingTierId, listParams)
}

func (c Client) AssignFeatureToPricingTier(featureId string, pricingTierId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.warrantClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeFeature,
		ObjectId:   featureId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypePricingTier,
			ObjectId:   pricingTierId,
		},
	})
}

func AssignFeatureToPricingTier(featureId string, pricingTierId string) (*warrant.Warrant, error) {
	return getClient().AssignFeatureToPricingTier(featureId, pricingTierId)
}

func (c Client) RemoveFeatureFromPricingTier(featureId string, pricingTierId string) error {
	return warrant.NewClient(c.warrantClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeFeature,
		ObjectId:   featureId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypePricingTier,
			ObjectId:   pricingTierId,
		},
	})
}

func RemoveFeatureFromPricingTier(featureId string, pricingTierId string) error {
	return getClient().RemoveFeatureFromPricingTier(featureId, pricingTierId)
}

func (c Client) ListFeaturesForTenant(tenantId string, listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s/features?%s", tenantId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var features []warrant.Feature
	err = json.Unmarshal([]byte(body), &features)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return features, nil
}

func ListFeaturesForTenant(tenantId string, listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	return getClient().ListFeaturesForTenant(tenantId, listParams)
}

func (c Client) AssignFeatureToTenant(featureId string, tenantId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.warrantClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeFeature,
		ObjectId:   featureId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   tenantId,
		},
	})
}

func AssignFeatureToTenant(featureId string, tenantId string) (*warrant.Warrant, error) {
	return getClient().AssignFeatureToTenant(featureId, tenantId)
}

func (c Client) RemoveFeatureFromTenant(featureId string, tenantId string) error {
	return warrant.NewClient(c.warrantClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeFeature,
		ObjectId:   featureId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   tenantId,
		},
	})
}

func RemoveFeatureFromTenant(featureId string, tenantId string) error {
	return getClient().RemoveFeatureFromTenant(featureId, tenantId)
}

func (c Client) ListFeaturesForUser(userId string, listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/features?%s", userId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var features []warrant.Feature
	err = json.Unmarshal([]byte(body), &features)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return features, nil
}

func ListFeaturesForUser(userId string, listParams *warrant.ListFeatureParams) ([]warrant.Feature, error) {
	return getClient().ListFeaturesForUser(userId, listParams)
}

func (c Client) AssignFeatureToUser(featureId string, userId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.warrantClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeFeature,
		ObjectId:   featureId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignFeatureToUser(featureId string, userId string) (*warrant.Warrant, error) {
	return getClient().AssignFeatureToUser(featureId, userId)
}

func (c Client) RemoveFeatureFromUser(featureId string, userId string) error {
	return warrant.NewClient(c.warrantClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeFeature,
		ObjectId:   featureId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemoveFeatureFromUser(featureId string, userId string) error {
	return getClient().RemoveFeatureFromUser(featureId, userId)
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
