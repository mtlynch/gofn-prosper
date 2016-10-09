package prosper

import (
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

type (
	// AccountParams contains the parameters to the Accounts API.
	AccountParams struct {
		// TODO(mtlynch): Implement support for "filters" and
		// "suppress_in_flight_gross" parameters.
	}

	// AccountInformation contains the information about the user's Prosper
	// account retrieved from the Account API.
	AccountInformation struct {
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

	// Accounter supports the Account interface for retrieving user account
	// information.
	Accounter interface {
		Account(AccountParams) (AccountInformation, error)
	}
)

// Account queries the Propser API for properties of the user's account,
// including balance information and note summaries. Accounts partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/accounts-api/
func (c defaultClient) Account(AccountParams) (AccountInformation, error) {
	rawResponse, err := c.rawClient.Accounts(thin.AccountParams{})
	if err != nil {
		return AccountInformation{}, err
	}
	return c.accountParser.Parse(rawResponse)
}
