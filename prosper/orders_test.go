package prosper

import (
	"reflect"
	"sort"
	"testing"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

var (
	gotBidRequest []thin.BidRequest
	gotOrderID    OrderID
)

func (c *mockRawClient) PlaceBid(br []thin.BidRequest) (thin.OrderResponse, error) {
	gotBidRequest = br
	return c.orderResponse, c.err
}

func (c *mockRawClient) OrderStatus(orderID string) (thin.OrderResponse, error) {
	gotOrderID = OrderID(orderID)
	return c.orderResponse, c.err
}

type mockOrderParser struct {
	gotOrderResponse thin.OrderResponse
	orderResponse    OrderResponse
	err              error
}

func (p *mockOrderParser) Parse(r thin.OrderResponse) (OrderResponse, error) {
	p.gotOrderResponse = r
	return p.orderResponse, p.err
}

const (
	listingIDA = 123
	listingIDB = 456
	bidAmountA = 25.0
	bidAmountB = 35.72
)

func TestPlaceBid(t *testing.T) {
	var tests = []struct {
		listingID           ListingNumber
		bidAmount           float64
		wantBidRequest      []thin.BidRequest
		rawOrderResponse    thin.OrderResponse
		clientErr           error
		parsedOrderResponse OrderResponse
		parserErr           error
		wantErr             error
	}{
		{
			listingID: listingIDA,
			bidAmount: bidAmountA,
			wantBidRequest: []thin.BidRequest{
				{
					ListingID: listingIDA,
					BidAmount: bidAmountA,
				},
			},
			rawOrderResponse: thin.OrderResponse{
				OrderID: "order_id_a",
			},
			parsedOrderResponse: OrderResponse{
				OrderID: "order_id_a",
			},
		},
		{
			listingID: listingIDB,
			bidAmount: bidAmountB,
			wantBidRequest: []thin.BidRequest{
				{
					ListingID: listingIDB,
					BidAmount: bidAmountB,
				},
			},
			rawOrderResponse: thin.OrderResponse{
				OrderID: "order_id_b",
			},
			parsedOrderResponse: OrderResponse{
				OrderID: "order_id_b",
			},
		},
		{
			listingID: listingIDA,
			bidAmount: bidAmountA,
			clientErr: errMockRawClientFail,
			wantErr:   errMockRawClientFail,
		},
		{
			listingID: listingIDA,
			bidAmount: bidAmountA,
			parserErr: errMockRawClientFail,
			wantErr:   errMockRawClientFail,
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
		got, err := c.PlaceBid(BidRequest{
			ListingID: tt.listingID,
			BidAmount: tt.bidAmount,
		})
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
	orderIDA = "orderA"
	orderIDB = "orderB"
)

func TestOrderStatus(t *testing.T) {
	var tests = []struct {
		orderID             OrderID
		rawOrderResponse    thin.OrderResponse
		clientErr           error
		parsedOrderResponse OrderResponse
		parserErr           error
		wantErr             error
	}{
		{
			orderID: orderIDA,
			rawOrderResponse: thin.OrderResponse{
				OrderID: orderIDA,
			},
			parsedOrderResponse: OrderResponse{
				OrderID: orderIDA,
			},
		},
		{
			orderID: orderIDB,
			rawOrderResponse: thin.OrderResponse{
				OrderID: orderIDB,
			},
			parsedOrderResponse: OrderResponse{
				OrderID: orderIDB,
			},
		},
		{
			orderID:   orderIDA,
			clientErr: errMockRawClientFail,
			wantErr:   errMockRawClientFail,
		},
		{
			orderID:   orderIDA,
			parserErr: errMockRawClientFail,
			wantErr:   errMockRawClientFail,
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
		got, err := c.OrderStatus(tt.orderID)
		if err != tt.wantErr {
			t.Errorf("unexpected error from OrderStatus. got: %v, want: %v", err, tt.wantErr)
		} else if tt.wantErr == nil {
			if gotOrderID != tt.orderID {
				t.Errorf("unexpected order ID. got: %+v, want: %+v", gotOrderID, tt.orderID)
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

func TestSortOrderIDs(t *testing.T) {
	var tests = []struct {
		unsorted OrderIDs
		want     OrderIDs
	}{
		{
			unsorted: OrderIDs{"id-a", "id-b", "id-c"},
			want:     OrderIDs{"id-a", "id-b", "id-c"},
		},
		{
			unsorted: OrderIDs{"id-c", "id-b", "id-a"},
			want:     OrderIDs{"id-a", "id-b", "id-c"},
		},
	}
	for _, tt := range tests {
		sort.Sort(tt.unsorted)
		if !reflect.DeepEqual(tt.unsorted, tt.want) {
			t.Errorf("sorting order IDs failed, got: %v, want: %v", tt.unsorted, tt.want)
		}
	}
}
