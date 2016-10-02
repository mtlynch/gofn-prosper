// Package thin is a very thin client implementation of the Prosper REST APIs.
// It differs from the higher level "prosper" package in that this package
// performs minimal parsing and type conversion on the raw JSON strings that the
// Prosper REST APIs return, while the "prosper" package parses the responses
// into native types.
package thin

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/mtlynch/gofn-prosper/prosper/auth"
)

const baseProsperUrl = "https://api.prosper.com/v1"

// Client is a client that communicates with the Prosper REST APIs.
type Client struct {
	baseUrl      string
	tokenManager auth.TokenManager
}

// NewClient creates a new Client instance with the given token manager.
func NewClient(t auth.TokenManager) Client {
	return Client{
		baseUrl:      baseProsperUrl,
		tokenManager: t,
	}
}

// DoRequest performs a single HTTP request against the Prosper server and
// returns the result of the request.
func (c Client) DoRequest(method, urlStr string, body io.Reader, response interface{}) error {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return err
	}
	accessToken, err := c.token()
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			msgCleaned := regexp.MustCompile(`\n\s*`).ReplaceAllString(string(body), " ")
			return errors.New("request failed: " + resp.Status + " -" + msgCleaned)
		} else {
			return errors.New("request failed: " + resp.Status)
		}
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}
	return nil
}

// RawApiHandler is a thin implementation of the Prosper REST APIs.
type RawApiHandler interface {
	Accounts() (AccountsResponse, error)
	Notes(offset, limit int) (NotesResponse, error)
	Search(SearchParams) (SearchResponse, error)
	PlaceBid([]BidRequest) (OrderResponse, error)
	OrderStatus(string) (OrderResponse, error)
}

func (c Client) token() (string, error) {
	token, err := c.tokenManager.Token()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}
