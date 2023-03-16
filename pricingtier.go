package warrant

const ObjectTypePricingTier = "pricing-tier"

type PricingTier struct {
	PricingTierId string `json:"pricingTierId"`
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
	PricingTierId string `json:"pricingTierId"`
}
