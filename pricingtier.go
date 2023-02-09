package warrant

type PricingTier struct {
	PricingTierId string `json:"pricingTierId"`
}

type ListPricingTierParams struct {
	ListParams
}

type PricingTierParams struct {
	PricingTierId string `json:"pricingTierId"`
}
