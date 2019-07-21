package services

import (
	"context"
	"log"
)

// Server ...
type Server struct{}

// Ping ...
func (s *Server) Ping(ctx context.Context, in *PingRequest) (*PingReply, error) {
	log.Printf("Received")
	return &PingReply{Name: "Hello "}, nil
}

// Login ...
func (s *Server) Login(ctx context.Context, in *LoginRequest) (*HelloReply, error) {
	log.Printf("Received")
	return &HelloReply{Message: "Hello "}, nil
}

// Register ...
func (s *Server) Register(ctx context.Context, in *RegisterRequest) (*HelloReply, error) {
	log.Printf("Received")
	return &HelloReply{Message: "Hello "}, nil
}
