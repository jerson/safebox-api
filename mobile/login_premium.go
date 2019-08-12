package safebox

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/services"
)

// LoginPremium ...
func (s *SafeBox) LoginPremium(username, password string) (*AuthResponse, error) {

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

	productResponse, err := client.HasProduct(context.Background(), &services.HasProductRequest{
		AccessToken: response.AccessToken,
		Slug:        "wearableaccess",
	})
	if productResponse == nil || !productResponse.Purchased {
		return nil, errors.New("premium required")
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
