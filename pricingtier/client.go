package pricingtier

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

func (c Client) Create(params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypePricingTier,
		RequestOptions: params.RequestOptions,
	}
	if params.PricingTierId != "" {
		objectParams.ObjectId = params.PricingTierId
	}
	if params.Meta != nil {
		objectParams.Meta = params.Meta
	}
	object, err := object.Create(&objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.PricingTier{
		PricingTierId: object.ObjectId,
		Meta:          object.Meta,
	}, nil
}

func Create(params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	return getClient().Create(params)
}

func (c Client) Get(pricingTierId string, params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypePricingTier,
		ObjectId:       pricingTierId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Get(warrant.ObjectTypePricingTier, pricingTierId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.PricingTier{
		PricingTierId: object.ObjectId,
		Meta:          object.Meta,
	}, nil
}

func Get(pricingTierId string, params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	return getClient().Get(pricingTierId, params)
}

func (c Client) Update(pricingTierId string, params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypePricingTier,
		ObjectId:       pricingTierId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Update(warrant.ObjectTypePricingTier, pricingTierId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.PricingTier{
		PricingTierId: object.ObjectId,
		Meta:          object.Meta,
	}, nil
}

func Update(pricingTierId string, params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	return getClient().Update(pricingTierId, params)
}

func (c Client) Delete(pricingTierId string) error {
	return object.Delete(warrant.ObjectTypePricingTier, pricingTierId)
}

func Delete(pricingTierId string) error {
	return getClient().Delete(pricingTierId)
}

func (c Client) ListPricingTiers(listParams *warrant.ListPricingTierParams) (warrant.ListResponse[warrant.PricingTier], error) {
	var pricingTiersListResponse warrant.ListResponse[warrant.PricingTier]

	objectsListResponse, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: listParams.ListParams,
		ObjectType: warrant.ObjectTypePricingTier,
	})
	if err != nil {
		return pricingTiersListResponse, err
	}

	users := make([]warrant.PricingTier, 0)
	for _, object := range objectsListResponse.Results {
		users = append(users, warrant.PricingTier{
			PricingTierId: object.ObjectId,
			Meta:          object.Meta,
		})
	}

	pricingTiersListResponse = warrant.ListResponse[warrant.PricingTier]{
		Results:    users,
		PrevCursor: objectsListResponse.PrevCursor,
		NextCursor: objectsListResponse.NextCursor,
	}

	return pricingTiersListResponse, nil
}

func ListPricingTiers(listParams *warrant.ListPricingTierParams) (warrant.ListResponse[warrant.PricingTier], error) {
	return getClient().ListPricingTiers(listParams)
}

func (c Client) ListPricingTiersForTenant(tenantId string, listParams *warrant.ListPricingTierParams) (warrant.ListResponse[warrant.PricingTier], error) {
	var pricingTiersListResponse warrant.ListResponse[warrant.PricingTier]

	queryResponse, err := warrant.Query(fmt.Sprintf("select pricing-tier where tenant:%s is *", tenantId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return pricingTiersListResponse, err
	}

	users := make([]warrant.PricingTier, 0)
	for _, queryResult := range queryResponse.Results {
		users = append(users, warrant.PricingTier{
			PricingTierId: queryResult.ObjectId,
			Meta:          queryResult.Meta,
		})
	}

	pricingTiersListResponse = warrant.ListResponse[warrant.PricingTier]{
		Results:    users,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return pricingTiersListResponse, nil
}

func ListPricingTiersForTenant(userId string, listParams *warrant.ListPricingTierParams) (warrant.ListResponse[warrant.PricingTier], error) {
	return getClient().ListPricingTiersForTenant(userId, listParams)
}

func (c Client) AssignPricingTierToTenant(pricingTierId string, tenantId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePricingTier,
		ObjectId:   pricingTierId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   tenantId,
		},
	})
}

func AssignPricingTierToTenant(pricingTierId string, tenantId string) (*warrant.Warrant, error) {
	return getClient().AssignPricingTierToTenant(pricingTierId, tenantId)
}

func (c Client) RemovePricingTierFromTenant(pricingTierId string, tenantId string) error {
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePricingTier,
		ObjectId:   pricingTierId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   tenantId,
		},
	})
}

func RemovePricingTierFromTenant(pricingTierId string, tenantId string) error {
	return getClient().RemovePricingTierFromTenant(pricingTierId, tenantId)
}

func (c Client) ListPricingTiersForUser(userId string, listParams *warrant.ListPricingTierParams) (warrant.ListResponse[warrant.PricingTier], error) {
	var pricingTiersListResponse warrant.ListResponse[warrant.PricingTier]

	queryResponse, err := warrant.Query(fmt.Sprintf("select pricing-tier where user:%s is *", userId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return pricingTiersListResponse, err
	}

	users := make([]warrant.PricingTier, 0)
	for _, queryResult := range queryResponse.Results {
		users = append(users, warrant.PricingTier{
			PricingTierId: queryResult.ObjectId,
			Meta:          queryResult.Meta,
		})
	}

	pricingTiersListResponse = warrant.ListResponse[warrant.PricingTier]{
		Results:    users,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return pricingTiersListResponse, nil
}

func ListPricingTiersForUser(userId string, listParams *warrant.ListPricingTierParams) (warrant.ListResponse[warrant.PricingTier], error) {
	return getClient().ListPricingTiersForUser(userId, listParams)
}

func (c Client) AssignPricingTierToUser(pricingTierId string, userId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePricingTier,
		ObjectId:   pricingTierId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignPricingTierToUser(pricingTierId string, userId string) (*warrant.Warrant, error) {
	return getClient().AssignPricingTierToUser(pricingTierId, userId)
}

func (c Client) RemovePricingTierFromUser(pricingTierId string, userId string) error {
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePricingTier,
		ObjectId:   pricingTierId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemovePricingTierFromUser(pricingTierId string, userId string) error {
	return getClient().RemovePricingTierFromUser(pricingTierId, userId)
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
