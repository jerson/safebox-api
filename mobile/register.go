package safebox

import (
	"context"
	"safebox.jerson.dev/api/modules/openpgp"
	"safebox.jerson.dev/api/services"
)

// Register ...
func (s *SafeBox) Register(username, password string) (*AuthResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := services.NewServicesClient(conn)

	options := &openpgp.Options{
		Name:       username,
		Passphrase: password,
	}
	pgp := openpgp.NewOpenPGP()
	keyPair, err := pgp.Generate(options)
	if err != nil {
		return nil, err
	}

	request := &services.RegisterRequest{
		Username:   username,
		PublicKey:  keyPair.PublicKey,
		PrivateKey: keyPair.PrivateKey,
	}
	response, err := client.Register(context.Background(), request)
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
