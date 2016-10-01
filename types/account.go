// Package types provides low-level structs for sending or receiving information
// to or from the Prosper APIs.
package types

import "time"

// AccountInformation contains the information about the user's Prosper account
// retrieved from the Account API.
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
