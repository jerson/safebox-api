package services

import (
	"context"
)

// HasProduct ...
func (s *Server) HasProduct(context context.Context, in *HasProductRequest) (*HasProductResponse, error) {
	return &HasProductResponse{}, nil
}
