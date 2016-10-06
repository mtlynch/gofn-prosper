package prosper

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

func (c *mockRawClient) Accounts(thin.AccountParams) (thin.AccountResponse, error) {
	return c.accountsResponse, c.err
}

type mockAccountParser struct {
	accountsResponseGot thin.AccountResponse
	accountInformation  AccountInformation
	err                 error
}

func (p *mockAccountParser) Parse(r thin.AccountResponse) (AccountInformation, error) {
	p.accountsResponseGot = r
	return p.accountInformation, p.err
}

func TestAccountSuccess(t *testing.T) {
	a := thin.AccountResponse{
		LastDepositAmount: 250,
		LastDepositDate:   "2015-10-05",
	}
	want := AccountInformation{
		LastDepositAmount: 250,
		LastDepositDate:   time.Date(2015, 10, 5, 0, 0, 0, 0, time.UTC),
	}
	parser := mockAccountParser{accountInformation: want}
	client := Client{
		rawClient:     &mockRawClient{accountsResponse: a},
		accountParser: &parser,
	}
	got, err := client.Account(AccountParams{})
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
	parser := mockAccountParser{}
	client := Client{
		rawClient:     &mockRawClient{err: errMockRawClientFail},
		accountParser: &parser,
	}
	_, err := client.Account(AccountParams{})
	if err != errMockRawClientFail {
		t.Errorf("Client.Account err got: %v, want: %v", err, errMockRawClientFail)
	}
	if !reflect.DeepEqual(parser.accountsResponseGot, thin.AccountResponse{}) {
		t.Errorf("Client.Account should not attempt to parse when raw client fails.")
	}
}

func TestAccountFailsWhenParserFails(t *testing.T) {
	parserErr := errors.New("mock parser error")
	parser := mockAccountParser{err: parserErr}
	client := Client{
		rawClient:     &mockRawClient{},
		accountParser: &parser,
	}
	_, err := client.Account(AccountParams{})
	if err != parserErr {
		t.Errorf("Client.Account err got: %v, want: %v", err, parserErr)
	}
}
