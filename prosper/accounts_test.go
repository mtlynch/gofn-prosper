package prosper

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

func (c *mockRawClient) Accounts() (thin.AccountsResponse, error) {
	return c.accountsResponse, c.err
}

type mockAccountsParser struct {
	accountsResponseGot thin.AccountsResponse
	accountInformation  types.AccountInformation
	err                 error
}

func (p *mockAccountsParser) Parse(r thin.AccountsResponse) (types.AccountInformation, error) {
	p.accountsResponseGot = r
	return p.accountInformation, p.err
}

func TestAccountSuccess(t *testing.T) {
	a := thin.AccountsResponse{
		LastDepositAmount: 250,
		LastDepositDate:   "2015-10-05",
	}
	want := types.AccountInformation{
		LastDepositAmount: 250,
		LastDepositDate:   time.Date(2015, 10, 5, 0, 0, 0, 0, time.UTC),
	}
	parser := mockAccountsParser{accountInformation: want}
	client := Client{
		rawClient: &mockRawClient{accountsResponse: a},
		ap:        &parser,
	}
	got, err := client.Account()
	if err != nil {
		t.Errorf("Client.Account failed with %v", err)
	}
	if got != want {
		t.Errorf("Client.Account got %#v, want %#v", got, want)
	}
	if parser.accountsResponseGot != a {
		t.Errorf("parser got: %v, want %v", parser.accountsResponseGot, a)
	}
}

func TestAccountFailsWhenRawClientFails(t *testing.T) {
	parser := mockAccountsParser{}
	client := Client{
		rawClient: &mockRawClient{err: mockRawClientErr},
		ap:        &parser,
	}
	_, err := client.Account()
	if err != mockRawClientErr {
		t.Errorf("Client.Account err got: %v, want: %v", err, mockRawClientErr)
	}
	if !reflect.DeepEqual(parser.accountsResponseGot, thin.AccountsResponse{}) {
		t.Errorf("Client.Account should not attempt to parse when raw client fails.")
	}
}

func TestAccountFailsWhenParserFails(t *testing.T) {
	parserErr := errors.New("mock parser error")
	parser := mockAccountsParser{err: parserErr}
	client := Client{
		rawClient: &mockRawClient{},
		ap:        &parser,
	}
	_, err := client.Account()
	if err != parserErr {
		t.Errorf("Client.Account err got: %v, want: %v", err, parserErr)
	}
}
