package safebox

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/services"
)

// Logout ...
func (s *SafeBox) Logout() (*LogoutResponse, error) {

	conn, err := s.dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if s.Session.response == nil {
		return nil, errors.New("must login first")
	}

	client := services.NewServicesClient(conn)
	request := &services.LogoutRequest{
		AccessToken: s.Session.response.AccessToken,
	}
	_, err = client.Logout(context.Background(), request)
	if err != nil {
		return nil, err
	}
	s.Session.logout()
	return &LogoutResponse{}, nil
}
