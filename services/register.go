package services

import (
	"context"
	"log"
)

// Register ...
func (s *Server) Register(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error) {
	log.Printf("Received")
	return &RegisterResponse{AccessToken: "dd"}, nil
}
