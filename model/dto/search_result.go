package dto

type Column struct {
	Name    string `json:"name"`
	Dataset string `json:"dataset"`
}

type SearchResultDto struct {
	itemType string
	item     interface{}
}

// NewBookDto is constructor.
func NewSearchResultDto() *SearchResultDto {
	return &SearchResultDto{}
}
