package prosper

import (
	"fmt"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
)

type noteParser interface {
	Parse(thin.NoteResult) (Note, error)
}

type defaultNoteParser struct{}

func (p defaultNoteParser) Parse(r thin.NoteResult) (Note, error) {
	originationDate, err := parseProsperDate(r.OriginationDate)
	if err != nil {
		return Note{}, err
	}
	nextPaymentDueDate, err := parseProsperDate(r.NextPaymentDueDate)
	if err != nil {
		return Note{}, err
	}
	defaultReason, err := parseDefaultReason(r.NoteDefaultReason)
	if err != nil {
		return Note{}, err
	}
	prosperRating, err := parseProsperRating(r.ProsperRating)
	if err != nil {
		return Note{}, err
	}
	noteStatus, err := parseNoteStatus(r.NoteStatus)
	if err != nil {
		return Note{}, err
	}
	return Note{
		AgeInMonths:                          r.AgeInMonths,
		AmountBorrowed:                       r.AmountBorrowed,
		BorrowerRate:                         r.BorrowerRate,
		DaysPastDue:                          r.DaysPastDue,
		DebtSaleProceedsReceivedProRataShare: r.DebtSaleProceedsReceivedProRataShare,
		InterestPaidProRataShare:             r.InterestPaidProRataShare,
		IsSold: r.IsSold,
		LateFeesPaidProRataShare:         r.LateFeesPaidProRataShare,
		ListingNumber:                    ListingNumber(r.ListingNumber),
		LoanNoteID:                       r.LoanNoteID,
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

func parseDefaultReason(defaultReason int64) (*DefaultReason, error) {
	if defaultReason == 0 {
		return nil, nil
	}
	if defaultReason < int64(DefaultReasonMin) || defaultReason > int64(DefaultReasonMax) {
		return nil, fmt.Errorf("default reason out of range: %d, expected %d-%d", defaultReason, DefaultReasonMin, DefaultReasonMax)
	}
	dr := DefaultReason(defaultReason)
	return &dr, nil
}

func parseProsperRating(rating string) (ProsperRating, error) {
	stringToRating := map[string]ProsperRating{
		"AA":  RatingAA,
		"A":   RatingA,
		"B":   RatingB,
		"C":   RatingC,
		"D":   RatingD,
		"E":   RatingE,
		"HR":  RatingHR,
		"N/A": RatingNA,
	}
	parsed, ok := stringToRating[rating]
	if !ok {
		return RatingNA, fmt.Errorf("unrecognized Prosper rating value: %s", rating)
	}
	return parsed, nil
}

func parseNoteStatus(noteStatus int64) (NoteStatus, error) {
	if noteStatus < int64(NoteStatusMin) || noteStatus > int64(NoteStatusMax) {
		return NoteStatusInvalid, fmt.Errorf("note status out of range: %d, expected %d-%d", noteStatus, NoteStatusMax, NoteStatusMax)
	}
	return NoteStatus(noteStatus), nil
}
