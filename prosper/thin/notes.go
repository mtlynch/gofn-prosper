package thin

import (
	"fmt"
	"strings"
)

type (
	// NotesParams contains the parameters to the Notes API.
	NotesParams struct {
		Offset int
		Limit  int
		// TODO(mtlynch): Implement support for the sort_by parameter.
	}

	// NoteResult contains response information about a single Propser note in
	// minimally parsed form.
	NoteResult struct {
		LoanNumber                           int64   `json:"loan_number"`
		AmountBorrowed                       float64 `json:"amount_borrowed"`
		BorrowerRate                         float64 `json:"borrower_rate"`
		Rating                               string  `json:"prosper_rating"`
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
		LoanNoteID                           string  `json:"loan_note_id"`
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

	// NotesResponse contains the full response from the Notes API in minimally
	// parsed form.
	NotesResponse struct {
		Result      []NoteResult `json:"result"`
		ResultCount int          `json:"result_count"`
		TotalCount  int          `json:"total_count"`
	}
)

// Notes returns a subset of the notes that the user owns. Notes partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/notes-api/
func (c defaultClient) Notes(p NotesParams) (response NotesResponse, err error) {
	q := notesParamsToQueryString(p)
	url := fmt.Sprintf("%s/notes/?%s", c.baseURL, q)
	err = c.DoRequest("GET", url, nil, &response)
	if err != nil {
		return NotesResponse{}, err
	}
	return response, nil
}

func notesParamsToQueryString(p NotesParams) string {
	var clauses []string
	if p.Offset != 0 {
		clauses = append(clauses, fmt.Sprintf("offset=%d", p.Offset))
	}
	if p.Limit != 0 {
		clauses = append(clauses, fmt.Sprintf("limit=%d", p.Limit))
	}
	return strings.Join(clauses, "&")
}
