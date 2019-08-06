package services

import (
	"context"
	"errors"
	"safebox.jerson.dev/api/models"
	appContext "safebox.jerson.dev/api/modules/context"
	"safebox.jerson.dev/api/repositories"
	"time"
)

// BuyProduct ...
func (s *Server) BuyProduct(context context.Context, in *BuyProductRequest) (*BuyProductResponse, error) {

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

	purchaseInput := models.Purchase{
		UserID:    user.ID,
		ProductID: product.ID,
		Payload:   in.Payload,
	}
	repository := repositories.NewPurchaseRepository(ctx)
	purchase, err := repository.Create(purchaseInput)
	if err != nil {
		log.Error(err)
		return nil, errors.New("there was a problem, try again later")
	}

	return &BuyProductResponse{
		Id:   purchase.ID,
		Date: purchase.Date.Format(time.RFC3339),
	}, nil
}
