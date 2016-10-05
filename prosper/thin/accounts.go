package thin

type (

	// AccountsParams specifies the optional parameters to the Prosper accounts
	// API.
	AccountsParams struct {
		// TODO(mtlynch): Implement support for "filters" and
		// "suppress_in_flight_gross" parameters.
	}

	// AccountsResponse contains the response from the Accounts API in minimally
	// parsed form. This struct omits a few fields. See:
	// https://github.com/mtlynch/gofn-prosper/issues/13
	AccountsResponse struct {
		AvailableCashBalance                float64 `json:"available_cash_balance"`
		PendingInvestmentsPrimaryMarket     float64 `json:"pending_investments_primary_market"`
		PendingInvestmentsSecondaryMarket   float64 `json:"pending_investments_secondary_market"`
		PendingQuickInvestOrders            float64 `json:"pending_quick_invest_orders"`
		TotalPrincipalReceivedOnActiveNotes float64 `json:"total_principal_received_on_active_notes"`
		TotalAmountInvestedOnActiveNotes    float64 `json:"total_amount_invested_on_active_notes"`
		OutstandingPrincipalOnActiveNotes   float64 `json:"outstanding_principal_on_active_notes"`
		TotalAccountValue                   float64 `json:"total_account_value"`
		InflightGross                       float64 `json:"inflight_gross"`
		LastDepositAmount                   float64 `json:"last_deposit_amount"`
		LastDepositDate                     string  `json:"last_deposit_date"`
		LastWithdrawAmount                  float64 `json:"last_withdraw_amount"`
		LastWithdrawDate                    string  `json:"last_withdraw_date"`
		ExternalUserID                      string  `json:"external_user_id"`
	}
)

// Accounts queries the Propser API for properties of the user's account,
// including balance information and note summaries. Accounts partially
// implements the REST API described at:
// https://developers.prosper.com/docs/investor/accounts-api/
func (c Client) Accounts(AccountsParams) (response AccountsResponse, err error) {
	err = c.DoRequest("GET", c.baseURL+"/accounts/prosper/", nil, &response)
	if err != nil {
		return AccountsResponse{}, err
	}
	return response, nil
}
