package prosper

import (
	"log"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

// IncomeRange represents the income range for the borrower associated with a
// loan. Possible values correspond to the values defined under the income_range
// attribute documented at:
// https://developers.prosper.com/docs/investor/searchlistings-api/
type IncomeRange int8

// Set of possible IncomeRange values.
const (
	NotDisplayed       IncomeRange = 0
	ZeroIncome         IncomeRange = 1
	Between0And25k     IncomeRange = 2
	Between25kAnd50k   IncomeRange = 3
	Between50kAnd75k   IncomeRange = 4
	Between75kAnd100k  IncomeRange = 5
	Over100k           IncomeRange = 6
	NotEmployed        IncomeRange = 7
	IncomeRangeMin     IncomeRange = NotDisplayed
	IncomeRangeMax     IncomeRange = NotEmployed
	IncomeRangeInvalid IncomeRange = -1
)

// FicoScore represents the FICO credit score of the borrower associated with a
// loan. Possible values correspond to the values defined under the fico_score
// attribute documented at:
// https://developers.prosper.com/docs/investor/searchlistings-api/
type FicoScore int8

// Set of possible FicoScore values.
const (
	Below600 FicoScore = iota
	Between600And619
	Between620And639
	Between640And659
	Between660And679
	Between680And699
	Between700And719
	Between720And739
	Between740And759
	Between760And779
	Between780And799
	Between800And819
	Between820And850
	FicoScoreInvalid
)

// ListingStatus represents the status of a loan listing. Possible values
// correspond to the values defined under the listing_status attribute
// documented at:
// https://developers.prosper.com/docs/investor/searchlistings-api/
type ListingStatus int8

// Set of possible ListingStatus values.
const (
	ListingActive                    ListingStatus = 2
	ListingWithdrawn                 ListingStatus = 4
	ListingExpired                   ListingStatus = 5
	ListingCompleted                 ListingStatus = 6
	ListingCancelled                 ListingStatus = 7
	ListingPendingReviewOrAcceptance ListingStatus = 8
	ListingStatusMin                 ListingStatus = ListingActive
	ListingStatusMax                 ListingStatus = ListingPendingReviewOrAcceptance
	ListingStatusUnknown             ListingStatus = -1
)

// ListingNumber represents the unique identifier associated with a listing.
type ListingNumber int64

// Listing represents the information associated with a listing, as retrieved
// from the Search API, documented here:
// https://developers.prosper.com/docs/investor/searchlistings-api/
type Listing struct {
	AmountDelinquent                          float64
	AmountFunded                              float64
	AmountParticipation                       float64
	AmountRemaining                           float64
	BankcardUtilization                       float64
	BorrowerApr                               float64
	BorrowerCity                              string
	BorrowerRate                              float64
	BorrowerState                             string
	CreditLinesLast7Years                     int64
	CreditPullDate                            time.Time
	CurrentCreditLines                        int64
	CurrentDelinquencies                      int64
	DelinquenciesLast7Years                   int64
	DelinquenciesOver30Days                   int64
	DelinquenciesOver60Days                   int64
	DelinquenciesOver90Days                   int64
	DtiWprosperLoan                           float64
	EffectiveYield                            float64
	EmploymentStatusDescription               string
	EstimatedLossRate                         float64
	EstimatedReturn                           float64
	FicoScore                                 FicoScore
	FirstRecordedCreditLine                   time.Time
	FundingThreshold                          float64
	IncomeRange                               IncomeRange
	IncomeRangeDescription                    string
	IncomeVerifiable                          bool
	InquiriesLast6Months                      int64
	InstallmentBalance                        float64
	InvestmentTypeDescription                 string
	InvestmentTypeID                          int64 //TODO: Parse this
	IsHomeowner                               bool
	LastUpdatedDate                           time.Time
	LenderIndicator                           int64 //TODO: Parse this
	LenderYield                               float64
	ListingAmount                             float64
	ListingCategoryID                         int64
	ListingCreationDate                       time.Time
	ListingEndDate                            time.Time
	ListingMonthlyPayment                     float64
	ListingNumber                             ListingNumber
	ListingStartDate                          time.Time
	ListingStatus                             ListingStatus
	ListingStatusReason                       string
	ListingTerm                               int64
	ListingTitle                              string
	MaxPriorProsperLoan                       float64
	MemberKey                                 string
	MinPriorProsperLoan                       float64
	MonthlyDebt                               float64
	MonthsEmployed                            int64
	NowDelinquentDerog                        int64
	Occupation                                string
	OldestTradeOpenDate                       time.Time
	OpenCreditLines                           int64
	PartialFundingIndicator                   bool
	PercentFunded                             float64
	PriorProsperLoanEarliestPayOff            int64
	PriorProsperLoans                         int64
	PriorProsperLoans31dpd                    int64
	PriorProsperLoans61dpd                    int64
	PriorProsperLoansActive                   int64
	PriorProsperLoansBalanceOutstanding       float64
	PriorProsperLoansCyclesBilled             int64
	PriorProsperLoansLateCycles               int64
	PriorProsperLoansLatePaymentsOneMonthPlus int64
	PriorProsperLoansOntimePayments           int64
	PriorProsperLoansPrincipalBorrowed        float64
	PriorProsperLoansPrincipalOutstanding     float64
	ProsperRating                             ProsperRating
	ProsperScore                              int64 //TODO: Parse this
	PublicRecordsLast10Years                  int64
	PublicRecordsLast12Months                 int64
	RealEstateBalance                         float64
	RealEstatePayment                         float64
	RevolvingAvailablePercent                 float64
	RevolvingBalance                          float64
	SatisfactoryAccounts                      int64
	ScoreX                                    string //TODO: Parse this
	ScoreXChange                              string //TODO: Parse this
	StatedMonthlyIncome                       float64
	TotalInquiries                            int64
	TotalOpenRevolvingAccounts                int64
	TotalTradeItems                           int64
	VerificationStage                         int64 //TODO: Parse this
	WasDelinquentDerog                        int64
	WholeLoanEndDate                          time.Time
	WholeLoanStartDate                        time.Time
}

type (
	// SearchFilter specifies a filter for the types of listings to retrieve in
	// the Search function.
	SearchFilter struct {
		EstimatedReturn                           types.Float64Range
		IncomeRange                               []IncomeRange
		InquiriesLast6Months                      types.Int32Range
		PriorProsperLoansLatePaymentsOneMonthPlus types.Int32Range
		PriorProsperLoansBalanceOutstanding       types.Float64Range
		DtiWprosperLoan                           types.Float64Range
		ProsperRating                             []ProsperRating
		ListingStartDate                          types.TimeRange
		ListingStatus                             []ListingStatus
	}

	// SearchParams specifies parameters to the Search.
	SearchParams struct {
		Offset                  int
		Limit                   int
		ExcludeListingsInvested bool
		Filter                  SearchFilter
	}

	// SearchResponse represents the full response from the Search API, documented
	// at: https://developers.prosper.com/docs/investor/searchlistings-api/
	SearchResponse struct {
		Results     []Listing
		ResultCount int
		TotalCount  int
	}

	// ListingSearcher is an interface that supports the Search API for active
	// Prosper listings.
	ListingSearcher interface {
		Search(SearchParams) (SearchResponse, error)
	}
)

// Search queries Prosper for current listings that match specified search
// parameters. Search implements the REST API described at:
// https://developers.prosper.com/docs/investor/searchlistings-api/
func (c Client) Search(p SearchParams) (response SearchResponse, err error) {
	rawResponse, err := c.rawClient.Search(searchParamsToThinType(p))
	if err != nil {
		return SearchResponse{}, err
	}
	var results []Listing
	for _, lRaw := range rawResponse.Results {
		l, err := c.listingParser.Parse(lRaw)
		if err != nil {
			log.Printf("failed to parse listing. err: %v, listing: %+v", err, lRaw)
			return SearchResponse{}, err
		}
		results = append(results, l)
	}
	return SearchResponse{
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

func prosperRatingToString(r ProsperRating) string {
	ratingToString := map[ProsperRating]string{
		RatingAA: "AA",
		RatingA:  "A",
		RatingB:  "B",
		RatingC:  "C",
		RatingD:  "D",
		RatingE:  "E",
		RatingHR: "HR",
		RatingNA: "N/A",
	}
	s, ok := ratingToString[r]
	if !ok {
		panic("failed to convert prosper rating")
	}
	return s
}
