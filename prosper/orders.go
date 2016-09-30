package prosper

import (
	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type BidPlacer interface {
	PlaceBid(types.ListingNumber, float64) (types.OrderResponse, error)
}

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

type OrderStatusQuerier interface {
	OrderStatus(orderId types.OrderId) (types.OrderResponse, error)
}

func (c Client) OrderStatus(orderId types.OrderId) (types.OrderResponse, error) {
	rawResponse, err := c.rawClient.OrderStatus(string(orderId))
	if err != nil {
		return types.OrderResponse{}, err
	}
	return c.orderParser.Parse(rawResponse)
}
