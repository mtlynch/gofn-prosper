package types

// SearchResponse represents the full response from the Search API, documented
// at: https://developers.prosper.com/docs/investor/searchlistings-api/
type SearchResponse struct {
	Results     []Listing
	ResultCount int
	TotalCount  int
}
