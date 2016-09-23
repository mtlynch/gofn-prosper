package types

import "time"

type BidStatusValue int8

const (
	Pending BidStatusValue = iota
	Invested
	Expired
)

type BidResult int8

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

type BidRequest struct {
	ListingId ListingNumber
	BidAmount float64
}

type BidStatus struct {
	BidRequest
	Status          BidStatusValue
	Result          BidResult
	BidAmountPlaced float64
}

type OrderStatus int8

const (
	OrderInProgress OrderStatus = iota
	OrderCompleted
)

type OrderId string

type OrderResponse struct {
	OrderId     OrderId
	BidStatus   []BidStatus
	OrderStatus OrderStatus
	OrderDate   time.Time
}
