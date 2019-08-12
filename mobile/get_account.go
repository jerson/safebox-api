package safebox

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/modules/openpgp"
	"safebox.jerson.dev/api/services"
)

// GetAccount ...
func (s *SafeBox) GetAccount(id int64) (*AccountResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if s.Session.response == nil {
		return nil, errors.New("must login first")
	}
	if s.Session.password == "" {
		return nil, errors.New("must set password first")
	}

	client := services.NewServicesClient(conn)

	request := &services.AccountRequest{
		AccessToken: s.Session.response.AccessToken,
		Id:          id,
	}
	response, err := client.GetAccount(context.Background(), request)
	if err != nil {
		return nil, err
	}

	pgp := openpgp.NewOpenPGP()
	decodedPassword, err := pgp.Decrypt(response.Account.Password, s.Session.response.KeyPair.PrivateKey, s.Session.password)
	if err != nil {
		return nil, err
	}
	response.Account.Password = decodedPassword
	return &AccountResponse{Account: &Account{
		Label:    response.Account.Label,
		Username: response.Account.Username,
		Password: response.Account.Password,
		Hint:     response.Account.Hint,
	}}, nil
}
