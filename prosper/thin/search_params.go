package thin

import (
	"fmt"
	"strings"
	"time"

	"github.com/mtlynch/gofn-prosper/types"
)

func stringsToClauseValues(name string, s []string) string {
	joined := strings.Join(s, ",")
	return fmt.Sprintf("%s=%s", name, joined)
}

func intsToClauseValues(name string, ints []int) string {
	converted := []string{}
	for _, val := range ints {
		converted = append(converted, fmt.Sprintf("%d", val))
	}
	return stringsToClauseValues(name, converted)
}

func float64RangeToClauses(name string, r types.Float64Range) (clauses []string) {
	if r.Min != nil {
		clauses = append(clauses, fmt.Sprintf("%s_min=%.4f", name, *r.Min))
	}
	if r.Max != nil {
		clauses = append(clauses, fmt.Sprintf("%s_max=%.4f", name, *r.Max))
	}
	return clauses
}

func int32RangeToClauses(name string, r types.Int32Range) (clauses []string) {
	if r.Min != nil {
		clauses = append(clauses, fmt.Sprintf("%s_min=%d", name, *r.Min))
	}
	if r.Max != nil {
		clauses = append(clauses, fmt.Sprintf("%s_max=%d", name, *r.Max))
	}
	return clauses
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02+15:04:05")
}

func timeRangeToClauses(name string, r types.TimeRange) (clauses []string) {
	if r.Min != nil {
		clauses = append(clauses, fmt.Sprintf("%s_min=%s", name, formatTime(*r.Min)))
	}
	if r.Max != nil {
		clauses = append(clauses, fmt.Sprintf("%s_max=%s", name, formatTime(*r.Max)))
	}
	return clauses
}

func searchParamsToQueryString(p SearchParams) string {
	var clauses []string
	if p.Offset != 0 {
		clauses = append(clauses, fmt.Sprintf("offset=%d", p.Offset))
	}
	if p.Limit != 0 {
		clauses = append(clauses, fmt.Sprintf("limit=%d", p.Limit))
	}
	if p.ExcludeListingsInvested {
		clauses = append(clauses, "exclude_listings_invested=true")
	}
	if len(p.Filter.IncomeRange) > 0 {
		var rangeValues []string
		for _, v := range p.Filter.IncomeRange {
			rangeValues = append(rangeValues, fmt.Sprintf("%d", v))
		}
		clauses = append(clauses, stringsToClauseValues("income_range", rangeValues))
	}
	if len(p.Filter.Rating) > 0 {
		clauses = append(clauses, stringsToClauseValues("prosper_rating", p.Filter.Rating))
	}
	if len(p.Filter.ListingStatus) > 0 {
		clauses = append(clauses, intsToClauseValues("listing_status", p.Filter.ListingStatus))
	}
	clauses = append(clauses, float64RangeToClauses("estimated_return", p.Filter.EstimatedReturn)...)
	clauses = append(clauses, int32RangeToClauses("inquiries_last6_months", p.Filter.InquiriesLast6Months)...)
	clauses = append(clauses, int32RangeToClauses("prior_prosper_loans_late_payments_one_month_plus", p.Filter.PriorProsperLoansLatePaymentsOneMonthPlus)...)
	clauses = append(clauses, float64RangeToClauses("prior_prosper_loans_balance_outstanding", p.Filter.PriorProsperLoansBalanceOutstanding)...)
	clauses = append(clauses, float64RangeToClauses("dti_wprosper_loan", p.Filter.DtiWprosperLoan)...)
	clauses = append(clauses, timeRangeToClauses("listing_start_date", p.Filter.ListingStartDate)...)

	return strings.Join(clauses, "&")
}
