package thin

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/mtlynch/gofn-prosper/auth"
)

const baseProsperUrl = "https://api.prosper.com/v1"

type Client struct {
	baseUrl      string
	tokenManager auth.TokenManager
}

func NewClient(t auth.TokenManager) Client {
	return Client{
		baseUrl:      baseProsperUrl,
		tokenManager: t,
	}
}

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
