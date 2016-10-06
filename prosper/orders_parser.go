package prosper

import (
	"fmt"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

type orderParser interface {
	Parse(thin.OrderResponse) (OrderResponse, error)
}

type defaultOrderParser struct{}

func (p defaultOrderParser) Parse(r thin.OrderResponse) (OrderResponse, error) {
	orderDate, err := parseProsperTime(r.OrderDate)
	if err != nil {
		return OrderResponse{}, err
	}
	bidStatus, err := parseBidStatusSlice(r.BidStatus)
	if err != nil {
		return OrderResponse{}, err
	}
	orderStatus, err := parseOrderStatus(r.OrderStatus)
	if err != nil {
		return OrderResponse{}, err
	}
	return OrderResponse{
		OrderID:     OrderID(r.OrderID),
		BidStatus:   bidStatus,
		OrderStatus: orderStatus,
		OrderDate:   orderDate,
	}, nil
}

func parseBidStatusSlice(status []thin.BidStatus) (parsed []BidStatus, err error) {
	for _, s := range status {
		sParsed, err := parseBidStatus(s)
		if err != nil {
			return []BidStatus{}, err
		}
		parsed = append(parsed, sParsed)
	}
	return parsed, nil
}

func parseBidStatus(s thin.BidStatus) (BidStatus, error) {
	bidStatus, err := parseBidStatusValue(s.Status)
	if err != nil {
		return BidStatus{}, err
	}
	result, err := parseBidResult(s.BidResult)
	if err != nil {
		return BidStatus{}, err
	}
	return BidStatus{
		BidRequest: BidRequest{
			ListingID: ListingNumber(s.ListingID),
			BidAmount: s.BidAmount,
		},
		Status:          bidStatus,
		Result:          result,
		BidAmountPlaced: s.BidAmountPlaced,
	}, nil
}

func parseBidStatusValue(status string) (BidStatusValue, error) {
	stringToBidStatusValue := map[string]BidStatusValue{
		"PENDING":  Pending,
		"INVESTED": Invested,
		"EXPIRED":  Expired,
	}
	parsed, ok := stringToBidStatusValue[status]
	if !ok {
		return parsed, fmt.Errorf("unrecognized bid status value: %s", status)
	}
	return parsed, nil
}

func parseBidResult(result string) (BidResult, error) {
	stringToBidResult := map[string]BidResult{
		"":                                   NoBidResult,
		"NONE":                               NoBidResult,
		"AMOUNT_BID_TOO_HIGH":                AmountBidTooHigh,
		"AMOUNT_BID_TOO_LOW":                 AmountBidTooLow,
		"BID_FAILED":                         BidFailed,
		"BID_SUCCEEDED":                      BidSucceeded,
		"CANNOT_BID_ON_SELF":                 CannotBidOnSelf,
		"INSUFFICIENT_FUNDS":                 InsufficientFunds,
		"INTERNAL_ERROR":                     InternalError,
		"INVESTMENT_ORDER_ALREADY_PROCESSED": InvestmentOrderAlreadyProcessed,
		"LENDER_NOT_ELIGIBLE_TO_BID":         LenderNotEligibleToBid,
		"LISTING_NOT_BIDDABLE":               ListingNotBiddable,
		"SUITABILITY_REQUIREMENTS_NOT_MET":   SuitabilityRequirementsNotMet,
		"PARTIAL_BID_SUCCEEDED":              PartialBidSucceeded,
	}
	parsed, ok := stringToBidResult[result]
	if !ok {
		return parsed, fmt.Errorf("unrecognized bid result value: %s", result)
	}
	return parsed, nil
}

func parseOrderStatus(status string) (OrderStatus, error) {
	stringToOrderStatus := map[string]OrderStatus{
		"IN_PROGRESS": OrderInProgress,
		"COMPLETED":   OrderCompleted,
	}
	parsed, ok := stringToOrderStatus[status]
	if !ok {
		return parsed, fmt.Errorf("unrecognized order status value: %s", status)
	}
	return parsed, nil
}
