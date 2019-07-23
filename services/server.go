package services

import (
	"context"
	"log"
)

// Server ...
type Server struct{}

// Ping ...
func (s *Server) Ping(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	log.Printf("Received")
	return &PingResponse{Name: "Hello "}, nil
}