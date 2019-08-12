package safebox

import (
	"context"
	"safebox.jerson.dev/api/services"
)

// Login ...
func (s *SafeBox) Login(username, password string) (*AuthResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := services.NewServicesClient(conn)

	request := &services.LoginRequest{
		Username: username,
		Password: password,
	}
	response, err := client.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}
	s.Session.login(response)
	s.Session.setPassword(password)
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
