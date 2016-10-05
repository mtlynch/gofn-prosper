package prosper

import "github.com/mtlynch/gofn-prosper/prosper/thin"
import "github.com/mtlynch/gofn-prosper/types"

type (
	// AccountsParams contains the parameters to the Accounts API.
	AccountsParams struct {
		// TODO(mtlynch): Implement support for "filters" and
		// "suppress_in_flight_gross" parameters.
	}

	// Accounter supports the Account interface for retrieving user account
	// information.
	Accounter interface {
		Account(AccountsParams) (types.AccountInformation, error)
	}
)

// Account queries the Propser API for properties of the user's account,
// including balance information and note summaries. Accounts partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/accounts-api/
func (c Client) Account(AccountsParams) (types.AccountInformation, error) {
	rawResponse, err := c.rawClient.Accounts(thin.AccountsParams{})
	if err != nil {
		return types.AccountInformation{}, err
	}
	return c.ap.Parse(rawResponse)
}
