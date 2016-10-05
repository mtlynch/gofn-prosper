package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type accountParser interface {
	Parse(thin.AccountResponse) (types.AccountInformation, error)
}

type defaultAccountParser struct{}

func (a defaultAccountParser) Parse(r thin.AccountResponse) (types.AccountInformation, error) {
	lastDepositDate, err := parseProsperDate(r.LastDepositDate)
	if err != nil {
		return types.AccountInformation{}, err
	}
	lastWithdrawDate, err := parseProsperDate(r.LastWithdrawDate)
	if err != nil {
		return types.AccountInformation{}, err
	}
	return types.AccountInformation{
		AvailableCashBalance:                r.AvailableCashBalance,
		TotalPrincipalReceivedOnActiveNotes: r.TotalPrincipalReceivedOnActiveNotes,
		OutstandingPrincipalOnActiveNotes:   r.OutstandingPrincipalOnActiveNotes,
		LastWithdrawAmount:                  r.LastWithdrawAmount,
		LastDepositAmount:                   r.LastDepositAmount,
		LastDepositDate:                     lastDepositDate,
		PendingInvestmentsPrimaryMarket:     r.PendingInvestmentsPrimaryMarket,
		PendingInvestmentsSecondaryMarket:   r.PendingInvestmentsSecondaryMarket,
		PendingQuickInvestOrders:            r.PendingQuickInvestOrders,
		TotalAmountInvestedOnActiveNotes:    r.TotalAmountInvestedOnActiveNotes,
		TotalAccountValue:                   r.TotalAccountValue,
		InflightGross:                       r.InflightGross,
		LastWithdrawDate:                    lastWithdrawDate,
	}, nil
}
