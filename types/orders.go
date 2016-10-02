package types

import "time"

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
	ListingId ListingNumber
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

const (
	OrderInProgress OrderStatus = iota
	OrderCompleted
)

// OrderId is the unique identifier associated with a Prosper order request.
type OrderId string

// OrderResponse represents the response from the Prosper Order APIs, defined
// at: https://developers.prosper.com/docs/investor/orders-api/
type OrderResponse struct {
	OrderId     OrderId
	BidStatus   []BidStatus
	OrderStatus OrderStatus
	OrderDate   time.Time
}
