package prosper

import (
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

// BidStatusValue represents the status of an order. The values correspond to
// the values of the bid_status attribute defined at:
// https://developers.prosper.com/docs/investor/orders-api/
type BidStatusValue int8

// Set of possible BidStatusValue values.
const (
	Pending BidStatusValue = iota
	Invested
	Expired
)

// BidResult represents the result status of an order. The values correspond to
// the values of the bid_result attribute defined at:
// https://developers.prosper.com/docs/investor/orders-api/
type BidResult int8

// Set of possible BidResult values.
const (
	NoBidResult BidResult = iota
	AmountBidTooHigh
	AmountBidTooLow
	BidFailed
	BidSucceeded
	CannotBidOnSelf
	InsufficientFunds
	InternalError
	InvestmentOrderAlreadyProcessed
	LenderNotEligibleToBid
	ListingNotBiddable
	SuitabilityRequirementsNotMet
	PartialBidSucceeded
)

// BidRequest represents an order for a given Prosper listing.
type BidRequest struct {
	ListingID ListingNumber
	BidAmount float64
}

// BidStatus represents the status of a bid that has been placed for a listing.
type BidStatus struct {
	BidRequest
	Status          BidStatusValue
	Result          BidResult
	BidAmountPlaced float64
}

// OrderStatus represents the status of an order the user has placed for one or
// more listings. The values correspond to the values of the order_status
// attribute defined at:
// https://developers.prosper.com/docs/investor/orders-api/
type OrderStatus int8

// Set of possible OrderStatus values.
const (
	OrderInProgress OrderStatus = iota
	OrderCompleted
)

// OrderID is the unique identifier associated with a Prosper order request.
type OrderID string

// OrderIDs is a slice of OrderID objects.
type OrderIDs []OrderID

// Len returns the length of an OrderID slice.
func (slice OrderIDs) Len() int {
	return len(slice)
}

// Less returns whether OrderID i is less than OrderID j.
func (slice OrderIDs) Less(i, j int) bool {
	return slice[i] < slice[j]
}

// Swap swaps the position of two elements in an OrderIDs slice.
func (slice OrderIDs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// OrderResponse represents the response from the Prosper Order APIs, defined
// at: https://developers.prosper.com/docs/investor/orders-api/
type OrderResponse struct {
	OrderID     OrderID
	BidStatus   []BidStatus
	OrderStatus OrderStatus
	OrderDate   time.Time
}

// BidPlacer places a bid on the given listing for the requested amount.
type BidPlacer interface {
	PlaceBid(BidRequest) (OrderResponse, error)
}

// PlaceBid places a bid for the given listing at the given bid amount.
func (c defaultClient) PlaceBid(b BidRequest) (OrderResponse, error) {
	rawResponse, err := c.rawClient.PlaceBid([]thin.BidRequest{
		{
			ListingID: int64(b.ListingID),
			BidAmount: b.BidAmount,
		},
	})
	if err != nil {
		return OrderResponse{}, err
	}
	return c.orderParser.Parse(rawResponse)
}

// OrderStatusQuerier retrieves the status of a previously placed order.
type OrderStatusQuerier interface {
	OrderStatus(orderID OrderID) (OrderResponse, error)
}

// OrderStatus retrieves the status of the given Propser Order ID.
func (c defaultClient) OrderStatus(orderID OrderID) (OrderResponse, error) {
	rawResponse, err := c.rawClient.OrderStatus(string(orderID))
	if err != nil {
		return OrderResponse{}, err
	}
	return c.orderParser.Parse(rawResponse)
}
