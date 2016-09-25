package prosper

import "github.com/mtlynch/gofn-prosper/types"

// Accounter supports the Account interface for retrieving user account
// information.
type Accounter interface {
	Account() (types.AccountInformation, error)
}

// Accounts queries the Propser API for properties of the user's account,
// including balance information and note summaries. Accounts partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/accounts-api/
func (c Client) Account() (types.AccountInformation, error) {
	rawResponse, err := c.rawClient.Accounts()
	if err != nil {
		return types.AccountInformation{}, err
	}
	return c.ap.Parse(rawResponse)
}
