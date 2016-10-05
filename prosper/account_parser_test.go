package prosper

import (
	"reflect"
	"testing"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

func TestAccountParserParsesTimeCorrectly(t *testing.T) {
	got, err := defaultAccountParser{}.Parse(thin.AccountResponse{
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
	})
	want := types.AccountInformation{
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
		LastDepositDate:                     time.Date(2015, 10, 23, 0, 0, 0, 0, time.UTC),
		LastWithdrawAmount:                  5400,
		LastWithdrawDate:                    time.Date(2015, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	if err != nil {
		t.Errorf("accountParser.Parse failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("accountParser.Parse returned %#v, want %#v", got, want)
	}
}

func TestAccountParserParsesCorrectlyWhenNoWithdrawsExist(t *testing.T) {
	got, err := defaultAccountParser{}.Parse(thin.AccountResponse{
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
	})
	want := types.AccountInformation{
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
		LastDepositDate:                     time.Date(2015, 10, 23, 0, 0, 0, 0, time.UTC),
	}
	if err != nil {
		t.Errorf("accountParser.Parse failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("accountParser.Parse returned %#v, want %#v", got, want)
	}
}

func TestAccountParserParsesCorrectlyWhenNoDepositsExist(t *testing.T) {
	got, err := defaultAccountParser{}.Parse(thin.AccountResponse{
		AvailableCashBalance:                22139.89,
		PendingInvestmentsPrimaryMarket:     5700,
		PendingInvestmentsSecondaryMarket:   0,
		PendingQuickInvestOrders:            0,
		TotalPrincipalReceivedOnActiveNotes: 95460.28,
		TotalAmountInvestedOnActiveNotes:    394096.56,
		OutstandingPrincipalOnActiveNotes:   298636.28,
		TotalAccountValue:                   326476.16,
		InflightGross:                       12345.67,
		LastWithdrawAmount:                  5400,
		LastWithdrawDate:                    "2015-10-02",
	})
	want := types.AccountInformation{
		AvailableCashBalance:                22139.89,
		PendingInvestmentsPrimaryMarket:     5700,
		PendingInvestmentsSecondaryMarket:   0,
		PendingQuickInvestOrders:            0,
		TotalPrincipalReceivedOnActiveNotes: 95460.28,
		TotalAmountInvestedOnActiveNotes:    394096.56,
		OutstandingPrincipalOnActiveNotes:   298636.28,
		TotalAccountValue:                   326476.16,
		InflightGross:                       12345.67,
		LastWithdrawAmount:                  5400,
		LastWithdrawDate:                    time.Date(2015, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	if err != nil {
		t.Errorf("accountParser.Parse failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("accountParser.Parse returned %#v, want %#v", got, want)
	}
}

func TestAccountParserFailsOnInvalidLastDepositDate(t *testing.T) {
	_, err := defaultAccountParser{}.Parse(thin.AccountResponse{
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
		LastDepositDate:                     "23-10-2015",
		LastWithdrawAmount:                  5400,
		LastWithdrawDate:                    "2015-10-02",
	})
	if err == nil {
		t.Error("accountParser.Parse should fail when LastDepositDate is invalid")
	}
}

func TestAccountParserFailsOnInvalidLastWithdrawDate(t *testing.T) {
	_, err := defaultAccountParser{}.Parse(thin.AccountResponse{
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
		LastWithdrawDate:                    "02-10-2015",
	})
	if err == nil {
		t.Error("accountParser.Parse should fail when LastWithdrawDate is invalid")
	}
}
