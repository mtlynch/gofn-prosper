package types

import (
	"time"
)

type IncomeRange int8

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

type FicoScore int8

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

type ListingStatus int8

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

type ListingNumber int64

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
	InvestmentTypeId                          int64 //TODO: Parse this
	IsHomeowner                               bool
	LastUpdatedDate                           time.Time
	LenderIndicator                           int64 //TODO: Parse this
	LenderYield                               float64
	ListingAmount                             float64
	ListingCategoryId                         int64
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
