package warrant

type Feature struct {
	FeatureId string `json:"featureId"`
}

type ListFeatureParams struct {
	ListParams
}

type FeatureParams struct {
	FeatureId string `json:"featureId"`
}
