package prosper

import (
	"log"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type (
	// SearchFilter specifies a filter for the types of listings to retrieve in
	// the Search function.
	SearchFilter struct {
		EstimatedReturn                           types.Float64Range
		IncomeRange                               []types.IncomeRange
		InquiriesLast6Months                      types.Int32Range
		PriorProsperLoansLatePaymentsOneMonthPlus types.Int32Range
		PriorProsperLoansBalanceOutstanding       types.Float64Range
		DtiWprosperLoan                           types.Float64Range
		ProsperRating                             []types.ProsperRating
		ListingStartDate                          types.TimeRange
		ListingStatus                             []types.ListingStatus
	}

	// SearchParams specifies parameters to the Search.
	SearchParams struct {
		Offset                  int
		Limit                   int
		ExcludeListingsInvested bool
		Filter                  SearchFilter
	}

	// ListingSearcher is an interface that supports the Search API for active
	// Prosper listings.
	ListingSearcher interface {
		Search(SearchParams) (types.SearchResponse, error)
	}
)

// Search queries Prosper for current listings that match specified search
// parameters. Search implements the REST API described at:
// https://developers.prosper.com/docs/investor/searchlistings-api/
func (c Client) Search(p SearchParams) (response types.SearchResponse, err error) {
	rawResponse, err := c.rawClient.Search(searchParamsToThinType(p))
	if err != nil {
		return types.SearchResponse{}, err
	}
	var results []types.Listing
	for _, lRaw := range rawResponse.Results {
		l, err := c.listingParser.Parse(lRaw)
		if err != nil {
			log.Printf("failed to parse listing. err: %v, listing: %+v", err, lRaw)
			return types.SearchResponse{}, err
		}
		results = append(results, l)
	}
	return types.SearchResponse{
		Results:     results,
		ResultCount: rawResponse.ResultCount,
		TotalCount:  rawResponse.TotalCount,
	}, nil
}

func searchParamsToThinType(p SearchParams) thin.SearchParams {
	return thin.SearchParams{
		Offset: p.Offset,
		Limit:  p.Limit,
		ExcludeListingsInvested: p.ExcludeListingsInvested,
		Filter:                  searchFilterToThinType(p.Filter),
	}
}

func searchFilterToThinType(f SearchFilter) thin.SearchFilter {
	incomeRanges := []int8{}
	for _, incomeRange := range f.IncomeRange {
		incomeRanges = append(incomeRanges, int8(incomeRange))
	}
	ratings := []string{}
	for _, rating := range f.ProsperRating {
		ratings = append(ratings, prosperRatingToString(rating))
	}
	listingStatus := []int{}
	for _, status := range f.ListingStatus {
		listingStatus = append(listingStatus, int(status))
	}
	return thin.SearchFilter{
		EstimatedReturn:                           f.EstimatedReturn,
		IncomeRange:                               incomeRanges,
		InquiriesLast6Months:                      f.InquiriesLast6Months,
		PriorProsperLoansLatePaymentsOneMonthPlus: f.PriorProsperLoansLatePaymentsOneMonthPlus,
		PriorProsperLoansBalanceOutstanding:       f.PriorProsperLoansBalanceOutstanding,
		DtiWprosperLoan:                           f.DtiWprosperLoan,
		ProsperRating:                             ratings,
		ListingStartDate:                          f.ListingStartDate,
		ListingStatus:                             listingStatus,
	}
}

func prosperRatingToString(r types.ProsperRating) string {
	ratingToString := map[types.ProsperRating]string{
		types.RatingAA: "AA",
		types.RatingA:  "A",
		types.RatingB:  "B",
		types.RatingC:  "C",
		types.RatingD:  "D",
		types.RatingE:  "E",
		types.RatingHR: "HR",
		types.RatingNA: "N/A",
	}
	s, ok := ratingToString[r]
	if !ok {
		panic("failed to convert prosper rating")
	}
	return s
}
