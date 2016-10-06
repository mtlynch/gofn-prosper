package prosper

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mtlynch/gofn-prosper/prosper/thin"
	"github.com/mtlynch/gofn-prosper/types"
)

func (c *mockRawClient) Search(p thin.SearchParams) (thin.SearchResponse, error) {
	c.searchParams = p
	return c.searchResponse, c.err
}

type mockListingParser struct {
	searchResultsGot []thin.SearchResult
	listings         []Listing
	errs             []error
}

func (p *mockListingParser) Parse(r thin.SearchResult) (Listing, error) {
	p.searchResultsGot = append(p.searchResultsGot, r)
	var l Listing
	l, p.listings = p.listings[0], p.listings[1:]
	var err error
	err, p.errs = p.errs[0], p.errs[1:]
	return l, err
}

var (
	rawListingA             = thin.SearchResult{ListingNumber: 1234}
	rawListingB             = thin.SearchResult{ListingNumber: 4567}
	listingA                = Listing{ListingNumber: 1234}
	listingB                = Listing{ListingNumber: 4567}
	errMockListingParseFail = errors.New("mock listing parser error")
)

func rawSearchFilterEqual(a, b thin.SearchFilter) bool {
	if !types.Float64RangeEqual(a.EstimatedReturn, b.EstimatedReturn) {
		return false
	}
	if len(a.IncomeRange) != len(b.IncomeRange) {
		return false
	}
	for i := range a.IncomeRange {
		if a.IncomeRange[i] != b.IncomeRange[i] {
			return false
		}
	}
	if !types.Int32RangeEqual(a.InquiriesLast6Months, b.InquiriesLast6Months) {
		return false
	}
	if !types.Float64RangeEqual(a.DtiWprosperLoan, b.DtiWprosperLoan) {
		return false
	}
	if len(a.ProsperRating) != len(b.ProsperRating) {
		return false
	}
	for i := range a.ProsperRating {
		if a.ProsperRating[i] != b.ProsperRating[i] {
			return false
		}
	}
	if len(a.ListingStatus) != len(b.ListingStatus) {
		return false
	}
	for i := range a.ListingStatus {
		if a.ListingStatus[i] != b.ListingStatus[i] {
			return false
		}
	}
	return true
}

func rawSearchParamsEqual(a, b thin.SearchParams) bool {
	if a.Offset != b.Offset {
		return false
	}
	if a.Limit != b.Limit {
		return false
	}
	if a.ExcludeListingsInvested != b.ExcludeListingsInvested {
		return false
	}
	return rawSearchFilterEqual(a.Filter, b.Filter)
}

