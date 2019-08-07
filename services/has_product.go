package services

import (
	"context"
	"errors"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

// HasProduct ...
func (s *Server) HasProduct(context context.Context, in *HasProductRequest) (*HasProductResponse, error) {

	ctx := appContext.NewContext(context, "HasProduct")
	defer ctx.Close()

	log := ctx.GetLogger("RPC")

	user, err := getUserByToken(ctx, in.AccessToken)
	if err != nil {
		log.Error(err)
		return nil, errors.New("session has expired")
	}

	productRepository := repositories.NewProductRepository(ctx)
	product, err := productRepository.FindOneBySlug(in.Slug)
	if err != nil {
		log.Error(err)
		return nil, errors.New("product not found")
	}

	response := HasProductResponse{}
	repository := repositories.NewPurchaseRepository(ctx)
	purchase, err := repository.FindOneByUser(user.ID, product.ID)
	if purchase != nil {
		response.Purchased = true
		response.Date = purchase.Date.Format(time.RFC3339)
	} else {
		response.Purchased = false
	}

	if err != nil {
		log.Debug(err)
	}

	return &response, nil
}
