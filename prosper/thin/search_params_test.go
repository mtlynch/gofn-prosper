package thin

import (
	"testing"
	"time"

	"github.com/mtlynch/gofn-prosper/interval"
)

func TestSearchParamsToQueryString(t *testing.T) {
	var tests = []struct {
		p    SearchParams
		want string
	}{
		{
			p: SearchParams{
				Offset: 25,
			},
			want: "offset=25",
		},
		{
			p: SearchParams{
				Limit: 46,
			},
			want: "limit=46",
		},
		{
			p: SearchParams{
				ExcludeListingsInvested: true,
			},
			want: "exclude_listings_invested=true",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					EstimatedReturn: interval.Float64Range{
						Min: interval.CreateFloat64(0.05),
					},
				},
			},
			want: "estimated_return_min=0.0500",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					EstimatedReturn: interval.Float64Range{
						Min: interval.CreateFloat64(0.05),
						Max: interval.CreateFloat64(0.07),
					},
				},
			},
			want: "estimated_return_min=0.0500&estimated_return_max=0.0700",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					IncomeRange: []int8{2},
				},
			},
			want: "income_range=2",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					IncomeRange: []int8{2},
				},
			},
			want: "income_range=2",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					IncomeRange: []int8{2, 3},
				},
			},
			want: "income_range=2,3",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					IncomeRange: []int8{2, 3, 8},
				},
			},
			want: "income_range=2,3,8",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					InquiriesLast6Months: interval.Int32Range{
						Min: interval.CreateInt32(3),
					},
				},
			},
			want: "inquiries_last6_months_min=3",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					InquiriesLast6Months: interval.Int32Range{
						Max: interval.CreateInt32(2),
					},
				},
			},
			want: "inquiries_last6_months_max=2",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					PriorProsperLoansLatePaymentsOneMonthPlus: interval.Int32Range{
						Min: interval.CreateInt32(3),
					},
				},
			},
			want: "prior_prosper_loans_late_payments_one_month_plus_min=3",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					PriorProsperLoansLatePaymentsOneMonthPlus: interval.Int32Range{
						Max: interval.CreateInt32(5),
					},
				},
			},
			want: "prior_prosper_loans_late_payments_one_month_plus_max=5",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					PriorProsperLoansBalanceOutstanding: interval.Float64Range{
						Min: interval.CreateFloat64(3.2),
					},
				},
			},
			want: "prior_prosper_loans_balance_outstanding_min=3.2000",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					PriorProsperLoansBalanceOutstanding: interval.Float64Range{
						Max: interval.CreateFloat64(5.8),
					},
				},
			},
			want: "prior_prosper_loans_balance_outstanding_max=5.8000",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					DtiWprosperLoan: interval.Float64Range{
						Min: interval.CreateFloat64(0.04),
					},
				},
			},
			want: "dti_wprosper_loan_min=0.0400",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					DtiWprosperLoan: interval.Float64Range{
						Max: interval.CreateFloat64(0.04),
					},
				},
			},
			want: "dti_wprosper_loan_max=0.0400",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					Rating: []string{"AA"},
				},
			},
			want: "prosper_rating=AA",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					Rating: []string{"A", "C", "E"},
				},
			},
			want: "prosper_rating=A,C,E",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					ListingStatus: []int{2},
				},
			},
			want: "listing_status=2",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					ListingStatus: []int{2, 6, 7},
				},
			},
			want: "listing_status=2,6,7",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					ListingStartDate: interval.TimeRange{
						Min: interval.CreateTime(time.Date(2016, 2, 28, 11, 46, 5, 0, time.UTC)),
					},
				},
			},
			want: "listing_start_date_min=2016-02-28+11:46:05",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					ListingStartDate: interval.TimeRange{
						Max: interval.CreateTime(time.Date(2016, 2, 28, 11, 46, 5, 0, time.UTC)),
					},
				},
			},
			want: "listing_start_date_max=2016-02-28+11:46:05",
		},
		{
			p: SearchParams{
				Filter: SearchFilter{
					ListingStartDate: interval.TimeRange{
						Min: interval.CreateTime(time.Date(2016, 2, 28, 11, 46, 5, 0, time.UTC)),
						Max: interval.CreateTime(time.Date(2016, 2, 29, 11, 46, 5, 0, time.UTC)),
					},
				},
			},
			want: "listing_start_date_min=2016-02-28+11:46:05&listing_start_date_max=2016-02-29+11:46:05",
		},
	}
	for _, tt := range tests {
		got := searchParamsToQueryString(tt.p)
		if got != tt.want {
			t.Errorf("searchParamsToQueryString() got: %v, want: %v", got, tt.want)
		}
	}

}
