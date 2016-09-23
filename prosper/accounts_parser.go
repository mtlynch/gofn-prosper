package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type accountsParser interface {
	Parse(thin.AccountsResponse) (types.AccountInformation, error)
}

type defaultAccountsParser struct{}

func (a defaultAccountsParser) Parse(r thin.AccountsResponse) (types.AccountInformation, error) {
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
