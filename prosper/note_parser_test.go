package prosper

import (
	"reflect"
	"testing"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

func TestNoteParser(t *testing.T) {
	defaultReasonBankruptcy := types.Bankruptcy
	var tests = []struct {
		input         thin.NoteResult
		want          types.Note
		expectSuccess bool
		msg           string
	}{
		{
			input: thin.NoteResult{
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
				LoanNoteID:                           "7735-205",
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
			want: types.Note{
				LoanNumber:                           7735,
				AmountBorrowed:                       25000,
				BorrowerRate:                         0.1749,
				ProsperRating:                        types.RatingNA,
				Term:                                 36,
				AgeInMonths:                          100,
				OriginationDate:                      time.Date(2014, 02, 22, 0, 0, 0, 0, time.UTC),
				DaysPastDue:                          252,
				PrincipalBalanceProRataShare:         42.17326,
				InterestPaidProRataShare:             4.7373,
				NextPaymentDueDate:                   time.Date(2015, 07, 23, 0, 0, 0, 0, time.UTC),
				DebtSaleProceedsReceivedProRataShare: 0,
				LateFeesPaidProRataShare:             0,
				NextPaymentDueAmountProRataShare:     0,
				PrincipalPaidProRataShare:            7.82672,
				ServiceFeesPaidProRataShare:          -0.13544,
				ProsperFeesPaidProRataShare:          0,
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           types.Defaulted,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    &defaultReasonBankruptcy,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			expectSuccess: true,
			msg:           "valid note should parse successfully",
		},
		{
			input: thin.NoteResult{
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
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           3,
				NoteStatusDescription:                "DEFAULTED",
				IsSold:                               false,
			},
			want: types.Note{
				LoanNumber:                           7735,
				AmountBorrowed:                       25000,
				BorrowerRate:                         0.1749,
				ProsperRating:                        types.RatingNA,
				Term:                                 36,
				AgeInMonths:                          100,
				OriginationDate:                      time.Date(2014, 02, 22, 0, 0, 0, 0, time.UTC),
				DaysPastDue:                          252,
				PrincipalBalanceProRataShare:         42.17326,
				InterestPaidProRataShare:             4.7373,
				NextPaymentDueDate:                   time.Date(2015, 07, 23, 0, 0, 0, 0, time.UTC),
				DebtSaleProceedsReceivedProRataShare: 0,
				LateFeesPaidProRataShare:             0,
				NextPaymentDueAmountProRataShare:     0,
				PrincipalPaidProRataShare:            7.82672,
				ServiceFeesPaidProRataShare:          -0.13544,
				ProsperFeesPaidProRataShare:          0,
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           types.Defaulted,
				NoteStatusDescription:                "DEFAULTED",
				IsSold:                               false,
			},
			expectSuccess: true,
			msg:           "missing default reason fields should parse successfully",
		},
		{
			input: thin.NoteResult{
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
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           3,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    999999,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			expectSuccess: false,
			msg:           "invalid NoteDefaultReason should cause error",
		},
		{
			input: thin.NoteResult{
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
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           99999999,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    2,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			expectSuccess: false,
			msg:           "invalid NoteStatus should cause error",
		},
		{
			input: thin.NoteResult{
				LoanNumber:                           7735,
				AmountBorrowed:                       25000,
				BorrowerRate:                         0.1749,
				ProsperRating:                        "invalid rating",
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
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           5,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    2,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			expectSuccess: false,
			msg:           "invalid ProsperRating should cause error",
		},
		{
			input: thin.NoteResult{
				LoanNumber:                           7735,
				AmountBorrowed:                       25000,
				BorrowerRate:                         0.1749,
				ProsperRating:                        "N/A",
				Term:                                 36,
				AgeInMonths:                          100,
				OriginationDate:                      "invalid origination date",
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
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           5,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    2,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			expectSuccess: false,
			msg:           "invalid OriginationDate should cause error",
		},
		{
			input: thin.NoteResult{
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
				NextPaymentDueDate:                   "invalid next payment date",
				DebtSaleProceedsReceivedProRataShare: 0,
				LateFeesPaidProRataShare:             0,
				NextPaymentDueAmountProRataShare:     0,
				PrincipalPaidProRataShare:            7.82672,
				GroupLeaderAward:                     0,
				ServiceFeesPaidProRataShare:          -0.13544,
				ProsperFeesPaidProRataShare:          0,
				LoanNoteID:                           "7735-205",
				ListingNumber:                        93874,
				NoteOwnershipAmount:                  50,
				NoteSaleGrossAmountReceived:          0,
				NoteSaleFeesPaid:                     0,
				NoteStatus:                           5,
				NoteStatusDescription:                "DEFAULTED",
				NoteDefaultReason:                    2,
				NoteDefaultReasonDescription:         "Bankruptcy",
				IsSold: false,
			},
			expectSuccess: false,
			msg:           "invalid NextPaymentDueDate should cause error",
		},
	}
	for _, tt := range tests {
		got, err := defaultNoteParser{}.Parse(tt.input)
		if tt.expectSuccess && err != nil {
			t.Errorf("%s - expected successful parsing of %+v, got error: %v", tt.msg, tt.input, err)
		} else if !tt.expectSuccess && err == nil {
			t.Errorf("%s - expected failure for %+v, got nil", tt.msg, tt.input)
		}
		if tt.expectSuccess && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s - noteParser.Parse returned %#v, want %#v", tt.msg, got, tt.want)
		}
	}
}
