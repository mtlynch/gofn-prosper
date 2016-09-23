package thin

type AccountsResponse struct {
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
	ExternalUserId                      string  `json:"external_user_id"`
}
