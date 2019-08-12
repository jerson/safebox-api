package safebox

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/services"
)

// RefreshToken ...
func (s *SafeBox) RefreshToken() (*AuthResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if s.Session.response == nil {
		return nil, errors.New("must login first")
	}

	client := services.NewServicesClient(conn)

	request := &services.RefreshTokenRequest{
		AccessToken: s.Session.response.AccessToken,
	}
	response, err := client.RefreshToken(context.Background(), request)
	if err != nil {
		return nil, err
	}

	s.Session.login(response)
	return &AuthResponse{
		AccessToken: response.AccessToken,
		DateExpire:  response.DateExpire,
		Date:        response.Date,
		KeyPair: &KeyPairResponse{
			PublicKey:  response.KeyPair.PublicKey,
			PrivateKey: response.KeyPair.PrivateKey,
		},
	}, nil
}
