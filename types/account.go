package types

import "time"

type AccountInformation struct {
	AvailableCashBalance                float64
	TotalPrincipalReceivedOnActiveNotes float64
	OutstandingPrincipalOnActiveNotes   float64
	LastWithdrawAmount                  float64
	LastDepositAmount                   float64
	LastDepositDate                     time.Time
	PendingInvestmentsPrimaryMarket     float64
	PendingInvestmentsSecondaryMarket   float64
	PendingQuickInvestOrders            float64
	TotalAmountInvestedOnActiveNotes    float64
	TotalAccountValue                   float64
	InflightGross                       float64
	LastWithdrawDate                    time.Time
}
