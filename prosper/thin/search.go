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

	// SearchResult contains response information about a single Prosper listing
	// in minimally parsed form.
	SearchResult struct {
		AmountDelinquent                          float64 `json:"amount_delinquent"`
		AmountFunded                              float64 `json:"amount_funded"`
		AmountParticipation                       float64 `json:"amount_participation"`
		AmountRemaining                           float64 `json:"amount_remaining"`
		BankcardUtilization                       float64 `json:"bankcard_utilization"`
		BorrowerApr                               float64 `json:"borrower_apr"`
		BorrowerCity                              string  `json:"borrower_city"`
		BorrowerListingDescription                string  `json:"borrower_listing_description"`
		BorrowerMetropolitanArea                  string  `json:"borrower_metropolitan_area"`
		BorrowerRate                              float64 `json:"borrower_rate"`
		BorrowerState                             string  `json:"borrower_state"`
		ChannelCode                               string  `json:"channel_code"`
		CreditLinesLast7Years                     int64   `json:"credit_lines_last7_years"`
		CreditPullDate                            string  `json:"credit_pull_date"`
		CurrentCreditLines                        int64   `json:"current_credit_lines"`
		CurrentDelinquencies                      int64   `json:"current_delinquencies"`
		DelinquenciesLast7Years                   int64   `json:"delinquencies_last7_years"`
		DelinquenciesOver30Days                   int64   `json:"delinquencies_over30_days"`
		DelinquenciesOver60Days                   int64   `json:"delinquencies_over60_days"`
		DelinquenciesOver90Days                   int64   `json:"delinquencies_over90_days"`
		DtiWprosperLoan                           float64 `json:"dti_wprosper_loan"`
		EffectiveYield                            float64 `json:"effective_yield"`
		EmploymentStatusDescription               string  `json:"employment_status_description"`
		EstimatedLossRate                         float64 `json:"estimated_loss_rate"`
		EstimatedReturn                           float64 `json:"estimated_return"`
		FicoScore                                 string  `json:"fico_score"`
		FirstRecordedCreditLine                   string  `json:"first_recorded_credit_line"`
		FundingThreshold                          float64 `json:"funding_threshold"`
		GroupIndicator                            bool    `json:"group_indicator"`
		GroupName                                 string  `json:"group_name"`
		IncomeRange                               int64   `json:"income_range"`
		IncomeRangeDescription                    string  `json:"income_range_description"`
		IncomeVerifiable                          bool    `json:"income_verifiable"`
		InquiriesLast6Months                      int64   `json:"inquiries_last6_months"`
		InstallmentBalance                        float64 `json:"installment_balance"`
		InvestmentTypeDescription                 string  `json:"investment_type_description"`
		InvestmentTypeid                          int64   `json:"investment_typeid"`
		IsHomeowner                               bool    `json:"is_homeowner"`
		LastUpdatedDate                           string  `json:"last_updated_date"`
		LenderIndicator                           int64   `json:"lender_indicator"`
		LenderYield                               float64 `json:"lender_yield"`
		ListingAmount                             float64 `json:"listing_amount"`
		ListingCategoryId                         int64   `json:"listing_category_id"`
		ListingCreationDate                       string  `json:"listing_creation_date"`
		ListingEndDate                            string  `json:"listing_end_date"`
		ListingMonthlyPayment                     float64 `json:"listing_monthly_payment"`
		ListingNumber                             int64   `json:"listing_number"`
		ListingPurpose                            string  `json:"listing_purpose"`
		ListingStartDate                          string  `json:"listing_start_date"`
		ListingStatus                             int64   `json:"listing_status"`
		ListingStatusReason                       string  `json:"listing_status_reason"`
		ListingTerm                               int64   `json:"listing_term"`
		ListingTitle                              string  `json:"listing_title"`
		MaxPriorProsperLoan                       float64 `json:"max_prior_prosper_loan"`
		MemberKey                                 string  `json:"member_key"`
		MinPriorProsperLoan                       float64 `json:"min_prior_prosper_loan"`
		MonthlyDebt                               float64 `json:"monthly_debt"`
		MonthsEmployed                            int64   `json:"months_employed"`
		NowDelinquentDerog                        int64   `json:"now_delinquent_derog"`
		Occupation                                string  `json:"occupation"`
		OldestTradeOpenDate                       string  `json:"oldest_trade_open_date"`
		OpenCreditLines                           int64   `json:"open_credit_lines"`
		PartialFundingIndicator                   bool    `json:"partial_funding_indicator"`
		PercentFunded                             float64 `json:"percent_funded"`
		PriorProsperLoans                         int64   `json:"prior_prosper_loans"`
		PriorProsperLoanEarliestPayOff            int64   `json:"prior_prosper_loan_earliest_pay_off"`
		PriorProsperLoans31dpd                    int64   `json:"prior_prosper_loans31dpd"`
		PriorProsperLoans61dpd                    int64   `json:"prior_prosper_loans61dpd"`
		PriorProsperLoansActive                   int64   `json:"prior_prosper_loans_active"`
		PriorProsperLoansBalanceOutstanding       float64 `json:"prior_prosper_loans_balance_outstanding"`
		PriorProsperLoansCyclesBilled             int64   `json:"prior_prosper_loans_cycles_billed"`
		PriorProsperLoansLateCycles               int64   `json:"prior_prosper_loans_late_cycles"`
		PriorProsperLoansLatePaymentsOneMonthPlus int64   `json:"prior_prosper_loans_late_payments_one_month_plus"`
		PriorProsperLoansOntimePayments           int64   `json:"prior_prosper_loans_ontime_payments"`
		PriorProsperLoansPrincipalBorrowed        float64 `json:"prior_prosper_loans_principal_borrowed"`
		PriorProsperLoansPrincipalOutstanding     float64 `json:"prior_prosper_loans_principal_outstanding"`
		ProsperRating                             string  `json:"prosper_rating"`
		ProsperScore                              int64   `json:"prosper_score"`
		PublicRecordsLast10Years                  int64   `json:"public_records_last10_years"`
		PublicRecordsLast12Months                 int64   `json:"public_records_last12_months"`
		RealEstateBalance                         float64 `json:"real_estate_balance"`
		RealEstatePayment                         float64 `json:"real_estate_payment"`
		RevolvingAvailablePercent                 float64 `json:"revolving_available_percent"`
		RevolvingBalance                          float64 `json:"revolving_balance"`
		SatisfactoryAccounts                      int64   `json:"satisfactory_accounts"`
		Scorex                                    string  `json:"scorex"`
		ScorexChange                              string  `json:"scorex_change"`
		StatedMonthlyIncome                       float64 `json:"stated_monthly_income"`
		TotalInquiries                            int64   `json:"total_inquiries"`
		TotalOpenRevolvingAccounts                int64   `json:"total_open_revolving_accounts"`
		TotalTradeItems                           int64   `json:"total_trade_items"`
		VerificationStage                         int64   `json:"verification_stage"`
		WasDelinquentDerog                        int64   `json:"was_delinquent_derog"`
		WholeLoanEndDate                          string  `json:"whole_loan_end_date"`
		WholeLoanStartDate                        string  `json:"whole_loan_start_date"`
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
