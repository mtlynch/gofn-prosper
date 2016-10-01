package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

// BidPlacer places bids on listings for the requested amount.
type BidPlacer interface {
	PlaceBid(types.ListingNumber, float64) (types.OrderResponse, error)
}

// PlaceBid places a bid for the given listing at the given bid amount.
func (c Client) PlaceBid(listingId types.ListingNumber, bidAmount float64) (types.OrderResponse, error) {
	rawResponse, err := c.rawClient.PlaceBid([]thin.BidRequest{
		{
			ListingId: int64(listingId),
			BidAmount: bidAmount,
		},
	})
	if err != nil {
		return types.OrderResponse{}, err
	}
	return c.orderParser.Parse(rawResponse)
}

// OrderStatusQuerier retrieves the status of a previously placed order.
type OrderStatusQuerier interface {
	OrderStatus(orderId types.OrderId) (types.OrderResponse, error)
}

// OrderStatus retrieves the status of the given Propser Order ID.
func (c Client) OrderStatus(orderId types.OrderId) (types.OrderResponse, error) {
	rawResponse, err := c.rawClient.OrderStatus(string(orderId))
	if err != nil {
		return types.OrderResponse{}, err
	}
	return c.orderParser.Parse(rawResponse)
}
