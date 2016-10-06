package prosper

import (
	"fmt"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

type listingParser interface {
	Parse(thin.SearchResult) (Listing, error)
}

type defaultListingParser struct{}

func (p defaultListingParser) Parse(r thin.SearchResult) (Listing, error) {
	incomeRange, err := parseIncomeRange(r.IncomeRange)
	if err != nil {
		return Listing{}, err
	}
	listingStatus, err := parseListingStatus(r.ListingStatus)
	if err != nil {
		return Listing{}, err
	}
	ficoScore, err := parseFicoScore(r.FicoScore)
	if err != nil {
		return Listing{}, err
	}
	rating, err := parseRating(r.Rating)
	if err != nil {
		return Listing{}, err
	}
	oldestTradeOpenDate, err := parseProsperOldTime(r.OldestTradeOpenDate)
	if err != nil {
		return Listing{}, err
	}
	firstRecordedCreditLine, err := parseProsperTime(r.FirstRecordedCreditLine)
	if err != nil {
		return Listing{}, err
	}
	creditPullDate, err := parseProsperTime(r.CreditPullDate)
	if err != nil {
		return Listing{}, err
	}
	listingCreationDate, err := parseProsperTime(r.ListingCreationDate)
	if err != nil {
		return Listing{}, err
	}
	listingEndDate, err := parseProsperTime(r.ListingEndDate)
	if err != nil {
		return Listing{}, err
	}
	listingStartDate, err := parseProsperTime(r.ListingStartDate)
	if err != nil {
		return Listing{}, err
	}
	wholeLoanStartDate, err := parseProsperTime(r.WholeLoanStartDate)
	if err != nil {
		return Listing{}, err
	}
	wholeLoanEndDate, err := parseProsperTime(r.WholeLoanEndDate)
	if err != nil {
		return Listing{}, err
	}
	lastUpdatedDate, err := parseProsperTime(r.LastUpdatedDate)
	if err != nil {
		return Listing{}, err
	}
	return Listing{
		PriorProsperLoans:                         r.PriorProsperLoans,
		AmountDelinquent:                          r.AmountDelinquent,
		AmountParticipation:                       r.AmountParticipation,
		DelinquenciesOver60Days:                   r.DelinquenciesOver60Days,
		IncomeRange:                               incomeRange,
		ListingMonthlyPayment:                     r.ListingMonthlyPayment,
		OldestTradeOpenDate:                       oldestTradeOpenDate,
		PriorProsperLoansPrincipalOutstanding:     r.PriorProsperLoansPrincipalOutstanding,
		PublicRecordsLast12Months:                 r.PublicRecordsLast12Months,
		TotalOpenRevolvingAccounts:                r.TotalOpenRevolvingAccounts,
		VerificationStage:                         r.VerificationStage,
		ListingStatus:                             listingStatus,
		ListingTitle:                              r.ListingTitle,
		DelinquenciesLast7Years:                   r.DelinquenciesLast7Years,
		EmploymentStatusDescription:               r.EmploymentStatusDescription,
		TotalTradeItems:                           r.TotalTradeItems,
		BorrowerRate:                              r.BorrowerRate,
		IsHomeowner:                               r.IsHomeowner,
		LastUpdatedDate:                           lastUpdatedDate,
		ListingAmount:                             r.ListingAmount,
		ListingNumber:                             ListingNumber(r.ListingNumber),
		PriorProsperLoans61dpd:                    r.PriorProsperLoans61dpd,
		PriorProsperLoansPrincipalBorrowed:        r.PriorProsperLoansPrincipalBorrowed,
		WasDelinquentDerog:                        r.WasDelinquentDerog,
		BankcardUtilization:                       r.BankcardUtilization,
		InstallmentBalance:                        r.InstallmentBalance,
		InvestmentTypeID:                          r.InvestmentTypeid,
		ListingCategoryID:                         r.ListingCategoryID,
		BorrowerCity:                              r.BorrowerCity,
		BorrowerState:                             r.BorrowerState,
		IncomeRangeDescription:                    r.IncomeRangeDescription,
		RevolvingAvailablePercent:                 r.RevolvingAvailablePercent,
		CurrentCreditLines:                        r.CurrentCreditLines,
		DtiWprosperLoan:                           r.DtiWprosperLoan,
		FicoScore:                                 ficoScore,
		FirstRecordedCreditLine:                   firstRecordedCreditLine,
		RealEstateBalance:                         r.RealEstateBalance,
		SatisfactoryAccounts:                      r.SatisfactoryAccounts,
		InquiriesLast6Months:                      r.InquiriesLast6Months,
		LenderYield:                               r.LenderYield,
		MemberKey:                                 r.MemberKey,
		PriorProsperLoanEarliestPayOff:            r.PriorProsperLoanEarliestPayOff,
		PriorProsperLoansCyclesBilled:             r.PriorProsperLoansCyclesBilled,
		CurrentDelinquencies:                      r.CurrentDelinquencies,
		DelinquenciesOver30Days:                   r.DelinquenciesOver30Days,
		InvestmentTypeDescription:                 r.InvestmentTypeDescription,
		ListingStatusReason:                       r.ListingStatusReason,
		MonthlyDebt:                               r.MonthlyDebt,
		MonthsEmployed:                            r.MonthsEmployed,
		PartialFundingIndicator:                   r.PartialFundingIndicator,
		Rating:                                    rating,
		ProsperScore:                              r.ProsperScore,
		BorrowerApr:                               r.BorrowerApr,
		PriorProsperLoans31dpd:                    r.PriorProsperLoans31dpd,
		PriorProsperLoansLatePaymentsOneMonthPlus: r.PriorProsperLoansLatePaymentsOneMonthPlus,
		FundingThreshold:                          r.FundingThreshold,
		RealEstatePayment:                         r.RealEstatePayment,
		CreditLinesLast7Years:                     r.CreditLinesLast7Years,
		CreditPullDate:                            creditPullDate,
		PublicRecordsLast10Years:                  r.PublicRecordsLast10Years,
		RevolvingBalance:                          r.RevolvingBalance,
		ScoreX:                                    r.Scorex,
		ScoreXChange:                              r.ScorexChange,
		MinPriorProsperLoan:                       r.MinPriorProsperLoan,
		AmountFunded:                              r.AmountFunded,
		EffectiveYield:                            r.EffectiveYield,
		EstimatedLossRate:                         r.EstimatedLossRate,
		ListingCreationDate:                       listingCreationDate,
		ListingEndDate:                            listingEndDate,
		ListingStartDate:                          listingStartDate,
		Occupation:                                r.Occupation,
		PercentFunded:                             r.PercentFunded,
		PriorProsperLoansBalanceOutstanding:       r.PriorProsperLoansBalanceOutstanding,
		AmountRemaining:                           r.AmountRemaining,
		DelinquenciesOver90Days:                   r.DelinquenciesOver90Days,
		OpenCreditLines:                           r.OpenCreditLines,
		PriorProsperLoansActive:                   r.PriorProsperLoansActive,
		PriorProsperLoansLateCycles:               r.PriorProsperLoansLateCycles,
		ListingTerm:                               r.ListingTerm,
		PriorProsperLoansOntimePayments:           r.PriorProsperLoansOntimePayments,
		EstimatedReturn:                           r.EstimatedReturn,
		IncomeVerifiable:                          r.IncomeVerifiable,
		LenderIndicator:                           r.LenderIndicator,
		MaxPriorProsperLoan:                       r.MaxPriorProsperLoan,
		NowDelinquentDerog:                        r.NowDelinquentDerog,
		StatedMonthlyIncome:                       r.StatedMonthlyIncome,
		TotalInquiries:                            r.TotalInquiries,
		WholeLoanStartDate:                        wholeLoanStartDate,
		WholeLoanEndDate:                          wholeLoanEndDate,
	}, nil
}

func parseIncomeRange(incomeRange int64) (IncomeRange, error) {
	if incomeRange < int64(IncomeRangeMin) || incomeRange > int64(IncomeRangeMax) {
		return IncomeRangeInvalid, fmt.Errorf("income range out of range: %d, expected %d-%d", incomeRange, IncomeRangeMin, IncomeRangeMax)
	}
	return IncomeRange(incomeRange), nil
}

func parseListingStatus(listingStatus int64) (ListingStatus, error) {
	if listingStatus < int64(ListingStatusMin) || listingStatus > int64(ListingStatusMax) {
		return ListingStatusUnknown, fmt.Errorf("listing status out of range: %d, expected %d-%d", listingStatus, ListingStatusMin, ListingStatusMax)
	}
	return ListingStatus(listingStatus), nil
}

func parseFicoScore(ficoScore string) (FicoScore, error) {
	stringToScore := map[string]FicoScore{
		"<600":    Below600,
		"600-619": Between600And619,
		"620-639": Between620And639,
		"640-659": Between640And659,
		"660-679": Between660And679,
		"680-699": Between680And699,
		"700-719": Between700And719,
		"720-739": Between720And739,
		"740-759": Between740And759,
		"760-779": Between760And779,
		"780-799": Between780And799,
		"800-819": Between800And819,
		"820-850": Between820And850,
	}
	parsed, ok := stringToScore[ficoScore]
	if !ok {
		return FicoScoreInvalid, fmt.Errorf("unrecognized fico score: %s", ficoScore)
	}
	return parsed, nil
}
