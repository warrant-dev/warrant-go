package warrant

const ObjectTypePricingTier = "pricing-tier"

type PricingTier struct {
	PricingTierId string                 `json:"pricingTierId"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
}

func (pricingTier PricingTier) GetObjectType() string {
	return "pricing-tier"
}

func (pricingTier PricingTier) GetObjectId() string {
	return pricingTier.PricingTierId
}

type ListPricingTierParams struct {
	ListParams
}

type PricingTierParams struct {
	RequestOptions
	PricingTierId string                 `json:"pricingTierId"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
}
