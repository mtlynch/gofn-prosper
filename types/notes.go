package types

import "time"

// DefaultReason describes the reason a note went into default. The values
// correspond to the values of the note_default_reason attribute defined at:
// https://developers.prosper.com/docs/investor/notes-api/
type DefaultReason int64

// Set of possible DefaultReason values.
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

// NoteStatus represents the status of an owned note. The values correspond to
// the values of the note_status attribute defined at:
// https://developers.prosper.com/docs/investor/notes-api/
type NoteStatus int64

// Set of possible NoteStatus values.
const (
	OriginationDelayed     NoteStatus = 0
	Current                NoteStatus = 1
	Chargeoff              NoteStatus = 2
	Defaulted              NoteStatus = 3
	Completed              NoteStatus = 4
	FinalPaymentInProgress NoteStatus = 5 // deprecated
	Cancelled              NoteStatus = 6
	NoteStatusMin          NoteStatus = OriginationDelayed
	NoteStatusMax          NoteStatus = Cancelled
	NoteStatusInvalid      NoteStatus = -1
)

// ProsperRating represents the Prosper-assigned credit rating of a Prosper
// note. The values correspond to the values of prosper_rating defined at:
// https://developers.prosper.com/docs/investor/notes-api/
type ProsperRating int8

// Set of possible ProsperRating values.
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

// Note represents the information about an owned Prosper note, returned by the
// Notes API, described at:
// https://developers.prosper.com/docs/investor/notes-api/
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

// NotesResponse represents the full response from the Notes API, described at:
// https://developers.prosper.com/docs/investor/notes-api/
type NotesResponse struct {
	Result      []Note
	ResultCount int
	TotalCount  int
}
