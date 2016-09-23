package thin

type NoteResult struct {
	LoanNumber                           int64   `json:"loan_number"`
	AmountBorrowed                       float64 `json:"amount_borrowed"`
	BorrowerRate                         float64 `json:"borrower_rate"`
	ProsperRating                        string  `json:"prosper_rating"`
	Term                                 int64   `json:"term"`
	AgeInMonths                          int64   `json:"age_in_months"`
	OriginationDate                      string  `json:"origination_date"`
	DaysPastDue                          int64   `json:"days_past_due"`
	GroupLeaderAward                     int64   `json:"group_leader_award"`
	PrincipalBalanceProRataShare         float64 `json:"principal_balance_pro_rata_share"`
	ServiceFeesPaidProRataShare          float64 `json:"service_fees_paid_pro_rata_share"`
	PrincipalPaidProRataShare            float64 `json:"principal_paid_pro_rata_share"`
	InterestPaidProRataShare             float64 `json:"interest_paid_pro_rata_share"`
	ProsperFeesPaidProRataShare          float64 `json:"prosper_fees_paid_pro_rata_share"`
	LateFeesPaidProRataShare             float64 `json:"late_fees_paid_pro_rata_share"`
	DebtSaleProceedsReceivedProRataShare float64 `json:"debt_sale_proceeds_received_pro_rata_share"`
	NextPaymentDueAmountProRataShare     float64 `json:"next_payment_due_amount_pro_rata_share"`
	NextPaymentDueDate                   string  `json:"next_payment_due_date"`
	LoanNoteId                           string  `json:"loan_note_id"`
	ListingNumber                        int64   `json:"listing_number"`
	NoteOwnershipAmount                  float64 `json:"note_ownership_amount"`
	NoteSaleGrossAmountReceived          float64 `json:"note_sale_gross_amount_received"`
	NoteSaleFeesPaid                     float64 `json:"note_sale_fees_paid"`
	NoteStatus                           int64   `json:"note_status"`
	NoteStatusDescription                string  `json:"note_status_description"`
	NoteDefaultReason                    int64   `json:"note_default_reason"`
	NoteDefaultReasonDescription         string  `json:"note_default_reason_description"`
	IsSold                               bool    `json:"is_sold"`
}
