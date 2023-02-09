package pricingtier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/client"
)

type Client struct {
	warrantClient *client.WarrantClient
}

func (c Client) Create(params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	resp, err := c.warrantClient.MakeRequest("POST", "/v1/pricing-tiers", params)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var newPricingTier warrant.PricingTier
	err = json.Unmarshal([]byte(body), &newPricingTier)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &newPricingTier, nil
}

func Create(params *warrant.PricingTierParams) (*warrant.PricingTier, error) {
	return getClient().Create(params)
}

func (c Client) Get(pricingTierId string) (*warrant.PricingTier, error) {
	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/pricing-tiers/%s", pricingTierId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var foundPricingTier warrant.PricingTier
	err = json.Unmarshal([]byte(body), &foundPricingTier)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &foundPricingTier, nil
}

func Get(pricingTierId string) (*warrant.PricingTier, error) {
	return getClient().Get(pricingTierId)
}

func (c Client) Delete(pricingTierId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/pricing-tiers/%s", pricingTierId), nil)
	if err != nil {
		return err
	}
	return nil
}

func Delete(pricingTierId string) error {
	return getClient().Delete(pricingTierId)
}

func (c Client) ListPricingTiers(listParams *warrant.ListPricingTierParams) ([]warrant.PricingTier, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/pricing-tiers?%s", queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var permissions []warrant.PricingTier
	err = json.Unmarshal([]byte(body), &permissions)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return permissions, nil
}

func ListPricingTiers(listParams *warrant.ListPricingTierParams) ([]warrant.PricingTier, error) {
	return getClient().ListPricingTiers(listParams)
}

func (c Client) ListPricingTiersForTenant(tenantId string, listParams *warrant.ListPricingTierParams) ([]warrant.PricingTier, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/tenants/%s/pricing-tiers?%s", tenantId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var pricingTiers []warrant.PricingTier
	err = json.Unmarshal([]byte(body), &pricingTiers)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return pricingTiers, nil
}

func ListPricingTiersForTenant(userId string, listParams *warrant.ListPricingTierParams) ([]warrant.PricingTier, error) {
	return getClient().ListPricingTiersForTenant(userId, listParams)
}

func (c Client) AssignPricingTierToTenant(pricingTierId string, tenantId string) (*warrant.PricingTier, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/tenants/%s/pricing-tiers/%s", tenantId, pricingTierId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var assignedPricingTier warrant.PricingTier
	err = json.Unmarshal([]byte(body), &assignedPricingTier)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &assignedPricingTier, nil
}

func AssignPricingTierToTenant(pricingTierId string, tenantId string) (*warrant.PricingTier, error) {
	return getClient().AssignPricingTierToTenant(pricingTierId, tenantId)
}

func (c Client) RemovePricingTierFromTenant(pricingTierId string, tenantId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/tenants/%s/pricing-tiers/%s", tenantId, pricingTierId), nil)
	if err != nil {
		return err
	}
	return nil
}

func RemovePricingTierFromTenant(pricingTierId string, tenantId string) error {
	return getClient().RemovePricingTierFromTenant(pricingTierId, tenantId)
}

func (c Client) ListPricingTiersForUser(userId string, listParams *warrant.ListPricingTierParams) ([]warrant.PricingTier, error) {
	queryParams, err := query.Values(listParams)
	if err != nil {
		return nil, client.WrapError("Could not parse listParams", err)
	}

	resp, err := c.warrantClient.MakeRequest("GET", fmt.Sprintf("/v1/users/%s/pricing-tiers?%s", userId, queryParams.Encode()), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var pricingTiers []warrant.PricingTier
	err = json.Unmarshal([]byte(body), &pricingTiers)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return pricingTiers, nil
}

func ListPricingTiersForUser(userId string, listParams *warrant.ListPricingTierParams) ([]warrant.PricingTier, error) {
	return getClient().ListPricingTiersForUser(userId, listParams)
}

func (c Client) AssignPricingTierToUser(pricingTierId string, userId string) (*warrant.PricingTier, error) {
	resp, err := c.warrantClient.MakeRequest("POST", fmt.Sprintf("/v1/users/%s/pricing-tiers/%s", userId, pricingTierId), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, client.WrapError("Error reading response", err)
	}
	var assignedPricingTier warrant.PricingTier
	err = json.Unmarshal([]byte(body), &assignedPricingTier)
	if err != nil {
		return nil, client.WrapError("Invalid response from server", err)
	}
	return &assignedPricingTier, nil
}

func AssignPricingTierToUser(pricingTierId string, userId string) (*warrant.PricingTier, error) {
	return getClient().AssignPricingTierToUser(pricingTierId, userId)
}

func (c Client) RemovePricingTierFromUser(pricingTierId string, userId string) error {
	_, err := c.warrantClient.MakeRequest("DELETE", fmt.Sprintf("/v1/users/%s/pricing-tiers/%s", userId, pricingTierId), nil)
	if err != nil {
		return err
	}
	return nil
}

func RemovePricingTierFromUser(pricingTierId string, userId string) error {
	return getClient().RemovePricingTierFromUser(pricingTierId, userId)
}

func getClient() Client {
	if warrant.ApiKey == "" {
		panic("You must provide an ApiKey to initialize the Warrant Client")
	}

	config := client.ClientConfig{
		ApiKey:            warrant.ApiKey,
		AuthorizeEndpoint: warrant.AuthorizeEndpoint,
	}

	return Client{
		&client.WarrantClient{
			HttpClient: http.DefaultClient,
			Config:     config,
		},
	}
}
