package feature

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

func (c Client) Create(params *warrant.FeatureParams) (*warrant.Feature, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeFeature,
		RequestOptions: params.RequestOptions,
	}
	if params.FeatureId != "" {
		objectParams.ObjectId = params.FeatureId
	}
	if params.Meta != nil {
		objectParams.Meta = params.Meta
	}
	object, err := object.Create(&objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Feature{
		FeatureId: object.ObjectId,
		Meta:      object.Meta,
	}, nil
}

func Create(params *warrant.FeatureParams) (*warrant.Feature, error) {
	return getClient().Create(params)
}

func (c Client) Get(featureId string, params *warrant.FeatureParams) (*warrant.Feature, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeFeature,
		ObjectId:       featureId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Get(warrant.ObjectTypeFeature, featureId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Feature{
		FeatureId: object.ObjectId,
		Meta:      object.Meta,
	}, nil
}

func Get(featureId string, params *warrant.FeatureParams) (*warrant.Feature, error) {
	return getClient().Get(featureId, params)
}

func (c Client) Update(featureId string, params *warrant.FeatureParams) (*warrant.Feature, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeFeature,
		ObjectId:       featureId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Update(warrant.ObjectTypeFeature, featureId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Feature{
		FeatureId: object.ObjectId,
		Meta:      object.Meta,
	}, nil
}

func Update(featureId string, params *warrant.FeatureParams) (*warrant.Feature, error) {
	return getClient().Update(featureId, params)
}

func (c Client) Delete(featureId string) error {
	return object.Delete(warrant.ObjectTypeFeature, featureId)
}

func Delete(featureId string) error {
	return getClient().Delete(featureId)
}

func (c Client) ListFeatures(listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	var featuresListResponse warrant.ListResponse[warrant.Feature]

	objectsListResponse, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: listParams.ListParams,
		ObjectType: warrant.ObjectTypeFeature,
	})
	if err != nil {
		return featuresListResponse, err
	}

	features := make([]warrant.Feature, 0)
	for _, object := range objectsListResponse.Results {
		features = append(features, warrant.Feature{
			FeatureId: object.ObjectId,
			Meta:      object.Meta,
		})
	}

	featuresListResponse = warrant.ListResponse[warrant.Feature]{
		Results:    features,
		PrevCursor: objectsListResponse.PrevCursor,
		NextCursor: objectsListResponse.NextCursor,
	}

	return featuresListResponse, nil
}

func ListFeatures(listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	return getClient().ListFeatures(listParams)
}

func (c Client) ListFeaturesForPricingTier(pricingTierId string, listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	var featuresListResponse warrant.ListResponse[warrant.Feature]

	queryResponse, err := warrant.Query(fmt.Sprintf("select feature where pricing-tier:%s is *", pricingTierId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return featuresListResponse, err
	}

	features := make([]warrant.Feature, 0)
	for _, queryResult := range queryResponse.Results {
		features = append(features, warrant.Feature{
			FeatureId: queryResult.ObjectId,
			Meta:      queryResult.Meta,
		})
	}

	featuresListResponse = warrant.ListResponse[warrant.Feature]{
		Results:    features,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return featuresListResponse, nil
}

func ListFeaturesForPricingTier(pricingTierId string, listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	return getClient().ListFeaturesForPricingTier(pricingTierId, listParams)
}

func (c Client) AssignFeatureToPricingTier(featureId string, pricingTierId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
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
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
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

func (c Client) ListFeaturesForTenant(tenantId string, listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	var featuresListResponse warrant.ListResponse[warrant.Feature]

	queryResponse, err := warrant.Query(fmt.Sprintf("select feature where tenant:%s is *", tenantId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return featuresListResponse, err
	}

	features := make([]warrant.Feature, 0)
	for _, queryResult := range queryResponse.Results {
		features = append(features, warrant.Feature{
			FeatureId: queryResult.ObjectId,
			Meta:      queryResult.Meta,
		})
	}

	featuresListResponse = warrant.ListResponse[warrant.Feature]{
		Results:    features,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return featuresListResponse, nil
}

func ListFeaturesForTenant(tenantId string, listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	return getClient().ListFeaturesForTenant(tenantId, listParams)
}

func (c Client) AssignFeatureToTenant(featureId string, tenantId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
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
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
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

func (c Client) ListFeaturesForUser(userId string, listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	var featuresListResponse warrant.ListResponse[warrant.Feature]

	queryResponse, err := warrant.Query(fmt.Sprintf("select feature where user:%s is *", userId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return featuresListResponse, err
	}

	features := make([]warrant.Feature, 0)
	for _, queryResult := range queryResponse.Results {
		features = append(features, warrant.Feature{
			FeatureId: queryResult.ObjectId,
			Meta:      queryResult.Meta,
		})
	}

	featuresListResponse = warrant.ListResponse[warrant.Feature]{
		Results:    features,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return featuresListResponse, nil
}

func ListFeaturesForUser(userId string, listParams *warrant.ListFeatureParams) (warrant.ListResponse[warrant.Feature], error) {
	return getClient().ListFeaturesForUser(userId, listParams)
}

func (c Client) AssignFeatureToUser(featureId string, userId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
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
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
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