func TestSearch(t *testing.T) {
	var tests = []struct {
		searchParams        SearchParams
		wantRawSearchParams thin.SearchParams
		rawSearchResponse   thin.SearchResponse
		rawClientErr        error
		parsedListings      []Listing
		parseErrors         []error
		want                SearchResponse
		wantErr             error
		msg                 string
	}{
		{
			rawClientErr: errMockRawClientFail,
			wantErr:      errMockRawClientFail,
			msg:          "search should fail when raw client fails",
		},
		{
			rawSearchResponse: thin.SearchResponse{
				Results:     []thin.SearchResult{rawListingA},
				ResultCount: 1,
				TotalCount:  1,
			},
			parsedListings: []Listing{listingA},
			parseErrors:    []error{nil},
			want: SearchResponse{
				Results:     []Listing{listingA},
				ResultCount: 1,
				TotalCount:  1,
			},
			msg: "parsing a single result with no errors should succeed",
		},
		{
			searchParams: SearchParams{
				Offset: 25,
				Limit:  50,
				ExcludeListingsInvested: true,
				Filter: SearchFilter{
					EstimatedReturn:      types.NewFloat64Range(0.0, 0.2),
					IncomeRange:          []IncomeRange{ZeroIncome, Between0And25k},
					InquiriesLast6Months: types.NewInt32Range(1, 5),
					DtiWprosperLoan:      types.NewFloat64Range(0.0, 0.4),
					ProsperRating:        []ProsperRating{RatingA, RatingC},
					ListingStatus:        []ListingStatus{ListingActive, ListingExpired},
				},
			},
			wantRawSearchParams: thin.SearchParams{
				Offset: 25,
				Limit:  50,
				ExcludeListingsInvested: true,
				Filter: thin.SearchFilter{
					EstimatedReturn:      types.NewFloat64Range(0.0, 0.2),
					IncomeRange:          []int8{1, 2},
					InquiriesLast6Months: types.NewInt32Range(1, 5),
					DtiWprosperLoan:      types.NewFloat64Range(0.0, 0.4),
					ProsperRating:        []string{"A", "C"},
					ListingStatus:        []int{2, 5},
				},
			},
			rawSearchResponse: thin.SearchResponse{
				Results:     []thin.SearchResult{rawListingA},
				ResultCount: 1,
				TotalCount:  1,
			},
			parsedListings: []Listing{listingA},
			parseErrors:    []error{nil},
			want: SearchResponse{
				Results:     []Listing{listingA},
				ResultCount: 1,
				TotalCount:  1,
			},
			msg: "parsing a single result from search parameters should succeed",
		},
		{
			rawSearchResponse: thin.SearchResponse{
				Results:     []thin.SearchResult{rawListingA, rawListingB},
				ResultCount: 2,
				TotalCount:  2,
			},
			parsedListings: []Listing{listingA, listingB},
			parseErrors:    []error{nil, nil},
			want: SearchResponse{
				Results:     []Listing{listingA, listingB},
				ResultCount: 2,
				TotalCount:  2,
			},
			msg: "parsing multiple results with no errors should succeed",
		},
		{
			rawSearchResponse: thin.SearchResponse{
				Results:     []thin.SearchResult{rawListingA},
				ResultCount: 1,
				TotalCount:  1,
			},
			parsedListings: []Listing{listingA},
			parseErrors:    []error{errMockListingParseFail},
			wantErr:        errMockListingParseFail,
			msg:            "parsing a single result with a parse error should fail",
		},
		{
			rawSearchResponse: thin.SearchResponse{
				Results:     []thin.SearchResult{rawListingA, {}, rawListingB},
				ResultCount: 3,
				TotalCount:  3,
			},
			parsedListings: []Listing{listingA, {}, listingB},
			parseErrors:    []error{nil, errMockListingParseFail, nil},
			wantErr:        errMockListingParseFail,
			msg:            "a single parser error among successful parses should fail",
		},
	}
	for _, tt := range tests {
		rawClient := mockRawClient{
			searchResponse: tt.rawSearchResponse,
			err:            tt.rawClientErr,
		}
		lp := mockListingParser{
			listings: tt.parsedListings,
			errs:     tt.parseErrors,
		}
		c := Client{
			rawClient:     &rawClient,
			listingParser: &lp,
		}
		got, err := c.Search(tt.searchParams)
		if err != tt.wantErr {
			t.Errorf("%s: Client.Search got unexpected error. got %v, want %v", tt.msg, err, tt.wantErr)
		} else if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: Client.Search got %#v, want %#v", tt.msg, got, tt.want)
		} else if err == nil {
			if !rawSearchParamsEqual(rawClient.searchParams, tt.wantRawSearchParams) {
				t.Errorf("%s: unexpected conversion of search parameters. got %+v, want %+v", tt.msg, rawClient.searchParams, tt.wantRawSearchParams)
			}
			if !reflect.DeepEqual(tt.rawSearchResponse.Results, lp.searchResultsGot) {
				t.Errorf("%s: listing parser got unexpected value. got %v, want %v", tt.msg, lp.searchResultsGot, tt.rawSearchResponse.Results)
			}
		}
	}
}
