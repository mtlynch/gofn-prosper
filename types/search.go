package types

type SearchResponse struct {
	Results     []Listing
	ResultCount int
	TotalCount  int
}
