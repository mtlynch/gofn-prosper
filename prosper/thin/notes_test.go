package thin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNotesSuccessfulResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/notes/",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testFormValues(t, r, values{
				"offset": "0",
				"limit":  "25",
			})
			fmt.Fprint(w, `{
   "result": [
   {
      "loan_number": 7735,
      "amount_borrowed": 25000,
      "borrower_rate": 0.1749,
      "prosper_rating": "N/A",
      "term": 36,
      "age_in_months": 100,
      "origination_date": "2014-02-22",
      "days_past_due": 252,
      "principal_balance_pro_rata_share": 42.17326,
      "service_fees_paid_pro_rata_share": -0.13544,
      "principal_paid_pro_rata_share": 7.82672,
      "interest_paid_pro_rata_share": 4.7373,
      "prosper_fees_paid_pro_rata_share": 0,
      "late_fees_paid_pro_rata_share": 0,
      "debt_sale_proceeds_received_pro_rata_share": 0,
      "next_payment_due_amount_pro_rata_share": 0,
      "next_payment_due_date": "2015-07-23",
      "loan_note_id": "7735-205",
      "listing_number": 93874,
      "note_ownership_amount": 50,
      "note_sale_gross_amount_received": 0,
      "note_sale_fees_paid": 0,
      "note_status": 3,
      "note_status_description": "DEFAULTED",
      "note_default_reason": 2,
      "note_default_reason_description": "Bankruptcy",
      "is_sold": false
   },
   {
      "loan_number": 7772,
      "amount_borrowed": 15000,
      "borrower_rate": 0.169,
      "prosper_rating": "N/A",
      "term": 36,
      "age_in_months": 100,
      "origination_date": "2014-02-23",
      "days_past_due": 245,
      "principal_balance_pro_rata_share": 0,
      "service_fees_paid_pro_rata_share": -0.019167,
      "principal_paid_pro_rata_share": 1.1119,
      "interest_paid_pro_rata_share": 0.648233,
      "prosper_fees_paid_pro_rata_share": 0,
      "late_fees_paid_pro_rata_share": 0,
      "debt_sale_proceeds_received_pro_rata_share": 4.270933,
      "next_payment_due_amount_pro_rata_share": 0,
      "next_payment_due_date": "2015-07-23",
      "loan_note_id": "7772-64",
      "listing_number": 92569,
      "note_ownership_amount": 50,
      "note_sale_gross_amount_received": 0,
      "note_sale_fees_paid": 0,
      "note_status": 3,
      "note_status_description": "DEFAULTED",
      "note_default_reason": 1,
      "note_default_reason_description": "Delinquency",
      "is_sold": false
   }
   ],
   "result_count": 2,
   "total_count": 2
}`)
		},
	)

	client := Client{
		baseUrl:      server.URL,
		tokenManager: mockTokenManager{},
	}
	got, err := client.Notes(0, 25)
	if err != nil {
		t.Fatalf("client.Notes failed: %v", err)
	}

	want := NotesResponse{
		Result: []NoteResult{
			{
				LoanNumber:                           7735,
				AmountBorrowed:                       25000,
				BorrowerRate:                         0.1749,
				ProsperRating:                        "N/A",
				Term:                                 36,
				AgeInMonths:                          100,
				OriginationDate:                      "2014-02-22",
				DaysPastDue:                          252,
				PrincipalBalanceProRataShare:         42.17326,
				InterestPaidProRataShare:             4.7373,
				NextPaymentDueDate:                   "2015-07-23",
				DebtSaleProceedsReceivedProRataShare: 0,
				LateFeesPaidProRataShare:             0,
				NextPaymentDueAmountProRataShare:     0,
				PrincipalPaidProRataShare:            7.82672,
				GroupLeaderAward:                     0,
				ServiceFeesPaidProRataShare:          -0.13544,
				ProsperFeesPaidProRataShare:          0,
				LoanNoteId:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           3,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    2,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			{
				LoanNumber:                           7772,
				AmountBorrowed:                       15000,
				BorrowerRate:                         0.169,
				ProsperRating:                        "N/A",
				Term:                                 36,
				AgeInMonths:                          100,
				OriginationDate:                      "2014-02-23",
				DaysPastDue:                          245,
				PrincipalBalanceProRataShare:         0,
				InterestPaidProRataShare:             0.648233,
				NextPaymentDueDate:                   "2015-07-23",
				LateFeesPaidProRataShare:             0,
				DebtSaleProceedsReceivedProRataShare: 4.270933,
				NextPaymentDueAmountProRataShare:     0,
				PrincipalPaidProRataShare:            1.1119,
				GroupLeaderAward:                     0,
				ServiceFeesPaidProRataShare:          -0.019167,
				ProsperFeesPaidProRataShare:          0,
				LoanNoteId:                           "7772-64",
				ListingNumber:                        92569,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           3,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    1,
				NoteDefaultReasonDescription:         "Delinquency",
				IsSold: false,
			},
		},
		ResultCount: 2,
		TotalCount:  2,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("client.Search returned %#v, want %#v", got, want)
	}
}

func TestNotesErrorResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/notes/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mock server error: request failed")
		},
	)
	client := Client{
		baseUrl:      server.URL,
		tokenManager: mockTokenManager{},
	}
	_, err := client.Notes(0, 25)
	if err == nil {
		t.Fatal("client.Notes should fail when server returns error")
	}
}
