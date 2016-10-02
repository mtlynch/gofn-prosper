package prosper

import (
	"fmt"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type orderParser interface {
	Parse(thin.OrderResponse) (types.OrderResponse, error)
}

type defaultOrderParser struct{}

func (p defaultOrderParser) Parse(r thin.OrderResponse) (types.OrderResponse, error) {
	orderDate, err := parseProsperTime(r.OrderDate)
	if err != nil {
		return types.OrderResponse{}, err
	}
	bidStatus, err := parseBidStatusSlice(r.BidStatus)
	if err != nil {
		return types.OrderResponse{}, err
	}
	orderStatus, err := parseOrderStatus(r.OrderStatus)
	if err != nil {
		return types.OrderResponse{}, err
	}
	return types.OrderResponse{
		OrderID:     types.OrderID(r.OrderID),
		BidStatus:   bidStatus,
		OrderStatus: orderStatus,
		OrderDate:   orderDate,
	}, nil
}

func parseBidStatusSlice(status []thin.BidStatus) (parsed []types.BidStatus, err error) {
	for _, s := range status {
		sParsed, err := parseBidStatus(s)
		if err != nil {
			return []types.BidStatus{}, err
		}
		parsed = append(parsed, sParsed)
	}
	return parsed, nil
}

func parseBidStatus(s thin.BidStatus) (types.BidStatus, error) {
	bidStatus, err := parseBidStatusValue(s.Status)
	if err != nil {
		return types.BidStatus{}, err
	}
	result, err := parseBidResult(s.BidResult)
	if err != nil {
		return types.BidStatus{}, err
	}
	return types.BidStatus{
		BidRequest: types.BidRequest{
			ListingID: types.ListingNumber(s.ListingID),
			BidAmount: s.BidAmount,
		},
		Status:          bidStatus,
		Result:          result,
		BidAmountPlaced: s.BidAmountPlaced,
	}, nil
}

func parseBidStatusValue(status string) (types.BidStatusValue, error) {
	stringToBidStatusValue := map[string]types.BidStatusValue{
		"PENDING":  types.Pending,
		"INVESTED": types.Invested,
		"EXPIRED":  types.Expired,
	}
	parsed, ok := stringToBidStatusValue[status]
	if !ok {
		return parsed, fmt.Errorf("unrecognized bid status value: %s", status)
	}
	return parsed, nil
}

func parseBidResult(result string) (types.BidResult, error) {
	stringToBidResult := map[string]types.BidResult{
		"":                                   types.NoBidResult,
		"NONE":                               types.NoBidResult,
		"AMOUNT_BID_TOO_HIGH":                types.AmountBidTooHigh,
		"AMOUNT_BID_TOO_LOW":                 types.AmountBidTooLow,
		"BID_FAILED":                         types.BidFailed,
		"BID_SUCCEEDED":                      types.BidSucceeded,
		"CANNOT_BID_ON_SELF":                 types.CannotBidOnSelf,
		"INSUFFICIENT_FUNDS":                 types.InsufficientFunds,
		"INTERNAL_ERROR":                     types.InternalError,
		"INVESTMENT_ORDER_ALREADY_PROCESSED": types.InvestmentOrderAlreadyProcessed,
		"LENDER_NOT_ELIGIBLE_TO_BID":         types.LenderNotEligibleToBid,
		"LISTING_NOT_BIDDABLE":               types.ListingNotBiddable,
		"SUITABILITY_REQUIREMENTS_NOT_MET":   types.SuitabilityRequirementsNotMet,
		"PARTIAL_BID_SUCCEEDED":              types.PartialBidSucceeded,
	}
	parsed, ok := stringToBidResult[result]
	if !ok {
		return parsed, fmt.Errorf("unrecognized bid result value: %s", result)
	}
	return parsed, nil
}

func parseOrderStatus(status string) (types.OrderStatus, error) {
	stringToOrderStatus := map[string]types.OrderStatus{
		"IN_PROGRESS": types.OrderInProgress,
		"COMPLETED":   types.OrderCompleted,
	}
	parsed, ok := stringToOrderStatus[status]
	if !ok {
		return parsed, fmt.Errorf("unrecognized order status value: %s", status)
	}
	return parsed, nil
}
