package prosper

import "github.com/mtlynch/gofn-prosper/types"

type Accounter interface {
	Account() (types.AccountInformation, error)
}

func (c Client) Account() (types.AccountInformation, error) {
	rawResponse, err := c.rawClient.Accounts()
	if err != nil {
		return types.AccountInformation{}, err
	}
	return c.ap.Parse(rawResponse)
}
