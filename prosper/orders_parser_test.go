package prosper

import (
	"reflect"
	"testing"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

func TestOrderResponseParser(t *testing.T) {
	var tests = []struct {
		input         thin.OrderResponse
		want          types.OrderResponse
		expectSuccess bool
		msg           string
	}{
		{
			input: thin.OrderResponse{
				OrderId: "eba54767-d3d6-4b91-a0ba-cafaeb551f63",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 4891804,
							BidAmount: 25.0,
						},
						Status: "PENDING",
					},
				},
				OrderStatus: "IN_PROGRESS",
				OrderDate:   "2016-03-25 00:18:04 +0000",
			},
			want: types.OrderResponse{
				OrderId: "eba54767-d3d6-4b91-a0ba-cafaeb551f63",
				BidStatus: []types.BidStatus{
					{
						BidRequest: types.BidRequest{
							ListingId: 4891804,
							BidAmount: 25.0,
						},
						Status: types.Pending,
					},
				},
				OrderStatus: types.OrderInProgress,
				OrderDate:   time.Date(2016, 03, 25, 0, 18, 04, 0, time.UTC),
			},
			expectSuccess: true,
			msg:           "order should parse successfully when bid result is missing",
		},
		{
			input: thin.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 2211270,
							BidAmount: 100.0,
						},
						Status:          "INVESTED",
						BidResult:       "BID_SUCCEEDED",
						BidAmountPlaced: 100.0,
					},
				},
				OrderStatus: "COMPLETED",
				OrderDate:   "2015-09-17 19:54:58 +0000",
			},
			want: types.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []types.BidStatus{
					{
						BidRequest: types.BidRequest{
							ListingId: 2211270,
							BidAmount: 100.0,
						},
						BidAmountPlaced: 100.0,
						Status:          types.Invested,
						Result:          types.BidSucceeded,
					},
				},
				OrderStatus: types.OrderCompleted,
				OrderDate:   time.Date(2015, 9, 17, 19, 54, 58, 0, time.UTC),
			},
			expectSuccess: true,
			msg:           "valid completed order should parse successfully",
		},
		{
			input: thin.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 3211275,
							BidAmount: 50.0,
						},
						Status:          "EXPIRED",
						BidResult:       "INSUFFICIENT_FUNDS",
						BidAmountPlaced: 0.0,
					},
				},
				OrderStatus: "COMPLETED",
				OrderDate:   "2016-01-24 12:32:05 +0000",
			},
			want: types.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []types.BidStatus{
					{
						BidRequest: types.BidRequest{
							ListingId: 3211275,
							BidAmount: 50.0,
						},
						BidAmountPlaced: 0.0,
						Status:          types.Expired,
						Result:          types.InsufficientFunds,
					},
				},
				OrderStatus: types.OrderCompleted,
				OrderDate:   time.Date(2016, 1, 24, 12, 32, 5, 0, time.UTC),
			},
			expectSuccess: true,
			msg:           "valid failed order should parse successfully",
		},
		{
			input: thin.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 2211270,
							BidAmount: 100.0,
						},
						Status:          "invalid bid status",
						BidResult:       "BID_SUCCEEDED",
						BidAmountPlaced: 100.0,
					},
				},
				OrderStatus: "COMPLETED",
				OrderDate:   "2015-09-17 19:54:58 +0000",
			},
			expectSuccess: false,
			msg:           "invalid bid status field should cause error",
		},
		{
			input: thin.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 2211270,
							BidAmount: 100.0,
						},
						Status:          "INVESTED",
						BidResult:       "BID_SUCCEEDED",
						BidAmountPlaced: 100.0,
					},
				},
				OrderStatus: "COMPLETED",
				OrderDate:   "invalid order date",
			},
			expectSuccess: false,
			msg:           "invalid OrderDate should cause error",
		},
		{
			input: thin.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 2211270,
							BidAmount: 100.0,
						},
						Status:          "INVESTED",
						BidResult:       "BID_SUCCEEDED",
						BidAmountPlaced: 100.0,
					},
				},
				OrderStatus: "invalid order status",
				OrderDate:   "2015-09-17 19:54:58 +0000",
			},
			expectSuccess: false,
			msg:           "invalid OrderStatus should cause error",
		},
		{
			input: thin.OrderResponse{
				OrderId: "90cf709d-81d6-416a-89f2-ba6ab8146ef2",
				BidStatus: []thin.BidStatus{
					{
						BidRequest: thin.BidRequest{
							ListingId: 2211270,
							BidAmount: 100.0,
						},
						Status:          "INVESTED",
						BidResult:       "invalid bid result",
						BidAmountPlaced: 100.0,
					},
				},
				OrderStatus: "COMPLETED",
				OrderDate:   "2015-09-17 19:54:58 +0000",
			},
			expectSuccess: false,
			msg:           "invalid BidResult should cause error",
		},
	}
	for _, tt := range tests {
		got, err := defaultOrderParser{}.Parse(tt.input)
		if tt.expectSuccess && err != nil {
			t.Errorf("%s - expected successful parsing of %+v, got error: %v", tt.msg, tt.input, err)
		} else if !tt.expectSuccess && err == nil {
			t.Errorf("%s - expected failure for %+v, got nil", tt.msg, tt.input)
		}
		if tt.expectSuccess && !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s - defaultOrderParser.Parse returned %#v, want %#v", tt.msg, got, tt.want)
		}
	}
}
