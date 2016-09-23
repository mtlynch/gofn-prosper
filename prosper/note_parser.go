package prosper

import (
	"fmt"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

type noteParser interface {
	Parse(thin.NoteResult) (types.Note, error)
}

type defaultNoteParser struct{}

func (p defaultNoteParser) Parse(r thin.NoteResult) (types.Note, error) {
	originationDate, err := parseProsperDate(r.OriginationDate)
	if err != nil {
		return types.Note{}, err
	}
	nextPaymentDueDate, err := parseProsperDate(r.NextPaymentDueDate)
	if err != nil {
		return types.Note{}, err
	}
	defaultReason, err := parseDefaultReason(r.NoteDefaultReason)
	if err != nil {
		return types.Note{}, err
	}
	prosperRating, err := parseProsperRating(r.ProsperRating)
	if err != nil {
		return types.Note{}, err
	}
	noteStatus, err := parseNoteStatus(r.NoteStatus)
	if err != nil {
		return types.Note{}, err
	}
	return types.Note{
		AgeInMonths:                          r.AgeInMonths,
		AmountBorrowed:                       r.AmountBorrowed,
		BorrowerRate:                         r.BorrowerRate,
		DaysPastDue:                          r.DaysPastDue,
		DebtSaleProceedsReceivedProRataShare: r.DebtSaleProceedsReceivedProRataShare,
		InterestPaidProRataShare:             r.InterestPaidProRataShare,
		IsSold: r.IsSold,
		LateFeesPaidProRataShare:         r.LateFeesPaidProRataShare,
		ListingNumber:                    types.ListingNumber(r.ListingNumber),
		LoanNoteId:                       r.LoanNoteId,
		LoanNumber:                       r.LoanNumber,
		NextPaymentDueAmountProRataShare: r.NextPaymentDueAmountProRataShare,
		NextPaymentDueDate:               nextPaymentDueDate,
		NoteDefaultReasonDescription:     r.NoteDefaultReasonDescription,
		NoteDefaultReason:                defaultReason,
		NoteOwnershipAmount:              r.NoteOwnershipAmount,
		NoteSaleFeesPaid:                 r.NoteSaleFeesPaid,
		NoteSaleGrossAmountReceived:      r.NoteSaleGrossAmountReceived,
		NoteStatusDescription:            r.NoteStatusDescription,
		NoteStatus:                       noteStatus,
		OriginationDate:                  originationDate,
		PrincipalBalanceProRataShare:     r.PrincipalBalanceProRataShare,
		PrincipalPaidProRataShare:        r.PrincipalPaidProRataShare,
		ProsperFeesPaidProRataShare:      r.ProsperFeesPaidProRataShare,
		ProsperRating:                    prosperRating,
		ServiceFeesPaidProRataShare:      r.ServiceFeesPaidProRataShare,
		Term: r.Term,
	}, nil
}

func parseDefaultReason(defaultReason int64) (*types.DefaultReason, error) {
	if defaultReason == 0 {
		return nil, nil
	}
	if defaultReason < int64(types.DefaultReasonMin) || defaultReason > int64(types.DefaultReasonMax) {
		return nil, fmt.Errorf("default reason out of range: %d, expected %d-%d", defaultReason, types.DefaultReasonMin, types.DefaultReasonMax)
	}
	dr := types.DefaultReason(defaultReason)
	return &dr, nil
}

func parseProsperRating(rating string) (types.ProsperRating, error) {
	stringToRating := map[string]types.ProsperRating{
		"AA":  types.RatingAA,
		"A":   types.RatingA,
		"B":   types.RatingB,
		"C":   types.RatingC,
		"D":   types.RatingD,
		"E":   types.RatingE,
		"HR":  types.RatingHR,
		"N/A": types.RatingNA,
	}
	parsed, ok := stringToRating[rating]
	if !ok {
		return types.RatingNA, fmt.Errorf("unrecognized Prosper rating value: %s", rating)
	}
	return parsed, nil
}

func parseNoteStatus(noteStatus int64) (types.NoteStatus, error) {
	if noteStatus < int64(types.NoteStatusMin) || noteStatus > int64(types.NoteStatusMax) {
		return types.NoteStatusInvalid, fmt.Errorf("note status out of range: %d, expected %d-%d", noteStatus, types.NoteStatusMax, types.NoteStatusMax)
	}
	return types.NoteStatus(noteStatus), nil
}
