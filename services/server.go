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

// AddAccount ...
func (s *Server) AddAccount(ctx context.Context, in *AccountAddRequest) (*AccountAddResponse, error) {
	log.Printf("Received")
	return &AccountAddResponse{}, nil
}

// GetAccounts ...
func (s *Server) GetAccounts(ctx context.Context, in *AccountsRequest) (*AccountsResponse, error) {
	log.Printf("Received")
	return &AccountsResponse{Accounts: []*AccountSingle{}}, nil
}

// GetAccount ...
func (s *Server) GetAccount(ctx context.Context, in *AccountRequest) (*AccountResponse, error) {
	log.Printf("Received")
	return &AccountResponse{Account: &Account{}}, nil
}

// BuyProduct ...
func (s *Server) BuyProduct(ctx context.Context, in *BuyProductRequest) (*BuyProductResponse, error) {
	log.Printf("Received")
	return &BuyProductResponse{Success: true}, nil
}
