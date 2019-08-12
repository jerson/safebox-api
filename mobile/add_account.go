package safebox

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/modules/openpgp"
	"safebox.jerson.dev/api/services"
)

// AddAccount ...
func (s *SafeBox) AddAccount(label, hint, username, password string) (*AddAccountResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if s.Session.response == nil {
		return nil, errors.New("must login first")
	}

	client := services.NewServicesClient(conn)

	pgp := openpgp.NewOpenPGP()
	key, err := pgp.Encrypt(password, s.Session.response.KeyPair.PublicKey)
	if err != nil {
		return nil, err
	}

	account := &services.Account{
		Label:    label,
		Username: username,
		Password: key,
		Hint:     hint,
	}
	request := &services.AddAccountRequest{
		AccessToken: s.Session.response.AccessToken,
		Account:     account,
	}
	response, err := client.AddAccount(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return &AddAccountResponse{ID: response.Id}, nil
}
