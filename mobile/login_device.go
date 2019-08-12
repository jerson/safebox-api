package safebox

import (
	"context"
	"safebox.jerson.dev/api/services"
)

// LoginDevice ...
func (s *SafeBox) LoginDevice(publicKey string) (*AuthResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := services.NewServicesClient(conn)

	request := &services.LoginDeviceRequest{
		PublicKey: publicKey,
	}
	response, err := client.LoginWithDevice(context.Background(), request)
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
