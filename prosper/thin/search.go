package thin

import "github.com/mtlynch/gofn-prosper/types"

type (
	// SearchFilter specifies a filter for the types of listings to retrieve in
	// the Search function.
	SearchFilter struct {
		EstimatedReturn                           types.Float64Range
		IncomeRange                               []int8
		InquiriesLast6Months                      types.Int32Range
		PriorProsperLoansLatePaymentsOneMonthPlus types.Int32Range
		PriorProsperLoansBalanceOutstanding       types.Float64Range
		DtiWprosperLoan                           types.Float64Range
		ProsperRating                             []string
		ListingStartDate                          types.TimeRange
		ListingStatus                             []int
	}

	// SearchParams specifies parameters to the Search.
	SearchParams struct {
		Offset                  int
		Limit                   int
		ExcludeListingsInvested bool
		Filter                  SearchFilter
	}

	// SearchResponse contains the full response from the Search API in minimally
	// parsed form.
	SearchResponse struct {
		Results     []SearchResult `json:"result"`
		ResultCount int            `json:"result_count"`
		TotalCount  int            `json:"total_count"`
	}
)

// Search queries Prosper for current listings that match specified search
// parameters. Search implements the REST API described at:
// https://developers.prosper.com/docs/investor/searchlistings-api/
func (c Client) Search(p SearchParams) (response SearchResponse, err error) {
	queryString := searchParamsToQueryString(p)
	err = c.DoRequest("GET", c.baseUrl+"/search/listings/?"+queryString, nil, &response)
	if err != nil {
		return SearchResponse{}, err
	}
	return response, nil
}
