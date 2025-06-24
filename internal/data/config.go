package data

type Models struct {
	URL              URL
	URLCreateRequest URLCreateRequest
}

func NewModels() *Models {
	return &Models{
		URL:              URL{},
		URLCreateRequest: URLCreateRequest{},
	}
}
