package warrant

const ObjectTypeFeature = "feature"

type Feature struct {
	FeatureId string `json:"featureId"`
}

func (feature Feature) GetObjectType() string {
	return "feature"
}

func (feature Feature) GetObjectId() string {
	return feature.FeatureId
}

type ListFeatureParams struct {
	ListParams
}

type FeatureParams struct {
	RequestOptions
	FeatureId string `json:"featureId"`
}
