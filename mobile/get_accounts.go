package safebox

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/services"
)

// GetAccounts ...
func (s *SafeBox) GetAccounts() (*AccountSingleCollection, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if s.Session.response == nil {
		return nil, errors.New("must login first")
	}

	client := services.NewServicesClient(conn)

	request := &services.AccountsRequest{
		AccessToken: s.Session.response.AccessToken,
	}
	response, err := client.GetAccounts(context.Background(), request)
	if err != nil {
		return nil, err
	}
	var accounts []*AccountSingle

	for _, account := range response.Accounts {
		accounts = append(accounts, &AccountSingle{
			ID:       account.Id,
			Label:    account.Label,
			Username: account.Username,
			Hint:     account.Hint,
		})
	}

	return NewAccountSingleCollection(accounts), nil
}
