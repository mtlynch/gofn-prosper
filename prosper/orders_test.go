package prosper

import (
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

var (
	gotBidRequest []thin.BidRequest
	gotOrderId    types.OrderId
)

func (c *mockRawClient) PlaceBid(br []thin.BidRequest) (thin.OrderResponse, error) {
	gotBidRequest = br
	return c.orderResponse, c.err
}

func (c *mockRawClient) OrderStatus(orderId string) (thin.OrderResponse, error) {
	gotOrderId = types.OrderId(orderId)
	return c.orderResponse, c.err
}

type mockOrderParser struct {
	gotOrderResponse thin.OrderResponse
	orderResponse    types.OrderResponse
	err              error
}

func (p *mockOrderParser) Parse(r thin.OrderResponse) (types.OrderResponse, error) {
	p.gotOrderResponse = r
	return p.orderResponse, p.err
}

const (
	listingIdA = 123
	listingIdB = 456
	bidAmountA = 25.0
	bidAmountB = 35.72
)

func TestPlaceBid(t *testing.T) {
	var tests = []struct {
		listingId           types.ListingNumber
		bidAmount           float64
		wantBidRequest      []thin.BidRequest
		rawOrderResponse    thin.OrderResponse
		clientErr           error
		parsedOrderResponse types.OrderResponse
		parserErr           error
		wantErr             error
	}{
		{
			listingId: listingIdA,
			bidAmount: bidAmountA,
			wantBidRequest: []thin.BidRequest{
				{
					ListingId: listingIdA,
					BidAmount: bidAmountA,
				},
			},
			rawOrderResponse: thin.OrderResponse{
				OrderId: "order_id_a",
			},
			parsedOrderResponse: types.OrderResponse{
				OrderId: "order_id_a",
			},
		},
		{
			listingId: listingIdB,
			bidAmount: bidAmountB,
			wantBidRequest: []thin.BidRequest{
				{
					ListingId: listingIdB,
					BidAmount: bidAmountB,
				},
			},
			rawOrderResponse: thin.OrderResponse{
				OrderId: "order_id_b",
			},
			parsedOrderResponse: types.OrderResponse{
				OrderId: "order_id_b",
			},
		},
		{
			listingId: listingIdA,
			bidAmount: bidAmountA,
			clientErr: mockRawClientErr,
			wantErr:   mockRawClientErr,
		},
		{
			listingId: listingIdA,
			bidAmount: bidAmountA,
			parserErr: mockRawClientErr,
			wantErr:   mockRawClientErr,
		},
	}
	for _, tt := range tests {
		mockRawClient := mockRawClient{
			orderResponse: tt.rawOrderResponse,
			err:           tt.clientErr,
		}
		mockParser := mockOrderParser{
			orderResponse: tt.parsedOrderResponse,
			err:           tt.parserErr,
		}
		c := Client{
			rawClient:   &mockRawClient,
			orderParser: &mockParser,
		}
		got, err := c.PlaceBid(tt.listingId, tt.bidAmount)
		if err != tt.wantErr {
			t.Errorf("unexpected error from PlaceBid. got: %v, want: %v", err, tt.wantErr)
		} else if tt.wantErr == nil {
			if !reflect.DeepEqual(gotBidRequest, tt.wantBidRequest) {
				t.Errorf("unexpected bid request. got: %+v, want: %+v", gotBidRequest, tt.wantBidRequest)
			}
			if !reflect.DeepEqual(mockParser.gotOrderResponse, tt.rawOrderResponse) {
				t.Errorf("unexpected order response sent to parser. got: %+v, want: %+v", mockParser.gotOrderResponse, tt.rawOrderResponse)
			}
			if !reflect.DeepEqual(got, tt.parsedOrderResponse) {
				t.Errorf("unexpected parsed order response. got: %+v, want: %+v", got, tt.parsedOrderResponse)
			}
		}
	}
}

const (
	orderIdA = "orderA"
	orderIdB = "orderB"
)

func TestOrderStatus(t *testing.T) {
	var tests = []struct {
		orderId             types.OrderId
		rawOrderResponse    thin.OrderResponse
		clientErr           error
		parsedOrderResponse types.OrderResponse
		parserErr           error
		wantErr             error
	}{
		{
			orderId: orderIdA,
			rawOrderResponse: thin.OrderResponse{
				OrderId: orderIdA,
			},
			parsedOrderResponse: types.OrderResponse{
				OrderId: orderIdA,
			},
		},
		{
			orderId: orderIdB,
			rawOrderResponse: thin.OrderResponse{
				OrderId: orderIdB,
			},
			parsedOrderResponse: types.OrderResponse{
				OrderId: orderIdB,
			},
		},
		{
			orderId:   orderIdA,
			clientErr: mockRawClientErr,
			wantErr:   mockRawClientErr,
		},
		{
			orderId:   orderIdA,
			parserErr: mockRawClientErr,
			wantErr:   mockRawClientErr,
		},
	}
	for _, tt := range tests {
		mockRawClient := mockRawClient{
			orderResponse: tt.rawOrderResponse,
			err:           tt.clientErr,
		}
		mockParser := mockOrderParser{
			orderResponse: tt.parsedOrderResponse,
			err:           tt.parserErr,
		}
		c := Client{
			rawClient:   &mockRawClient,
			orderParser: &mockParser,
		}
		got, err := c.OrderStatus(tt.orderId)
		if err != tt.wantErr {
			t.Errorf("unexpected error from OrderStatus. got: %v, want: %v", err, tt.wantErr)
		} else if tt.wantErr == nil {
			if gotOrderId != tt.orderId {
				t.Errorf("unexpected order ID. got: %+v, want: %+v", gotOrderId, tt.orderId)
			}
			if !reflect.DeepEqual(mockParser.gotOrderResponse, tt.rawOrderResponse) {
				t.Errorf("unexpected order response sent to parser. got: %+v, want: %+v", mockParser.gotOrderResponse, tt.rawOrderResponse)
			}
			if !reflect.DeepEqual(got, tt.parsedOrderResponse) {
				t.Errorf("unexpected parsed order response. got: %+v, want: %+v", got, tt.parsedOrderResponse)
			}
		}
	}
}
