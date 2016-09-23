package thin

func (c Client) Accounts() (response AccountsResponse, err error) {
	err = c.DoRequest("GET", c.baseUrl+"/accounts/prosper/", nil, &response)
	if err != nil {
		return AccountsResponse{}, err
	}
	return response, nil
}
