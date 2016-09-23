package types

import "time"

type DefaultReason int64

const (
	Delinquency      DefaultReason = 1
	Bankruptcy       DefaultReason = 2
	Deceased         DefaultReason = 3
	Repurchased      DefaultReason = 4
	PaidInFull       DefaultReason = 5
	SettledInFull    DefaultReason = 6
	Sold             DefaultReason = 7
	DefaultReasonMin DefaultReason = Delinquency
	DefaultReasonMax DefaultReason = Sold
)

type NoteStatus int64

const (
	// deprecated
	OriginationDelayed NoteStatus = 0
	Current            NoteStatus = 1
	Chargeoff          NoteStatus = 2
	Defaulted          NoteStatus = 3
	Completed          NoteStatus = 4
	// deprecated
	//FinalPaymentInProgress NoteStatus = 5
	Cancelled         NoteStatus = 6
	NoteStatusMin     NoteStatus = OriginationDelayed
	NoteStatusMax     NoteStatus = Cancelled
	NoteStatusInvalid NoteStatus = -1
)

type ProsperRating int8

const (
	RatingAA ProsperRating = iota
	RatingA
	RatingB
	RatingC
	RatingD
	RatingE
	RatingHR
	RatingNA
)

type Note struct {
	AgeInMonths                          int64
	AmountBorrowed                       float64
	BorrowerRate                         float64
	DaysPastDue                          int64
	DebtSaleProceedsReceivedProRataShare float64
	InterestPaidProRataShare             float64
	IsSold                               bool
	LateFeesPaidProRataShare             float64
	ListingNumber                        ListingNumber
	LoanNoteId                           string
	LoanNumber                           int64
	NextPaymentDueAmountProRataShare     float64
	NextPaymentDueDate                   time.Time
	NoteDefaultReasonDescription         string
	NoteDefaultReason                    *DefaultReason
	NoteOwnershipAmount                  float64
	NoteSaleFeesPaid                     float64
	NoteSaleGrossAmountReceived          float64
	NoteStatusDescription                string
	NoteStatus                           NoteStatus
	OriginationDate                      time.Time
	PrincipalBalanceProRataShare         float64
	PrincipalPaidProRataShare            float64
	ProsperFeesPaidProRataShare          float64
	ProsperRating                        ProsperRating
	ServiceFeesPaidProRataShare          float64
	Term                                 int64
}

type NotesResponse struct {
	Result      []Note
	ResultCount int
	TotalCount  int
}
