package services

import (
	"context"
)

// BuyProduct ...
func (s *Server) BuyProduct(context context.Context, in *BuyProductRequest) (*BuyProductResponse, error) {
	return &BuyProductResponse{Success: true}, nil
}
