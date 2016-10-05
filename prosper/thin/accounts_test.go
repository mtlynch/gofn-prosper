package thin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountsSuccessfulResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/accounts/prosper/",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `{
				"available_cash_balance": 22139.89,
				"pending_investments_primary_market": 5700,
				"pending_investments_secondary_market": 0,
				"pending_quick_invest_orders": 0,
				"total_principal_received_on_active_notes": 95460.28,
				"total_amount_invested_on_active_notes": 394096.56,
				"outstanding_principal_on_active_notes": 298636.28,
				"total_account_value": 326476.16,
				"inflight_gross": 12345.67,
				"last_deposit_amount": 20000,
				"last_deposit_date": "2015-10-23",
				"last_withdraw_amount": 5400,
				"last_withdraw_date": "2015-10-02"
			}`)
		},
	)

	client := Client{
		baseURL:      server.URL,
		tokenManager: mockTokenManager{},
	}
	got, err := client.Accounts(AccountsParams{})
	if err != nil {
		t.Fatalf("client.Accounts failed: %v", err)
	}

	want := AccountsResponse{
		AvailableCashBalance:                22139.89,
		PendingInvestmentsPrimaryMarket:     5700,
		PendingInvestmentsSecondaryMarket:   0,
		PendingQuickInvestOrders:            0,
		TotalPrincipalReceivedOnActiveNotes: 95460.28,
		TotalAmountInvestedOnActiveNotes:    394096.56,
		OutstandingPrincipalOnActiveNotes:   298636.28,
		TotalAccountValue:                   326476.16,
		InflightGross:                       12345.67,
		LastDepositAmount:                   20000,
		LastDepositDate:                     "2015-10-23",
		LastWithdrawAmount:                  5400,
		LastWithdrawDate:                    "2015-10-02",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("client.Accounts returned %#v, want %#v", got, want)
	}
}

func TestAccountsErrorResponse(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/accounts/prosper/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "mock server error: request failed")
		},
	)
	client := Client{
		baseURL:      server.URL,
		tokenManager: mockTokenManager{},
	}
	_, err := client.Accounts(AccountsParams{})
	if err == nil {
		t.Fatal("client.Accounts should fail when server returns error")
	}
}
