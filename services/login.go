package services

import (
	"context"
	"log"
)

// Login ...
func (s *Server) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	log.Printf("Received")
	return &LoginResponse{AccessToken: "dd"}, nil
}
